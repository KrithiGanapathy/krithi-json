package krithiJson

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func post(url string, json string) {
	http.Post(url, "application/json", bytes.NewBuffer([]byte(json)))
}

func postByte(url string, json []byte) {
	http.Post(url, "application/json", bytes.NewBuffer(json))
}

func get(url string) string {
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func main() {

}
