package pagination

type Pagination struct {
	Total       uint32 `json:"total"`
	PerPage     uint32 `json:"per_page"`
	CurrentPage uint32 `json:"current_page"`
	LastPage    uint32 `json:"last_page"`
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
	From        uint32 `json:"from"`
	To          uint32 `json:"to"`
}
