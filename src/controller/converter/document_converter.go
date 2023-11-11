package converter

import (
	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
	"github.com/pecolynx/golang-structure/src/controller/entity"
	"github.com/pecolynx/golang-structure/src/domain/model"
)

func ToDocumentHTTPEntity(document model.DocumentModel) (entity.DocumentResponseHTTPEntity, error) {
	e := entity.DocumentResponseHTTPEntity{
		BaseModel: entity.BaseModel{
			Version:   document.GetVersion(),
			CreatedBy: document.GetCreatedBy(),
			UpdatedBy: document.GetUpdatedBy(),
		},
		ID:      document.GetDocumentID().String(),
		Name:    document.GetName(),
		Content: document.GetContent(),
	}

	if err := libdomain.Validator.Struct(e); err != nil {
		return entity.DocumentResponseHTTPEntity{}, liberrors.Errorf("libD.Validator.Struct. err: %w", err)
	}

	return e, nil
}

// func ToDocumentAddParameter(param *entity.DocumentAddParameter) (domain.DocumentAddParameter, error) {
// 	domainParam, err := domain.NewDocumentAddParameter(domain.ProblemTypeName(param.ProblemType), param.Name, domain.Lang2JA, param.QuestionText, map[string]string{
// 		"audioEnabled": "false",
// 	})

// 	if err != nil {
// 		return nil, liberrors.Errorf(". err: %w", err)
// 	}

// 	return domainParam, nil
// }

// func ToDocumentUpdateParameter(param *entity.DocumentUpdateParameter) (domain.DocumentUpdateParameter, error) {
// 	domainParam, err := domain.NewDocumentUpdateParameter(param.Name, param.QuestionText)
// 	if err != nil {
// 		return nil, liberrors.Errorf(". err: %w", err)
// 	}

// 	return domainParam, nil
// }
