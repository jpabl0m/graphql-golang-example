package graphql

import (
	"context"

	"github.com/stobita/graphql-golang-example/internal/usecase"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	usecase usecase.InputPort
}

func NewResolver(u usecase.InputPort) *Resolver {
	return &Resolver{
		usecase: u,
	}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddPost(ctx context.Context, input AddPostInput) (*AddPostPayload, error) {
	i := &usecase.PostInputData{
		Title:   input.Title,
		Content: input.Content,
	}
	result, err := r.usecase.CreatePost(i)
	if err != nil {
		return nil, err
	}
	o := &AddPostPayload{
		&Post{
			Title:   result.Title,
			Content: result.Content,
		},
	}
	return o, nil
}

func (r *mutationResolver) AddComment(ctx context.Context, input AddCommentInput) (*AddCommentPayload, error) {
	i := &usecase.CommentInputData{
		Content: input.Content,
	}
	result, err := r.usecase.AddComment(i)
	if err != nil {
		return nil, err
	}
	o := &AddCommentPayload{
		&Comment{
			Content: result.Content,
		},
	}
	return o, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Viewer(ctx context.Context) (*User, error) {
	result, err := r.usecase.GetUser()
	if err != nil {
		return nil, err
	}
	o := &User{
		Name: result.Name,
	}
	return o, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*Post, error) {
	result, err := r.usecase.GetPosts()
	if err != nil {
		return nil, err
	}
	var o []*Post
	for _, v := range result {
		o = append(o, &Post{
			Title: v.Title,
		})
	}
	return o, nil
}
