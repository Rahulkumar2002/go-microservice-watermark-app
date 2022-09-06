package watermark

import (
	"context"

	"github.com/Rahulkumar2002/go-microservice-watermark-app/internal"
)

type Service interface {
	// To get all the documents.
	Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
	// It will return the status of the document for the passed ticet ID.
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	// It will watermark the document by the given mark.
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	// It will add the document and return the title ID.
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	// It will return the status of the service.
	ServiceStatus(ctx context.Context) (int, error)
}
