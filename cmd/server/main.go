package main

import (
	"github.com/spf13/viper"
	"github.com/zekroTJA/visitor-count/internal/config"
	"github.com/zekroTJA/visitor-count/internal/database"
	"github.com/zekroTJA/visitor-count/internal/log"
	"github.com/zekroTJA/visitor-count/internal/web"
)

func main() {
	config.Init()

	log.SetLogLevel(viper.GetInt("log.level"))

	db := new(database.MySql)
	err := db.Connect(viper.GetString("db.dsn"))
	if err != nil {
		log.Log.Fatalf("failed connecting to database: %s", err.Error())
	}
	log.Log.Infof("connected to database")

	ws := web.New(db)
	addr := viper.GetString("ws.addr")
	log.Log.Infof("web server running on %s", addr)
	if err = ws.ListenAndServeBlocking(addr); err != nil {
		log.Log.Fatalf("failed exposing web server: %s", err.Error())
	}
}
