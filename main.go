package main

import (
	"fmt"
	"net/http"
)

const body = `
<html>
  <head>
    <meta name="go-import" content="%s git https://%s/%s/%s">
  </head>
  <body></body>
</html>`

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf(body))
}

func main() {
	http.HandleFunc("/", proxyHandler)
	http.ListenAndServe(":8080", nil)
}
