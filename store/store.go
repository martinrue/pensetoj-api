package store

// Summary holds full summary of data in the store.
type Summary struct {
	LikesTotal   int            `json:"likesTotal"`
	ListensTotal int            `json:"listensTotal"`
	Likes        map[string]int `json:"likes"`
}

// Store describes a generic store.
type Store interface {
	AddAction(t string, id string, ip string) bool
	GetSummary() *Summary
	Size() (string, error)
}
