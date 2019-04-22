package lib

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ParseURL validate url input, return the url with the protocol scheme, if the url is invalid error will be returned
func ParseURL(urlString string) (string, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}

	return url.String(), nil
}

// GetURLResponce make http request and print the md4 of the responce
func GetURLResponce(address string) {
	URLString, err := ParseURL(address)
	if err != nil {
		fmt.Println(address, "invalid url")
		return
	}

	resp, err := http.Get(URLString)
	if err != nil {
		fmt.Println(URLString, "failed to fetch url")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	data := []byte(body)
	fmt.Printf("%s %x\n", URLString, md5.Sum(data))
}
