package cloudflare

// Done returns true for the last page and false otherwise.
func (p ResultInfo) Done() bool {
	// A little hacky but if the response body is lacking a defined `ResultInfo`
	// object the page will be 1 however the counts will be empty so if we have
	// that response, we just assume this is the only page.
	if p.Page == 1 && p.TotalPages == 0 {
		return true
	}

	return p.Page > 1 && p.Page > p.TotalPages
}

// Next advances the page of a paginated API response, but does not fetch the
// next page of results.
func (p ResultInfo) Next() ResultInfo {
	// A little hacky but if the response body is lacking a defined `ResultInfo`
	// object the page will be 1 however the counts will be empty so if we have
	// that response, we just assume this is the only page.
	if p.Page == 1 && p.TotalPages == 0 {
		return p
	}

	// This shouldn't happen normally however, when it does just return the
	// current page.
	if p.Page > p.TotalPages {
		return p
	}

	p.Page++
	return p
}

// HasMorePages returns whether there is another page of results after the
// current one.
func (p ResultInfo) HasMorePages() bool {
	if p.TotalPages == 0 {
		return false
	}

	return p.Page >= 1 && p.Page < p.TotalPages
}
