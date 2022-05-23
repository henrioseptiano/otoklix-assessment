package repositories

import (
	"log"

	"gorm.io/gorm"
	"otoklix-assessment.com/entities"
)

type PostRepositories struct {
	DB *gorm.DB
}

func New(db *gorm.DB) PostRepositories {
	return PostRepositories{db}
}

func (pr PostRepositories) CreatePost(data *entities.Post) *entities.Post {
	if err := pr.DB.Create(data).Error; err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil
	}
	return data
}

func (pr PostRepositories) UpdatePost(id int64, post *entities.Post) *entities.Post {
	res := pr.DB.Model(&entities.Post{}).Where("id = ?", id).Updates(&post)
	if err := res.Error; err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil
	}
	return post
}
func (pr PostRepositories) DeletePost(id int64) *entities.Post {
	getData := pr.ShowPost(id)
	if err := pr.DB.Unscoped().Where("id = ?", id).Delete(&entities.Post{}).Error; err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil
	}
	return getData
}
func (pr PostRepositories) ShowPost(id int64) *entities.Post {
	data := new(entities.Post)
	pr.DB.Model(&entities.Post{}).Where("id = ?", id).Find(&data)
	return data
}
func (pr PostRepositories) GetPosts() *[]entities.Post {
	data := make([]entities.Post, 0)
	pr.DB.Model(&entities.Post{}).Find(&data)
	return &data
}
