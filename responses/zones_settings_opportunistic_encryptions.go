package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type OpportunisticEncryptionListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result OpportunisticEncryption `json:"result"`
	JSON   OpportunisticEncryptionListResponseJSON
}

type OpportunisticEncryptionListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OpportunisticEncryptionListResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *OpportunisticEncryptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OpportunisticEncryptionUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result OpportunisticEncryption `json:"result"`
	JSON   OpportunisticEncryptionUpdateResponseJSON
}

type OpportunisticEncryptionUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OpportunisticEncryptionUpdateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *OpportunisticEncryptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
