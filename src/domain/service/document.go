//go:generate mockery --output mock --name Document
package service

import (
	"context"

	libdomain "github.com/pecolynx/golang-structure/lib/domain"
	liberrors "github.com/pecolynx/golang-structure/lib/errors"

	"github.com/pecolynx/golang-structure/src/domain/model"
)

type Document interface {
	model.DocumentModel
}

type document struct {
	model.DocumentModel
}

func NewDocument(ctx context.Context, documentModel model.DocumentModel) (Document, error) {
	m := &document{
		DocumentModel: documentModel,
	}

	if err := libdomain.Validator.Struct(m); err != nil {
		return nil, liberrors.Errorf("libD.Validator.Struct. err: %w", err)
	}

	return m, nil
}

func (m *document) GetDocumentModel() model.DocumentModel {
	return m.DocumentModel
}
