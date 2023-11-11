//go:generate mockery --output mock --name StudentUsecaseDocument
package usecase

import (
	"github.com/pecolynx/golang-structure/src/domain/service"
)

const DefaultPageNo = 1
const DefaultPageSize = 10

type UserUsecaseDocument interface {
	// FindDocuments(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID) (domain.DocumentSearchResult, error)

	// FindDocumentByID(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, workBookID domain.DocumentID) (domain.DocumentModel, error)

	// AddDocument(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, parameter domain.DocumentAddParameter) (domain.DocumentID, error)

	// UpdateDocument(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, documentID domain.DocumentID, version int, parameter domain.DocumentUpdateParameter) error

	// RemoveDocument(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, documentID domain.DocumentID, version int) error
}

type studentUsecaseDocument struct {
	transactionManager service.TransactionManager
	// pf          service.ProcessorFactory
	// findStudentFunc FindStudentFunc
}

func NewUserUsecaseDocument(transactionManager service.TransactionManager,

// pf service.ProcessorFactory
) UserUsecaseDocument {
	return &studentUsecaseDocument{
		transactionManager: transactionManager,
		// pf:          pf,
		// findStudentFunc: findStudentFunc,
	}
}

// func (s *studentUsecaseDocument) FindDocuments(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID) (domain.DocumentSearchResult, error) {
// 	var result domain.DocumentSearchResult
// 	fn := func(student service.Student) error {
// 		condition, err := domain.NewDocumentSearchCondition(DefaultPageNo, DefaultPageSize, []userD.SpaceID{})
// 		if err != nil {
// 			return liberrors.Errorf("service.NewDocumentSearchCondition. err: %w", err)
// 		}

// 		tmpResult, err := student.FindDocumentsFromPersonalSpace(ctx, condition)
// 		if err != nil {
// 			return liberrors.Errorf("student.FindDocumentsFromPersonalSpace. err: %w", err)
// 		}

// 		result = tmpResult
// 		return nil
// 	}

// 	if err := s.studentHandle(ctx, organizationID, operatorID, fn); err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func (s *studentUsecaseDocument) FindDocumentByID(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, workBookID domain.DocumentID) (domain.DocumentModel, error) {
// 	var result domain.DocumentModel
// 	fn := func(student service.Student) error {
// 		tmpResult, err := student.FindDocumentByID(ctx, workBookID)
// 		if err != nil {
// 			return liberrors.Errorf("student.FindDocumentByID. err: %w", err)
// 		}
// 		result = tmpResult
// 		return nil
// 	}

// 	if err := s.studentHandle(ctx, organizationID, operatorID, fn); err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func (s *studentUsecaseDocument) AddDocument(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, parameter domain.DocumentAddParameter) (domain.DocumentID, error) {
// 	var addedDocumentID domain.DocumentID
// 	fn := func(student service.Student) error {
// 		tmpAddedDocumentID, err := student.AddDocumentToPersonalSpace(ctx, parameter)
// 		if err != nil {
// 			return liberrors.Errorf("student.AddDocumentToPersonalSpace. err: %w", err)
// 		}
// 		addedDocumentID = tmpAddedDocumentID
// 		return nil
// 	}

// 	if err := s.studentHandle(ctx, organizationID, operatorID, fn); err != nil {
// 		return 0, err
// 	}

// 	return addedDocumentID, nil
// }

// func (s *studentUsecaseDocument) UpdateDocument(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, documentID domain.DocumentID, version int, parameter domain.DocumentUpdateParameter) error {
// 	fn := func(student service.Student) error {
// 		if err := student.UpdateDocument(ctx, documentID, version, parameter); err != nil {
// 			return liberrors.Errorf("student.UpdateDocument. err: %w", err)
// 		}
// 		return nil
// 	}
// 	return s.studentHandle(ctx, organizationID, operatorID, fn)
// }

// func (s *studentUsecaseDocument) RemoveDocument(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, documentID domain.DocumentID, version int) error {
// 	fn := func(student service.Student) error {
// 		if err := student.RemoveDocument(ctx, documentID, version); err != nil {
// 			return liberrors.Errorf("student.RemoveDocument. err: %w", err)
// 		}
// 		return nil
// 	}
// 	return s.studentHandle(ctx, organizationID, operatorID, fn)
// }

// func (s *studentUsecaseDocument) studentHandle(ctx context.Context, organizationID userD.OrganizationID, operatorID userD.AppUserID, fn func(service.Student) error) error {
// 	if err := s.transaction.Do(ctx, func(rf service.RepositoryFactory) error {
// 		student, err := s.findStudentFunc(ctx, rf, organizationID, operatorID)
// 		if err != nil {
// 			return liberrors.Errorf("usecase.FindStudent. err: %w", err)
// 		}
// 		return fn(student)
// 	}); err != nil {
// 		return liberrors.Errorf("studentHandle. err: %w", err)
// 	}
// 	return nil
// }
