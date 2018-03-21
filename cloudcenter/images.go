package cloudcenter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	validator "gopkg.in/validator.v2"
)

type ImageAPIResponse struct {
	Resource      *string `json:"resource,omitempty"`
	Size          *int64  `json:"size,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	TotalElements *int64  `json:"totalElements,omitempty"`
	TotalPages    *int64  `json:"totalPages,omitempty"`
	Images        []Image `json:"images,omitempty"`
}

type Image struct {
	Id                *string       `json:"id,omitempty"`
	TenantId          *int64        `json:"tenantId,omitempty"`
	Resource          *string       `json:"resource,omitempty"`
	Perms             *[]string     `json:"perms,omitempty"`
	Name              *string       `json:"name,omitempty" validate:"nonzero"`
	InternalImageName *string       `json:"internalImageName,omitempty"`
	Description       *string       `json:"description,omitempty"`
	Visibility        *string       `json:"visibility,omitempty"`
	ImageType         *string       `json:"imageType,omitempty" validate:"nonzero"`
	OSName            *string       `json:"osName,omitempty" validate:"nonzero"`
	Tags              *[]string     `json:"tags,omitempty"`
	Enabled           *bool         `json:"enabled,omitempty"`
	SystemImage       *bool         `json:"systemImage,omitempty"`
	NumOfNICs         *int64        `json:"numOfNics,omitempty"`
	AttachCount       *int64        `json:"count,omitempty"`
	Details           *ImageDetails `json:"detail,omitempty"`
}

type ImageDetails struct {
	Count       *int64        `json:"count,omitempty"`
	CloudImages *[]CloudImage `json:"cloudImages,omitempty"`
}

type CloudImage struct {
	Id                   *string     `json:"id,omitempty"`
	Resource             *string     `json:"resource,omitempty"`
	Perms                *[]string   `json:"perms,omitempty"`
	RegionId             *string     `json:"regionId,omitempty"`
	CloudProviderImageId *string     `json:"cloudProviderImageId,omitempty"`
	LaunchUserName       *string     `json:"launchUserName,omitempty"`
	ImageId              *string     `json:"imageId,omitempty"`
	GrantAndRevoke       *bool       `json:"grantAndRevoke,omitempty"`
	ImageCloudAccountId  *int64      `json:"imageCloudAccountId,omitempty"`
	Resources            *[]Resource `json:"resources,omitempty"`
	Mappings             *[]Mapping  `json:"mappings,omitempty"`
}

func (s *Client) GetImages(tenantId int) ([]Image, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/images?detail=true")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data ImageAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	image := data.Images
	return image, nil
}

func (s *Client) GetImage(tenantId int, imageId int) (*Image, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/images/" + strconv.Itoa(imageId) + "?detail=true")

	var data Image

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	image := &data

	return image, nil
}

func (s *Client) AddImage(image *Image) (*Image, error) {

	var data Image

	if errs := validator.Validate(image); errs != nil {
		return nil, errs
	}

	imageTenantId := int(*image.TenantId)

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(imageTenantId) + "/images")

	j, err := json.Marshal(image)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	image = &data

	return image, nil

}

func (s *Client) UpdateImage(image *Image) (*Image, error) {

	var data Image

	if errs := validator.Validate(image); errs != nil {
		return nil, errs
	}

	if nonzero(image.Id) {
		return nil, errors.New("Image.Id is missing")
	}

	imageTenantId := int(*image.TenantId)
	imageId := *image.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(imageTenantId) + "/images/" + imageId)

	j, err := json.Marshal(image)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	image = &data

	return image, nil
}

func (s *Client) DeleteImage(tenantId int, imageId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/images/" + strconv.Itoa(imageId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
