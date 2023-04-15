package cloudflare

import (
	"math"
)

// Done returns true for the last page and false otherwise.
func (p ResultInfo) Done() bool {
	return p.Page > 1 && p.Page > p.TotalPages
}

// Next advances the page of a paginated API response, but does not fetch the
// next page of results.
func (p ResultInfo) Next() ResultInfo {
	p.Page++
	return p
}

// HasMorePages returns whether there is another page of results after the
// current one.
func (p ResultInfo) HasMorePages() bool {
	return p.Page > 1 && p.Page < p.TotalPages
}

// DoneCount returns true for the last page and false otherwise.
// This function uses the Total and PerPage fields to generate TotalPages
// then calls Done()
func (p ResultInfo) DoneCount() bool {
	p.TotalPages = int(math.Ceil(float64(p.Total) / float64(p.PerPage)))
	return p.Done()
}
