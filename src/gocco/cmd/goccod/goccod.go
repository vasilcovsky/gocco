package main

import (
	"fmt"
	"gocco"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func githubRawURL(blobURL string) string {
	return "https://raw.githubusercontent.com/" + strings.Replace(blobURL, "/blob/", "/", 1)
}

func githubDownload(url string) (*gocco.SourceFile, error) {
	client := &http.Client{}

	res, err := client.Get(githubRawURL(url))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, http.ErrMissingFile
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	g := &gocco.SourceFile{
		Path:    url,
		Content: content,
		ETag:    res.Header.Get("ETag"),
		Expires: res.Header.Get("Expires"),
	}

	return g, nil
}

func goccoHandler(w http.ResponseWriter, r *http.Request) {
	content, err := githubDownload(r.RequestURI)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("ETag", content.ETag)
	w.Header().Set("Expires", content.Expires)
	w.Write(gocco.GenerateDocumentation(content))
}

// let's Go!
func main() {
	http.HandleFunc("/", goccoHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
