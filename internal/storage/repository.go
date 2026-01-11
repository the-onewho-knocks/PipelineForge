package storage

import "context"

type RepositoryRecord struct {
	Author      string
	Name        string
	URL         string
	Description string
	Language    string
	Stars       int
	Forks       int
	TodayStars  int
}

type Repository interface {
	Save(ctx context.Context, repo RepositoryRecord)
	List(ctx context.Context, limit int) ([]RepositoryRecord, error)
	Clear(ctx context.Context) error
}
