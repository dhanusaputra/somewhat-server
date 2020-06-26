package cert

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

var (
	demoKeyPair  *tls.Certificate
	demoCertPool *x509.CertPool
)

// GetCert returns a certicicate pair and pool
func GetCert() (*tls.Certificate, *x509.CertPool) {
	serverCrt, err := ioutil.ReadFile("cert/server.pem")
	if err != nil {
		log.Fatal(err)
	}
	serverKey, err := ioutil.ReadFile("cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	pair, err := tls.X509KeyPair(serverCrt, serverKey)
	if err != nil {
		log.Fatal(err)
	}
	demoKeyPair = &pair
	demoCertPool = x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM(serverCrt)
	if !ok {
		log.Fatal("bad certs")
	}

	return demoKeyPair, demoCertPool
}
