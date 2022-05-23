package repositories

import (
	"github.com/stretchr/testify/mock"
	"otoklix-assessment.com/entities"
)

type PostsRepoMock struct {
	Mock mock.Mock
}

func (pr PostsRepoMock) CreatePost(data *entities.Post) *entities.Post {
	arguments := pr.Mock.Called(data)
	if arguments.Get(0) == nil {
		return nil
	}
	postData := arguments.Get(0).(entities.Post)
	return &postData
}

func (pr PostsRepoMock) ShowPost(id int64) *entities.Post {
	arguments := pr.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}
	postData := arguments.Get(0).(entities.Post)
	return &postData
}

func (pr PostsRepoMock) UpdatePost(id int64, data *entities.Post) *entities.Post {
	arguments := pr.Mock.Called(id, data)
	if arguments.Get(0) == nil {
		return nil
	}
	postData := arguments.Get(0).(entities.Post)
	return &postData
}

func (pr PostsRepoMock) DeletePost(id int64) *entities.Post {
	arguments := pr.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}
	postData := arguments.Get(0).(entities.Post)
	return &postData
}

func (pr PostsRepoMock) GetPosts() *[]entities.Post {
	arguments := pr.Mock.Called()
	if arguments.Get(0) == nil {
		return nil
	}
	postData := arguments.Get(0).([]entities.Post)
	return &postData
}
