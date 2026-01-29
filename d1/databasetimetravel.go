// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// DatabaseTimeTravelService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabaseTimeTravelService] method instead.
type DatabaseTimeTravelService struct {
	Options []option.RequestOption
}

// NewDatabaseTimeTravelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDatabaseTimeTravelService(opts ...option.RequestOption) (r *DatabaseTimeTravelService) {
	r = &DatabaseTimeTravelService{}
	r.Options = opts
	return
}

// Retrieves the current bookmark, or the nearest bookmark at or before a provided
// timestamp. Bookmarks can be used with the restore endpoint to revert the
// database to a previous point in time.
func (r *DatabaseTimeTravelService) GetBookmark(ctx context.Context, databaseID string, params DatabaseTimeTravelGetBookmarkParams, opts ...option.RequestOption) (res *DatabaseTimeTravelGetBookmarkResponse, err error) {
	var env DatabaseTimeTravelGetBookmarkResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/time_travel/bookmark", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Restores a D1 database to a previous point in time either via a bookmark or a
// timestamp.
func (r *DatabaseTimeTravelService) Restore(ctx context.Context, databaseID string, params DatabaseTimeTravelRestoreParams, opts ...option.RequestOption) (res *DatabaseTimeTravelRestoreResponse, err error) {
	var env DatabaseTimeTravelRestoreResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/time_travel/restore", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DatabaseTimeTravelGetBookmarkResponse struct {
	// A bookmark representing a specific state of the database at a specific point in
	// time.
	Bookmark string                                    `json:"bookmark"`
	JSON     databaseTimeTravelGetBookmarkResponseJSON `json:"-"`
}

// databaseTimeTravelGetBookmarkResponseJSON contains the JSON metadata for the
// struct [DatabaseTimeTravelGetBookmarkResponse]
type databaseTimeTravelGetBookmarkResponseJSON struct {
	Bookmark    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseTimeTravelGetBookmarkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseTimeTravelGetBookmarkResponseJSON) RawJSON() string {
	return r.raw
}

// Response from a time travel restore operation.
type DatabaseTimeTravelRestoreResponse struct {
	// The new bookmark representing the state of the database after the restore
	// operation.
	Bookmark string `json:"bookmark"`
	// A message describing the result of the restore operation.
	Message string `json:"message"`
	// The bookmark representing the state of the database before the restore
	// operation. Can be used to undo the restore if needed.
	PreviousBookmark string                                `json:"previous_bookmark"`
	JSON             databaseTimeTravelRestoreResponseJSON `json:"-"`
}

// databaseTimeTravelRestoreResponseJSON contains the JSON metadata for the struct
// [DatabaseTimeTravelRestoreResponse]
type databaseTimeTravelRestoreResponseJSON struct {
	Bookmark         apijson.Field
	Message          apijson.Field
	PreviousBookmark apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DatabaseTimeTravelRestoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseTimeTravelRestoreResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseTimeTravelGetBookmarkParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional ISO 8601 timestamp. If provided, returns the nearest available
	// bookmark at or before this timestamp. If omitted, returns the current bookmark.
	Timestamp param.Field[time.Time] `query:"timestamp" format:"date-time"`
}

// URLQuery serializes [DatabaseTimeTravelGetBookmarkParams]'s query parameters as
// `url.Values`.
func (r DatabaseTimeTravelGetBookmarkParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DatabaseTimeTravelGetBookmarkResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []shared.ResponseInfo                 `json:"messages,required"`
	Result   DatabaseTimeTravelGetBookmarkResponse `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseTimeTravelGetBookmarkResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseTimeTravelGetBookmarkResponseEnvelopeJSON    `json:"-"`
}

// databaseTimeTravelGetBookmarkResponseEnvelopeJSON contains the JSON metadata for
// the struct [DatabaseTimeTravelGetBookmarkResponseEnvelope]
type databaseTimeTravelGetBookmarkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseTimeTravelGetBookmarkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseTimeTravelGetBookmarkResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseTimeTravelGetBookmarkResponseEnvelopeSuccess bool

const (
	DatabaseTimeTravelGetBookmarkResponseEnvelopeSuccessTrue DatabaseTimeTravelGetBookmarkResponseEnvelopeSuccess = true
)

func (r DatabaseTimeTravelGetBookmarkResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseTimeTravelGetBookmarkResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseTimeTravelRestoreParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A bookmark to restore the database to. Required if `timestamp` is not provided.
	Bookmark param.Field[string] `query:"bookmark"`
	// An ISO 8601 timestamp to restore the database to. Required if `bookmark` is not
	// provided.
	Timestamp param.Field[time.Time] `query:"timestamp" format:"date-time"`
}

// URLQuery serializes [DatabaseTimeTravelRestoreParams]'s query parameters as
// `url.Values`.
func (r DatabaseTimeTravelRestoreParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DatabaseTimeTravelRestoreResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Response from a time travel restore operation.
	Result DatabaseTimeTravelRestoreResponse `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseTimeTravelRestoreResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseTimeTravelRestoreResponseEnvelopeJSON    `json:"-"`
}

// databaseTimeTravelRestoreResponseEnvelopeJSON contains the JSON metadata for the
// struct [DatabaseTimeTravelRestoreResponseEnvelope]
type databaseTimeTravelRestoreResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseTimeTravelRestoreResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseTimeTravelRestoreResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseTimeTravelRestoreResponseEnvelopeSuccess bool

const (
	DatabaseTimeTravelRestoreResponseEnvelopeSuccessTrue DatabaseTimeTravelRestoreResponseEnvelopeSuccess = true
)

func (r DatabaseTimeTravelRestoreResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseTimeTravelRestoreResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
