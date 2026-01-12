package storage

import "context"

type TrendingRepo struct {
	Author     string
	Name       string
	URL        string
	Language   string
	Stars      int
	Forks      int
	TodayStars int
	ScrapedAt  string
}

type Repository interface {
	SaveTrendingRepo(ctx context.Context, repo TrendingRepo) error
	Close()
}
