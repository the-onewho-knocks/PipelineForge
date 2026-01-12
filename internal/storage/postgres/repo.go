package postgres

import (
	"context"
	"pipelineforge/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{db: db}
}

func (r *Repo) SaveTrendingRepo(ctx context.Context, repo storage.TrendingRepo) error {
	query := `
		INSERT INTO trending_repositories (
			author, name, url, language,
			stars, forks, today_stars, scraped_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (author, name, scraped_at)
		DO NOTHING
	`

	_, err := r.db.Exec(
		ctx,
		query,
		repo.Author,
		repo.Name,
		repo.URL,
		repo.Language,
		repo.Stars,
		repo.Forks,
		repo.TodayStars,
		repo.ScrapedAt,
	)

	return err
}

func (r *Repo) Close() {
	r.db.Close()
}
