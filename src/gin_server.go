package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	server, err := newTLSServer(os.Args[1], ":"+os.Args[4])
	if err != nil {
		fmt.Println("Failed to new server:", err)
		return
	}

	router := gin.Default()
	router.GET("/todos", func(c *gin.Context) {
		fmt.Println("Get todo list")
		c.String(200, "[]")
	})
	server.Handler = router

	err = server.ListenAndServeTLS(os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println("Failed to listen and serve:", err)
	}
}

func newTLSServer(ca, addr string) (*http.Server, error) {
	caCert, err := ioutil.ReadFile(ca)
	if err != nil {
		fmt.Println("Failed to read CA.cert:", err)
		return nil, err
	}

	var pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	var server = http.Server{
		Addr: addr,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert, //验证客户端证书
		},
	}
	return &server, nil
}
