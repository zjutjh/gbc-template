package register

import (
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/zjutjh/jhgo/config"
	"github.com/zjutjh/jhgo/swagger"

	"app/api"
	"app/middleware"
)

func Route(router *gin.Engine) {
	router.Use(middleware.Cors())

	r := router.Group(routePrefix())
	{
		routeBase(r, router)

		// 注册业务逻辑接口

	}
}

func routePrefix() string {
	return "/api"
}

func routeBase(r *gin.RouterGroup, router *gin.Engine) {
	// OpenAPI/Swagger 文档生成
	if slices.Contains([]string{config.AppEnvDev, config.AppEnvTest}, config.AppEnv()) {
		r.GET("/swagger.json", swagger.DocumentHandler(router))
	}

	// 健康检查
	r.GET("/health", api.HealthHandler())
}
