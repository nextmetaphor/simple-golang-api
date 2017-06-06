package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

const listenPort = 8080

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		encodableRequest := struct {
			Method           string
			URL              *url.URL
			Proto            string
			ProtoMajor       int
			ProtoMinor       int
			Header           http.Header
			Body             io.ReadCloser
			ContentLength    int64
			TransferEncoding []string
			Close            bool
			Host             string
			Form             url.Values
			PostForm         url.Values
			MultipartForm    *multipart.Form
			Trailer          http.Header
			RemoteAddr       string
			RequestURI       string
			TLS              *tls.ConnectionState
			Response         *http.Response
		}{
			request.Method,
			request.URL,
			request.Proto,
			request.ProtoMajor,
			request.ProtoMinor,
			request.Header,
			request.Body,
			request.ContentLength,
			request.TransferEncoding,
			request.Close,
			request.Host,
			request.Form,
			request.PostForm,
			request.MultipartForm,
			request.Trailer,
			request.RemoteAddr,
			request.RequestURI,
			request.TLS,
			request.Response,
		}

		if encodeErr := json.NewEncoder(writer).Encode(encodableRequest); encodeErr != nil {
			log.Print(encodeErr)
		}

	})

	log.Println("simple-golang-api listening on port " + strconv.Itoa(listenPort) + "...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(listenPort), nil))
}
