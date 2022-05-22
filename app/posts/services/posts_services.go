package services

import "otoklix-assessment.com/app/posts"

type PostServices struct {
	PostRepo posts.PostRepositoriesInterface
}

func New(postRepo posts.PostRepositoriesInterface) PostServices {
	return PostServices{postRepo}
}

func (pr *PostServices) CreatePost() {}

func (pr *PostServices) UpdatePost() {}
func (pr *PostServices) DeletePost() {}
func (pr *PostServices) ShowPost()   {}
func (pr *PostServices) GetPosts()   {}
