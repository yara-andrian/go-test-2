package main

import (
	"bytes"
	"net/http"
	"time"

	"test/config"
	"test/logger"
	"test/router"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	logger.Info("HELLO")
	var addr bytes.Buffer
	addr.WriteString("0.0.0.0:")
	addr.WriteString(viper.GetString("port"))
	server := &http.Server{
		Handler:      router.Create(),
		Addr:         addr.String(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// server := router.Create()
	logger.Infof("Attempting to listen on http://%s", addr.String())
	logger.Fatal(server.ListenAndServe())
}
