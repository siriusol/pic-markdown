package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/siriusol/pic-markdown/utils"
)

func TestConfig(t *testing.T) {
	fmt.Println(os.Getwd())

	path := "F:/STUDY/study-go/src/github.com/siriusol/pic-markdown/conf/config.txt"
	fmt.Println(utils.Get(path, "access_key"))
	fmt.Println(utils.Get(path, "secret_key"))
	fmt.Println(utils.Get(path, "host"))
	fmt.Println(utils.Get(path, "bucket"))
}
