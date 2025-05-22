package services

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"

	"github.com/lubie-placki-be/configs"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var opt = options.GridFSBucket().SetName("images")

func UploadImage(file multipart.File) (any, error) {
	bucket := configs.Client.Database("images").GridFSBucket(opt)

	filename := bson.NewObjectID()
	uploadStream, err := bucket.OpenUploadStream(
		context.TODO(), filename.String(),
	)
	if err != nil {
		return "", err
	}
	defer uploadStream.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if _, err = uploadStream.Write(fileContent); err != nil {
		return "", err
	}

	if err := uploadStream.Close(); err != nil {
		return "", err
	}

	return uploadStream.FileID, nil
}

func DownloadImage(id string) (*bytes.Buffer, error) {
	bucket := configs.Client.Database("images").GridFSBucket(opt)

	fileId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	fileBuffer := bytes.NewBuffer(nil)
	if _, err := bucket.DownloadToStream(context.TODO(), fileId, fileBuffer); err != nil {
		return &bytes.Buffer{}, err
	}

	return fileBuffer, nil
}
