package bootstrap

import (
	"github.com/thinhhuynh/golang-gin/api/controllers"
	"github.com/thinhhuynh/golang-gin/api/middlewares"
	"github.com/thinhhuynh/golang-gin/api/routes"
	"github.com/thinhhuynh/golang-gin/lib"
	"github.com/thinhhuynh/golang-gin/repository"
	"github.com/thinhhuynh/golang-gin/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
