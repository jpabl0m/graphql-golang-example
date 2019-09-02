// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type AddCommentInput struct {
	PostID  string `json:"postId"`
	Content string `json:"content"`
}

type AddCommentPayload struct {
	Comment *Comment `json:"comment"`
}

type AddPostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AddPostPayload struct {
	Post *Post `json:"post"`
}
