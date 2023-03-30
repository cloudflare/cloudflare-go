package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type SslRecommenderListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result SslRecommender `json:"result"`
	JSON   SslRecommenderListResponseJSON
}

type SslRecommenderListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SslRecommenderListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *SslRecommenderListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SslRecommenderUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result SslRecommender `json:"result"`
	JSON   SslRecommenderUpdateResponseJSON
}

type SslRecommenderUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SslRecommenderUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *SslRecommenderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
