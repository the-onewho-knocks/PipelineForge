package scraper

import (
	"strings"
	"testing"
)

const sampleHTML = `
<article class="Box-row">
	<h2>
		<a href="/golang/go"> golang / go </a>
	</h2>
	<p>A programming language</p>

	<span itemprop="programmingLanguage">Go</span>

	<div class="f6">
		<a href="/golang/go/stargazers">110,000</a>
		<a href="/golang/go/forks">17,000</a>
	</div>

	<span>1,234 stars today</span>
</article>
`

func TestScrapeFromReader(t *testing.T) {
	repos, err := scrapeFromReader(strings.NewReader(sampleHTML))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repos) != 1 {
		t.Fatalf("expected 1 repo, got %d", len(repos))
	}

	r := repos[0]

	if r.Author != "golang" {
		t.Errorf("author mismatch: %s", r.Author)
	}

	if r.Name != "go" {
		t.Errorf("name mismatch: %s", r.Name)
	}

	if r.Language != "Go" {
		t.Errorf("language mismatch: %s", r.Language)
	}

	if r.Stars != 110000 {
		t.Errorf("stars mismatch: %d", r.Stars)
	}

	if r.Forks != 17000 {
		t.Errorf("forks mismatch: %d", r.Forks)
	}

	if r.TodayStars != 1234 {
		t.Errorf("today stars mismatch: %d", r.TodayStars)
	}
}
