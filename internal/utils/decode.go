package utils

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"fmt"
	"io"

	http "github.com/vimbing/fhttp"

	"github.com/andybalholm/brotli"
)

func GetResponseBody(r *http.Response) (string, error) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	decodedBody, err := DecodeFhttp(r.Header, body)

	if err != nil {
		return "", err
	}

	return decodedBody, nil
}

func DecodeFhttp(headers http.Header, body []byte) (string, error) {
	defer func() (string, error) {
		if err := recover(); err != nil {
			return "", errors.New("asd")
		}
		return "", nil
	}()

	var encoding string

	if len(headers["Content-Encoding"]) == 0 {
		encoding = "NAN"
	} else {
		encoding = headers["Content-Encoding"][0]
	}

	if encoding == "br" {
		decodedBody, err := unBrotliData(body)

		if err != nil {
			return "", err
		}

		return string(decodedBody), nil
	} else if encoding == "deflate" {
		decodedBody, err := enflateData(body)

		if err != nil {
			return "", err
		}

		return string(decodedBody), nil
	} else if encoding == "gzip" {
		decodedBody, err := gUnzipData(body)

		if err != nil {
			return "", err
		}

		return string(decodedBody), nil
	} else {
		return (string(body)), nil
	}
}

func unBrotliData(data []byte) (resData []byte, err error) {
	br := brotli.NewReader(bytes.NewReader(data))
	respBody, err := io.ReadAll(br)
	return respBody, err
}

func enflateData(data []byte) (resData []byte, err error) {
	zr, _ := zlib.NewReader(bytes.NewReader(data))
	defer zr.Close()
	enflated, err := io.ReadAll(zr)
	return enflated, err
}

func gUnzipData(data []byte) (resData []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("")
		}
	}()
	gz, _ := gzip.NewReader(bytes.NewReader(data))
	defer gz.Close()
	respBody, err := io.ReadAll(gz)
	return respBody, err
}
