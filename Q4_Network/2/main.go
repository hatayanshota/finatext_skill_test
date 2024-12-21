package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

const (
	caCertFile = "ca_cert.pem"
	certsDir   = "./certs"
)

func main() {
	caCertBytes, err := os.ReadFile(caCertFile)
	if err != nil {
		log.Fatalf("failed to read source cert: %v", err)
	}

	caCert, err := parseCertificate(string(caCertBytes))
	if err != nil {
		log.Fatalf("failed to parse source cert: %v", err)
	}

	certFiles, err := os.ReadDir(certsDir)
	if err != nil {
		log.Fatalf("failed to read cert dir: %v", err)
	}

	for _, certFile := range certFiles {
		certBytes, err := os.ReadFile(certsDir + "/" + certFile.Name())
		if err != nil {
			log.Fatalf("failed to read cert file: %v", err)
		}

		cert, err := parseCertificate(string(certBytes))
		if err != nil {
			log.Fatalf("failed to parse cert: %v", err)
		}

		if string(cert.AuthorityKeyId) == string(caCert.SubjectKeyId) {
			log.Printf("result cert file: %v", certFile.Name())
		}
	}
}

// parseCertificate PEM形式の証明書をパースする
func parseCertificate(pemData string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, fmt.Errorf("pem block is empty")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}
