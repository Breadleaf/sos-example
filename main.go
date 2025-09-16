package main

import (
	"fmt"
	"os"
	"bufio"

	sos "github.com/breadleaf/sos/pkg/http/client"
)

func main() {
	client := sos.NewClient("http://0.0.0.0:8080")

	// list all buckets
	buckets, err := client.ListBuckets()
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"[sos client] failed to list buckets: %v\n",
			err,
		)
	}

	// print the buckets
	for idx, bucket := range buckets {
		fmt.Printf("buckets[%d] = %s\n", idx, bucket)
	}

	// list all objects in files bucket
	bucketName := "files"
	files, err := client.ListObjects(bucketName)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"[sos client] failed to list objects in bucket(%s): %v\n",
			bucketName,
			err,
		)
	}

	// print files in files bucket
	for idx, file := range files {
		fmt.Printf("files[%d] = %s\n", idx, file)
	}

	// read file from files bucket
	objectName := "documents/LICENSE"
	object, err := client.GetObject(bucketName, objectName)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"[sos client] failed to get object(%s) from bucket(%s): %v\n",
			objectName,
			bucketName,
			err,
		)
	}
	defer object.Close()

	// print object content
	scanner := bufio.NewScanner(object)
	idx := 0
	for scanner.Scan() {
		fmt.Printf("%s[%03d] = %s\n", objectName, idx, scanner.Text())
		idx += 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(
			os.Stderr,
			"[sos client] bufio scanner error: %v\n",
			err,
		)
	}
}
