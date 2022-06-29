package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

var (
	// ErrMissingUploadURL is for when a URL is required but missing
	ErrMissingUploadURL = errors.New("required url missing")
	// ErrMissingMaxDuration is for when MaxDuration is required but missing
	ErrMissingMaxDuration = errors.New("required max duration missing")
)

// StreamVideo represents a stream video
type StreamVideo struct {
	AllowedOrigins        []string               `json:"allowedOrigins,omitempty"`
	Created               *time.Time             `json:"created,omitempty"`
	Duration              int                    `json:"duration,omitempty"`
	Input                 StreamVideoInput       `json:"input,omitempty"`
	MaxDurationSeconds    int                    `json:"maxDurationSeconds,omitempty"`
	Meta                  map[string]interface{} `json:"meta,omitempty"`
	Modified              *time.Time             `json:"modified,omitempty"`
	UploadExpiry          *time.Time             `json:"uploadExpiry,omitempty"`
	Playback              StreamVideoPlayback    `json:"playback,omitempty"`
	Preview               string                 `json:"preview,omitempty"`
	ReadyToStream         bool                   `json:"readyToStream,omitempty"`
	RequireSignedURLs     bool                   `json:"requireSignedURLs,omitempty"`
	Size                  int                    `json:"size,omitempty"`
	Status                StreamVideoStatus      `json:"status,omitempty"`
	Thumbnail             string                 `json:"thumbnail,omitempty"`
	ThumbnailTimestampPct float64                `json:"thumbnailTimestampPct,omitempty"`
	UID                   string                 `json:"uid,omitempty"`
	Creator               string                 `json:"creator,omitempty"`
	LiveInput             string                 `json:"liveInput,omitempty"`
	Uploaded              *time.Time             `json:"uploaded,omitempty"`
	Watermark             StreamVideoWatermark   `json:"watermark,omitempty"`
	NFT                   StreamVideoNFT         `json:"nft,omitempty"`
}

// StreamVideoInput represents the input parameters of a stream video
type StreamVideoInput struct {
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}

// StreamVideoPlayback represents the playback URLs for a video
type StreamVideoPlayback struct {
	HLS  string `json:"hls,omitempty"`
	Dash string `json:"dash,omitempty"`
}

// StreamVideoStatus represents the status of a stream video
type StreamVideoStatus struct {
	State           string `json:"state,omitempty"`
	PctComplete     int    `json:"pctComplete,omitempty"`
	ErrorReasonCode string `json:"errorReasonCode,omitempty"`
	ErrorReasonText string `json:"errorReasonText,omitempty"`
}

// StreamVideoWatermark represents a watermark for a stream video
type StreamVideoWatermark struct {
	UID            string     `json:"uid,omitempty"`
	Size           int        `json:"size,omitempty"`
	Height         int        `json:"height,omitempty"`
	Width          int        `json:"width,omitempty"`
	Created        *time.Time `json:"created,omitempty"`
	DownloadedFrom string     `json:"downloadedFrom,omitempty"`
	Name           string     `json:"name,omitempty"`
	Opacity        float64    `json:"opacity,omitempty"`
	Padding        float64    `json:"padding,omitempty"`
	Scale          float64    `json:"scale,omitempty"`
	Position       string     `json:"position,omitempty"`
}

// StreamVideoNFT represents a NFT for a stream video
type StreamVideoNFT struct {
	Contract string `json:"contract,omitempty"`
	Token    int    `json:"token,omitempty"`
}

// UploadVideoOptions are base upload options used by both upload from URL and create video
type UploadVideoOptions struct {
	Creator               string                  `json:"creator,omitempty"`
	ThumbnailTimestampPct float64                 `json:"thumbnailTimestampPct,omitempty"`
	AllowedOrigins        []string                `json:"allowedOrigins,omitempty"`
	RequiredSignedURLs    bool                    `json:"requiredSignedURLs,omitempty"`
	Watermark             UploadVideoURLWatermark `json:"watermark,omitempty"`
}

// StreamUploadFromURLOptions are the parameters used when uploading a video from URL
type StreamUploadFromURLOptions struct {
	URL string `json:"url"`
	UploadVideoOptions
}

