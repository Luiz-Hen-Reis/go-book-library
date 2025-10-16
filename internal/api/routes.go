package api

import "github.com/go-chi/chi/v5"

func (api *Api) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func (r chi.Router)  {
			r.Get("/authors", api.handleListAuthors)
			r.Get("/authors/{id}", api.handleGetAuthorByID)
			r.Post("/authors", api.handleCreateAuthor)
			r.Delete("/authors/{id}", api.handleDeleteAuthorByID)
		})
	})
}