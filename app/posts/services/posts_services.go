package services

import (
	"otoklix-assessment.com/app/posts"
	"otoklix-assessment.com/entities"
	"otoklix-assessment.com/models"
)

type PostServices struct {
	PostRepo posts.PostRepositoriesInterface
}

func New(postRepo posts.PostRepositoriesInterface) PostServices {
	return PostServices{postRepo}
}

func (pr PostServices) CreatePost(createPost *models.ReqPost) *entities.Post {
	post := new(entities.Post)
	post.Title = createPost.Title
	post.Content = createPost.Content
	return pr.PostRepo.CreatePost(post)
}

func (pr PostServices) UpdatePost(id int64, updatePost *models.ReqPost) *entities.Post {
	post := pr.PostRepo.ShowPost(id)
	post.Title = updatePost.Title
	post.Content = updatePost.Content
	return pr.PostRepo.UpdatePost(id, post)
}
func (pr PostServices) DeletePost(id int64) *entities.Post {
	return pr.PostRepo.DeletePost(id)
}
func (pr PostServices) ShowPost(id int64) *entities.Post {
	return pr.PostRepo.ShowPost(id)
}
func (pr PostServices) GetPosts() *[]entities.Post {
	return pr.PostRepo.GetPosts()
}
