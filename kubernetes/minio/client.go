package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint        = "192.168.197.128:9000"
	accessKeyID     = "miniothomas"
	secretAccessKey = "miniothomas123"

	minioClient *minio.Client
)

func init() {
	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func uploadFile(ctx context.Context, filePath, bucket, objectName string) error {
	res, err := minioClient.FPutObject(ctx, bucket, objectName, filePath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully uploaded %s of size %+v\n", objectName, res)
	return nil
}

func upload(ctx context.Context, bucket, objectName string, data []byte) error {
	buf := bytes.NewReader(data)
	res, err := minioClient.PutObject(ctx, bucket, objectName, buf, buf.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Successfully uploaded %s of size %+v\n", objectName, res)
	return nil
}

func getObject(ctx context.Context, bucket, objectName string) ([]byte, error) {
	object, err := minioClient.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer object.Close()

	return ioutil.ReadAll(object)
}

func createBucket(ctx context.Context, bucket string) error {
	err := minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		if exists, err2 := minioClient.BucketExists(ctx, bucket); err2 == nil && exists {
			log.Printf("We already own %s\n", bucket)
			return nil
		} else {
			log.Fatal(err)
			return err
		}
	}
	log.Printf("Successfully created %s\n", bucket)
	return nil
}

func do1() {
	createBucket(context.Background(), "bucket-01")
	for i := 1; i <= 1000; i++ {
		str := fmt.Sprintf("%d-%d-%d-%d", i, i, i, i)
		upload(context.Background(), "bucket-01", fmt.Sprintf("file-%d.txt", i), []byte(str))
	}
}

func do2() {
	bucket := "bucket-01"
	for i := 1; i <= 1000; i++ {
		objectName := fmt.Sprintf("file-%d.txt", i)
		data, err := getObject(context.Background(), bucket, objectName)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("download success, bucket: %s, objectName: %s, data: %s", bucket, objectName, data)
	}
}

func do3() {
	bucket := "bucket-01"
	for i := 1; i <= 1000; i++ {
		objectName := fmt.Sprintf("file-%d.txt", i)
		storePath := filepath.Join("download", objectName)
		if err := minioClient.FGetObject(context.Background(), bucket, objectName, storePath, minio.GetObjectOptions{}); err != nil {
			log.Fatal(err)
		}
		log.Printf("download success, bucket: %s, objectName: %s", bucket, objectName)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
