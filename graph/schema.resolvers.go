package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"blog-post-api/graph/model"

	"github.com/google/uuid"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.PostInput) (*model.Post, error) {
	context := GetContext(ctx)
	Id := uuid.NewString()

	post := model.Post{
		ID:       Id,
		Tile:     &input.Tile,
		Content:  &input.Content,
		AuthorID: &input.Author,
	}

	err := context.DB.Create(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input model.PostInput) (*model.Post, error) {
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}
	context := GetContext(ctx)
	var post model.Post
	err := context.DB.Model(&post).Where("id = ?", id).Updates(&model.Post{
		Tile:     &input.Tile,
		Content:  &input.Content,
		AuthorID: &input.Author,
	}).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*model.Post, error) {
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}

	context := GetContext(ctx)
	var post model.Post

	err := context.DB.Where("id = ?", id).Delete(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// AddAuthor is the resolver for the addAuthor field.
func (r *mutationResolver) AddAuthor(ctx context.Context, input model.AuthorInput) (*model.Author, error) {
	context := GetContext(ctx)
	Id := uuid.NewString()
	author := model.Author{ID: Id, FirstName: input.FirstName, LastName: input.LastName}

	err := context.DB.Create(&author).Error
	if err != nil {
		return nil, err
	}

	return &author, nil
}

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id string, input model.AuthorInput) (*model.Author, error) {
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}
	context := GetContext(ctx)
	var author model.Author
	err := context.DB.Model(&author).Where("id = ?", id).Updates(&model.Author{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}).Error

	if err != nil {
		return nil, err
	}

	return &author, nil
}

// DeleteAuthor is the resolver for the deleteAuthor field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}

	context := GetContext(ctx)
	var author model.Author

	err := context.DB.Where("id = ?", id).Delete(&author).Error
	if err != nil {
		return nil, err
	}

	return &author, nil
}

// AddComment is the resolver for the AddComment field.
func (r *mutationResolver) AddComment(ctx context.Context, input model.CommentInput) (*model.Comment, error) {
	context := GetContext(ctx)
	Id := uuid.NewString()
	comment := model.Comment{
		ID:      Id,
		Content: input.Content,
		Post:    input.Post,
	}

	err := context.DB.Create(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.CommentInput) (*model.Comment, error) {
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}
	context := GetContext(ctx)
	var comment model.Comment
	err := context.DB.Model(&comment).Where("id = ?", id).Updates(&model.Comment{
		Content: input.Content,
		Post:    input.Post,
	}).Error

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*model.Comment, error) {
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}

	context := GetContext(ctx)
	var comment model.Comment

	err := context.DB.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// AddLike is the resolver for the AddLike field.
func (r *mutationResolver) AddLike(ctx context.Context, input model.LikeInput) (*model.Like, error) {
	context := GetContext(ctx)
	Id := uuid.NewString()
	like := model.Like{
		ID:    Id,
		Liked: input.Liked,
		Post:  input.Post,
	}

	err := context.DB.Create(&like).Error
	if err != nil {
		return nil, err
	}

	return &like, nil
}

// PostLikesCount is the resolver for the postLikesCount field.
func (r *queryResolver) PostLikesCount(ctx context.Context, id string) (*int64, error) {
	context := GetContext(ctx)
	var count int64
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}
	err := context.DB.Model(&model.Like{}).Where("post = ? AND liked = ?", id, true).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &count, nil
}

// PostCommentsCount is the resolver for the postCommentsCount field.
func (r *queryResolver) PostCommentsCount(ctx context.Context, id string) (*int64, error) {
	context := GetContext(ctx)
	var count int64
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}
	err := context.DB.Model(&model.Comment{}).Where("post = ?", id).Count(&count).Error
	if err != nil {
		return nil, err
	}

	return &count, nil
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	context := GetContext(ctx)
	var authors []*model.Author
	// Get all records
	err := context.DB.Find(&authors).Error
	if err != nil {
		return nil, err
	}
	return authors, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	context := GetContext(ctx)
	var posts []*model.Post
	// Get all records
	err := context.DB.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	if id == string("") {
		return nil, fmt.Errorf("id is required")
	}

	context := GetContext(ctx)
	var post model.Post
	err := context.DB.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
