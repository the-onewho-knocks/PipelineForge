package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Worker
	MessagesConsumed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "pipelineforge",
			Subsystem: "worker",
			Name:      "messages_consumed_total",
			Help:      "Total number of RabbitMQ messages consumed",
		},
	)

	MessageProcessingDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "pipelineforge",
			Subsystem: "worker",
			Name:      "message_processing_seconds",
			Help:      "Time taken to process a single message",
		},
	)

	// DB
	DBInserts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "pipelineforge",
			Subsystem: "postgres",
			Name:      "inserts_total",
			Help:      "Total number of successful DB inserts",
		},
	)

	DBErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "pipelineforge",
			Subsystem: "postgres",
			Name:      "errors_total",
			Help:      "Total number of DB errors",
		},
	)
)

// GitHub metrics
var (
	GithubRepoStars = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "pipelineforge",
			Subsystem: "github",
			Name:      "repo_stars",
			Help:      "Stars of a GitHub repository",
		},
		[]string{"author", "repo", "language", "url"},
	)

	GithubRepoForks = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "pipelineforge",
			Subsystem: "github",
			Name:      "repo_forks",
			Help:      "Forks of a GitHub repository",
		},
		[]string{"author", "repo", "language"},
	)

	GithubRepoTodayStars = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "pipelineforge",
			Subsystem: "github",
			Name:      "repo_today_stars",
			Help:      "Stars gained today",
		},
		[]string{"author", "repo", "language"},
	)

	GithubReposTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "pipelineforge",
			Subsystem: "github",
			Name:      "repos_total",
			Help:      "Repositories processed per scrape",
		},
		[]string{"language"},
	)

	GithubStarsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "pipelineforge",
			Subsystem: "github",
			Name:      "stars_total",
			Help:      "Total stars per language per scrape",
		},
		[]string{"language"},
	)

	GithubLastScrapedAt = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "pipelineforge",
			Subsystem: "github",
			Name:      "last_scraped_timestamp",
			Help:      "Last successful scrape time",
		},
	)
)

func Register() {
	prometheus.MustRegister(
		MessagesConsumed,
		MessageProcessingDuration,
		DBInserts,
		DBErrors,
		GithubRepoStars,
		GithubRepoForks,
		GithubRepoTodayStars,
		GithubReposTotal,
		GithubStarsTotal,
		GithubLastScrapedAt,
	)
}

func ResetGithubMetrics() {
	GithubRepoStars.Reset()
	GithubRepoForks.Reset()
	GithubRepoTodayStars.Reset()
	GithubReposTotal.Reset()
	GithubStarsTotal.Reset()
}

func RecordGithubRepo(author, repo, lang, url string, stars, forks, today int) {
	GithubRepoStars.WithLabelValues(author, repo, lang, url).Set(float64(stars))
	GithubRepoForks.WithLabelValues(author, repo, lang).Set(float64(forks))
	GithubRepoTodayStars.WithLabelValues(author, repo, lang).Set(float64(today))
	GithubReposTotal.WithLabelValues(lang).Inc()
	GithubStarsTotal.WithLabelValues(lang).Add(float64(stars))
}

func MarkScrapeDone() {
	GithubLastScrapedAt.Set(float64(time.Now().Unix()))
}
