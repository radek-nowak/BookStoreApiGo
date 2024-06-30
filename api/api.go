package api

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/oapi-codegen/runtime/types"
)

type Api struct{}

// GetBooks implements StrictServerInterface.
func (a *Api) GetBooks(ctx context.Context, request GetBooksRequestObject) (GetBooksResponseObject, error) {
	author := ""
	date := types.Date{Time: time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)}
	return GetBooks200JSONResponse{
		{
			Author:        &author,
			Id:            1,
			PublishedDate: &date,
			Title:         new(string),
		},
	}, nil
}

func NewApp() *fiber.App {
	app := fiber.New()
	api := &Api{}

	server := NewStrictHandler(api, nil)
	RegisterHandlers(app, server)

	return app
}
