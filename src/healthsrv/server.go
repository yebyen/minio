package healthsrv

import (
	"fmt"
	"net/http"

	minio "github.com/minio/minio-go/v7"
)

// DefaultHost and DefaultPort for minio server
const (
	DefaultHost = "0.0.0.0"
	DefaultPort = 8082
)

// Start starts the healthcheck server on $host:$port and blocks. It only returns if the server fails, with the indicative error
func Start(host string, port int, minioClient *minio.Client) error {
	mux := http.NewServeMux()
	mux.Handle("/healthz", healthZHandler(minioClient))

	hostStr := fmt.Sprintf("%s:%d", host, port)
	return http.ListenAndServe(hostStr, mux)
}
