package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func CreateOrLoadCerts() (certLoc string, keyLoc string) {
	certLoc = "./certs/cert.pub"
	keyLoc = "./certs/priv.key"

	if !FileExists(certLoc) && !FileExists(keyLoc) {
		Logger("CERTIFICATE HELPER", "Server certificate is not defined, generating self-signed one", false)
		// Check if certs folder exist
		if _, err := os.Stat("./certs"); os.IsNotExist(err) {
			Logger("CERTIFICATE HELPER", "Folder not exist, creating...", false)
			os.Mkdir("./certs", 0755)
		}

		key, err := rsa.GenerateKey(rand.Reader, 4096)
		if err != nil {
			Logger("CERTIFICATE HELPER", "Cannot generate the rsa key", true)
		}
		keyBytes := x509.MarshalPKCS1PrivateKey(key)
		// PEM encoding of private key
		keyPEM := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: keyBytes,
			},
		)

		subject := pkix.Name{
			CommonName:   os.Getenv("APP_NAME"),
			Country:      []string{"ID"},
			Locality:     []string{"GREATER JAKARTA"},
			Province:     []string{"GREATER JAKARTA"},
			Organization: []string{"CORPORATION"},
		}

		issuer := pkix.Name{
			CommonName:   os.Getenv("APP_NAME"),
			Country:      []string{"ID"},
			Locality:     []string{"GREATER JAKARTA"},
			Province:     []string{"GREATER JAKARTA"},
			Organization: []string{"CORPORATION"},
		}

		template := x509.Certificate{
			SerialNumber:          big.NewInt(0),
			Subject:               subject,
			Issuer:                issuer,
			SignatureAlgorithm:    x509.SHA384WithRSA,
			NotBefore:             time.Now(),
			NotAfter:              time.Now().Add(365 * 24 * 10 * time.Hour),
			BasicConstraintsValid: true,
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyAgreement | x509.KeyUsageKeyEncipherment | x509.KeyUsageDataEncipherment,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		}

		derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
		if err != nil {
			Logger("CERTIFICATE HELPER", "Cannot generate the certificate", true)

		}

		//pem encoding of certificate
		certPem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "CERTIFICATE",
				Bytes: derBytes,
			},
		)

		// Write priv key
		WriteFile(keyLoc, keyPEM)
		// Write pub certs
		WriteFile(certLoc, certPem)
	}

	return certLoc, keyLoc
}
