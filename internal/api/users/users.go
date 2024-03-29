package users

import (
	"github.com/go-chi/chi"
	"github.com/sanpezlo/chess/internal/resources/user"
	"github.com/sanpezlo/chess/internal/services/auth"
	"go.uber.org/fx"
)

type controller struct {
	repository user.Repository
}

func New(repository user.Repository) *controller {
	return &controller{repository}
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(func(r chi.Router, c *controller) {
		rtr := chi.NewRouter()
		r.Mount("/users", rtr)

		rtr.With(auth.MustBeAuthenticated).Get("/self", c.self)

		rtr.Get("/{id}", c.get)

		rtr.Get("/", c.list)

		rtr.Patch("/{id}", c.patch)

	}),
)
