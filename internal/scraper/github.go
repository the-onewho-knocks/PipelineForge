package scraper

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Repository struct {
	Author      string
	Name        string
	URL         string
	Description string
	Language    string
	Stars       int
	Forks       int
	TodayStars  int
}

var (
	spaceRe       = regexp.MustCompile(`\s+`)
	todayStarsRe = regexp.MustCompile(`([\d,]+)`)
)

func cleanText(text string) string {
	return strings.TrimSpace(spaceRe.ReplaceAllString(text, " "))
}

func parseNumber(text string) int {
	text = strings.ReplaceAll(text, ",", "")
	n, _ := strconv.Atoi(text)
	return n
}

// scrapeFromReader allows testing without HTTP
func scrapeFromReader(r io.Reader) ([]Repository, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var repos []Repository

	doc.Find("article.Box-row").Each(func(_ int, s *goquery.Selection) {
		var repo Repository

		link := s.Find("h2 a")
		href, ok := link.Attr("href")
		if !ok {
			return
		}

		parts := strings.Split(strings.TrimPrefix(href, "/"), "/")
		if len(parts) < 2 {
			return
		}

		repo.Author = parts[0]
		repo.Name = parts[1]
		repo.URL = "https://github.com" + href

		repo.Description = cleanText(s.Find("p").First().Text())
		repo.Language = cleanText(
			s.Find("span[itemprop='programmingLanguage']").Text(),
		)

		s.Find("a").Each(func(_ int, a *goquery.Selection) {
			href, _ := a.Attr("href")
			text := cleanText(a.Text())

			switch {
			case strings.HasSuffix(href, "/stargazers"):
				repo.Stars = parseNumber(text)
			case strings.HasSuffix(href, "/forks"):
				repo.Forks = parseNumber(text)
			}
		})

		s.Find("span").Each(func(_ int, span *goquery.Selection) {
			text := cleanText(span.Text())
			if strings.Contains(text, "stars today") {
				if m := todayStarsRe.FindStringSubmatch(text); len(m) > 1 {
					repo.TodayStars = parseNumber(m[1])
				}
			}
		})

		if repo.Author == "" || repo.Name == "" {
			return
		}

		repos = append(repos, repo)
	})

	return repos, nil
}

func ScrapeTrendingRepos(ctx context.Context, timeRange string) ([]Repository, error) {
	url := fmt.Sprintf("https://github.com/trending?since=%s", timeRange)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "PipelineForge/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return scrapeFromReader(resp.Body)
}
