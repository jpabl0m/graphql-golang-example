package usecase

// Usecase is application usecase
type Usecase struct{}

func New() *Usecase {
	return &Usecase{}
}

// InputPort is input port
type InputPort interface {
	CreatePost(*PostInputData) (*PostOutputData, error)
	AddComment(*CommentInputData) (*CommentOutputData, error)
	GetUser() (*UserOutputData, error)
	GetPosts() ([]*PostOutputData, error)
	GetUserPosts(userID string) ([]*PostOutputData, error)
	GetPostAuthor(postID string) (*UserOutputData, error)
	GetPostComments(postID string) ([]*CommentOutputData, error)
	GetCommentPost(commentID string) (*PostOutputData, error)
}

// PostInputData is input data of post
type PostInputData struct {
	Title   string
	Content string
}

// CommentInputData is input data of comment
type CommentInputData struct {
	PostID  string
	Content string
}

type UserOutputData struct {
	Name string
}

// PostOutputData is output data of post
type PostOutputData struct {
	Title   string
	Content string
}

// CommentOutputData is output data of comment
type CommentOutputData struct {
	Content string
}

// CreatePost create new post
func (u *Usecase) CreatePost(*PostInputData) (*PostOutputData, error) {
	panic("not implemented")
}

// AddComment add comment to post
func (u *Usecase) AddComment(*CommentInputData) (*CommentOutputData, error) {
	panic("not implemented")
}

func (u *Usecase) GetUser() (*UserOutputData, error) {
	panic("not implemented")
}

func (u *Usecase) GetPosts() ([]*PostOutputData, error) {
	panic("not implemented")
}

func (u *Usecase) GetUserPosts(userID string) ([]*PostOutputData, error) {
	panic("not implemented")
}

func (u *Usecase) GetPostAuthor(postID string) (*UserOutputData, error) {
	panic("not implemented")
}

func (u *Usecase) GetPostComments(postID string) ([]*CommentOutputData, error) {
	panic("not implemented")
}

func (u *Usecase) GetCommentPost(commentID string) (*PostOutputData, error) {
	panic("not implemented")

}
