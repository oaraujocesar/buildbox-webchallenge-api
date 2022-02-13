package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	arg := CreatePostParams{
		ImageUrl: "https://avatars.githubusercontent.com/u/61256606?s=400&u=87690975a21547ae6f62be49567d2d712383fa16&v=4",
		Name:     "César O. Araújo",
		Message:  "This is a test message",
	}

	post, err := testQueries.CreatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.ImageUrl, post.ImageUrl)
	require.Equal(t, arg.Name, post.Name)
	require.Equal(t, arg.Message, post.Message)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)

	return post
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	createdPost := createRandomPost(t)

	post, err := testQueries.GetPost(context.Background(), createdPost.ID)

	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, createdPost.ImageUrl, post.ImageUrl)
	require.Equal(t, createdPost.Name, post.Name)
	require.Equal(t, createdPost.Message, post.Message)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.CreatedAt)
}

func TestListPosts(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}

	arg := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListPosts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestDeletePost(t *testing.T) {
	post := createRandomPost(t)

	err := testQueries.DeletePost(context.Background(), post.ID)

	require.NoError(t, err)

	deletedPost, err := testQueries.GetPost(context.Background(), post.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedPost)
}
