//go:generate mockery --output mock --name AppUserModel
package model

import (
	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
)

type AppUserID uint

type AppUserModel interface {
	BaseModel
	GetAppUserID() AppUserID
	GetLoginID() string
	GetUsername() string
}

type appUserModel struct {
	BaseModel
	AppUserID AppUserID `validate:"gte=0"`
	LoginID   string    `validate:"required"`
	Username  string    `validate:"required"`
}

func NewAppUserModel(model BaseModel, appUserID AppUserID, loginID, username string) (AppUserModel, error) {
	m := &appUserModel{
		BaseModel: model,
		AppUserID: appUserID,
		LoginID:   loginID,
		Username:  username,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *appUserModel) GetAppUserID() AppUserID {
	return m.AppUserID
}

func (m *appUserModel) GetLoginID() string {
	return m.LoginID
}

func (m *appUserModel) GetUsername() string {
	return m.Username
}
