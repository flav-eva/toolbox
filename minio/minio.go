package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

// Init initializes the minio client
func Init(minioAccessKeyID, minioSecretAccessKey, minioEndpoint string, minioSecure bool) {
	var err error
	creds := credentials.NewStaticV4(minioAccessKeyID, minioSecretAccessKey, "")
	minioClient, err = minio.New(minioEndpoint, &minio.Options{
		Creds:  creds,
		Secure: minioSecure,
	})
	if err != nil {
		panic(err)
	}
}

// Client returns the minio client
func Client() *minio.Client {
	return minioClient
}
