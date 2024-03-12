package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
)

type Email struct {
	Date    string
	From    string
	To      string
	Subject string
	Body    string
}

func Index(filename string, limit int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	tarReader := tar.NewReader(gzipReader)
	i := 0
	for {
		if i == limit {
			break
		}

		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		// If it's not a file, then continue
		if header.Typeflag != tar.TypeReg {
			continue
		}

		message, err := mail.ReadMessage(tarReader)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}

		date := message.Header.Get("Date")
		from := message.Header.Get("From")
		to := message.Header.Get("To")
		subject := message.Header.Get("Subject")
		body, err := io.ReadAll(message.Body)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}

		ZINCSEARCH_ENDPOINT := os.Getenv("ZINCSEARCH_ENDPOINT")
		ZINCSEARCH_USER := os.Getenv("ZINCSEARCH_USER")
		ZINCSEARCH_PASSWORD := os.Getenv("ZINCSEARCH_PASSWORD")

		i++
		email, err := json.Marshal(Email{Date: date, From: from, To: to, Subject: subject, Body: string(body)})
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.NewRequest("POST", ZINCSEARCH_ENDPOINT, bytes.NewReader(email))
		if err != nil {
			fmt.Println("Error", err)
			continue
		}

		req.SetBasicAuth(ZINCSEARCH_USER, ZINCSEARCH_PASSWORD)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Println("Error", resp.StatusCode)
			continue
		}
		fmt.Println("Document indexed")
	}

	fmt.Println("All documents indexed")
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("File name argument not provided")
	}

	fmt.Println("Starting...\n-")
	filename := os.Args[1]
	Index(filename, -1)
}
