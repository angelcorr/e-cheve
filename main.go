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

func main() {
	fmt.Println("Starting...\n-")

	file, err := os.Open("enron_mail_20110402.tgz")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	gR, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	tR := tar.NewReader(gR)
	for {
		h, err := tR.Next()
		if err == io.EOF {
			break
		}

		if h.Typeflag != tar.TypeReg {
			continue
		}

		pM, err := mail.ReadMessage(tR)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}

		mH := pM.Header
		d := mH.Get("Date")
		f := mH.Get("From")
		t := mH.Get("To")
		s := mH.Get("Subject")

		b, err := io.ReadAll(pM.Body)
		if err != nil {
			log.Fatal(err)
		}

		zE := os.Getenv("ZINCSEARCH_ENDPOINT")
		zU := os.Getenv("ZINCSEARCH_USER")
		zP := os.Getenv("ZINCSEARCH_PASSWORD")

		e, err := json.Marshal(Email{Date: d, From: f, To: t, Subject: s, Body: string(b)})
		req, err := http.NewRequest("POST", zE, bytes.NewReader(e))
		if err != nil {
			log.Fatal(err)
		}
		req.SetBasicAuth(zU, zP)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		log.Println(resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
}
