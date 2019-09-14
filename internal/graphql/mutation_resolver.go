package graphql

import (
	"context"

	"github.com/stobita/graphql-golang-example/internal/usecase"
)

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

func (r *mutationResolver) SignUp(ctx context.Context, input SignUpInput) (*SignUpPayload, error) {
	i := &usecase.UserInputData{
		Email:    input.Email,
		Password: input.Password,
	}
	if err := r.usecase.SignUp(i); err != nil {
		return &SignUpPayload{Result: false}, err
	}
	return &SignUpPayload{Result: true}, nil
}

func (r *mutationResolver) SignIn(ctx context.Context, input SignInInput) (*SignInPayload, error) {
	i := &usecase.UserInputData{
		Email:    input.Email,
		Password: input.Password,
	}
	if err := r.usecase.SignUp(i); err != nil {
		return &SignInPayload{Result: false}, err
	}
	return &SignInPayload{Result: true}, nil
}
