// author: ashing
// time: 2020/7/12 2:51 下午
// mail: axingfly@gmail.com
// Less is more.

package util

import (
	"crypto/tls"
	"io/ioutil"
	"log"

	"golang.org/x/net/http2"
)

func GetTLSConfig(certPemPath, certKeyPath string) *tls.Config {
	var certKeyPair *tls.Certificate
	cert, _ := ioutil.ReadFile(certPemPath)
	key, _ := ioutil.ReadFile(certKeyPath)

	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		log.Printf("TLS KeyPair err: %v\n", err)
		return nil
	}

	certKeyPair = &pair

	return &tls.Config{
		Certificates: []tls.Certificate{*certKeyPair},
		NextProtos:   []string{http2.NextProtoTLS},
	}
}
