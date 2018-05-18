package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Please input: <CA.cert> <Client.cert> <Client.key> <Server Addr>")
		os.Exit(-1)
	}

	caCert, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read ca.cert:", err)
		os.Exit(-1)
	}

	var pool = x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	cliCert, err := tls.LoadX509KeyPair(os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println("Failed to load client cert:", err)
		os.Exit(-1)
	}

	var client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{cliCert},
			},
		},
	}

	resp, err := client.Get(os.Args[4])
	if err != nil {
		fmt.Println("Failed to get server:", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		os.Exit(-1)
	}
	fmt.Println("Response:\n\t", string(body))
}
