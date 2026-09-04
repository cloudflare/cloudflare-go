// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// ProtocolService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProtocolService] method instead.
type ProtocolService struct {
	Options []option.RequestOption
}

// NewProtocolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProtocolService(opts ...option.RequestOption) (r *ProtocolService) {
	r = &ProtocolService{}
	r.Options = opts
	return
}

// Retrieves a list of Spectrum application protocols available for a zone.
func (r *ProtocolService) List(ctx context.Context, query ProtocolListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProtocolListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/spectrum/protocols", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Retrieves a list of Spectrum application protocols available for a zone.
func (r *ProtocolService) ListAutoPaging(ctx context.Context, query ProtocolListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProtocolListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type ProtocolListResponse struct {
	// The full name of the application protocol.
	Description string `json:"description" api:"required"`
	// The short name of the application protocol.
	Name string `json:"name" api:"required"`
	// The available listening ports for the given protocol.
	Ports []int64 `json:"ports" api:"required"`
	// The transport layer protocol used by the application protocol
	Transport string                   `json:"transport" api:"required"`
	JSON      protocolListResponseJSON `json:"-"`
}

// protocolListResponseJSON contains the JSON metadata for the struct
// [ProtocolListResponse]
type protocolListResponseJSON struct {
	Description apijson.Field
	Name        apijson.Field
	Ports       apijson.Field
	Transport   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r protocolListResponseJSON) RawJSON() string {
	return r.raw
}

type ProtocolListParams struct {
	// Zone identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}
