package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type SslRecommenderListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result fields.Field[SslRecommender] `json:"result"`
}

// MarshalJSON serializes SslRecommenderListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SslRecommenderListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SslRecommenderListResponse) String() (result string) {
	return fmt.Sprintf("&SslRecommenderListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SslRecommenderUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Result fields.Field[SslRecommender] `json:"result"`
}

// MarshalJSON serializes SslRecommenderUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SslRecommenderUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SslRecommenderUpdateResponse) String() (result string) {
	return fmt.Sprintf("&SslRecommenderUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SslRecommenderUpdateParams struct {
	// Enrollment in the SSL/TLS Recommender service which tries to detect and
	// recommend (by sending periodic emails) the most secure SSL/TLS setting your
	// origin servers support.
	Value fields.Field[SslRecommender] `json:"value,required"`
}

// MarshalJSON serializes SslRecommenderUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SslRecommenderUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SslRecommenderUpdateParams) String() (result string) {
	return fmt.Sprintf("&SslRecommenderUpdateParams{Value:%s}", r.Value)
}
