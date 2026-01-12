package queue

import "time"

// TrendingRepoMessage is the event sent from scraper → workers
type TrendingRepoMessage struct {
	Author     string    `json:"author"`
	Name       string    `json:"name"`
	URL        string    `json:"url"`
	Language   string    `json:"language"`
	Stars      int       `json:"stars"`
	Forks      int       `json:"forks"`
	TodayStars int       `json:"today_stars"`
	ScrapedAt  time.Time `json:"scraped_at"`
}
