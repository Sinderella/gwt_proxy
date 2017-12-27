package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	flag.Parse()
	setCA(caCert, caKey)
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.Verbose = *verbose

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx)(*http.Request, *http.Response) {
			log.Println(r.Host)

			for k, v := range r.Header{
				log.Printf("%s: %v\n", k, v)
			}

			return r, nil
		})

	log.Fatal(http.ListenAndServe(*addr, proxy))
}