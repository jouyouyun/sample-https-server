package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type handlerInfo struct{}

func (handler *handlerInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI, r.Method)
	fmt.Fprintln(w, "<span style=\"color: red; align: center;\">Hello, world!</span>")
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Please input: <CA.cert> <Server.cert> <Server.key> <port>")
		os.Exit(-1)
	}

	caCert, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read CA.cert:", err)
		os.Exit(-1)
	}

	var pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	var server = http.Server{
		Addr:    ":" + os.Args[4],
		Handler: &handlerInfo{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert, //验证客户端证书
		},
	}

	err = server.ListenAndServeTLS(os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println("Failed to start tls server:", err)
		os.Exit(-1)
	}
}
