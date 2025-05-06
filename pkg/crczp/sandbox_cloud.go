package crczp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Image struct {
	OsDistro        string   `json:"os_distro" tfsdk:"os_distro"`
	OsType          string   `json:"os_type" tfsdk:"os_type"`
	DiskFormat      string   `json:"disk_format" tfsdk:"disk_format"`
	ContainerFormat string   `json:"container_format" tfsdk:"container_format"`
	Visibility      string   `json:"visibility" tfsdk:"visibility"`
	Size            float64  `json:"size" tfsdk:"size"`
	Status          string   `json:"status" tfsdk:"status"`
	MinRam          int      `json:"min_ram" tfsdk:"min_ram"`
	MinDisk         int      `json:"min_disk" tfsdk:"min_disk"`
	CreatedAt       string   `json:"created_at" tfsdk:"created_at"`
	UpdatedAt       string   `json:"updated_at" tfsdk:"updated_at"`
	Tags            []string `json:"tags" tfsdk:"tags"`
	DefaultUser     string   `json:"default_user" tfsdk:"default_user"`
	Name            string   `json:"name" tfsdk:"name"`
	OwnerSpecified  struct{} `json:"owner_specified" tfsdk:"owner_specified"`
}

func (c *Client) GetImages(ctx context.Context) ([]Image, error) {
	return c.GetImagesByPage(ctx, 1, 10000)
}

func (c *Client) GetImagesByPage(ctx context.Context, page int64, pageSize int64) ([]Image, error) {
	query := url.Values{}
	query.Add("page", strconv.FormatInt(page, 10))
	query.Add("page_size", strconv.FormatInt(pageSize, 10))
	query.Add("cached", boolToString(false))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/sandbox-service/api/v1/images?%s", c.Endpoint, query.Encode()), nil)
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequestWithRetry(req, http.StatusOK, "sandbox images", "")
	if err != nil {
		return nil, err
	}

	images := Pagination[[]Image]{}
	err = json.Unmarshal(body, &images)
	if err != nil {
		return nil, err
	}

	return images.Results, nil
}
