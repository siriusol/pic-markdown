package utils

import (
	"context"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
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
	nowPath, err := os.Getwd()
	if err != nil {
		log.Printf("[GenerateUploadToken] [Getwd] error:%v", err)
		return ""
	}
	configPath := nowPath + "\\conf\\config.txt"
	accessKey := Get(configPath, "access_key")
	secretKey := Get(configPath, "secret_key")
	bucket := Get(configPath, "bucket")
	mac := qbox.NewMac(accessKey, secretKey)

	putPolicy := storage.PutPolicy{
		Scope: bucket,
		// Expires: 7200,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// putPolicy.Expires = 7200 // 有效期以秒为单位
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func UploadFile(file string) (cloudUrl string) {

	log.Printf("Upload file %s start...", file)
	index := strings.LastIndex(file, "\\")
	key := file[index+1:] + "?time=" + time.Now().Format("2006-01-02#15:04:05")
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

	encodeKey := url.QueryEscape(ret.Key)

	return "qip9mf1yt.hb-bkt.clouddn.com/" + encodeKey
}
