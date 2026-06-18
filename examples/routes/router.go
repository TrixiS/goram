package routes

import (
	"context"

	"github.com/TrixiS/goram"
	"github.com/TrixiS/goram/cbdata"
	"github.com/TrixiS/goram/examples/markups"
	"github.com/TrixiS/goram/filters"
	"github.com/TrixiS/goram/handlers"
)

func CreateRouter(adminUserID int64) *handlers.Router {
	checkAdmin := func(id int64) (bool, error) {
		return adminUserID == id, nil
	}

	rootRouter := handlers.NewRouter(handlers.RouterOptions{Name: "root"})
	adminGroup := rootRouter.Group(handlers.RouterOptions{Name: "admin"})

	if adminUserID != 0 {
		adminGroup.
			FilterMessage(func(ctx context.Context, bot *goram.Bot, message *goram.Message, data handlers.Data) (bool, error) {
				return checkAdmin(message.From.ID)
			}).
			FilterCallbackQuery(func(ctx context.Context, bot *goram.Bot, query *goram.CallbackQuery, data handlers.Data) (bool, error) {
				return checkAdmin(query.From.ID)
			})
	}

	adminGroup.
		Message(Start, filters.Text("/start")).
		CallbackQuery(
			HelloQuery,
			cbdata.FilterFunc(
				markups.Prefix,
				func(data markups.CbData) bool { return data.Type == markups.Hello },
			),
		).
		CallbackQuery(
			WorldQuery,
			cbdata.FilterFunc(
				markups.Prefix,
				func(data markups.CbData) bool { return data.Type == markups.World },
			),
		)

	return rootRouter
}
