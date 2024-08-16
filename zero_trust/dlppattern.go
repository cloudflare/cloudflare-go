// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DLPPatternService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPPatternService] method instead.
type DLPPatternService struct {
	Options []option.RequestOption
}

// NewDLPPatternService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPPatternService(opts ...option.RequestOption) (r *DLPPatternService) {
	r = &DLPPatternService{}
	r.Options = opts
	return
}

// Validates whether this pattern is a valid regular expression. Rejects it if the
// regular expression is too complex or can match an unbounded-length string. The
// regex will be rejected if it uses `*` or `+`. Bound the maximum number of
// characters that can be matched using a range, e.g. `{1,100}`.
func (r *DLPPatternService) Validate(ctx context.Context, params DLPPatternValidateParams, opts ...option.RequestOption) (res *DLPPatternValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/patterns/validate", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type DLPPatternValidateResponse struct {
	Valid bool                           `json:"valid,required"`
	JSON  dlpPatternValidateResponseJSON `json:"-"`
}

// dlpPatternValidateResponseJSON contains the JSON metadata for the struct
// [DLPPatternValidateResponse]
type dlpPatternValidateResponseJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPatternValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPatternValidateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPPatternValidateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Regex     param.Field[string] `json:"regex,required"`
}

func (r DLPPatternValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
