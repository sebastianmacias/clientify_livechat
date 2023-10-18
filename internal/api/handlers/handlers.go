// Code generated by go run -tags scripts scripts/handlers/gen_handlers.go; DO NOT EDIT.
package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sebastianmacias/clientify_livechat/internal/api"
	"github.com/sebastianmacias/clientify_livechat/internal/api/handlers/auth"
	"github.com/sebastianmacias/clientify_livechat/internal/api/handlers/common"
	"github.com/sebastianmacias/clientify_livechat/internal/api/handlers/push"
)

func AttachAllRoutes(s *api.Server) {
	// attach our routes
	s.Router.Routes = []*echo.Route{
		auth.GetUserInfoRoute(s),
		auth.PostChangePasswordRoute(s),
		auth.PostForgotPasswordCompleteRoute(s),
		auth.PostForgotPasswordRoute(s),
		auth.PostLoginRoute(s),
		auth.PostLogoutRoute(s),
		auth.PostRefreshRoute(s),
		auth.PostRegisterRoute(s),
		common.GetHealthyRoute(s),
		common.GetReadyRoute(s),
		common.GetSwaggerRoute(s),
		common.GetVersionRoute(s),
		push.GetPushTestRoute(s),
		push.PostUpdatePushTokenRoute(s),
	}
}
