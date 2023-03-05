package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/crazyvan/learn-go-with-tests/blogposts"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("whoops, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	firstBody := `Title: Post 1
Description: Hai 1
Tags: tdd, go
---
B
L
M`

	secondBody := `Title: Post 2
Description: Hai 2
Tags: rust, borrow-checked
---
Hello
from
the other side.`

	fs := fstest.MapFS{
		"hello world. md": {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Hai 1",
		Tags:        []string{"tdd", "go"},
		Body: `B
L
M`,
	}

	assertPost(t, got, want)
}

func TestReturnsFSError(t *testing.T) {
	_, err := blogposts.NewPostsFromFS(StubFailingFS{})

	if err == nil {
		t.Fatal("wanted error from FS")
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
