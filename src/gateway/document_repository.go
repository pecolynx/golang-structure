package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"math"
	"time"

	// casbinquery "github.com/pecolynx/casbin-query"

	// "github.com/kujilabo/cocotola/cocotola-api/src/app/gateway/casbinquery"
	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
	liblog "github.com/pecolynx/golang-structure/lib/log"
	"github.com/pecolynx/golang-structure/src/domain/model"
	"github.com/pecolynx/golang-structure/src/domain/service"
	"github.com/pecolynx/golang-structure/src/log"
	usermodel "github.com/pecolynx/golang-structure/user/domain/model"
	"gorm.io/gorm"
	// "github.com/kujilabo/cocotola/cocotola-api/src/app/gateway/casbinquery"
)

func f() {
	var Id int  // want "NG"
	println(Id) // want "NG"
}

type documentEntity struct {
	ID        string
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uint
	UpdatedBy uint
	Title     string
	Content   string
}

func (e *documentEntity) TableName() string {
	return "document"
}

func jsonToStringMap(s string) (map[string]string, error) {
	var m map[string]string
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return nil, liberrors.Errorf("json.Unmarshal. err: %w", err)
	}
	return m, nil
}

func stringMapToJSON(m map[string]string) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", liberrors.Errorf("json.Marshal. err: %w", err)
	}
	return string(b), nil
}

func (e *documentEntity) toDocumentModel() (model.DocumentModel, error) {
	baseModel, err := usermodel.NewBaseModel(e.Version, e.CreatedAt, e.UpdatedAt, e.CreatedBy, e.UpdatedBy)
	if err != nil {
		return nil, liberrors.Errorf("userD.NewModel. err: %w", err)
	}

	documetnID, err := model.NewDocumentID(e.ID)
	if err != nil {
		return nil, liberrors.Errorf("failed to NewDocumentID. value: %s, err: %w", e.ID, err)
	}

	document, err := model.NewDocumentModel(baseModel, documetnID, e.Title, e.Content)
	if err != nil {
		return nil, liberrors.Errorf("failed to NewDocument. entity: %+v, err: %w", e, err)
	}
	return document, nil
}

type documentRepository struct {
	driverName string
	db         *gorm.DB
	// rf           service.RepositoryFactory
	// pf           service.ProcessorFactory
	// problemTypes ProblemTypes
}

func newDocumentRepository(ctx context.Context, driverName string,
	//pf service.ProcessorFactory,
	db *gorm.DB,
	// rf service.RepositoryFactory, problemTypes ProblemTypes
) service.DocumentRepository {
	return &documentRepository{
		driverName: driverName,
		db:         db,
		// rf:           rf,
		// pf:           pf,
		// problemTypes: problemTypes,
	}
}

func (r *documentRepository) FindPrivateDocuments(ctx context.Context, operator model.DocumentWriterModel, param model.DocumentSearchCondition) (model.DocumentSearchResult, error) {
	ctx, span := tracer.Start(ctx, "documentRepository.FindPersonalDocuments")
	defer span.End()
	ctx = liblog.WithLoggerName(ctx, log.AppGatewayLoggerContextKey)
	logger := liblog.GetLoggerFromContext(ctx, log.AppGatewayLoggerContextKey)
	logger.DebugContext(ctx, "documentRepository.FindPrivateDocuments", slog.Uint64("OperatorID", uint64(operator.GetAppUserID())))

	if param == nil {
		return nil, libdomain.ErrInvalidArgument
	}

	limit := param.GetPageSize()
	offset := (param.GetPageNo() - 1) * param.GetPageSize()
	documentEntities := []documentEntity{}

	// objectColumnName := "name"
	// subQuery, err := casbinquery.QueryObject(r.db, r.driverName, model.DocumentObjectPrefix, objectColumnName, "user_"+strconv.Itoa(int(operator.GetID())), "read")
	// if err != nil {
	// 	return nil, liberrors.Errorf("casbinquery.QueryObject. err: %w", err)
	// }

	// if result := r.db.Model(&documentEntity{}).
	// 	Joins("inner join (?) AS t3 ON `document`.`id`= t3."+objectColumnName, subQuery).
	// 	Order("`document`.`name`").Limit(limit).Offset(offset).
	// 	Scan(&documentEntities); result.Error != nil {
	// 	return nil, result.Error
	// }

	// results := make([]model.DocumentModel, len(documentEntities))
	// priv := userD.NewPrivileges([]userD.RBACAction{model.PrivilegeRead})
	// for i, e := range documentEntities {
	// 	problemType, err := r.problemTypes.ToProblemType(e.ProblemTypeID)
	// 	if err != nil {
	// 		return nil, liberrors.Errorf("r.problemTypes.ToProblemType. err: %w", err)
	// 	}
	// 	w, err := e.toDocumentModel(r.rf, r.pf, operator, problemType, priv)
	// 	if err != nil {
	// 		return nil, liberrors.Errorf("toDocumentModel. err: %w", err)
	// 	}
	// 	results[i] = w
	// }

	where := func() *gorm.DB {
		return r.db.Where("created_by = ?", uint(operator.GetAppUserID()))
	}

	if result := where().Limit(limit).Offset(offset).Find(&documentEntities); result.Error != nil {
		return nil, result.Error
	}

	var count int64
	if result := where().Model(&documentEntity{}).Count(&count); result.Error != nil {
		return nil, result.Error
	}

	results := make([]model.DocumentModel, len(documentEntities))
	for i, e := range documentEntities {
		m, err := e.toDocumentModel()
		if err != nil {
			return nil, liberrors.Errorf("toDocumentModel. err: %w", err)
		}
		results[i] = m
	}
	if count > math.MaxInt32 {
		return nil, errors.New("overflow")
	}

	documents, err := model.NewDocumentSearchResult(int(count), results)
	if err != nil {
		return nil, liberrors.Errorf(". err: %w", err)
	}

	return documents, nil
}

