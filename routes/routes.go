package routes

import (
	"github.com/gofiber/fiber/v2"
	Post "otoklik-assessment.com/app/posts/controller"
	PostService "otoklik-assessment.com/app/posts/services"
	PostRepo "otoklix-assessment.com/app/posts/repositories"
)

func PostRoutes(app *fiber.App) {
	postRepo := PostRepo.New()
	postService := PostService.New(postRepo)

	post := &Post.PostController{postService}

	app.Get("/posts", post.GetPosts)
	app.Get("/posts/1", post.ShowPost)
	app.Post("/posts", post.CreatePost)
	app.Put("/posts/:id", post.UpdatePost)
	app.Delete("/posts/:id", post.DeletePost)
}
