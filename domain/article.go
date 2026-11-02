package domain

// Article represents a news article from an RSS feed
type Article struct {
	ID          string
	Title       string
	Description string
	Link        string
	Published   string
	FeedTitle   string
}
