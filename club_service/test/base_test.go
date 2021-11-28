package test

import (
	"github.com/andrewd92/timeclub/club_service/infrastructure/connection"
	"github.com/andrewd92/timeclub/club_service/server"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	connection.SetTestEnvironment()
	_ = os.Setenv("VIPER_CONFIG_NAME", "test")
	go server.StartApplication()
	time.Sleep(50 * time.Millisecond)
	os.Exit(m.Run())
}
