//go:generate mockery --output mock --name Model
package model

import (
	"time"

	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
)

type BaseModel interface {
	GetVersion() int
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetCreatedBy() uint
	GetUpdatedBy() uint
}

type baseModel struct {
	// ID        uint `validate:"gte=0"`
	Version   int `validate:"required,gte=1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uint `validate:"gte=0"`
	UpdatedBy uint `validate:"gte=0"`
}

func NewBaseModel(version int, createdAt, updatedAt time.Time, createdBy, updatedBy uint) (BaseModel, error) {
	m := &baseModel{
		// ID:        id,
		Version:   version,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libD.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *baseModel) GetVersion() int {
	return m.Version
}

func (m *baseModel) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *baseModel) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

func (m *baseModel) GetCreatedBy() uint {
	return m.CreatedBy
}

func (m *baseModel) GetUpdatedBy() uint {
	return m.UpdatedBy
}
