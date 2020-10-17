package route_service

import "github.com/gofiber/fiber"

func  InitRoutes(app *fiber.App) {
	// 首先使用Group来包裹路由
	api := app.Group("/api")

	// 这里的userRoutes为同级目录中的处理用户的增删改查的函数
	userRoutes(api.Group("/users"))
}
