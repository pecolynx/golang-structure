//go:generate mockery --output mock --name StudentModel
package model

import (
	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"
	usermodel "github.com/pecolynx/golang-structure/user/domain/model"
)

type DocumentWriterModel interface {
	usermodel.AppUserModel
	GetAppUserID() usermodel.AppUserID
}

type documentWriterModel struct {
	usermodel.AppUserModel
	// AppUserID usermodel.AppUserID
}

func NewDocumentWriterModel(appUserModel usermodel.AppUserModel) (DocumentWriterModel, error) {
	m := &documentWriterModel{
		AppUserModel: appUserModel,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libdomain.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *documentWriterModel) GetAppUserID() usermodel.AppUserID {
	return m.AppUserModel.GetAppUserID()
}
