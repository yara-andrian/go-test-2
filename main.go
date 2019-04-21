package main

import (
	"bytes"

	"go_test/logger"

	"github.com/spf13/viper"
)

func main() {
	logger.Info("HELLO")
	var addr bytes.Buffer
	addr.WriteString("0.0.0.0:")
	addr.WriteString(viper.GetString("port"))
	logger.Infof("Attempting to listen on http://%s", addr.String())
}
