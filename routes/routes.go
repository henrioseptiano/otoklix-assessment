package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	Post "otoklix-assessment.com/app/posts/controller"
	PostRepo "otoklix-assessment.com/app/posts/repositories"
	PostService "otoklix-assessment.com/app/posts/services"
)

func PostRoutes(app *fiber.App, db *gorm.DB) {
	postRepo := PostRepo.New(db)
	postService := PostService.New(postRepo)

	post := &Post.PostController{postService}

	app.Get("/posts", post.GetPosts)
	app.Get("/posts/:id", post.ShowPost)
	app.Post("/posts", post.CreatePost)
	app.Put("/posts/:id", post.UpdatePost)
	app.Delete("/posts/:id", post.DeletePost)
}
