package utils

import (
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"log"
)

// 自定义返回值结构体
type PutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func GenerateUploadToken() string {
	accessKey := "accessKey"
	secretKey := "secretKey"
	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "bucket"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
		// Expires: 7200,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// putPolicy.Expires = 7200 // 有效期以秒为单位
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func UploadFile() {
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	upToken := GenerateUploadToken()
	key := "myKey"
	localFile := "E:/test/test.txt"
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		log.Fatal("Upload error!")
	}
	log.Printf("File:%v\n", ret)
}
