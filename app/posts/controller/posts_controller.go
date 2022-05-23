package controller

import (
	"github.com/gofiber/fiber/v2"
	"otoklix-assessment.com/app/posts"
	"otoklix-assessment.com/models"
)

type PostController struct {
	PostServices posts.PostServicesInterface
}

func (pr PostController) CreatePost(c *fiber.Ctx) error {
	req := models.ReqPost{}

	if err := c.BodyParser(&req); err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}

	res := pr.PostServices.CreatePost(&req)
	if res == nil {
		return models.ResponseError(c, "Cannot Create Post", 401)
	}

	return models.ResponseJSON(c, res)
}

func (pr PostController) UpdatePost(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return models.ResponseError(c, "Invalid ID", 401)
	}

	req := models.ReqPost{}
	if err := c.BodyParser(&req); err != nil {
		return models.ResponseError(c, err.Error(), 401)
	}

	res := pr.PostServices.UpdatePost(int64(postId), &req)
	if res == nil {
		return models.ResponseError(c, "Cannot Create Post", 401)
	}
	return models.ResponseJSON(c, res)
}
func (pr PostController) DeletePost(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return models.ResponseError(c, "Invalid ID", 401)
	}

	res := pr.PostServices.DeletePost(int64(postId))
	if res == nil {
		return models.ResponseError(c, "Cannot Delete Post", 401)
	}

	return models.ResponseJSON(c, res)
}
func (pr PostController) ShowPost(c *fiber.Ctx) error {
	postId, err := c.ParamsInt("id")
	if err != nil {
		return models.ResponseError(c, "Invalid ID", 401)
	}
	return models.ResponseJSON(c, pr.PostServices.ShowPost(int64(postId)))
}
func (pr PostController) GetPosts(c *fiber.Ctx) error {
	data := pr.PostServices.GetPosts()
	return models.ResponseJSON(c, data)
}
