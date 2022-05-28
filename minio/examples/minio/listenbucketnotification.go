//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/zhangdapeng520/zdpgo_ftp/minio"
)

func main() {
	// Note: YOUR-ACCESSKEYID, YOUR-SECRETACCESSKEY and my-bucketname are
	// dummy values, please replace them with original values.

	// Requests are always secure (HTTPS) by default. Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().

	// New returns an Amazon S3 compatible client object. API compatibility (v2 or v4) is automatically
	// determined based on the Endpoint value.
	minioClient, err := minio.New("play.min.io", "YOUR-ACCESS", "YOUR-SECRET", true)
	if err != nil {
		log.Fatalln(err)
	}

	// s3Client.TraceOn(os.Stderr)

	// Create a done channel to control 'ListenBucketNotification' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	// Listen for bucket notifications on "mybucket" filtered by prefix, suffix and events.
	for notificationInfo := range minioClient.ListenBucketNotification("YOUR-BUCKET", "PREFIX", "SUFFIX", []string{
		"s3:ObjectCreated:*",
		"s3:ObjectAccessed:*",
		"s3:ObjectRemoved:*",
	}, doneCh) {
		if notificationInfo.Err != nil {
			log.Fatalln(notificationInfo.Err)
		}
		log.Println(notificationInfo)
	}
}
