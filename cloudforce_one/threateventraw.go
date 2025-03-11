// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ThreatEventRawService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventRawService] method instead.
type ThreatEventRawService struct {
	Options []option.RequestOption
}

// NewThreatEventRawService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventRawService(opts ...option.RequestOption) (r *ThreatEventRawService) {
	r = &ThreatEventRawService{}
	r.Options = opts
	return
}

// Updates a raw event
func (r *ThreatEventRawService) Update(ctx context.Context, accountID float64, eventID string, rawID string, body ThreatEventRawUpdateParams, opts ...option.RequestOption) (res *ThreatEventRawUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	if rawID == "" {
		err = errors.New("missing required rawId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/raw/%s", accountID, eventID, rawID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Reads data for a raw event
func (r *ThreatEventRawService) Get(ctx context.Context, accountID float64, eventID string, rawID string, opts ...option.RequestOption) (res *ThreatEventRawGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	if rawID == "" {
		err = errors.New("missing required rawId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%v/cloudforce-one/events/%s/raw/%s", accountID, eventID, rawID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventRawUpdateResponse struct {
	ID   string                           `json:"id,required"`
	Data interface{}                      `json:"data,required"`
	JSON threatEventRawUpdateResponseJSON `json:"-"`
}

// threatEventRawUpdateResponseJSON contains the JSON metadata for the struct
// [ThreatEventRawUpdateResponse]
type threatEventRawUpdateResponseJSON struct {
	ID          apijson.Field
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRawUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRawUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRawGetResponse struct {
	ID        string                        `json:"id,required"`
	AccountID float64                       `json:"accountId,required"`
	Created   string                        `json:"created,required"`
	Data      interface{}                   `json:"data,required"`
	Source    string                        `json:"source,required"`
	TLP       string                        `json:"tlp,required"`
	JSON      threatEventRawGetResponseJSON `json:"-"`
}

// threatEventRawGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventRawGetResponse]
type threatEventRawGetResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Created     apijson.Field
	Data        apijson.Field
	Source      apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventRawGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRawGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRawUpdateParams struct {
	Data   param.Field[interface{}] `json:"data"`
	Source param.Field[string]      `json:"source"`
	TLP    param.Field[string]      `json:"tlp"`
}

func (r ThreatEventRawUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
