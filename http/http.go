package http

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"time"
)

func Get_http_status(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	fmt.Println("status:", resp.StatusCode)
	return resp.StatusCode, nil
}

func Get_http_ping(url string) (time.Duration, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	start := time.Now()

	resp, err := client.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	fmt.Println("elapsed time:", elapsed)
	return elapsed, nil
}

func Get_https_cert(url string) ([]*x509.Certificate, error) {
	conn, err := tls.Dial("tcp", url+":443", nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Println("Expiration date:", cert.NotAfter.Format(time.RFC3339))
	}
	return certs, nil
}
