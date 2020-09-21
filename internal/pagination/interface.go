package pagination

import sq "github.com/Masterminds/squirrel"

type PaginationInterface interface {
	BuildPagination(query sq.SelectBuilder, limit, page uint64) (res Pagination, err error)
}
