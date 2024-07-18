package setup

import (
	"fmt"

	"digitao.io/website/app"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func SetupObjstorage(configuration *app.Configuration) (*minio.Client, error) {
	endpoint := fmt.Sprintf("%s:%d", configuration.Objstorage.Host, configuration.Objstorage.Port)
	minioClient, err := minio.New(
		endpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				configuration.Objstorage.User,
				configuration.Objstorage.Password,
				"",
			),
			Secure: false,
		},
	)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
