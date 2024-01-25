// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
)

type PageResultInfo struct {
	Count      interface{}        `json:"count"`
	Page       interface{}        `json:"page"`
	PerPage    interface{}        `json:"per_page"`
	TotalCount interface{}        `json:"total_count"`
	JSON       pageResultInfoJSON `json:"-"`
}

// pageResultInfoJSON contains the JSON metadata for the struct [PageResultInfo]
type pageResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type Page[T any] struct {
	Result     []T            `json:"result"`
	ResultInfo PageResultInfo `json:"result_info"`
	JSON       pageJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// pageJSON contains the JSON metadata for the struct [Page[T]]
type pageJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Page[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *Page[T]) GetNextPage() (res *Page[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Page[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type PageAutoPager[T any] struct {
	page *Page[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageAutoPager[T any](page *Page[T], err error) *PageAutoPager[T] {
	return &PageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result) == 0 {
			return false
		}
	}
	r.cur = r.page.Result[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageAutoPager[T]) Index() int {
	return r.run
}
