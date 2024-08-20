package service

type PaginatedBooks struct {
	Total        int    `json:"total"`
	PerPage      int    `json:"per_page"`
	CurrentPage  int    `json:"current_page"`
	LastPage     int    `json:"last_page"`
	FirstPageURL string `json:"first_page_url"`
	LastPageURL  string `json:"last_page_url"`
	NextPageURL  string `json:"next_page_url"`
	PrevPageURL  string `json:"prev_page_url"`
	From         int    `json:"from"`
	To           int    `json:"to"`
	Data         []Book `json:"data"`
}
