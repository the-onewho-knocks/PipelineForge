package worker

import (
	"context"
	"time"

	"pipelineforge/internal/metrics"
	"pipelineforge/internal/queue"
	"pipelineforge/internal/storage"
)

type Handler struct {
	repo storage.Repository
}

func NewHandler(repo storage.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context, msg queue.TrendingRepoMessage) error {
	metrics.RecordGithubRepo(
		msg.Author,
		msg.Name,
		msg.Language,
		msg.URL,
		msg.Stars,
		msg.Forks,
		msg.TodayStars,
	)

	err := h.repo.SaveTrendingRepo(ctx, storage.TrendingRepo{
		Author:     msg.Author,
		Name:       msg.Name,
		URL:        msg.URL,
		Language:   msg.Language,
		Stars:      msg.Stars,
		Forks:      msg.Forks,
		TodayStars: msg.TodayStars,
		ScrapedAt:  msg.ScrapedAt.Format(time.RFC3339),
	})

	if err != nil {
		metrics.DBErrors.Inc()
		return err
	}

	metrics.DBInserts.Inc()
	metrics.MarkScrapeDone()
	return nil
}
