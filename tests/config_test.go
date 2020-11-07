package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/siriusol/pic-markdown/utils"
)

func TestConfig(t *testing.T) {
	fmt.Println(os.Getwd())

	path := "../conf/config.txt"
	fmt.Println(utils.GetValueByConfig(path, "access_key"))
	fmt.Println(utils.GetValueByConfig(path, "secret_key"))
	fmt.Println(utils.GetValueByConfig(path, "host"))
	fmt.Println(utils.GetValueByConfig(path, "bucket"))
}
