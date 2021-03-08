package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var body = `
<html>
<head></head>
<body>
	<h1>Hello world!</h1>
	From '%s%s' on '%s'
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hn, _ := os.Hostname()
		log.Printf("%s: %s%s", hn, r.URL.Host, r.URL.Path)
		fmt.Fprintf(w, body, r.Host, r.URL.Path, hn)
	})
	http.ListenAndServe(":8080", nil)
}
