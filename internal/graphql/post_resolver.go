package graphql

import "context"

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *Post) (*User, error) {
	result, err := r.usecase.GetPostAuthor(obj.ID)
	if err != nil {
		return nil, err
	}
	o := &User{
		Name: result.Name,
	}
	return o, nil
}

func (r *postResolver) Comments(ctx context.Context, obj *Post) ([]*Comment, error) {
	result, err := r.usecase.GetPostComments(obj.ID)
	if err != nil {
		return nil, err
	}
	var o []*Comment
	for _, v := range result {
		o = append(o, &Comment{
			Content: v.Content,
		})
	}
	return o, nil
}
