package pagination

import (
	"database/sql"
	"math"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

type Pagination struct {
	con         *sql.DB
	Total       uint32 `json:"total"`
	PerPage     uint32 `json:"per_page"`
	CurrentPage uint32 `json:"current_page"`
	LastPage    uint32 `json:"last_page"`
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
	From        uint32 `json:"from"`
	To          uint32 `json:"to"`
}

func New(db *sql.DB) *Pagination {
	return &Pagination{
		con: db,
	}
}

func (p *Pagination) BuildPagination(query sq.SelectBuilder, limit, page uint64) (res Pagination, err error) {
	// total = get sum querySQL
	// per_page = get queryURL ? default
	// Current page = get page ? 1
	// Last page = total / per_page (ceil)
	// next_page_url = current_page + 1
	// prev_page_url = current_page != 1 ? current_page - 1 : null
	// from :
	//  -> if current_page == 1 ? 1
	//  -> if current_page == 2 ? per_page + 1
	//  -> if current_page > 2 ? per_page * (current_page - 1) + 1
	// to :
	//  -> if current_page == 1 ? per_page
	//  -? if current_page > 1 ? per_page * current_page
	var total uint32
	pagQuery := query
	pagQuery = query.Columns("count(*) as total_count")
	pagQuery = pagQuery.RemoveLimit()
	pagQuery = pagQuery.RemoveOffset()
	pagQuery = pagQuery.RunWith(p.con)
	err = pagQuery.QueryRow().Scan(&total)

	if err != nil {
		return Pagination{}, err
	}
	// Set Total
	res.Total = total
	// Set Total
	res.Total = uint32(total)
	// Set PerPage
	// if filter.limit is null per_page default is 10
	if res.PerPage = 10; limit != uint64(0) {
		res.PerPage = uint32(limit)
	}

	// Set Current Page
	// if filter more than 1 current page same with filter.page
	if res.CurrentPage = 1; page > 1 {
		res.CurrentPage = uint32(page)
	}

	// Set LastPage
	calcLastPage := float64(total) / float64(res.PerPage)
	// Ceil returns the least integer value greater than or equal to x.
	res.LastPage = uint32(math.Ceil(calcLastPage))

	// Set NextPage URL
	res.NextPageUrl = strconv.Itoa(int(res.CurrentPage + uint32(1)))

	// Set PrevPage URL
	if res.PrevPageUrl = "1"; res.CurrentPage > 1 {
		sum := res.CurrentPage - 1
		res.PrevPageUrl = strconv.Itoa(int(sum))
	}

	// Set From
	// If current page == 2 just add 1
	if res.From = 1; res.CurrentPage == 2 {
		res.From = res.PerPage + 1
	}
	// if current page > 2, limit * currentpage - 1 and + 1 for initial first number
	if res.CurrentPage > 2 {
		res.From = res.PerPage*(res.CurrentPage-1) + 1
	}

	// Set To
	// ex. Total 17
	// Last page is 2
	// Page 1 == 10
	// Page 2 == 7
	// if general
	if res.To = res.PerPage; res.CurrentPage > 1 {
		res.To = res.PerPage * res.CurrentPage
	}
	// for get difference if currentPage == lastPage -> total - (limit * (lastPage - 1))
	if res.CurrentPage == res.LastPage {
		difference := res.Total - (res.PerPage * (res.LastPage - 1))
		// and then sum (limit * lastPage - 1) + difference
		res.To = (res.PerPage * (res.LastPage - 1)) + difference
	}
	return res, nil
}
