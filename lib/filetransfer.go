package main

import (
  "log"
  "os"
  "github.com/minio/minio-go"
  "github.com/joho/godotenv"
)

const (  useSSL   =   true
)

func main() {

  // Load env vars
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
  secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
  bucket = :=  os.Getenv("S3_BUCKET")
  // endpoint := "os.Getenv("S3_ENDPOINT")"
  endpoint := "s3.amazon.com"

  // Initialize minio client object.
  minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
  if err != nil {
      log.Fatalln(err)
  }

  n, err := minioClient.PutObject(bucket)

  log.Println("%v", minioClient)
}


