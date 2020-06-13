package entity

type Discussion struct {
	ID        uint32 `json:"id"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
}
