CREATE TABLE IF NOT EXISTS trending_repositories (
    id SERIAL PRIMARY KEY,

    author TEXT NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,

    language TEXT,
    stars INTEGER NOT NULL,
    forks INTEGER NOT NULL,
    today_stars INTEGER NOT NULL,

    scraped_at TIMESTAMPTZ NOT NULL,

    created_at TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT unique_repo_day UNIQUE (author, name, scraped_at)
);

CREATE INDEX idx_trending_scraped_at
ON trending_repositories (scraped_at);

CREATE INDEX idx_trending_language
ON trending_repositories (language);
