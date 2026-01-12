package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pipelineforge/internal/config"
	"pipelineforge/internal/metrics"
	"pipelineforge/internal/queue"
	"pipelineforge/internal/storage/postgres"
	"pipelineforge/internal/worker"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()

	db, err := postgres.NewPool(ctx, cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.New(db)
	defer repo.Close()

	rmq, err := queue.NewRabbitMQ(cfg.RabbitURL, cfg.QueueName)
	if err != nil {
		log.Fatal(err)
	}
	defer rmq.Close()

	metrics.Register()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("metrics exposed on :2112/metrics")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	handler := worker.NewHandler(repo)
	consumer := worker.NewConsumer(rmq, handler)

	if err := consumer.Start(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("worker started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
