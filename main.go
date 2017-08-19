package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const body = `
<html>
  <head>
    <meta name="go-import" content="%s/%s %s %s">
  </head>
  <body></body>
</html>
`

var (
	vcs    string
	pkgFmt string
	listen string
	host   string
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	pkgRoot := strings.TrimPrefix(r.URL.Path, "/")
	pkgParts := strings.Split(pkgRoot, "/")
	if len(pkgParts) >= 2 {
		pkgRoot = strings.Join(pkgParts[:2], "/")
	}
	repoRoot := fmt.Sprintf(pkgFmt, pkgRoot)
	fmt.Fprint(w, fmt.Sprintf(body, host, pkgRoot, vcs, repoRoot))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	flag.StringVar(&vcs, "vcs", "git", "Version control system to be used by 'go get'.")
	flag.StringVar(&pkgFmt, "pkgfmt", "", `Format for the package source of the package to be fetched e.g. git+ssh://git@private.repo.com/%s.git.`)
	flag.StringVar(&listen, "listen", ":8080", "Address to listen on.")
	flag.StringVar(&host, "host", "", "Hostname where service is running e.g. go.pkg.io.")
	flag.Parse()

	http.HandleFunc("/", proxyHandler)
	http.HandleFunc("/healthz", healthHandler)

	if strings.HasPrefix(listen, ":") {
		listen = "0.0.0.0" + listen
	}
	log.Printf("Listening on http://%s", listen)
	http.ListenAndServe(listen, nil)
}