// StreamCreateVideoOptions are parameters used when creating a video
type StreamCreateVideoOptions struct {
	MaxDurationSeconds int        `json:"maxDurationSeconds,omitempty"`
	Expiry             *time.Time `json:"expiry,omitempty"`
	UploadVideoOptions
}

// UploadVideoURLWatermark represents UID of an existing watermark
type UploadVideoURLWatermark struct {
	UID string `json:"uid,omitempty"`
}

// StreamVideoCreate represents parameters returned after creating a video
type StreamVideoCreate struct {
	UploadURL string               `json:"uploadURL,omitempty"`
	UID       string               `json:"uid,omitempty"`
	Watermark StreamVideoWatermark `json:"watermark,omitempty"`
}

// StreamListOptions represents parameters used when listing stream videos
type StreamListOptions struct {
	After         *time.Time `url:"after,omitempty"`
	Before        *time.Time `url:"before,omitempty"`
	Creator       string     `url:"creator,omitempty"`
	IncludeCounts bool       `url:"include_counts,omitempty"`
	Search        string     `url:"search,omitempty"`
	Limit         int        `url:"limit,omitempty"`
	Asc           bool       `url:"asc,omitempty"`
	Status        string     `url:"status,omitempty"`
}

// StreamSignedURLOptions represent parameters used when creating a signed URL
type StreamSignedURLOptions struct {
	ID           string            `json:"id,omitempty"`
	PEM          string            `json:"pem,omitempty"`
	EXP          int               `json:"exp,omitempty"`
	NBF          int               `json:"nbf,omitempty"`
	Downloadable bool              `json:"downloadable,omitempty"`
	AccessRules  StreamAccessRules `json:"accessRules,omitempty"`
}

// StreamVideoResponse represents an API response of a stream video
type StreamVideoResponse struct {
	Response
	Result StreamVideo `json:"result,omitempty"`
}

// StreamVideoCreateResponse represents an API response of creating a stream video
type StreamVideoCreateResponse struct {
	Response
	Result StreamVideoCreate `json:"result,omitempty"`
}

// StreamListResponse represents the API response from a StreamListRequest
type StreamListResponse struct {
	Response
	Result []StreamVideo `json:"result,omitempty"`
	Total  string        `json:"total,omitempty"`
	Range  string        `json:"range,omitempty"`
}

// StreamSignedURLResponse represents an API response for a signed URL
type StreamSignedURLResponse struct {
	Response
	Result struct {
		Token string `json:"token,omitempty"`
	}
}

// StreamAccessRules represents the accessRules when creating a signed URL
//I feel like should be implemented somewhere already but couldn't find it
// accessRules under https://api.cloudflare.com/#stream-videos-create-a-signed-url-token-for-a-video
type StreamAccessRules []map[string]interface{}

// StreamUploadFromURL send a video URL to it will be downloaded and made available on Stream.
//
// API Reference: https://api.cloudflare.com/#stream-videos-upload-a-video-from-a-url
func (api *API) StreamUploadFromURL(ctx context.Context, accountID string, options StreamUploadFromURLOptions) (StreamVideo, error) {
	if options.URL == "" {
		return StreamVideo{}, ErrMissingUploadURL
	}
	uri := fmt.Sprintf("/accounts/%s/stream/copy", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, options)
	if err != nil {
		return StreamVideo{}, err
	}
	var r StreamVideoResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return StreamVideo{}, err
	}
	return r.Result, nil
}

// StreamUploadVideoFile uploads a video from a path to the file
//
// API Reference: https://api.cloudflare.com/#stream-videos-upload-a-video-using-a-single-http-request
func (api *API) StreamUploadVideoFile(ctx context.Context, accountID string, filePath string) (StreamVideo, error) {
	uri := fmt.Sprintf("/accounts/%s/stream", accountID)

	// Create new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return StreamVideo{}, err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return StreamVideo{}, err
	}
	if _, err := io.Copy(formFile, file); err != nil {
		return StreamVideo{}, err
	}
	if err := writer.Close(); err != nil {
		return StreamVideo{}, err
	}
	res, err := api.makeRequestWithAuthTypeAndHeaders(ctx, http.MethodPost, uri, body, api.authType, http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{writer.FormDataContentType()},
	})
	if err != nil {
		return StreamVideo{}, err
	}
	var streamVideoResponse StreamVideoResponse
	if err := json.Unmarshal(res, &streamVideoResponse); err != nil {
		return StreamVideo{}, err
	}
	return streamVideoResponse.Result, nil
}

