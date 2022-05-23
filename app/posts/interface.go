package posts

import (
	"otoklix-assessment.com/entities"
	"otoklix-assessment.com/models"
)

type PostServicesInterface interface {
	CreatePost(*models.ReqPost) *entities.Post
	UpdatePost(int64, *models.ReqPost) *entities.Post
	DeletePost(int64) *entities.Post
	ShowPost(int64) *entities.Post
	GetPosts() *[]entities.Post
}

type PostRepositoriesInterface interface {
	CreatePost(*entities.Post) *entities.Post
	UpdatePost(int64, *entities.Post) *entities.Post
	DeletePost(int64) *entities.Post
	ShowPost(int64) *entities.Post
	GetPosts() *[]entities.Post
}
