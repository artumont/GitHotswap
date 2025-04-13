package githotswaputils

import (
	"github.com/artumont/GitHotswap/internal/handlers"
	"github.com/artumont/GitHotswap/internal/router"
)

func RegistryInit(r *router.Router) {
	r.RegisterHandler(
		"profile",
		handlers.NewProfileHandler(r.Cfg),
	)
}
