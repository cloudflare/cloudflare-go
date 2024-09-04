// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
  "bytes"
  "context"
  "errors"
  "fmt"
  "mime/multipart"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v2/internal/apiform"
  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/shared"
)

// IndicatorFeedSnapshotService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndicatorFeedSnapshotService] method instead.
type IndicatorFeedSnapshotService struct {
Options []option.RequestOption
}

// NewIndicatorFeedSnapshotService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIndicatorFeedSnapshotService(opts ...option.RequestOption) (r *IndicatorFeedSnapshotService) {
  r = &IndicatorFeedSnapshotService{}
  r.Options = opts
  return
}

// Update indicator feed data
func (r *IndicatorFeedSnapshotService) Update(ctx context.Context, feedID int64, params IndicatorFeedSnapshotUpdateParams, opts ...option.RequestOption) (res *IndicatorFeedSnapshotUpdateResponse, err error) {
  var env IndicatorFeedSnapshotUpdateResponseEnvelope
  opts = append(r.Options[:], opts...)
  if params.AccountID.Value == "" {
    err = errors.New("missing required account_id parameter")
    return
  }
  path := fmt.Sprintf("accounts/%s/intel/indicator-feeds/%v/snapshot", params.AccountID, feedID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type IndicatorFeedSnapshotUpdateResponse struct {
// Feed id
FileID int64 `json:"file_id"`
// Name of the file unified in our system
Filename string `json:"filename"`
// Current status of upload, should be unified
Status string `json:"status"`
JSON indicatorFeedSnapshotUpdateResponseJSON `json:"-"`
}

// indicatorFeedSnapshotUpdateResponseJSON contains the JSON metadata for the
// struct [IndicatorFeedSnapshotUpdateResponse]
type indicatorFeedSnapshotUpdateResponseJSON struct {
FileID apijson.Field
Filename apijson.Field
Status apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedSnapshotUpdateResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedSnapshotUpdateResponseJSON) RawJSON() (string) {
  return r.raw
}

type IndicatorFeedSnapshotUpdateParams struct {
// Identifier
AccountID param.Field[string] `path:"account_id,required"`
// The file to upload
Source param.Field[string] `json:"source"`
}

func (r IndicatorFeedSnapshotUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
  buf := bytes.NewBuffer(nil)
  writer := multipart.NewWriter(buf)
  err = apiform.MarshalRoot(r, writer)
  if err != nil {
    writer.Close()
    return nil, "", err
  }
  err = writer.Close()
  if err != nil {
    return nil, "", err
  }
  return buf.Bytes(), writer.FormDataContentType(), nil
}

type IndicatorFeedSnapshotUpdateResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// Whether the API call was successful
Success IndicatorFeedSnapshotUpdateResponseEnvelopeSuccess `json:"success,required"`
Result IndicatorFeedSnapshotUpdateResponse `json:"result"`
JSON indicatorFeedSnapshotUpdateResponseEnvelopeJSON `json:"-"`
}

// indicatorFeedSnapshotUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [IndicatorFeedSnapshotUpdateResponseEnvelope]
type indicatorFeedSnapshotUpdateResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Success apijson.Field
Result apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedSnapshotUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedSnapshotUpdateResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type IndicatorFeedSnapshotUpdateResponseEnvelopeSuccess bool

const (
  IndicatorFeedSnapshotUpdateResponseEnvelopeSuccessTrue IndicatorFeedSnapshotUpdateResponseEnvelopeSuccess = true
)

func (r IndicatorFeedSnapshotUpdateResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case IndicatorFeedSnapshotUpdateResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}
