# PipeLineForge – GitHub Trending Observability Pipeline

This repository hosts PipelineForge, a production-grade data ingestion and observability backend built using Go.
The system continuously scrapes GitHub Trending repositories, processes them asynchronously using RabbitMQ, persists structured data into PostgreSQL, and exposes real-time metrics, logs, and alerts via Prometheus and Grafana.

The project demonstrates how real-world data pipelines are built, monitored, and operated in production environments.

## The Project Demonstrates
1) Clean modular Go architecture
2) Decoupled producer–consumer pipeline using RabbitMQ
3) Asynchronous job processing with backpressure handling
4) PostgreSQL persistence for processed records
5) First-class observability with metrics, logs, and alerts
6) Custom Prometheus metrics instrumentation
7) Grafana dashboards for real-time visibility
8) Alerting rules for abnormal system behavior
9) Context-aware concurrency using goroutines
10) Production-focused backend and SRE practices

# Table of Contents
- [Architecture Diagram](#Architecture-Diagram)
- [Core Design Principles](#core-design-principles)
- [Technology Stack](#technology-stack)
- [System Components](#system-components)
- [API Design & Routes](#api-design--routes)
- [Getting Started](#getting-started)
- [Author](#author)
- [License](#license)

# Architecture Diagram:-
Below is a high-level overview of the system architecture:

<img width="1425" height="434" alt="pipelineforge arch" src="https://github.com/user-attachments/assets/d4587570-be1b-4ddd-a740-c56176df2f28" />

# Core Design Principles:-
1) Separation of Concerns – Scraping, processing, and storage are isolated
2) Loose Coupling – Services communicate only via the message queue
3) Backpressure Safety – Queue absorbs traffic spikes
4) Observability First – Metrics and alerts are built-in
5) Fail-Safe Design – Graceful shutdown and error isolation
6) Scalable by Default – Components can scale independently

# Technology Stack
1) Go
2) RabbitMQ
3) PostgreSQL
4) Prometheus
5) Grafana

# System Components:-
### Scraper Service (Producer)
1) Scrapes GitHub Trending repositories
2) Extracts repository metadata (name, author, stars, language, URL)
3) Publishes raw data to RabbitMQ
4) Exposes Prometheus metrics

### Message Queue (RabbitMQ)
1) Buffers scraped repository data
2) Decouples ingestion from processing
3) Enables asynchronous and scalable workflows

### Worker Service (Consumer)
1) Consumes messages from RabbitMQ
2) Validates and processes repository data
3) Inserts structured records into PostgreSQL
4) Exposes detailed processing metrics

### Database (PostgreSQL)
1) Stores processed repository data
2) Optimized for write-heavy workloads
3) Ensures durability and consistency

### Data Flow
1) Scraper fetches GitHub Trending repositories
2) Data is published to RabbitMQ
3) Worker consumes messages asynchronously
4) Processed data is stored in PostgreSQL
5) Metrics are scraped by Prometheus
6) Grafana visualizes system health

## Observability
### Dashboards
https://cdn.discordapp.com/attachments/1460311264327631012/1460597009361932381/image.png?ex=69677e77&is=69662cf7&hm=fc1d1b31055dd1a44ec94dac8db027229519bce5978fe9088c9a498d45005e97&
1) Scraper throughput
2) Queue consumption rate
3) Database write rate
4) System health overview

### Alerts
1) High database insert rate
2) Consumer lag detection
3) Abnormal error spikes

