package entity

import (
	"time"

	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
	usermodel "github.com/pecolynx/golang-structure/user/domain/model"
)

type BaseModel struct {
	// ID        uint      `json:"id" validate:"gte=0"`
	Version   int       `json:"version" validate:"required,gte=1"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy uint      `json:"createdBy" validate:"gte=0"`
	UpdatedBy uint      `json:"updatedBy" validate:"gte=0"`
}

func NewBaseModel(model usermodel.BaseModel) (BaseModel, error) {
	m := BaseModel{
		Version:   model.GetVersion(),
		CreatedAt: model.GetCreatedAt(),
		UpdatedAt: model.GetUpdatedAt(),
		CreatedBy: model.GetCreatedBy(),
		UpdatedBy: model.GetUpdatedBy(),
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return BaseModel{}, liberrors.Errorf("libD.Validator.Struct. err: %w", err)
	}

	return m, nil
}

type DocumentResponseHTTPEntity struct {
	BaseModel
	ID      string `json:"id"`
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type DocumentAddParameter struct {
	Name string `json:"name" binding:"required"`
}

type DocumentUpdateParameter struct {
	Name string `json:"name" binding:"required"`
}
