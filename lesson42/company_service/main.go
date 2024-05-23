package main

import (
	"fmt"
	"path/filepath"
	"project/api"
	"project/api/handlers"
	cf "project/config"
	"project/config/logger"
	"project/postgresql"
	"project/postgresql/managers"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path

	db, err := postgresql.ConnectDB(config)
	cf.CheckErr(err)
	defer db.Close()

	cm := managers.NewCompanyManager(db)
	rcm := managers.NewRecruiterManager(db)
	vm := managers.NewVacancyManager(db)
	im := managers.NewInterviewManager(db)

	h := handlers.NewHTTPHandler(cm, rcm, vm, im, *logger)
	r := api.NewGin(h)

	fmt.Printf("Server started on port %s\n", config.HTTPPort)
	logger.INFO.Println("Server started on port " + config.HTTPPort)
	err = r.Run(config.HTTPPort)
	cf.CheckErr(err)
}
