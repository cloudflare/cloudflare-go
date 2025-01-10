// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// CaptionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCaptionService] method instead.
type CaptionService struct {
	Options  []option.RequestOption
	Language *CaptionLanguageService
}

// NewCaptionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCaptionService(opts ...option.RequestOption) (r *CaptionService) {
	r = &CaptionService{}
	r.Options = opts
	r.Language = NewCaptionLanguageService(opts...)
	return
}

// Lists the available captions or subtitles for a specific video.
func (r *CaptionService) Get(ctx context.Context, identifier string, query CaptionGetParams, opts ...option.RequestOption) (res *[]Caption, err error) {
	var env CaptionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Caption struct {
	// Whether the caption was generated via AI.
	Generated bool `json:"generated"`
	// The language label displayed in the native language to users.
	Label string `json:"label"`
	// The language tag in BCP 47 format.
	Language string `json:"language"`
	// The status of a generated caption.
	Status CaptionStatus `json:"status"`
	JSON   captionJSON   `json:"-"`
}

// captionJSON contains the JSON metadata for the struct [Caption]
type captionJSON struct {
	Generated   apijson.Field
	Label       apijson.Field
	Language    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Caption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionJSON) RawJSON() string {
	return r.raw
}

// The status of a generated caption.
type CaptionStatus string

const (
	CaptionStatusReady      CaptionStatus = "ready"
	CaptionStatusInprogress CaptionStatus = "inprogress"
	CaptionStatusError      CaptionStatus = "error"
)

func (r CaptionStatus) IsKnown() bool {
	switch r {
	case CaptionStatusReady, CaptionStatusInprogress, CaptionStatusError:
		return true
	}
	return false
}

type CaptionGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type CaptionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success CaptionGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []Caption                         `json:"result"`
	JSON    captionGetResponseEnvelopeJSON    `json:"-"`
}

// captionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CaptionGetResponseEnvelope]
type captionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CaptionGetResponseEnvelopeSuccess bool

const (
	CaptionGetResponseEnvelopeSuccessTrue CaptionGetResponseEnvelopeSuccess = true
)

func (r CaptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
