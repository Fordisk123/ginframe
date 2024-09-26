package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	gerrors "github.com/pkg/errors"
	"io"
	"strings"
)

type S3Client struct {
	AccessKey  string           `json:"access_key"`
	SecretKey  string           `json:"secret_key"`
	EndPoint   string           `json:"end_point"`
	DisableSsl bool             `json:"disable_ssl"`
	Region     string           `json:"region"`
	Session    *session.Session `json:"session"`
}

func NewS3FileClient(AccessKey, SecretKey, Region, EndPoint string, DisableSsl bool) (*S3Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(AccessKey, SecretKey, ""),
		Endpoint:         aws.String(EndPoint),
		Region:           aws.String(Region),
		DisableSSL:       aws.Bool(DisableSsl),
		S3ForcePathStyle: aws.Bool(true), //virtual-host style方式，不要修改
	})
	if err != nil {
		return nil, err
	}
	return &S3Client{
		AccessKey:  AccessKey,
		SecretKey:  SecretKey,
		EndPoint:   EndPoint,
		DisableSsl: DisableSsl,
		Region:     Region,
		Session:    sess,
	}, nil
}

func splitBucketNameAndFileName(fullPathFileName string) (string, string) {
	index := strings.Index(fullPathFileName, "/")
	secondIndex := strings.Index(fullPathFileName[index+1:], "/")
	return fullPathFileName[1 : secondIndex+1], fullPathFileName[secondIndex+1:]
}

func (s *S3Client) Upload(ctx context.Context, fileName string, reader io.ReadSeekCloser) error {
	defer func() {
		if reader != nil {
			reader.Close()
		}
	}()
	select {
	case <-ctx.Done():
		return fmt.Errorf("upload %s timeout", fileName)
	default:
		bucketName, filename := splitBucketNameAndFileName(fileName)
		if bucketName == "" || filename == "" {
			return fmt.Errorf("%s filename is invalid of s3", fileName)
		}

		svc := s3.New(s.Session)
		_, err := svc.HeadBucket(&s3.HeadBucketInput{
			Bucket: &bucketName,
		})
		if err != nil && strings.Contains(err.Error(), "The specified bucket does not exist") {
			_, err := svc.CreateBucket(&s3.CreateBucketInput{
				Bucket: &bucketName,
			})
			if err != nil {
				return gerrors.Wrapf(err, "create %s bucket of %s failed", bucketName, fileName)
			}
		}
		if err != nil {
			return gerrors.Wrapf(err, "get %s bucket of %s failed", bucketName, fileName)
		}
		_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(filename),
			Body:   reader,
		})
		return err
	}
}

func (s *S3Client) DownLoad(ctx context.Context, fileName string) (io.ReadCloser, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("download %s timeout", fileName)
	default:
		bucketName, filename := splitBucketNameAndFileName(fileName)
		if bucketName == "" || filename == "" {
			return nil, fmt.Errorf("%s filename is invalid of s3", fileName)
		}
		svc := s3.New(s.Session)
		object, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(filename),
		})
		if err != nil {
			return nil, err
		}
		return object.Body, nil
	}
}
