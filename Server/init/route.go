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
    route.SetAttachmentRoute(route.Route)
	route.SetAuthRoute(route.Route)
	route.SetCategoryRoute(route.Route)
	route.SetChallengeRoute(route.Route)
	route.SetCTFRoute(route.Route)
	route.SetCTFChallengeRoute(route.Route)
	route.SetCTFTeamRoute(route.Route)
    route.SetCTFTeamUserRoute(route.Route)
    route.SetCTFUserRoute(route.Route)
    route.SetDockerNodeRoute(route.Route)
	route.SetGroupRoute(route.Route)
	route.SetGroupChallengeRoute(route.Route)
	route.SetGroupUserRoute(route.Route)
    route.SetImageRoute(route.Route)
    route.SetUserRoute(route.Route)
    route.SetUserChallengeRoute(route.Route)
	route.SetUserContainerRoute(route.Route)
    route.SetUserCTFRoute(route.Route)
	route.SetUserGroupRoute(route.Route)
	route.SetUserGroupChallengeRoute(route.Route)
	route.SetUserGroupContainerRoute(route.Route)
}
