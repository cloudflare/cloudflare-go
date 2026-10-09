// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package k2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// StreamSubscriptionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamSubscriptionService] method instead.
type StreamSubscriptionService struct {
	Options []option.RequestOption
}

// NewStreamSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStreamSubscriptionService(opts ...option.RequestOption) (r *StreamSubscriptionService) {
	r = &StreamSubscriptionService{}
	r.Options = opts
	return
}

// Lists every subscription on one stream, oldest first. Lag uses committed
// positions, not reserved read positions. A failed tail observation preserves
// subscription metadata and returns unavailable lag.
func (r *StreamSubscriptionService) List(ctx context.Context, streamID string, query StreamSubscriptionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[StreamSubscriptionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/k2/streams/%s/subscriptions", query.AccountID, streamID)
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

// Lists every subscription on one stream, oldest first. Lag uses committed
// positions, not reserved read positions. A failed tail observation preserves
// subscription metadata and returns unavailable lag.
func (r *StreamSubscriptionService) ListAutoPaging(ctx context.Context, streamID string, query StreamSubscriptionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[StreamSubscriptionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, streamID, query, opts...))
}

type StreamSubscriptionListResponse struct {
	ID         string                                `json:"id" api:"required"`
	CreatedAt  time.Time                             `json:"created_at" api:"required" format:"date-time"`
	Lag        StreamSubscriptionListResponseLag     `json:"lag" api:"required"`
	ModifiedAt time.Time                             `json:"modified_at" api:"required" format:"date-time"`
	Name       string                                `json:"name" api:"required"`
	StartAt    StreamSubscriptionListResponseStartAt `json:"start_at" api:"required"`
	JSON       streamSubscriptionListResponseJSON    `json:"-"`
}

// streamSubscriptionListResponseJSON contains the JSON metadata for the struct
// [StreamSubscriptionListResponse]
type streamSubscriptionListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Lag         apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	StartAt     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamSubscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type StreamSubscriptionListResponseLag struct {
	// Decimal-string distance from the subscription's committed position to the
	// observed exclusive stream tail, in records. Includes in-flight records and may
	// include expired records or records acknowledged beyond an earlier gap. Null
	// unless status is available.
	Records string `json:"records" api:"required,nullable"`
	// Unsupported means the subscription or stream topology cannot be measured;
	// unavailable means a required observation failed or was inconsistent.
	Status StreamSubscriptionListResponseLagStatus `json:"status" api:"required"`
	JSON   streamSubscriptionListResponseLagJSON   `json:"-"`
}

// streamSubscriptionListResponseLagJSON contains the JSON metadata for the struct
// [StreamSubscriptionListResponseLag]
type streamSubscriptionListResponseLagJSON struct {
	Records     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamSubscriptionListResponseLag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamSubscriptionListResponseLagJSON) RawJSON() string {
	return r.raw
}

// Unsupported means the subscription or stream topology cannot be measured;
// unavailable means a required observation failed or was inconsistent.
type StreamSubscriptionListResponseLagStatus string

const (
	StreamSubscriptionListResponseLagStatusAvailable   StreamSubscriptionListResponseLagStatus = "available"
	StreamSubscriptionListResponseLagStatusUnsupported StreamSubscriptionListResponseLagStatus = "unsupported"
	StreamSubscriptionListResponseLagStatusUnavailable StreamSubscriptionListResponseLagStatus = "unavailable"
)

func (r StreamSubscriptionListResponseLagStatus) IsKnown() bool {
	switch r {
	case StreamSubscriptionListResponseLagStatusAvailable, StreamSubscriptionListResponseLagStatusUnsupported, StreamSubscriptionListResponseLagStatusUnavailable:
		return true
	}
	return false
}

type StreamSubscriptionListResponseStartAt struct {
	Type StreamSubscriptionListResponseStartAtType `json:"type" api:"required"`
	JSON streamSubscriptionListResponseStartAtJSON `json:"-"`
}

// streamSubscriptionListResponseStartAtJSON contains the JSON metadata for the
// struct [StreamSubscriptionListResponseStartAt]
type streamSubscriptionListResponseStartAtJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamSubscriptionListResponseStartAt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamSubscriptionListResponseStartAtJSON) RawJSON() string {
	return r.raw
}

type StreamSubscriptionListResponseStartAtType string

const (
	StreamSubscriptionListResponseStartAtTypeEarliest StreamSubscriptionListResponseStartAtType = "earliest"
	StreamSubscriptionListResponseStartAtTypeLatest   StreamSubscriptionListResponseStartAtType = "latest"
)

func (r StreamSubscriptionListResponseStartAtType) IsKnown() bool {
	switch r {
	case StreamSubscriptionListResponseStartAtTypeEarliest, StreamSubscriptionListResponseStartAtTypeLatest:
		return true
	}
	return false
}

type StreamSubscriptionListParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
