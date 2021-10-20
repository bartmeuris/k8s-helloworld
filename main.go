package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var head = `
<html>
<head></head>
<body>
`

var body = `
	<h1>Hello world!</h1>
	From '%s%s' on '%s'
	<pre>
`
var foot = `
	</pre>
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hn, _ := os.Hostname()
		log.Printf("%s: %s%s", hn, r.URL.Host, r.URL.Path)

		w.Write([]byte(head))
		fmt.Fprintf(w, body, r.Host, r.URL.Path, hn)
		if b, err := httputil.DumpRequest(r, true); err == nil {
			w.Write(b)
		}
		w.Write([]byte(foot))
	})
	http.ListenAndServe(":8080", nil)
}
