package core

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// Secret is a wrapper type around secret data.
type Secret map[string]string

// NewSecret retrieves a new secret from the AWS secret manager and reads
// the data to the Secrets type in the form of map[string]string.
func NewSecret(name string) (*Secret, error) {
	svc := secretsmanager.New(session.New())
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return nil, fmt.Errorf("%s: %s", aerr.Message(), aerr.Error())
		}

		return nil, err
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			return nil, fmt.Errorf("base64 Decode Error: %v", err)
		}
		secretString = string(decodedBinarySecretBytes[:len])
	}

	var data map[string]string
	err = json.Unmarshal([]byte(secretString), &data)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal secret: %v", err)
	}

	s := Secret(data)
	return &s, nil
}

// RSAPublicKey returns an *rsa.PublicKey using the data from the
// Secret with the given key.
func (s Secret) RSAPublicKey(key string) (*rsa.PublicKey, error) {
	log.Println("------------------- FORMATTED DATA --------------------")
	log.Printf(s[key])
	data := []byte(formatRSAPublicKeyData(s[key]))
	log.Println("------------------- FORMATTED DATA --------------------")
	log.Printf(formatRSAPublicKeyData(s[key]))
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("invalid key format")
	}

	pk, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse public key: %v", err)
	}

	return pk, nil
}

// RSAPrivateKey returns an *rsa.PrivateKey using the data from the
// Secret with the given key.
func (s Secret) RSAPrivateKey(key string) (*rsa.PrivateKey, error) {
	log.Println("------------------- FORMATTED DATA --------------------")
	log.Printf(s[key])
	data := []byte(formatRSAPrivateKeyData(s[key]))
	log.Println("------------------- FORMATTED DATA --------------------")
	log.Printf(formatRSAPrivateKeyData(s[key]))
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("invalid key format")
	}

	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse public key: %v", err)
	}

	return pk, nil
}

func formatRSAPrivateKeyData(in string) string {
	parts := strings.Split(in, " ")
	partCount := len(parts)

	data := strings.Join(parts[:4], " ")
	for i := 4; i < partCount-4; i++ {
		data += "\n" + parts[i]
	}
	return data + "\n" + strings.Join(parts[partCount-4:], " ")
}

func formatRSAPublicKeyData(in string) string {
	parts := strings.Split(in, " ")
	partCount := len(parts)

	data := strings.Join(parts[:3], " ")
	for i := 3; i < partCount-3; i++ {
		data += "\n" + parts[i]
	}
	return data + "\n" + strings.Join(parts[partCount-3:], " ")
}
