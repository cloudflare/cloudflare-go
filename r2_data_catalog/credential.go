// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_data_catalog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// CredentialService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialService] method instead.
type CredentialService struct {
	Options []option.RequestOption
}

// NewCredentialService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCredentialService(opts ...option.RequestOption) (r *CredentialService) {
	r = &CredentialService{}
	r.Options = opts
	return
}

// Store authentication credentials for a catalog. These credentials are used to
// authenticate with R2 storage when performing catalog operations.
func (r *CredentialService) New(ctx context.Context, bucketName string, params CredentialNewParams, opts ...option.RequestOption) (res *CredentialNewResponse, err error) {
	var env CredentialNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2-catalog/%s/credential", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CredentialNewResponse = interface{}

type CredentialNewParams struct {
	// Use this to identify the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Provides the Cloudflare API token for accessing R2.
	Token param.Field[string] `json:"token,required"`
}

func (r CredentialNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CredentialNewResponseEnvelope struct {
	// Contains errors if the API call was unsuccessful.
	Errors []CredentialNewResponseEnvelopeErrors `json:"errors,required"`
	// Contains informational messages.
	Messages []CredentialNewResponseEnvelopeMessages `json:"messages,required"`
	// Indicates whether the API call was successful.
	Success bool                              `json:"success,required"`
	Result  CredentialNewResponse             `json:"result,nullable"`
	JSON    credentialNewResponseEnvelopeJSON `json:"-"`
}

// credentialNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CredentialNewResponseEnvelope]
type credentialNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CredentialNewResponseEnvelopeErrors struct {
	// Specifies the error code.
	Code int64 `json:"code,required"`
	// Describes the error.
	Message string                                  `json:"message,required"`
	JSON    credentialNewResponseEnvelopeErrorsJSON `json:"-"`
}

// credentialNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CredentialNewResponseEnvelopeErrors]
type credentialNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CredentialNewResponseEnvelopeMessages struct {
	// Specifies the message code.
	Code int64 `json:"code,required"`
	// Contains the message text.
	Message string                                    `json:"message,required"`
	JSON    credentialNewResponseEnvelopeMessagesJSON `json:"-"`
}

// credentialNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CredentialNewResponseEnvelopeMessages]
type credentialNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CredentialNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r credentialNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
