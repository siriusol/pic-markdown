package utils

// 自定义返回值结构体
type PutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

// 七牛云配置结构体
type QiNiuConfig struct {
	AccessKey string
	SecretKey string
	Host      string
	Bucket    string
}
