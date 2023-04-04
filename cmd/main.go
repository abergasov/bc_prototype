package main

import (
	"bc_prototype/internal/config"
	"bc_prototype/internal/logger"
	"bc_prototype/internal/repository/contract"
	membersRepo "bc_prototype/internal/repository/members"
	"bc_prototype/internal/routes"
	circleService "bc_prototype/internal/service/circle"
	"bc_prototype/internal/storage/database"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

var (
	confFile = flag.String("config", "configs/app_conf.yml", "Configs file path")
	appHash  = os.Getenv("GIT_HASH")
)

func main() {
	flag.Parse()
	appLog, err := logger.NewAppLogger(appHash)
	if err != nil {
		log.Fatalf("unable to create logger: %s", err)
	}
	appLog.Info("app starting", zap.String("conf", *confFile))
	appConf, err := config.InitConf(*confFile)
	if err != nil {
		appLog.Fatal("unable to init config", err, zap.String("config", *confFile))
	}

	appLog.Info("create storage connections")
	dbConn, err := getDBConnect(appLog, &appConf.ConfigDB, appConf.MigratesFolder)
	if err != nil {
		appLog.Fatal("unable to connect to db", err, zap.String("host", appConf.ConfigDB.Address))
	}
	defer func() {
		if err = dbConn.Close(); err != nil {
			appLog.Fatal("unable to close db connection", err)
		}
	}()

	appLog.Info("init repositories")
	repoMembers := membersRepo.InitRepo(dbConn)
	repoContract := contract.InitRepo(dbConn)

	appLog.Info("init services")
	service := circleService.InitService(appLog, repoMembers, repoContract)

	appLog.Info("init http service")
	appHTTPServer := routes.InitAppRouter(appLog, service, fmt.Sprintf(":%d", appConf.AppPort))
	defer func() {
		if err = appHTTPServer.Stop(); err != nil {
			appLog.Fatal("unable to stop http service", err)
		}
	}()
	go func() {
		if err = appHTTPServer.Run(); err != nil {
			appLog.Fatal("unable to start http service", err)
		}
	}()

	// register app shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // This blocks the main thread until an interrupt is received
}

func getDBConnect(log logger.AppLogger, cnf *config.DBConf, migratesFolder string) (*database.DBConnect, error) {
	for i := 0; i < 5; i++ {
		dbConnect, err := database.InitDBConnect(cnf, migratesFolder)
		if err == nil {
			return dbConnect, nil
		}
		log.Error("can't connect to db", err, zap.Int("attempt", i))
		time.Sleep(time.Duration(i) * time.Second * 5)
	}
	return nil, fmt.Errorf("can't connect to db")
}
