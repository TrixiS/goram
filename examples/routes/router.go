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
	checkAdmin := func(id int64) bool {
		return adminUserID == 0 || adminUserID == id
	}

	router := handlers.
		NewRouter(handlers.RouterOptions{Name: "root"}).
		FilterMessage(
			func(ctx context.Context, bot *goram.Bot, message *goram.Message, data handlers.Data) (bool, error) {
				return checkAdmin(message.From.ID), nil
			},
		).
		FilterCallbackQuery(
			func(ctx context.Context, bot *goram.Bot, query *goram.CallbackQuery, data handlers.Data) (bool, error) {
				return checkAdmin(query.From.ID), nil
			},
		).
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

	return router
}
