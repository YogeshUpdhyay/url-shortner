package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/YogeshUpdhyay/url-shortener/internal/constants"
	"github.com/YogeshUpdhyay/url-shortener/internal/controllers"
	"github.com/YogeshUpdhyay/url-shortener/internal/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	// api routes
	apiRouter := router.Group(constants.ApiRoute)
	{
		// v1 routes
		v1Router := apiRouter.Group(constants.V1Route)
		{
			v1Router.POST(constants.ShortenUrlRoute, middlewares.AppAuthenticateMiddleware(), controllers.ShortenUrl)
			v1Router.POST(constants.DeleteUrlRoute, controllers.DeleteUrl)
			v1Router.POST(constants.AuthenticateRoute, controllers.Authenticate)
			v1Router.POST(constants.CreateAppRoute, middlewares.UserAuthenticateMiddleware(), controllers.CreateApp)
			v1Router.GET(constants.ListAppsRoute, middlewares.UserAuthenticateMiddleware(), controllers.HandleListApps)
		}
	}

	// ui routes
	router.GET(constants.RedirectUrlRoute, controllers.RedirectUrl)
	router.GET(constants.RedirectToAppRoute, controllers.RedirectToApp)
}
