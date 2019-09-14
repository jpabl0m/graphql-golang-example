package usecase

// Usecase is application usecase
type usecase struct{}

func New() Usecase {
	return &usecase{}
}

type contextKey string

const UserIDContextKey contextKey = "userID"

// InputPort is input port
type Usecase interface {
	CreatePost(*PostInputData) (*PostOutputData, error)
	AddComment(*CommentInputData) (*CommentOutputData, error)
	GetUser() (*UserOutputData, error)
	GetPosts() ([]*PostOutputData, error)
	GetUserPosts(userID string) ([]*PostOutputData, error)
	GetPostAuthor(postID string) (*UserOutputData, error)
	GetPostComments(postID string) ([]*CommentOutputData, error)
	GetCommentPost(commentID string) (*PostOutputData, error)
	SignUp(*UserInputData) error
	SignIn(*UserInputData) error
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

type UserInputData struct {
	Email    string
	Password string
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
func (u *usecase) CreatePost(*PostInputData) (*PostOutputData, error) {
	panic("not implemented")
}

// AddComment add comment to post
func (u *usecase) AddComment(*CommentInputData) (*CommentOutputData, error) {
	panic("not implemented")
}

func (u *usecase) GetUser() (*UserOutputData, error) {
	panic("not implemented")
}

func (u *usecase) GetPosts() ([]*PostOutputData, error) {
	panic("not implemented")
}

func (u *usecase) GetUserPosts(userID string) ([]*PostOutputData, error) {
	panic("not implemented")
}

func (u *usecase) GetPostAuthor(postID string) (*UserOutputData, error) {
	panic("not implemented")
}

func (u *usecase) GetPostComments(postID string) ([]*CommentOutputData, error) {
	panic("not implemented")
}

func (u *usecase) GetCommentPost(commentID string) (*PostOutputData, error) {
	panic("not implemented")

}

func (u *usecase) SignUp(*UserInputData) error {
	panic("not implemented")
}

func (u *usecase) SignIn(*UserInputData) error {
	panic("not implemented")
}
