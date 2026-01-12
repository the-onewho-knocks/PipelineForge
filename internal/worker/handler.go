package worker

import (
	"context"
	"pipelineforge/internal/queue"
	"pipelineforge/internal/storage"
	"time"
)

type Handler struct {
	repo storage.Repository
}

func NewHandler(repo storage.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context, msg queue.TrendingRepoMessage) error {
	return h.repo.SaveTrendingRepo(ctx, storage.TrendingRepo{
		Author:     msg.Author,
		Name:       msg.Name,
		URL:        msg.URL,
		Language:   msg.Language,
		Stars:      msg.Stars,
		Forks:      msg.Forks,
		TodayStars: msg.TodayStars,
		ScrapedAt:  msg.ScrapedAt.Format(time.RFC3339),
	})
}
