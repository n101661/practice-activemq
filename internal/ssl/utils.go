package ssl

import (
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

// NewCertPool creates a cert pool from system cert pool,
// and append custom cert, if any.
func NewCertPool(appendCertPath string) (*x509.CertPool, error) {
	pool, err := x509.SystemCertPool()
	if err != nil {
		log.Printf("[warn] failed to use system cert pool, create an empty pool instead: %v\n", err)
		pool = x509.NewCertPool()
	}

	if appendCertPath != "" {
		data, err := os.ReadFile(appendCertPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s file: %v", appendCertPath, err)
		}

		if pool.AppendCertsFromPEM(data) {
			log.Println("[info] appended custom cert")
		}
	}

	return pool, err
}
