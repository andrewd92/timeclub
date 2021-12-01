package main_flow

import (
	"github.com/andrewd92/timeclub/tests/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.InitConfig()

	os.Exit(m.Run())
}
