package postgres

import (
	"context"
	"pipelineforge/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repo {
	return &Repo{
		pool: pool,
	}
}

func (r *Repo) Save(ctx context.Context,
	repo storage.RepositoryRecord) error {

	_, err := r.pool.Exec(ctx, `
			insert into trending_repositories(
			author,
			name,
			url,
			description,
			language,
			stars,
			forks,
			today_stars
			)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
			ON CONFLICT (author, name) DO UPDATE SET
				stars = EXCLUDED.stars,
				forks = EXCLUDED.forks,
				today_stars = EXCLUDED.today_stars,
				description = EXCLUDED.description,
				language = EXCLUDED.language
		`,
		repo.Author,
		repo.Name,
		repo.URL,
		repo.Description,
		repo.Language,
		repo.Stars,
		repo.Forks,
		repo.TodayStars,
	)

	return err
}
