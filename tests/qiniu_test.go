package tests

import (
	"github.com/siriusol/pic-markdown/utils"
	"testing"
)

func TestGenerateUploadToken(t *testing.T) {
	// 每次输出都不同
	t.Log(utils.GenerateUploadToken())
}

func TestUploadFile(t *testing.T) {
	utils.UploadFile()
}
