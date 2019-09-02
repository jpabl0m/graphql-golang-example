package graphql

import "context"

type commentResolver struct{ *Resolver }

func (r *commentResolver) Post(ctx context.Context, obj *Comment) (*Post, error) {
	result, err := r.usecase.GetCommentPost(obj.ID)
	if err != nil {
		return nil, err
	}
	o := &Post{
		Title:   result.Title,
		Content: result.Content,
	}
	return o, nil
}
