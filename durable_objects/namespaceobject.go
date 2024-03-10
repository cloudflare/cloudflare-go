// File generated from our OpenAPI spec by Stainless.

package durable_objects

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// NamespaceObjectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNamespaceObjectService] method
// instead.
type NamespaceObjectService struct {
	Options []option.RequestOption
}

// NewNamespaceObjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceObjectService(opts ...option.RequestOption) (r *NamespaceObjectService) {
	r = &NamespaceObjectService{}
	r.Options = opts
	return
}

// Returns the Durable Objects in a given namespace.
func (r *NamespaceObjectService) List(ctx context.Context, id string, params NamespaceObjectListParams, opts ...option.RequestOption) (res *shared.CursorLimitPagination[WorkersObject], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces/%s/objects", params.AccountID, id)
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

// Returns the Durable Objects in a given namespace.
func (r *NamespaceObjectService) ListAutoPaging(ctx context.Context, id string, params NamespaceObjectListParams, opts ...option.RequestOption) *shared.CursorLimitPaginationAutoPager[WorkersObject] {
	return shared.NewCursorLimitPaginationAutoPager(r.List(ctx, id, params, opts...))
}

type WorkersObject struct {
	// ID of the Durable Object.
	ID string `json:"id"`
	// Whether the Durable Object has stored data.
	HasStoredData bool              `json:"hasStoredData"`
	JSON          workersObjectJSON `json:"-"`
}

// workersObjectJSON contains the JSON metadata for the struct [WorkersObject]
type workersObjectJSON struct {
	ID            apijson.Field
	HasStoredData apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersObjectJSON) RawJSON() string {
	return r.raw
}

type NamespaceObjectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of objects to return. The cursor attribute may be used to iterate
	// over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [NamespaceObjectListParams]'s query parameters as
// `url.Values`.
func (r NamespaceObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
