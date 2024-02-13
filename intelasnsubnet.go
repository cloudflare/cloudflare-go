// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IntelAsnSubnetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelAsnSubnetService] method
// instead.
type IntelAsnSubnetService struct {
	Options []option.RequestOption
}

// NewIntelAsnSubnetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelAsnSubnetService(opts ...option.RequestOption) (r *IntelAsnSubnetService) {
	r = &IntelAsnSubnetService{}
	r.Options = opts
	return
}

// Get ASN Subnets
func (r *IntelAsnSubnetService) List(ctx context.Context, accountID string, asn int64, opts ...option.RequestOption) (res *IntelAsnSubnetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/asn/%v/subnets", accountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IntelAsnSubnetListResponse struct {
	Asn int64 `json:"asn"`
	// Total results returned based on your search parameters.
	Count        float64 `json:"count"`
	IPCountTotal int64   `json:"ip_count_total"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64                        `json:"per_page"`
	Subnets []string                       `json:"subnets"`
	JSON    intelAsnSubnetListResponseJSON `json:"-"`
}

// intelAsnSubnetListResponseJSON contains the JSON metadata for the struct
// [IntelAsnSubnetListResponse]
type intelAsnSubnetListResponseJSON struct {
	Asn          apijson.Field
	Count        apijson.Field
	IPCountTotal apijson.Field
	Page         apijson.Field
	PerPage      apijson.Field
	Subnets      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IntelAsnSubnetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