// StreamCreateVideoDirectURL creates a video and returns an authenticated URL
//
// API Reference: https://api.cloudflare.com/#stream-videos-create-a-video-and-get-authenticated-direct-upload-url
func (api *API) StreamCreateVideoDirectURL(ctx context.Context, accountID string, options StreamCreateVideoOptions) (StreamVideoCreate, error) {
	if options.MaxDurationSeconds == 0 {
		return StreamVideoCreate{}, ErrMissingMaxDuration
	}
	uri := fmt.Sprintf("/accounts/%s/stream/direct_upload", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, options)
	if err != nil {
		return StreamVideoCreate{}, err
	}
	var streamVideoCreateResponse StreamVideoCreateResponse
	if err := json.Unmarshal(res, &streamVideoCreateResponse); err != nil {
		return StreamVideoCreate{}, err
	}
	return streamVideoCreateResponse.Result, nil
}

// StreamListVideos list videos currently in stream
//
// API reference: https://api.cloudflare.com/#stream-videos-list-videos
func (api *API) StreamListVideos(ctx context.Context, accountID string, options StreamListOptions) ([]StreamVideo, error) {
	uri := buildURI(fmt.Sprintf("/accounts/%s/stream", accountID), options)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []StreamVideo{}, err
	}
	var streamListResponse StreamListResponse
	if err := json.Unmarshal(res, &streamListResponse); err != nil {
		return []StreamVideo{}, err
	}
	return streamListResponse.Result, nil
}

// Skipped: https://api.cloudflare.com/#stream-videos-initiate-a-video-upload-using-tus

// StreamGetVideo gets the details for a specific video
//
// API Reference: https://api.cloudflare.com/#stream-videos-video-details
func (api *API) StreamGetVideo(ctx context.Context, accountID string, videoID string) (StreamVideo, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/%s", accountID, videoID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return StreamVideo{}, err
	}
	var streamVideoResponse StreamVideoResponse
	if err := json.Unmarshal(res, &streamVideoResponse); err != nil {
		return StreamVideo{}, err
	}
	return streamVideoResponse.Result, nil
}

// StreamEmbedHTML gets an HTML fragment to embed on a web page
//
// API Reference: https://api.cloudflare.com/#stream-videos-embed-code-html
func (api *API) StreamEmbedHTML(ctx context.Context, accountID string, videoID string) (string, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/%s/embed", accountID, videoID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return "", err
	}
	return string(res), nil
}

// StreamDeleteVideo deletes a video
//
// API Reference: https://api.cloudflare.com/#stream-videos-delete-video
func (api *API) StreamDeleteVideo(ctx context.Context, accountID string, videoID string) error {
	uri := fmt.Sprintf("/accounts/%s/stream/%s", accountID, videoID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil

}

// StreamAssociateNFT associates a video to a token and contract address
//
// API Reference: https://api.cloudflare.com/#stream-videos-associate-video-to-an-nft
func (api *API) StreamAssociateNFT(ctx context.Context, accountID string, videoID string, options StreamVideoNFT) (StreamVideo, error) {
	uri := fmt.Sprintf("/accounts/%s/stream/%s", accountID, videoID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, options)
	if err != nil {
		return StreamVideo{}, err
	}
	var streamVideoResponse StreamVideoResponse
	if err := json.Unmarshal(res, &streamVideoResponse); err != nil {
		return StreamVideo{}, err
	}
	return streamVideoResponse.Result, nil
}

//StreamCreateSignedURL creates a signed URL token for a video
//
// API Reference: https://api.cloudflare.com/#stream-videos-associate-video-to-an-nft
func (api *API) StreamCreateSignedURL(ctx context.Context, accountID string, videoID string, options StreamSignedURLOptions) (string, error) {

	uri := fmt.Sprintf("/accounts/%s/stream/%s/token", accountID, videoID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, options)

	if err != nil {
		return "", err
	}
	var streamSignedResponse StreamSignedURLResponse
	if err := json.Unmarshal(res, &streamSignedResponse); err != nil {
		return "", err
	}
	return streamSignedResponse.Result.Token, nil
}
