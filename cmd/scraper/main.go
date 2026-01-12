package main

import (
	"context"
	"log"
	"pipelineforge/internal/config"
	"pipelineforge/internal/producer"
	"pipelineforge/internal/queue"
	"pipelineforge/internal/scraper"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg := config.Load()

	// RabbitMQ
	rmq, err := queue.NewRabbitMQ(cfg.RabbitURL, cfg.QueueName)
	if err != nil {
		log.Fatal("rabbitmq:", err)
	}
	defer rmq.Close()

	prod := producer.New(rmq)

	log.Println("scraping github trending repositories...")

	repos, err := scraper.ScrapeTrendingRepos(ctx, "daily")
	if err != nil {
		log.Fatal("scrape error:", err)
	}

	published := 0
	for _, r := range repos {
		msg := queue.TrendingRepoMessage{
			Author:     r.Author,
			Name:       r.Name,
			URL:        r.URL,
			Language:   r.Language,
			Stars:      r.Stars,
			Forks:      r.Forks,
			TodayStars: r.TodayStars,
			ScrapedAt:  time.Now().UTC(),
		}

		if err := prod.Publish(ctx, msg); err != nil {
			log.Println("publish failed:", err)
			continue
		}
		published++
	}

	log.Printf("scraper finished: published %d/%d repositories\n", published, len(repos))
}
