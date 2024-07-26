package init

import (
	"github.com/CoolCoolTomato/MatrilxArena/Server/middleware"
	"github.com/CoolCoolTomato/MatrilxArena/Server/route"
	"github.com/gin-gonic/gin"
)

func RouteInit() {
	route.Route = gin.Default()
	route.Route.Use(middleware.CORSMiddleware())
	route.Route.Use(middleware.I18n())
	route.SetAuthRoute(route.Route)
	route.SetUserRoute(route.Route)
	route.SetImageRoute(route.Route)
	route.SetCategoryRoute(route.Route)
	route.SetChallengeRoute(route.Route)
	route.SetGroupRoute(route.Route)
	route.SetGroupChallengeRoute(route.Route)
	route.SetGroupUserRoute(route.Route)
	route.SetDockerNodeRoute(route.Route)
	route.SetUserContainerRoute(route.Route)
	route.SetUserChallengeRoute(route.Route)
	route.SetUserGroupRoute(route.Route)
    route.SetUserGroupChallengeRoute(route.Route)
    route.SetUserGroupContainerRoute(route.Route)
	route.SetAttachmentRoute(route.Route)
}
