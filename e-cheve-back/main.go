package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/emails", getEmails)

	log.Println("Server listening in port 3000...")
	http.ListenAndServe(":3000", r)
}

func getEmails(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	if term == "" {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	jsonTerm, err := json.Marshal(term)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	query := fmt.Sprintf(`{
		"search_type": "match",
		"query": {
				"term": %v,
				"start_time": "2024-01-01T00:00:00Z",
				"end_time": "2024-12-31T00:00:00Z"
		},
		"from": 0,
		"max_results": 10,
		"_source": []
	}`, string(jsonTerm))

	ZINCSEARCH_ENDPOINT := os.Getenv("ZINCSEARCH_ENDPOINT")
	ZINCSEARCH_USER := os.Getenv("ZINCSEARCH_USER")
	ZINCSEARCH_PASSWORD := os.Getenv("ZINCSEARCH_PASSWORD")

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/api/emails/_search", ZINCSEARCH_ENDPOINT), strings.NewReader(query))
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	req.SetBasicAuth(ZINCSEARCH_USER, ZINCSEARCH_PASSWORD)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
