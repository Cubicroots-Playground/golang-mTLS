package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func createClientWithCustomCert() *http.Client {
	caCert, err := os.ReadFile("ca-cert.pem")
	if err != nil {
		slog.Error("failed to read ca cert file", "error", err.Error())
	}
	clientCert, err := tls.LoadX509KeyPair("client/cert/cert.pem", "client/cert/key.pem")
	if err != nil {
		slog.Error("failed loading cert", "error", err.Error())
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      certPool,
				Certificates: []tls.Certificate{clientCert},
			},
		},
	}

}

func main() {
	client := createClientWithCustomCert()

	r, err := client.Get("https://localhost:8101")
	if err != nil {
		slog.Error("request failed", "error", err.Error())
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("reading body failed", "error", err.Error())
	}

	fmt.Printf("%s\n", body)
}