// func (r *documentRepository) getAllDocumentRoles(documentID model.DocumentID) []userD.RBACRole {
// 	return []userD.RBACRole{model.NewDocumentWriter(documentID), model.NewDocumentReader(documentID)}
// }

// func (r *documentRepository) getAllDocumentPrivileges() []userD.RBACAction {
// 	return []userD.RBACAction{model.PrivilegeRead, model.PrivilegeUpdate, model.PrivilegeRemove}
// }

// func (r *documentRepository) checkPrivileges(e *casbin.Enforcer, userObject userD.RBACUser, documentObject userD.RBACObject, privs []userD.RBACAction) (userD.Privileges, error) {
// 	actions := make([]userD.RBACAction, 0)
// 	for _, priv := range privs {
// 		ok, err := e.Enforce(string(userObject), string(documentObject), string(priv))
// 		if err != nil {
// 			return nil, liberrors.Errorf("e.Enforce. err: %w", err)
// 		}
// 		if ok {
// 			actions = append(actions, priv)
// 		}
// 	}
// 	return userD.NewPrivileges(actions), nil
// }

// // func (r *documentRepository) canReadDocument(operator userD.AppUser, documentID model.DocumentID) error {
// // 	objectColumnName := "name"
// // 	object := model.DocumentObjectPrefix + strconv.Itoa(int(uint(documentID)))
// // 	subject := "user_" + strconv.Itoa(int(operator.GetID()))
// // 	casbinQuery, err := casbinquery.FindObject(r.db, r.driverName, object, objectColumnName, subject, "read")
// // 	if err != nil {
// // 		return err
// // 	}
// // 	var name string
// // 	if result := casbinQuery.First(&name); result.Error != nil {
// // 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// // 			return model.ErrDocumentPermissionDenied
// // 		}
// // 		return result.Error
// // 	}
// // 	return nil
// // }

// func (r *documentRepository) FindDocumentByID(ctx context.Context, operator model.StudentModel, documentID model.DocumentID) (service.Document, error) {
// 	ctx, span := tracer.Start(ctx, "documentRepository.FindDocumentByID")
// 	defer span.End()

// 	documentEntity := documentEntity{}
// 	if result := r.db.
// 		Where("organization_id = ?", uint(operator.GetOrganizationID())).
// 		Where("id = ?", uint(documentID)).
// 		First(&documentEntity); result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, service.ErrDocumentNotFound
// 		}
// 		return nil, result.Error
// 	}

// 	priv, err := r.getPrivileges(ctx, operator, model.DocumentID(documentEntity.ID))
// 	if err != nil {
// 		return nil, liberrors.Errorf("getPrivileges. err: %w", err)
// 	}
// 	if !priv.HasPrivilege(model.PrivilegeRead) {
// 		return nil, liberrors.Errorf("AppUser(%d) has not privilege(read). err: %w", uint(operator.GetOrganizationID()), service.ErrDocumentPermissionDenied)
// 	}

// 	logger := log.FromContext(ctx)
// 	logger.Infof("ownerId: %d, operatorId: %d", documentEntity.OwnerID, operator.GetID())

// 	problemType, err := r.problemTypes.ToProblemType(documentEntity.ProblemTypeID)
// 	if err != nil {
// 		return nil, liberrors.Errorf("r.problemTypes.ToProblemType. err: %w", err)
// 	}

