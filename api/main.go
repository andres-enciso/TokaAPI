package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"TokaAPI/internal/auth"
	"TokaAPI/internal/db"
	"TokaAPI/internal/tasks"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func requestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" {
			id = time.Now().UTC().Format("20060102150405.000000000")
		}
		c.Set("reqID", id)
		c.Writer.Header().Set("X-Request-ID", id)
		c.Next()
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("porque lo eliminas si ya estaba en el repo")
	}

	database := db.Connect()
	auth.EnsureAdminUser(database,
		os.Getenv("ADMIN_USER"),
		os.Getenv("ADMIN_PASS"),
	)

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), requestID())

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	protected := r.Group("/tasks", auth.BasicAuthMiddleware(database))
	tasks.RegisterRoutes(protected, database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{Addr: ":" + port, Handler: r}

	go func() {
		log.Printf("escuchando e n :%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
