package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Image represents a Cloudflare Image.
type Image struct {
	ID                string                 `json:"id"`
	Filename          string                 `json:"filename"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	RequireSignedURLs bool                   `json:"requireSignedURLs"`
	Variants          []string               `json:"variants"`
	Uploaded          time.Time              `json:"uploaded"`
}

// ImageUploadRequest is the data required for an Image Upload request.
type ImageUploadRequest struct {
	File              io.ReadCloser
	Name              string
	RequireSignedURLs bool
	Metadata          map[string]interface{}
}

// write writes the image upload data to a multipart writer, so
// it can be used in an HTTP request.
func (b ImageUploadRequest) write(mpw *multipart.Writer) error {
	if b.File == nil {
		return errors.New("a file to upload must be specified")
	}
	name := b.Name
	part, err := mpw.CreateFormFile("file", name)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, b.File)
	if err != nil {
		_ = b.File.Close()
		return err
	}
	_ = b.File.Close()

	// According to the Cloudflare docs, this field defaults to false.
	// For simplicity, we will only send it if the value is true, however
	// if the default is changed to true, this logic will need to be updated.
	if b.RequireSignedURLs {
		err = mpw.WriteField("requireSignedURLs", "true")
		if err != nil {
			return err
		}
	}

	if b.Metadata != nil {
		part, err = mpw.CreateFormField("metadata")
		if err != nil {
			return err
		}
		err = json.NewEncoder(part).Encode(b.Metadata)
		if err != nil {
			return err
		}
	}

	return nil
}

// ImageUploadResponse is the API response for an image upload.
type ImageUploadResponse struct {
	Result Image `json:"result"`
	Response
}

// ImagesListResponse is the API response for listing all images.
type ImagesListResponse struct {
	Result struct {
		Images []Image `json:"images"`
	} `json:"result"`
	Response
}

// ImagesStatsResponse is the API response for image stats.
type ImagesStatsResponse struct {
	Result struct {
		Count ImagesStatsCount `json:"count"`
	} `json:"result"`
	Response
}

// ImagesStatsCount is the stats attached to a ImagesStatsResponse.
type ImagesStatsCount struct {
	Current int64 `json:"current"`
	Allowed int64 `json:"allowed"`
}

// UploadImage uploads a single image.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-upload-an-image-using-a-single-http-request
func (api *API) UploadImage(ctx context.Context, accountID string, upload ImageUploadRequest) (Image, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1", accountID)

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	if err := upload.write(w); err != nil {
		_ = w.Close()
		return Image{}, errors.Wrap(err, "error writing multipart body")
	}
	_ = w.Close()

	res, err := api.makeRequestWithAuthTypeAndHeaders(
		ctx,
		http.MethodPost,
		uri,
		body,
		api.authType,
		http.Header{
			"Accept":       []string{"application/json"},
			"Content-Type": []string{w.FormDataContentType()},
		},
	)
	if err != nil {
		return Image{}, err
	}

	var imageUploadResponse ImageUploadResponse
	err = json.Unmarshal(res, &imageUploadResponse)
	if err != nil {
		return Image{}, errors.Wrap(err, errUnmarshalError)
	}
	return imageUploadResponse.Result, nil
}

// Images lists all images.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-list-images
func (api *API) Images(ctx context.Context, accountID string, pageOpts PaginationOptions) ([]Image, error) {
	v := url.Values{}
	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}
	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	uri := fmt.Sprintf("/accounts/%s/images/v1", accountID)
	if len(v) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Image{}, err
	}

	var imagesListResponse ImagesListResponse
	err = json.Unmarshal(res, &imagesListResponse)
	if err != nil {
		return []Image{}, errors.Wrap(err, errUnmarshalError)
	}
	return imagesListResponse.Result.Images, nil
}

// ImagesStats gets stats for Cloudflare Images.
//
// API Reference: https://api.cloudflare.com/#cloudflare-images-images-usage-statistics
func (api *API) ImagesStats(ctx context.Context, accountID string) (ImagesStatsCount, error) {
	uri := fmt.Sprintf("/accounts/%s/images/v1/stats", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ImagesStatsCount{}, err
	}

	var imagesStatsResponse ImagesStatsResponse
	err = json.Unmarshal(res, &imagesStatsResponse)
	if err != nil {
		return ImagesStatsCount{}, errors.Wrap(err, errUnmarshalError)
	}
	return imagesStatsResponse.Result.Count, nil
}
