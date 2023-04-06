package http

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func Get_http_status(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("status:", resp.StatusCode)
	return nil
}

func Check_http_ping(url string) error {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	start := time.Now()

	resp, err := client.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	fmt.Println("elapsed time:", elapsed)
	return nil
}

func Get_https_cert(url string) error {
	conn, err := tls.Dial("tcp", url+":443", nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Println("Expiration date:", cert.NotAfter.Format(time.RFC3339))
	}
	return nil
}
