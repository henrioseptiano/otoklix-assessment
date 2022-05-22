package repositories

type PostRepositories struct{}

func New() PostRepositories {
	return PostRepositories{}
}

func (pr *PostRepositories) CreatePost() {}

func (pr *PostRepositories) UpdatePost() {}
func (pr *PostRepositories) DeletePost() {}
func (pr *PostRepositories) ShowPost()   {}
func (pr *PostRepositories) GetPosts()   {}
