package api

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	csrfMiddleware := csrf.Protect(
		[]byte(os.Getenv("HABIT_MANAGER_KEY")),
		csrf.Secure(false),
	)

	api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/csrftoken", api.handleGetCSRFToken)
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("loginuser", api.handleLoginUser)

				r.Group(func(r chi.Router) {
					r.Use(api.AuthMiddleware)
					r.Post("logoutuser", api.handleLogoutUser)
				})

			})

			r.Route("/categories", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(api.AuthMiddleware)
					r.Post("/", api.handleCreateCategory)
					r.Post("/getCategory", api.handleGetCategoryByName)
				})
			})

			r.Route("/habits", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(api.AuthMiddleware)
					r.Post("/", api.handleCreateHabit)
				})
			})
		})
	})
}
