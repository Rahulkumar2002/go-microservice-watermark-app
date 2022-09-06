package watermark

import (
	"context"
	"net/http"
	"os"

	"github.com/Rahulkumar2002/go-microservice-watermark-app/internal"
	"github.com/go-kit/kit/log"
	"github.com/lithammer/shortuuid/v3"
)

type watermarkService struct{}

func NewService() Service {
	return &watermarkService{}
}

func (w *watermarkService) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	// Query the database using the filters and return the list of documents
	// return error if the filter(key) is invalid and also return error if no item found

	doc := internal.Document{
		Content: "book",
		Title:   "Harry Potter and Half Blood Prince",
		Author:  "J.K. Rowlings",
		Topic:   "Fiction and Magic",
	}

	return []internal.Document{doc}, nil
}

func (w *watermarkService) Status(ctx context.Context, titledID string) (internal.Status, error) {
	// Query database using the ticketID and return the document info
	// return error if the tickedID is invalid or no document exists for the tickedID

	return internal.InProgress, nil

}

func (w *watermarkService) Watermark(ctx context.Context, titleID, mark string) (int, error) {
	// update the database entry with watermark field as non empty
	// first check if the watermark status is not already in InProgress , Started or Finished state
	// If yes , then return invalid request
	// return error if no item found using the ticketID

	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	// add the document entry in the database by calling the database service
	// return error if the doc is invalid and/or the database invalid entry error

	newTickedId := shortuuid.New()
	return newTickedId, nil
}

func (w *watermarkService) ServiceStatus(ctx context.Context) (int, error) {
	logger.Log("Checking the Service health ... ")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts ", log.DefaultTimestampUTC)
}
