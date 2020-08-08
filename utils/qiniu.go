package utils

import (
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"log"
	"strings"
	"time"
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
	accessKey := "t48lashcsN_X6hG4PXiF7ITYPwyBrkAhJw4QLfWQ"
	secretKey := "lE1TFvATyOJ2Wo6EF-CTy3ZANZMYXanTJGAlzzOM"
	mac := qbox.NewMac(accessKey, secretKey)
	bucket := "ther-pic"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
		// Expires: 7200,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// putPolicy.Expires = 7200 // 有效期以秒为单位
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func UploadFile(file string) {

	log.Printf("Upload file %s start...", file)
	index := strings.LastIndex(file, "\\")
	key := file[index+1:] + "?time=" + time.Now().Format("2006-01-02 15:04:05")
	log.Println(key)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	upToken := GenerateUploadToken()
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, file, &putExtra)
	if err != nil {
		log.Fatal("Upload error:", err)
	}
	log.Printf("File:%v\n", ret)

	log.Printf("Upload file %s Success!", file)
}
