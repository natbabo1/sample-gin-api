package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natbabo1/sample-gin-api/internal/book/handler"
	"github.com/natbabo1/sample-gin-api/internal/book/repo"
	"github.com/natbabo1/sample-gin-api/internal/book/service"
	"github.com/natbabo1/sample-gin-api/internal/shared/config"
	"github.com/natbabo1/sample-gin-api/internal/shared/db"
	"github.com/natbabo1/sample-gin-api/internal/shared/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	logr := logger.New(cfg.Env)
	sqlDB, err := db.New(cfg.DB.DSN, cfg.DB.MaxIdleConns, cfg.DB.MaxOpenConns, 200*time.Millisecond)
	if err != nil {
		log.Fatal("db connect", zap.Error(err))
	}

	bookRepo := repo.New(sqlDB)
	bookSvc := service.New(bookRepo)
	bookHdl := handler.New(bookSvc)

	r := gin.New()
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")
	bookHdl.Register(api.Group("/books"))

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logr.Fatal("server error", zap.Error(err))
		}
	}()
	logr.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logr.Fatal("shutdown", zap.Error((err)))
	}
	logr.Info("server existed")
}
