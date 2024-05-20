package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"src/ex02/db"
	"src/ex02/handlers"
	"src/ex02/middleware"
)

func main() {
	db.Init()
	defer db.DB.Close()
	r := gin.Default()
	r.Use(middleware.RateLimit)

	r.SetFuncMap(template.FuncMap{
		"safeHTML": handlers.SafeHTML,
	})

	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", handlers.GetPosts)
	r.GET("/post/:id", handlers.GetPost)
	r.GET("/admin", handlers.AdminPanel)
	r.POST("/admin", handlers.AdminLogin)
	r.POST("/admin/post", handlers.CreatePost)
	r.GET("/admin/edit/:id", handlers.EditPost)
	r.POST("/admin/edit/:id", handlers.UpdatePost)
	r.POST("/admin/delete/:id", handlers.DeletePost)

	r.Static("/web/css/", "./web/css/")
	r.Static("/web/images", "./web/images")

	r.Run(":8888")
}
