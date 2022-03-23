package main

import (
	"github.com/betaapskaita/beta-server/handler"
	"github.com/betaapskaita/beta-server/libs/databases"
	"github.com/betaapskaita/beta-server/libs/rpcs"
	"github.com/betaapskaita/beta-server/libs/web"
	"github.com/betaapskaita/beta-server/repositories"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func getDBWithInitialData() databases.Database {
	db := databases.NewMariaDB("dainius:nimda@tcp(localhost:3306)/coreacc_dev?charset=utf8mb4&parseTime=True&loc=Local")

	return db
}

func main() {

	db := getDBWithInitialData()

	accountRepository := repositories.NewAccountRepository(db)

	rpcServer := rpcs.NewServer()

	handler.NewGreeterServer(rpcServer, accountRepository)

	web.Start(":9000", rpcServer)
}
