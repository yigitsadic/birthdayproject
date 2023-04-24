package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/yigitsadic/birthday-app-api/internal/auth"
)

// MountRoutes mounts all routes to the given router.
func (s *Server) MountRoutes(router *chi.Mux) {
	router.Route("/sessions", func(r chi.Router) {
		r.Post("/create", s.HandleCreateSession)
		r.Post("/refresh", s.HandleRefreshSession)
	})

	protectedGroup := router.Group(nil)
	protectedGroup.Use(auth.InjectUser(s.JWTStore))
	protectedGroup.Use(auth.AuthGuard)

	protectedGroup.Route("/users/{id}", func(r chi.Router) {
		r.Get("/", s.HandleUserDetail)
		r.Put("/", s.HandleUserUpdate)
	})

	protectedGroup.Route("/companies/{id}", func(r chi.Router) {
		r.Get("/", s.HandleCompanyDetail)
		r.Put("/", s.HandleCompanyUpdate)
	})

	protectedGroup.Route("/companies/{companyId}/employees", func(r chi.Router) {
		r.Get("/", s.EmployeeListHandler)
		r.Post("/", s.EmployeeCreateHandler)
		//
		r.Route("/{id}", func(rr chi.Router) {
			rr.Get("/", s.EmployeeDetailHandler)
			rr.Put("/", s.EmployeeUpdateHandler)
			rr.Delete("/", s.EmployeeDeleteHandler)
		})
	})
}
