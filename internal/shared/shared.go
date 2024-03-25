// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
)

type ErrorData struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	JSON    errorDataJSON `json:"-"`
}

// errorDataJSON contains the JSON metadata for the struct [ErrorData]
type errorDataJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorDataJSON) RawJSON() string {
	return r.raw
}
