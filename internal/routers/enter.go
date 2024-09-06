package routers

import (
	"github.com/nhh57/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/nhh57/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manager.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
