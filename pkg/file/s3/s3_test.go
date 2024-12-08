package s3

import (
	"context"
	"io"
	"os"
	"testing"
)

func TestS3(t *testing.T) {

	s3, err := NewS3FileClient("minio", "4553283@wch", "us-east-1", "118.31.20.93:9000", true)
	if err != nil {
		panic(err)
	}

	open, err := os.Open("/tmp/123.txt")

	if err != nil {
		panic(err)
	}

	err = s3.Upload(context.Background(), "/public/123.txt", open)

	if err != nil {
		panic(err)
	}

}

func TestS3Download(t *testing.T) {

	s3, err := NewS3FileClient("minio", "4553283@wch", "us-east-1", "118.31.20.93:9000", true)
	if err != nil {
		panic(err)
	}

	load, err := s3.Download(context.Background(), "/public/123.txt")
	if err != nil {
		panic(err)
	}

	bb, err := io.ReadAll(load)

	if err != nil {
		panic(err)
	}

	println(string(bb))

	load.Close()

}

func TestS3Delete(t *testing.T) {

	s3, err := NewS3FileClient("minio", "4553283@wch", "us-east-1", "118.31.20.93:9000", true)
	if err != nil {
		panic(err)
	}

	err = s3.Delete(context.Background(), "/public/123.txt")
	if err != nil {
		panic(err)
	}

}
