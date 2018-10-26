package types

import (
	"strconv"
	"github.com/graphql-go/graphql"
)

import . "app/models"

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Email, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	UserType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetPostByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("posts", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PostType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetPostsForUser(user.ID)
			}
			return []Post{}, nil
		},
	})
	UserType.AddFieldConfig("follower", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Follower ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetFollowerByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("followers", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetFollowersForUser(user.ID)
			}
			return []User{}, nil
		},
	})
	UserType.AddFieldConfig("followee", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Followee ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetFolloweeByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("followees", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetFolloweesForUser(user.ID)
			}
			return []User{}, nil
		},
	})
}


var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*Post); ok == true {
					return post.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*Post); ok == true {
					return post.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*Post); ok == true {
					return post.Body, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	PostType.AddFieldConfig("user", &graphql.Field{
		Type: graphql.NewNonNull(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*Post); ok == true {
				return GetUserByID(post.UserID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comment", &graphql.Field{
		Type: CommentType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*Post); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetCommentByIDAndPost(id, post.ID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comments", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CommentType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*Post); ok == true {
				return GetCommentsForPost(post.ID)
			}
			return []Comment{}, nil
		},
	})
}


var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*Comment); ok == true {
					return comment.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*Comment); ok == true {
					return comment.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*Comment); ok == true {
					return comment.Body, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	CommentType.AddFieldConfig("user", &graphql.Field{
		Type: UserType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if comment, ok := p.Source.(*Comment); ok == true {
				return GetUserByID(comment.UserID)
			}
			return nil, nil
		},
	})
	CommentType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			i := p.Args["id"].(string)
			id, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			return GetPostByID(id)
		},
	})
}