// 	documentModel, err := documentEntity.toDocumentModel(r.rf, r.pf, operator, problemType, priv)
// 	if err != nil {
// 		return nil, liberrors.Errorf("documentEntity.toDocumentModel. err: %w", err)
// 	}

// 	document, err := service.NewDocument(ctx, r.rf, r.pf, documentModel)
// 	if err != nil {
// 		return nil, liberrors.Errorf(". err: %w", err)
// 	}

// 	return document, nil
// }

// func (r *documentRepository) FindDocumentByName(ctx context.Context, operator userD.AppUserModel, spaceID userD.SpaceID, name string) (service.Document, error) {
// 	ctx, span := tracer.Start(ctx, "documentRepository.FindDocumentByName")
// 	defer span.End()

// 	documentEntity := documentEntity{}
// 	if result := r.db.
// 		Where("organization_id = ?", uint(operator.GetOrganizationID())).
// 		Where("space_id = ?", uint(spaceID)).
// 		Where("name = ?", name).
// 		First(&documentEntity); result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, service.ErrDocumentNotFound
// 		}
// 		return nil, result.Error
// 	}

// 	var priv userD.Privileges
// 	if spaceID == service.GetSystemSpaceID() {
// 		priv = userD.NewPrivileges([]userD.RBACAction{model.PrivilegeRead})
// 	} else {
// 		privTmp, err := r.getPrivileges(ctx, operator, model.DocumentID(documentEntity.ID))
// 		if err != nil {
// 			return nil, liberrors.Errorf("failed to checkPrivileges. err: %w", err)
// 		}
// 		if !privTmp.HasPrivilege(model.PrivilegeRead) {
// 			return nil, service.ErrDocumentPermissionDenied
// 		}
// 		priv = privTmp
// 	}

// 	logger := log.FromContext(ctx)
// 	logger.Infof("ownerId: %d, operatorId: %d", documentEntity.OwnerID, operator.GetID())

// 	problemType, err := r.problemTypes.ToProblemType(documentEntity.ProblemTypeID)
// 	if err != nil {
// 		return nil, liberrors.Errorf("r.problemTypes.ToProblemType. err: %w", err)
// 	}

// 	documentModel, err := documentEntity.toDocumentModel(r.rf, r.pf, operator, problemType, priv)
// 	if err != nil {
// 		return nil, liberrors.Errorf("documentEntity.toDocumentModel. err: %w", err)
// 	}

// 	document, err := service.NewDocument(ctx, r.rf, r.pf, documentModel)
// 	if err != nil {
// 		return nil, liberrors.Errorf(". err: %w", err)
// 	}

// 	return document, nil
// }

// func (r *documentRepository) getPrivileges(ctx context.Context, operator userD.AppUserModel, documentID model.DocumentID) (userD.Privileges, error) {
// 	userRf, err := r.rf.NewUserRepositoryFactory(ctx)
// 	if err != nil {
// 		return nil, liberrors.Errorf("r.rf.NewUserRepositoryFactory. err: %w", err)
// 	}

// 	rbacRepo := userRf.NewRBACRepository(ctx)

// 	documentRoles := r.getAllDocumentRoles(documentID)
// 	userObject := userD.NewUserObject(userD.AppUserID(operator.GetID()))
// 	e, err := rbacRepo.NewEnforcerWithRolesAndUsers(documentRoles, []userD.RBACUser{userObject})
// 	if err != nil {
// 		return nil, liberrors.Errorf("failed to NewEnforcerWithRolesAndUsers. err: %w", err)
// 	}
// 	documentObject := model.NewDocumentObject(documentID)
// 	privs := r.getAllDocumentPrivileges()
// 	return r.checkPrivileges(e, userObject, documentObject, privs)
// }

// func (r *documentRepository) AddDocument(ctx context.Context, operator userD.AppUserModel, spaceID userD.SpaceID, param model.DocumentAddParameter) (model.DocumentID, error) {
// 	_, span := tracer.Start(ctx, "documentRepository.AddDocument")
// 	defer span.End()

