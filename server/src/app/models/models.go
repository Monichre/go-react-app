package models
import (
	"database/sql"

)

var db *sql.DB


type User struct {
	ID    int
	Email string
}

func InsertUser(user *User) error {
	var id int
	err := db.QueryRow(`
		INSERT INTO users(email)
		VALUES ($1)
		RETURNING id
	`, user.Email).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func GetUserByID(id int) (*User, error) {
	var email string
	err := db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func RemoveUserByID(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

func Follow(followerID, followeeID int) error {
	_, err := db.Exec(`
		INSERT INTO followers(follower_id, followee_id)
		VALUES ($1, $2)
	`, followerID, followeeID)
	return err
}

func Unfollow(followerID, followeeID int) error {
	_, err := db.Exec(`
		DELETE FROM followers
		WHERE follower_id=$1
		AND followee_id=$2
	`, followerID, followeeID)
	return err
}

func GetFollowerByIDAndUser(id int, userID int) (*User, error) {
	var email string
	err := db.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.follower_id
		AND f.follower_id=$1
		AND f.followee_id=$2
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func GetFollowersForUser(id int) ([]*User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.followee_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, Email: email})
	}
	return users, nil
}

func GetFolloweeByIDAndUser(id int, userID int) (*User, error) {
	var email string
	err := db.QueryRow(`
		SELECT u.email
		FROM users AS u, followers AS f
		WHERE u.id = f.followee_id
		AND f.followee_id=$1
		AND f.follower_id=$2
		LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

func GetFolloweesForUser(id int) ([]*User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.email
		FROM users AS u, followers AS f
		WHERE u.id=f.follower_id
		AND f.follower_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, Email: email})
	}
	return users, nil
}



type Post struct {
	ID     int
	UserID int
	Title  string
	Body   string
}

func InsertPost(post *Post) error {
	var id int
	err := db.QueryRow(`
		INSERT INTO posts(user_id, title, body)
		VALUES ($1, $2, $3)
		RETURNING id
	`, post.UserID, post.Title, post.Body).Scan(&id)
	if err != nil {
		return err
	}
	post.ID = id
	return nil
}

func RemovePostByID(id int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id=$1", id)
	return err
}

func GetPostByID(id int) (*Post, error) {
	var (
		userID      int
		title, body string
	)
	err := db.QueryRow(`
		SELECT user_id, title, body
		FROM posts
		WHERE id=$1
	`, id).Scan(&userID, &title, &body)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:     id,
		UserID: userID,
		Title:  title,
		Body:   body,
	}, nil
}

func GetPostByIDAndUser(id, userID int) (*Post, error) {
	var title, body string
	err := db.QueryRow(`
		SELECT title, body
		FROM posts
		WHERE id=$1
		AND user_id=$2
	`, id, userID).Scan(&title, &body)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:     id,
		UserID: userID,
		Title:  title,
		Body:   body,
	}, nil
}

func GetPostsForUser(id int) ([]*Post, error) {
	rows, err := db.Query(`
		SELECT p.id, p.title, p.body
		FROM posts AS p
		WHERE p.user_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		posts       = []*Post{}
		pid         int
		title, body string
	)
	for rows.Next() {
		if err = rows.Scan(&pid, &title, &body); err != nil {
			return nil, err
		}
		posts = append(posts, &Post{ID: id, UserID: id, Title: title, Body: body})
	}
	return posts, nil
}



type Comment struct {
	ID     int
	UserID int
	PostID int
	Title  string
	Body   string
}

func InsertComment(comment *Comment) error {
	var id int
	err := db.QueryRow(`
		INSERT INTO comments(user_id, post_id, title, body)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, comment.UserID, comment.PostID, comment.Title, comment.Body).Scan(&id)
	if err != nil {
		return err
	}
	comment.ID = id
	return nil
}

func RemoveCommentByID(id int) error {
	_, err := db.Exec("DELETE FROM comments WHERE id=$1", id)
	return err
}

func GetCommentByIDAndPost(id int, postID int) (*Comment, error) {
	var (
		userID      int
		title, body string
	)
	err := db.QueryRow(`
		SELECT user_id, title, body
		FROM posts
		WHERE id=$1
		AND post_id=$2
	`, id, postID).Scan(&userID, &title, &body)
	if err != nil {
		return nil, err
	}
	return &Comment{
		ID:     id,
		UserID: userID,
		PostID: postID,
		Title:  title,
		Body:   body,
	}, nil
}

func GetCommentsForPost(id int) ([]*Comment, error) {
	rows, err := db.Query(`
		SELECT c.id, c.user_id, c.title, c.body
		FROM comments AS c
		WHERE c.post_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		comments    = []*Comment{}
		cid, userID int
		title, body string
	)
	for rows.Next() {
		if err = rows.Scan(&cid, &userID, &title, &body); err != nil {
			return nil, err
		}
		comments = append(comments, &Comment{
			ID:     cid,
			UserID: userID,
			PostID: id,
			Title:  title,
			Body:   body,
		})
	}
	return comments, nil
}