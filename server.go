package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var bs bytes.Buffer
		bs.WriteString(fmt.Sprintf("URL: %s%s\n", r.Host, r.URL))
		bs.WriteString(fmt.Sprintf("Method: %s\n", r.Method))
		bs.WriteString("Headers:\n")
		for key, value := range r.Header {
			bs.WriteString(fmt.Sprintf("  %s: %s\n", key, value))
		}
		body, err := ioutil.ReadAll(r.Body)
		if err == nil && len(body) > 0 {
			bs.WriteString("Body:\n")
			bs.WriteString(string(body))
		}

		bs.WriteString(fmt.Sprintf("Raw: %+v\n", r))
		log.Printf("Request:\n%s\n", bs.String())
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
