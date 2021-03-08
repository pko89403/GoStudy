package handler

import (
	"net/http"
	"testing"

	"go-chi-test/platform/mock_http"
	"go-chi-test/platform/newsfeed"
)

func TestNewsfeedPOST(t *testing.T) {
	feed := newsfeed.New()

	headers := http.Header{}
	headers.Add("content-type", "application/json")

	w := &mock_http.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_http.RequestBody(map[string]string{
		"title": "hello",
		"post":  "world",
	})

	handler := NewsfeedPost(feed)
	handler(w, r)

	result := w.GetBodyString()

	if result != "Good Job!" {
		t.Errorf("Handler did not complete")
	}

	if len(feed.GetAll()) != 1 {
		t.Errorf("Item did not add")
	}

	if feed.GetAll()[0].Title != "hello" {
		t.Errorf("Item bad")
	}
}
