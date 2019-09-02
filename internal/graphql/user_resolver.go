package graphql

import "context"

type userResolver struct{ *Resolver }

func (r *userResolver) Posts(ctx context.Context, obj *User) ([]*Post, error) {
	result, err := r.usecase.GetUserPosts(obj.ID)
	if err != nil {
		return nil, err
	}
	var o []*Post
	for _, v := range result {
		o = append(o, &Post{
			Title: v.Title,
		})
	}
	return o, err
}
