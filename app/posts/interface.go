package posts

type PostServicesInterface interface {
	CreatePost()
	UpdatePost()
	DeletePost()
	ShowPost()
	GetPosts()
}

type PostRepositoriesInterface interface {
	CreatePost()
	UpdatePost()
	DeletePost()
	ShowPost()
	GetPosts()
}
