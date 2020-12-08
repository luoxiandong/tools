package tests

import (
	"github.com/luoxiandong/tools/log"
	"testing"
)

func TestLog(t *testing.T) {
	log.InitLogger("")

	log.Info("Test log format")
}
