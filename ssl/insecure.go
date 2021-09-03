package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

var (
	// Cert is a self signed certificate
	Cert tls.Certificate
	// CertPool contains the self signed certificate
	CertPool *x509.CertPool
)

func init() {
	var err error
	// Load Server cert
	Cert, err = tls.LoadX509KeyPair("ssl/server-cert.pem", "ssl/server-key.pem")
	if err != nil {
		log.Fatalln("Failed to parse key pair:", err)
	}
	Cert.Leaf, err = x509.ParseCertificate(Cert.Certificate[0])
	if err != nil {
		log.Fatalln("Failed to parse certificate:", err)
	}
    // Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("ssl/ca-cert.pem")
	if err != nil {
        log.Fatalln("Failed to read certificate CA:", err)
    }

	CertPool = x509.NewCertPool()
	CertPool.AddCert(Cert.Leaf)
	if !CertPool.AppendCertsFromPEM(pemServerCA) {
        log.Fatalln("failed to add server CA's certificate")
    }
}
