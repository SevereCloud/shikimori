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
