//go:generate mockery --output mock --name DocumentRepository
package service

import (
	"context"
	"errors"

	"github.com/pecolynx/golang-structure/src/domain/model"
)

var ErrDocumentNotFound = errors.New("workbook not found")
var ErrDocumentAlreadyExists = errors.New("workbook already exists")
var ErrDocumentPermissionDenied = errors.New("permission denied")

type DocumentRepository interface {
	FindPrivateDocuments(ctx context.Context, operator model.DocumentWriterModel, param model.DocumentSearchCondition) (model.DocumentSearchResult, error)

	// FindDocumentByID(ctx context.Context, operator domain.StudentModel, id domain.DocumentID) (Document, error)

	// FindDocumentByName(ctx context.Context, operator userD.AppUserModel, spaceID userD.SpaceID, name string) (Document, error)

	// AddDocument(ctx context.Context, operator userD.AppUserModel, spaceID userD.SpaceID, param domain.DocumentAddParameter) (domain.DocumentID, error)

	// UpdateDocument(ctx context.Context, operator domain.StudentModel, workbookID domain.DocumentID, version int, param domain.DocumentUpdateParameter) error

	// RemoveDocument(ctx context.Context, operator domain.StudentModel, workbookID domain.DocumentID, version int) error
}
