package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Post struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	Class    string `json:"type"`
	Date     string `json:"date"`
	Caption  string `json:"photo-caption"`
	PhotoUrl string `json:"photo-url-1280"`
}

func (p Post) downloadImage() {
	resp, err := http.Get(p.PhotoUrl)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("Trouble making GET photo request!")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Trouble reading response body!")
	}

	filename := path.Base(p.PhotoUrl)
	if filename == "" {
		log.Fatalf("Trouble deriving file name for %s", p.PhotoUrl)
	}

	err = ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		log.Fatal("Trouble creating file! -- ", err)
	}
}

type TumblrLog struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

type Tumblr struct {
	Blog  TumblrLog `json:"tumblrlog"`
	Posts []Post    `json"posts"`
}

func NewTumblr(url string, page int) Tumblr {
	contents := GetJson(url, page)

	var t Tumblr
	json.Unmarshal(contents, &t)
	return t
}

func (t Tumblr) DownloadImages() {
	for i, post := range t.Posts {
		fmt.Println("Post # ", i)
		fmt.Println(" ---> Caption: ", post.Caption)
		fmt.Println(" ---> Url : ", post.PhotoUrl)

		if post.Class != "photo" {
			fmt.Println(" ---> SKIPPING (not photo post)")
			continue
		}
		post.downloadImage()
		fmt.Println("")
	}
}

func GetJson(url string, page int) []byte {
	contents := restRequest(url, page)
	contents = filterContent(contents, "var tumblr_api_read = ", "", 1)
	contents = filterContent(contents, ";", "", -1)
	return contents
}

func filterContent(data []byte, orig string, target string, n int) []byte {
	c := string(data)
	c = strings.Replace(c, orig, target, n)
	return []byte(c)
}

func restRequest(url string, page int) []byte {
	if page != 1 {
		p := (page - 1) * 20
		url = fmt.Sprintf("%s/api/read/json?start=%d", url, p)
	} else {
		url = fmt.Sprintf("%s/api/read/json", url)
	}

	fmt.Println("Rest Request url: ", url)

	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("Trouble making REST GET request!")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Trouble reading JSON response body!")
	}

	return contents
}

func displayRawJson(contents []byte) {
	var out bytes.Buffer
	err := json.Indent(&out, contents, "", "	")
	if err != nil {
		log.Fatal("Trouble with json Indent!", err)
	}
	fmt.Println("")
	fmt.Println("---")
	out.WriteTo(os.Stdout)
	os.Exit(0)
}
func main() {

	pagePtr := flag.Int("page", 1, "blog page to download")
	rawJsonPtr := flag.Bool("raw", false, "dumpy raw json output for debugging")
	flag.Parse()

	url := flag.Arg(0)

	if url == "" {
		fmt.Fprintf(os.Stderr, "Please supply a tumblr url!\n")
		fmt.Fprintf(os.Stderr, "usage: %s [options] url\n", os.Args[0])
		os.Exit(0)
	}

	if *rawJsonPtr == true {
		contents := GetJson(url, *pagePtr)
		displayRawJson(contents)
		os.Exit(0)
	}

	t := NewTumblr(url, *pagePtr)
	fmt.Println("Blog Title: ", t.Blog.Title, "\n")
	t.DownloadImages()

}