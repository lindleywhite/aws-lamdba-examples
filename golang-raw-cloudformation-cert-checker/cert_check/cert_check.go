
package main

import (
	"crypto/tls"
	"net"
	"os"
	"strings"
	"errors"
	"log"
	"time"
	"fmt"
	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	dialer = &net.Dialer{Timeout: 5 * time.Second}
)

type CertificateResponse struct {
	Server string `json:server`
	Date string `json:certDate`
	ExpiresIn string `json:expiresIn`
}

type CertificateInvalid struct {
	Server string
}

func check(server string)  (certificate CertificateResponse, err error) {
	conn, err := tls.DialWithDialer(dialer, "tcp", server+":443", nil)
	if err != nil {
		log.Printf("%s Error: %v", server, err)
		panic(errors.New("Dialer Error"))
	}
	defer conn.Close()
	valid := conn.VerifyHostname(server)

	for _, c := range conn.ConnectionState().PeerCertificates {

		if valid == nil {
			certificate = CertificateResponse{server, c.NotAfter.Format("2006-01-02"), humanize.Time(c.NotAfter)}
			log.Printf("%v", certificate)
		} else {
			log.Printf("%s | %v\n", server, valid)
		}
		return certificate, nil
	}

	return
}

func main() {
	lambda.Start(check_cert)
}

func check_cert() (values string, err error){
	var s []CertificateResponse
	// collect list of server names
	names := getNames()

	for _, name := range names {
		certificate, err := check(name)

		if err == nil {
			s = append(s, certificate)
		}
	}
	log.Printf("%v", s)
	js, err := json.Marshal(s)

	values = fmt.Sprint(string(js))
	return values, nil
}

func getNames() (names []string) {

	domainString := os.Getenv("DOMAINS")

	names = strings.Split(domainString, ",")
	return
}