// 	problemTypeID, err := r.problemTypes.ToProblemTypeID(param.GetProblemType())
// 	if err != nil {
// 		return 0, liberrors.Errorf("unsupported problemType. problemType: %s", param.GetProblemType())
// 	}
// 	propertiesJSON, err := stringMapToJSON(param.GetProperties())
// 	if err != nil {
// 		return 0, liberrors.Errorf("stringMapToJSON. err: %w", err)
// 	}
// 	document := documentEntity{
// 		Version:        1,
// 		CreatedBy:      operator.GetID(),
// 		UpdatedBy:      operator.GetID(),
// 		OrganizationID: uint(operator.GetOrganizationID()),
// 		SpaceID:        uint(spaceID),
// 		OwnerID:        operator.GetID(),
// 		ProblemTypeID:  problemTypeID,
// 		Name:           param.GetName(),
// 		Lang2:          param.GetLang2().String(),
// 		QuestionText:   param.GetQuestionText(),
// 		Properties:     propertiesJSON,
// 	}
// 	if result := r.db.Create(&document); result.Error != nil {
// 		return 0, liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrDocumentAlreadyExists))
// 	}

// 	documentID := model.DocumentID(document.ID)

// 	userRf, err := r.rf.NewUserRepositoryFactory(ctx)
// 	if err != nil {
// 		return 0, liberrors.Errorf("r.rf.NewUserRepositoryFactory. err: %w", err)
// 	}

// 	rbacRepo := userRf.NewRBACRepository(ctx)
// 	userObject := userD.NewUserObject(userD.AppUserID(operator.GetID()))
// 	documentObject := model.NewDocumentObject(documentID)
// 	documentWriter := model.NewDocumentWriter(documentID)

// 	// the documentWriter role can read, update, remove
// 	if err := rbacRepo.AddNamedPolicy(documentWriter, documentObject, model.PrivilegeRead); err != nil {
// 		return 0, liberrors.Errorf("Failed to AddNamedPolicy. priv: read, err: %w", err)
// 	}
// 	if err := rbacRepo.AddNamedPolicy(documentWriter, documentObject, model.PrivilegeUpdate); err != nil {
// 		return 0, liberrors.Errorf("Failed to AddNamedPolicy. priv: update, err: %w", err)
// 	}
// 	if err := rbacRepo.AddNamedPolicy(documentWriter, documentObject, model.PrivilegeRemove); err != nil {
// 		return 0, liberrors.Errorf("Failed to AddNamedPolicy. priv: remove, err: %w", err)
// 	}

// 	// user is assigned the documentWriter role
// 	if err := rbacRepo.AddNamedGroupingPolicy(userObject, documentWriter); err != nil {
// 		return 0, liberrors.Errorf("Failed to AddNamedGroupingPolicy. err: %w", err)
// 	}

// 	// rbacRepo.NewEnforcerWithRolesAndUsers([]userD.RBACRole{documentWriter}, []userD.RBACUser{userObject})

// 	return documentID, nil
// }

// func (r *documentRepository) RemoveDocument(ctx context.Context, operator model.StudentModel, id model.DocumentID, version int) error {
// 	_, span := tracer.Start(ctx, "documentRepository.RemoveDocument")
// 	defer span.End()

// 	document := documentEntity{}
// 	if result := r.db.Where("organization_id = ? and id = ? and version = ?", operator.GetOrganizationID(), id, version).Delete(&document); result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return service.ErrDocumentNotFound
// 		}

// 		return result.Error
// 	}

// 	return nil
// }

// func (r *documentRepository) UpdateDocument(ctx context.Context, operator model.StudentModel, id model.DocumentID, version int, param model.DocumentUpdateParameter) error {
// 	_, span := tracer.Start(ctx, "documentRepository.UpdateDocument")
// 	defer span.End()

// 	if result := r.db.Model(&documentEntity{}).
// 		Where("organization_id = ?", uint(operator.GetOrganizationID())).
// 		Where("id = ?", uint(id)).
// 		Where("version = ?", version).
// 		Updates(map[string]interface{}{
// 			"name":          param.GetName(),
// 			"question_text": param.GetQuestionText(),
// 			"version":       gorm.Expr("version + 1"),
// 		}); result.Error != nil {
// 		return liberrors.Errorf(". err: %w", libG.ConvertDuplicatedError(result.Error, service.ErrDocumentAlreadyExists))
// 	}

// 	return nil
// }

// // func (r *documentRepository) ChangeSpace(ctx context.Context, operator model.AbstractStudent, id uint, spaceID uint) error {
// // 	result := r.db.Model(&documentEntity{}).Where(documentEntity{
// // 		OrganizationID: operator.OrganizationID(),
// // 		ID:             id,
// // 	}).Update(documentEntity{
// // 		SpaceID: spaceID,
// // 	})
// // 	if result.Error != nil {
// // 		return result.Error
// // 	}
// // 	if result.RowsAffected == 0 {
// // 		return model.NewDocumentNotFoundError(id)
// // 	}

// // 	return nil
// // }
