package db

import (
	"context"
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
