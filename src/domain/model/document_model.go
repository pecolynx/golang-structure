//go:generate mockery --output mock --name DocumentModel
package model

import (
	"fmt"
	"net/http"

	"github.com/oklog/ulid/v2"
	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
	usermodel "github.com/pecolynx/golang-structure/user/domain/model"
)

type DocumentID interface {
	String() string
}

type documentID struct {
	value ulid.ULID
}

func NewDocumentID(value string) (DocumentID, error) {
	ulidValue, err := ulid.Parse(value)
	if err != nil {
		return nil, err
	}

	return &documentID{
		value: ulidValue,
	}, nil
}

func (v *documentID) String() string {
	return v.value.String()
}

type DocumentModel interface {
	usermodel.BaseModel
	GetDocumentID() DocumentID
	// GetSpaceID() userD.SpaceID
	// GetOwnerID() userD.AppUserID
	GetName() string
	GetContent() string
	// GetLang2() Lang2
	// GetProblemType() ProblemTypeName
	// GetQuestionText() string
	// GetProperties() map[string]string
	// HasPrivilege(privilege userD.RBACAction) bool
}

type documentModel struct {
	usermodel.BaseModel
	// spaceID      userD.SpaceID    `validate:"required"`
	// ownerID      userD.AppUserID  `validate:"required"`
	// privileges   userD.Privileges `validate:"required"`
	DocumentID DocumentID
	Name       string `validate:"required"`
	Content    string `validate:"required"`
}

func NewDocumentModel(model usermodel.BaseModel,
	// spaceID userD.SpaceID, ownerID userD.AppUserID, privileges userD.Privileges,
	documentID DocumentID,
	name string, content string,
	//questsionText string, properties map[string]string
) (DocumentModel, error) {
	value := http.StatusAccepted
	fmt.Println(value)
	m := &documentModel{
		BaseModel: model,
		// spaceID:      spaceID,
		// ownerID:      ownerID,
		// privileges:   privileges,
		DocumentID: documentID,
		Name:       name,
		Content:    content,
		// QuestionText: questsionText,
		// Properties:   properties,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *documentModel) GetDocumentID() DocumentID {
	return m.DocumentID
}

// func (m *documentModel) GetSpaceID() userD.SpaceID {
// 	return m.spaceID
// }

// func (m *documentModel) GetOwnerID() userD.AppUserID {
// 	return m.ownerID
// }

func (m *documentModel) GetName() string {
	return m.Name
}

func (m *documentModel) GetContent() string {
	return m.Content
}

// func (m *documentModel) GetLang2() Lang2 {
// 	return m.Lang2
// }

// func (m *documentModel) GetProblemType() ProblemTypeName {
// 	return m.ProblemType
// }

// func (m *documentModel) GetQuestionText() string {
// 	return m.QuestionText
// }

// func (m *documentModel) GetProperties() map[string]string {
// 	return m.Properties
// }

// func (m *documentModel) HasPrivilege(privilege userD.RBACAction) bool {
// 	return m.privileges.HasPrivilege(privilege)
// }
