PipelineForge/
├── cmd/
│   ├── scraper/
│   │   └── main.go              # Scraper → RabbitMQ producer
│   │
│   └── worker/
│       └── main.go              # RabbitMQ consumer → Postgres
│
├── internal/
│   ├── scraper/
│   │   ├── github.go
│   │   └── github_test.go
│   │
│   ├── queue/
│   │   ├── message.go           # Shared RabbitMQ message contract
│   │   └── rabbitmq.go          # RabbitMQ connection + queue declare
│   │
│   ├── producer/
│   │   └── producer.go          # Publishes messages to RabbitMQ
│   │
│   ├── worker/
│   │   └── consumer.go          # Consumes RabbitMQ messages
│   │
│   ├── storage/
│   │   ├── repository.go        # Storage interface
│   │   └── postgres/
│   │       ├── db.go             # pgxpool setup
│   │       ├── repo.go           # Postgres repo implementation
│   │       └── repo_test.go
│   │
│   ├── metrics/
│   │   └── prometheus.go        # Prometheus metrics
│   │
│   └── config/
│       └── config.go             # Env/config loader
│
├── migrations/
│   └── 001_create_trending.sql
│
├── monitoring/
│   ├── prometheus/
│   │   └── prometheus.yml       # Prometheus config
│   └── grafana/
│       └── dashboards/
│           └── pipeline.json    # Grafana dashboard
├── scripts/
│   ├── run_postgres.sh          # Start Postgres locally
│   ├── run_rabbitmq.sh          # Start RabbitMQ locally
│   ├── run_prometheus.sh        # Start Prometheus
│   └── run_grafana.sh           # Start Grafana
│
├── .env.example
├── go.mod
├── go.sum
└── README.md
