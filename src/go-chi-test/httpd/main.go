package main

import (
	"fmt"
	"net/http"

	"go-chi-test/httpd/handler"
	"go-chi-test/platform/newsfeed"

	"github.com/go-chi/chi"
)

func main() {
	port := ":9000"
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		Title: "Hello",
		Post:  "World",
	})
	r := chi.NewRouter()

	// Get NewsFeed
	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	// Post NewsFeed
	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on port: " + port)
	http.ListenAndServe(port, r)
}
