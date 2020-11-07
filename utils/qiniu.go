package utils

import (
	"context"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// 获取七牛云配置
func GetQiNiuConfig(configPath string) QiNiuConfig {
	config := QiNiuConfig{}
	if exist := CheckFileExist(configPath); !exist {
		log.Fatalf("[GetQiNiuConfig] [CheckFileExist] file %s is not exist!", configPath)
		return config
	}
	config.AccessKey = GetValueByConfig(configPath, "access_key")
	config.SecretKey = GetValueByConfig(configPath, "secret_key")
	config.Host = GetValueByConfig(configPath, "host")
	config.Bucket = GetValueByConfig(configPath, "bucket")

	return config
}

// 生成上传 token
func GenerateUploadToken(config QiNiuConfig) string {
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
		// Expires: 7200,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// putPolicy.Expires = 7200 // 有效期以秒为单位
	upToken := putPolicy.UploadToken(mac)

	return upToken
}

func UploadFile(filePath string, config QiNiuConfig) (cloudUrl string, err error) {
	log.Printf("Upload file %s start...", filePath)
	index := strings.LastIndex(filePath, "\\")
	key := filePath[index+1:] + "-on" + time.Now().Format("2006-01-02at15-04-05")
	log.Println(key)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	upToken := GenerateUploadToken(config)
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		log.Fatal("Upload error:", err)
		return "", err
	}
	log.Printf("File:%+v\n", ret)
	log.Printf("Upload file %s Success!", filePath)
	encodeKey := url.QueryEscape(ret.Key)
	encodeKey = strings.ReplaceAll(encodeKey, "+", "%20")

	return config.Host + "/" + encodeKey, nil
}
