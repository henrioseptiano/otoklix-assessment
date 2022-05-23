package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"otoklix-assessment.com/app/posts/repositories"
	"otoklix-assessment.com/entities"
	"otoklix-assessment.com/models"
)

var postRepositories = &repositories.PostsRepoMock{Mock: mock.Mock{}}
var postServices = &PostServices{PostRepo: postRepositories}

func TestPostServices_CreatePost(t *testing.T) {
	post := entities.Post{
		Title:   "Ora Umum",
		Content: "Ora Umum ojo lali yo",
	}

	postRepositories.Mock.On("CreatePost", &post).Return(post)

	reqPost := models.ReqPost{
		Title:   "Ora Umum",
		Content: "Ora Umum ojo lali yo",
	}

	res := postServices.CreatePost(&reqPost)
	assert.NotNil(t, res)
}

func TestPostServices_UpdatePost(t *testing.T) {
	getPost := entities.Post{}
	postRepositories.Mock.On("ShowPost", int64(1)).Return(getPost)

	post := entities.Post{
		Title:   "Ora Umum2",
		Content: "Ora Umum ojo lali yo2",
	}

	postRepositories.Mock.On("UpdatePost", int64(1), &post).Return(post)

	reqPost := models.ReqPost{
		Title:   "Ora Umum2",
		Content: "Ora Umum ojo lali yo2",
	}

	res := postServices.UpdatePost(int64(1), &reqPost)
	assert.NotNil(t, res)
}

func TestPostServices_DeletePost(t *testing.T) {
	post := entities.Post{}
	postRepositories.Mock.On("DeletePost", int64(1)).Return(post)
	res := postServices.DeletePost(int64(1))
	assert.NotNil(t, res)
}

func TestPostServices_ShowPost(t *testing.T) {
	post := entities.Post{}
	postRepositories.Mock.On("ShowPost", int64(1)).Return(post)
	res := postServices.ShowPost(int64(1))
	assert.NotNil(t, res)
}

func TestPostServices_GetPosts(t *testing.T) {
	posts := make([]entities.Post, 0)
	postRepositories.Mock.On("GetPosts").Return(posts)
	res := postServices.GetPosts()
	assert.NotNil(t, res)
}
