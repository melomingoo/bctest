package main

import (
	"bc_melomingoo/common"
	"bc_melomingoo/handler"
	"bc_melomingoo/middleware"
	"bc_melomingoo/model"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"time"
)

var config *common.Config
var testDB *gorm.DB

func Init() bool {
	config = new(common.Config)
	err := configor.New(&configor.Config{}).Load(config, "env.toml")
	if err != nil {
		log.Fatalf("fail to configure [%v]\n", err)
		return false
	}

	var logErr error
	if logErr != nil {
		return false
	}

	var dbErr error
	testDB, dbErr = makeDBConnection(config.DB.DriverName, config.DB.DataSourceName, config.DB.MaxOpenConns, config.DB.MaxIdleConns)
	if dbErr != nil {
		return false
	}

	autoMigrate(testDB)
	return true
}

func main() {

	log.Println("main.go main start")
	defer log.Println("main.go main closed")

	if Init() == false {
		log.Fatal("init failed")
		return
	}

	router := mux.NewRouter()

	router.Use(
		//미들웨어
		middleware.CacheProxy,
		middleware.AuthProxy,
	)
	Handler(router)
	http.Handle("/", router)

	storeManager := common.NewStoreManager(testDB, config)
	storeManager.Start()

	port := config.Port

	log.Println(" ")
	log.Println("=================================================================")
	log.Println(fmt.Sprintf("Server starting on port %v", port))
	log.Println("=================================================================")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	log.Println("main exiting")

}

func Handler(router *mux.Router) {
	baseHandler := &handler.BaseHandler{
		Config: config,
		TestDB: testDB,
	}
	CoinHandler := (*handler.CoinHandler)(baseHandler)
	testHandler := (*handler.TestHandler)(baseHandler)
	router.HandleFunc("/testCheck", testHandler.TestHandlerCheck).Methods(http.MethodGet)
	router.HandleFunc("/coin/latest", CoinHandler.CoinHandlerList).Methods(http.MethodGet)
	router.HandleFunc("/coin/history", CoinHandler.CoinHandlerHistory).Methods(http.MethodGet)
	router.HandleFunc("/coin/change", CoinHandler.CoinHandlerChange).Methods(http.MethodGet)
	router.HandleFunc("/coin/latest3", CoinHandler.CoinHandlerList2).Methods(http.MethodGet)

}

func makeDBConnection(driverName string, databaseSource string,
	maxConn, maxIdleConn int) (*gorm.DB, error) {
	db, err := gorm.Open(driverName, databaseSource)
	if err != nil {
		return nil, err
	}
	db.DB().SetConnMaxLifetime(time.Duration(30) * time.Second)

	if maxConn > 0 {
		db.DB().SetMaxOpenConns(maxConn)
	} else {
		db.DB().SetMaxOpenConns(10)
	}

	// Max Idle Connection이 설정되었다면
	if maxIdleConn > 0 {
		// Max Connection 설정 여부 확인
		if maxConn > 0 {
			// Max Idle Connection이 Max Connection 보다 클 수 없으므로 크면 Max Connection으로 Max Idle Connection 설정
			if maxConn >= maxIdleConn {
				db.DB().SetMaxIdleConns(maxIdleConn)
				db.DB()
			} else {
				db.DB().SetMaxIdleConns(maxConn)
			}
		} else {
			// Max Connection 설정 없이 Max Idle Connection만 설정되면 의미가 없으므로 경고 출력 후 무시
		}
	} else {
		db.DB().SetMaxIdleConns(3)
	}
	return db, nil
}

func autoMigrate(testDB *gorm.DB) {
	log.Println("auto migration start...")
	testDB.AutoMigrate(&model.Test{}, &model.Coin{})
}
