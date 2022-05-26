package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type redirectPair struct {
	Received    string `json:"received"`
	Destination string `json:"destination"`
}

func getRedirectsFromFile(filename string) (map[string]string, error) {
	redirects := make(map[string]string)
	file, err := os.ReadFile(filename)
	if err != nil {
		return redirects, fmt.Errorf("failed to getRedirectsFromFile(%q):\n\t%s", filename, err.Error())
	}

	var rList []redirectPair
	json.Unmarshal(file, &rList)
	for _, r := range rList {
		redirects[r.Received] = r.Destination
	}

	return redirects, nil
}

func makeRedirectHandler(m map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if v, ok := m[r.URL.Path]; ok {
			http.Redirect(w, r, v, http.StatusFound)
		}
		fmt.Fprintf(w, "Redirect not specified =(")
	}
}

func getFilename() *string {
	filename := flag.String("f", "", "Name of file to read redirections list")
	flag.Parse()
	return filename
}

func main() {
	filename := getFilename()
	redirects, err := getRedirectsFromFile(*filename)
	if err != nil {
		log.Fatalf("Failed to get redirects from %v file:\n\t%v\n", *filename, err)
	}

	log.Printf("Detected map: %#v\n", redirects)
	if err := http.ListenAndServe(":8080", makeRedirectHandler((redirects))); err != nil {
		log.Fatal(err)
	}
}
