// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventGraphqlService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventGraphqlService] method instead.
type ThreatEventGraphqlService struct {
	Options []option.RequestOption
}

// NewThreatEventGraphqlService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventGraphqlService(opts ...option.RequestOption) (r *ThreatEventGraphqlService) {
	r = &ThreatEventGraphqlService{}
	r.Options = opts
	return
}

// Execute GraphQL aggregations over threat events. Supports multi-dimensional
// group-bys, optional date range filtering, and multi-dataset aggregation.
func (r *ThreatEventGraphqlService) New(ctx context.Context, body ThreatEventGraphqlNewParams, opts ...option.RequestOption) (res *ThreatEventGraphqlNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/graphql", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ThreatEventGraphqlNewResponse struct {
	Data   interface{}                       `json:"data" api:"nullable"`
	Errors []interface{}                     `json:"errors" api:"nullable"`
	JSON   threatEventGraphqlNewResponseJSON `json:"-"`
}

// threatEventGraphqlNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventGraphqlNewResponse]
type threatEventGraphqlNewResponseJSON struct {
	Data        apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventGraphqlNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventGraphqlNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventGraphqlNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
