//go:generate mockery --output mock --name RepositoryFactory
package service

import (
	"context"
)

type RepositoryFactory interface {
	NewDocumentRepository(ctx context.Context) DocumentRepository
}
