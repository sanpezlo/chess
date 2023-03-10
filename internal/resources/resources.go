package resources

import (
	"github.com/sanpezlo/chess/internal/resources/auth"
	"github.com/sanpezlo/chess/internal/resources/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	user.Module,
	auth.Module,
)
