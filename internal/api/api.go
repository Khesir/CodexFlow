package api

import "context"

type API struct {
	ctx context.Context
}

func NewApi() *API {
	return &API{}
}

func (a *API) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *API) Ping() string {
	return "pong"
}

func (a *API) GetTasks() []string {
	return []string{"S"}
}
