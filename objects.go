package shikimori

import "time"

type Image struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X96      string `json:"x96"`
	X48      string `json:"x48"`
}

type ExternalLink struct {
	ID         *int       `json:"id,omitempty"`
	Kind       string     `json:"kind"`
	URL        string     `json:"url"`
	Source     string     `json:"source"`
	EntryID    int        `json:"entry_id"`
	EntryType  string     `json:"entry_type"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	ImportedAt *time.Time `json:"imported_at,omitempty"`
}

type FranchiseLink struct {
	ID       int    `json:"id"`
	SourceID int    `json:"source_id"`
	TargetID int    `json:"target_id"`
	Source   int    `json:"source"`
	Target   int    `json:"target"`
	Weight   int    `json:"weight"`
	Relation string `json:"relation"`
}

type FranchiseNode struct {
	ID       int     `json:"id"`
	Date     int     `json:"date"`
	Name     string  `json:"name"`
	ImageURL string  `json:"image_url"`
	URL      string  `json:"url"`
	Year     *int    `json:"year"`
	Kind     *string `json:"kind"`
	Weight   int     `json:"weight"`
}

type Franchise struct {
	Links     []FranchiseLink `json:"links"`
	Nodes     []FranchiseNode `json:"nodes"`
	CurrentID int             `json:"current_id"`
}
