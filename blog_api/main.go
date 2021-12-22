package main

import (
	"api/conf"
	"api/db"
	"api/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	Db, _ := db.Db.DB()
	defer Db.Close()

	// 加載日誌
	log := utils.Log()

	gin.SetMode(conf.Conf.Server.Model)

	// 路由

	srv := &http.Server{
		Addr: conf.Conf.Server.Address,
	}

	go func() {
		// 啟動服務
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s \n", err)
		}
		log.Fatal("listen: %s \n", conf.Conf.Server.Address)
	}()

	quit := make(chan os.Signal)
	// 監聽消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
