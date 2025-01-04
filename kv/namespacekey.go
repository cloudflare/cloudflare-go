// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
)

// NamespaceKeyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceKeyService] method instead.
type NamespaceKeyService struct {
	Options []option.RequestOption
}

// NewNamespaceKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceKeyService(opts ...option.RequestOption) (r *NamespaceKeyService) {
	r = &NamespaceKeyService{}
	r.Options = opts
	return
}

// Lists a namespace's keys.
func (r *NamespaceKeyService) List(ctx context.Context, namespaceID string, params NamespaceKeyListParams, opts ...option.RequestOption) (res *pagination.CursorLimitPagination[Key], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/keys", params.AccountID, namespaceID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists a namespace's keys.
func (r *NamespaceKeyService) ListAutoPaging(ctx context.Context, namespaceID string, params NamespaceKeyListParams, opts ...option.RequestOption) *pagination.CursorLimitPaginationAutoPager[Key] {
	return pagination.NewCursorLimitPaginationAutoPager(r.List(ctx, namespaceID, params, opts...))
}

// A name for a value. A value stored under a given key may be retrieved via the
// same key.
type Key struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64 `json:"expiration"`
	// Arbitrary JSON that is associated with a key.
	Metadata map[string]interface{} `json:"metadata"`
	JSON     keyJSON                `json:"-"`
}

// keyJSON contains the JSON metadata for the struct [Key]
type keyJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Key) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyJSON) RawJSON() string {
	return r.raw
}

type NamespaceKeyListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the `cursors`
	// object in the `result_info` structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of keys to return. The cursor attribute may be used to iterate over
	// the next batch of keys if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
	// A string prefix used to filter down which keys will be returned. Exact matches
	// and any key names that begin with the prefix will be returned.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [NamespaceKeyListParams]'s query parameters as `url.Values`.
func (r NamespaceKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
