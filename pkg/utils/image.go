package utils

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

func ImageBase64() string {
	// bytes, err := ioutil.ReadFile("./file.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	resp, err := http.Get("https://issuer.iamza-sandbox.com/user.jpg")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)

	return base64Encoding
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
