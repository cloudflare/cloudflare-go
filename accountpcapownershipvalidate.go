// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountPcapOwnershipValidateService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPcapOwnershipValidateService] method instead.
type AccountPcapOwnershipValidateService struct {
	Options []option.RequestOption
}

// NewAccountPcapOwnershipValidateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPcapOwnershipValidateService(opts ...option.RequestOption) (r *AccountPcapOwnershipValidateService) {
	r = &AccountPcapOwnershipValidateService{}
	r.Options = opts
	return
}

// Validates buckets added to the packet captures API.
func (r *AccountPcapOwnershipValidateService) MagicPcapCollectionValidateBucketsForFullPacketCaptures(ctx context.Context, accountIdentifier string, body AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams, opts ...option.RequestOption) (res *PcapsOwnershipSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
