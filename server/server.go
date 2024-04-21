package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func createServerWithCustomCert() *http.Server {
	certRaw, err := os.ReadFile("ca-cert.pem")
	if err != nil {
		slog.Error("failed to read cert file", "error", err.Error())
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(certRaw)

	tlsConfig := &tls.Config{
		ClientCAs:  certPool,
		ClientAuth: tls.RequireAndVerifyClientCert, // This enforces the client to present its certificate.
	}

	return &http.Server{
		Addr:      ":8101",
		TLSConfig: tlsConfig,
	}

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	slog.Info("starting server")

	server := createServerWithCustomCert()

	// Register handlers.
	http.HandleFunc("/", helloWorld)

	// Listen & wait.
	err := server.ListenAndServeTLS("server/cert/cert.pem", "server/cert/key.pem")
	if err != nil {
		slog.Error("serving failed", "error", err.Error())
	}
	slog.Info("stopped server")
}
