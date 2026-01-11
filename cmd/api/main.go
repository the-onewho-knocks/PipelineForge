package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Repository struct {
	Author      string
	Name        string
	URL         string
	Description string
	Language    string
	Stars       string
	Forks       string
	TodayStars  string
}

func main() {
	repos, err := scrapeTrendingRepos("daily")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d trending repositories:\n\n", len(repos))
	for i, repo := range repos {
		fmt.Printf("--- Repository #%d ---\n", i+1)
		fmt.Printf("Name: %s/%s\n", repo.Author, repo.Name)
		fmt.Printf("URL: %s\n", repo.URL)
		fmt.Printf("Description: %s\n", repo.Description)
		fmt.Printf("Language: %s\n", repo.Language)
		fmt.Printf("Total Stars: %s\n", repo.Stars)
		fmt.Printf("Forks: %s\n", repo.Forks)
		fmt.Printf("Stars Today: %s\n", repo.TodayStars)
		fmt.Println()
	}
}

func cleanText(text string) string {
	// Remove extra whitespace and newlines
	re := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(re.ReplaceAllString(text, " "))
}

func scrapeTrendingRepos(timeRange string) ([]Repository, error) {
	// timeRange can be: "daily", "weekly", "monthly"
	url := fmt.Sprintf("https://github.com/trending?since=%s", timeRange)
	
	// Create HTTP client with user agent
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	
	// Set user agent to avoid being blocked
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching trending page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %v", err)
	}

	var repos []Repository

	doc.Find("article.Box-row").Each(func(i int, s *goquery.Selection) {
		repo := Repository{}

		// Get repository name and author
		nameLink := s.Find("h2 a")
		href, _ := nameLink.Attr("href")
		repo.URL = "https://github.com" + href
		
		parts := strings.Split(strings.TrimPrefix(href, "/"), "/")
		if len(parts) >= 2 {
			repo.Author = strings.TrimSpace(parts[0])
			repo.Name = strings.TrimSpace(parts[1])
		}

		// Get description
		repo.Description = cleanText(s.Find("p.col-9").Text())

		// Get language
		langSpan := s.Find("span[itemprop='programmingLanguage']")
		repo.Language = cleanText(langSpan.Text())

		// Get stats from the footer section
		// Find all links in the footer area
		s.Find("div.f6 a").Each(func(j int, link *goquery.Selection) {
			href, exists := link.Attr("href")
			if !exists {
				return
			}
			
			text := cleanText(link.Text())
			
			// Stars link
			if strings.HasSuffix(href, "/stargazers") {
				repo.Stars = text
			}
			
			// Forks link
			if strings.HasSuffix(href, "/forks") {
				repo.Forks = text
			}
		})

		// Get today's stars - look for the span with star icon in the right section
		s.Find("span.d-inline-block.float-sm-right").Each(func(j int, span *goquery.Selection) {
			// Check if this span contains the star icon
			if span.Find("svg.octicon-star").Length() > 0 {
				text := cleanText(span.Text())
				// This should be the "X stars today" text
				if strings.Contains(text, "star") {
					// Extract just the number
					re := regexp.MustCompile(`([\d,]+)`)
					matches := re.FindStringSubmatch(text)
					if len(matches) > 0 {
						repo.TodayStars = matches[1] + " stars today"
					}
				}
			}
		})

		// If TodayStars is still empty, try alternative selector
		if repo.TodayStars == "" {
			todaySpan := s.Find("span.float-sm-right")
			todaySpan.Each(func(j int, span *goquery.Selection) {
				if span.Find("svg.octicon-star").Length() > 0 {
					text := cleanText(span.Text())
					if text != "" && text != repo.Stars {
						repo.TodayStars = text
					}
				}
			})
		}

		repos = append(repos, repo)
	})

	return repos, nil
}