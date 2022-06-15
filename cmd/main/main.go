package main

import (
	"flag"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-redis/redis/v8"
	"github.com/lantu-dev/puki/pkg/hwcloud"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"seg/pkg/auth"
	authSetup "seg/pkg/auth/setup"
	"seg/pkg/base/rpc"
	"seg/pkg/socket"
	"seg/pkg/storage"
	testSetup "seg/pkg/test/setup"
)

var buildTag string = "dev"
var buildCommit string = "UNSET"

func main() {

	var err error
	address := flag.String("address", ":8001", "")
	flag.Parse()
	switch flag.Arg(0) {

	case "version":
		fmt.Printf("%s_%s", buildTag, buildCommit)
		return

	}

	log.SetLevel(log.TraceLevel)

	mrds, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mrds.Close()

	rds := redis.NewClient(&redis.Options{Addr: mrds.Addr()})

	var db *gorm.DB
	pgUrl := os.Getenv("PG_URL")
	if pgUrl != "" {
		db, err = gorm.Open(postgres.Open(pgUrl), &gorm.Config{})
		log.Info("using pg")
	} else {
		db, err = gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
		log.Info("using sqlite")
	}

	socket.Setup(db)

	reg := rpc.NewServiceRegistry("api")

	// 每新增一个模块 ( mod ) , 在这里新增一个 Setup 。
	storage.Setup()
	if err := authSetup.Setup(reg, db); err != nil {
		log.Fatal(err)
	}

	if err := testSetup.Setup(reg, db); err != nil {
		log.Fatal(err)
	}

	var smsSender *hwcloud.SMSSender = nil

	hwSmsEndpoint := os.Getenv("HWCLOUD_SMS_ENDPOINT")

	log.Infof("hw sms endpoint: %s", hwSmsEndpoint)
	if hwSmsEndpoint != "" {
		smsSender = &hwcloud.SMSSender{
			URL:       hwSmsEndpoint,
			AppKey:    os.Getenv("HWCLOUD_SMS_APP_KEY"),
			AppSecret: os.Getenv("HWCLOUD_SMS_APP_SECRET"),
			Channel:   os.Getenv("HWCLOUD_SMS_CHANNEL_ID"),
			Signature: os.Getenv("HWCLOUD_SMS_SIGN_NAME"),
		}
	}

	auth.SetupSMSLogin(func(phoneNumber int64, code string) error {
		if smsSender != nil {
			_, err := smsSender.SendMessage(os.Getenv("HWCLOUD_SMS_TEMPLATE_ID"), fmt.Sprintf("+%d", phoneNumber), code)
			if err != nil {
				log.Errorf("hwcloud send message %+v", err)
			}
		}
		spew.Dump(phoneNumber, code)
		return nil
	}, rds)

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/api/", reg)
	log.Infof("server listen @ %s", *address)
	log.Fatal(http.ListenAndServe(*address, nil))
}
