package controller

import "otoklix-assessment.com/app/posts"

type PostController struct {
	PostServices posts.PostServicesInterface
}

func (pr *PostController) CreatePost() {}

func (pr *PostController) UpdatePost() {}
func (pr *PostController) DeletePost() {}
func (pr *PostController) ShowPost()   {}
func (pr *PostController) GetPosts()   {}
