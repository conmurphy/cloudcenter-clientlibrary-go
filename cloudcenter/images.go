package cloudcenter

import "fmt"
import "net/http"

import "encoding/json"
import "strconv"
import "bytes"

type ImageAPIResponse struct {
	Resource      string  `json:"resource,omitempty"`
	Size          int     `json:"size,omitempty"`
	PageNumber    int     `json:"pageNumber,omitempty"`
	TotalElements int     `json:"totalElements,omitempty"`
	TotalPages    int     `json:"totalPages,omitempty"`
	Images        []Image `json:"images,omitempty"`
}

type Image struct {
	Id                string   `json:"id,omitempty"`
	ImageId           string   `json:"imageId,omitempty"`
	TenantId          int32    `json:"tenantId,omitempty"`
	Resource          string   `json:"resource,omitempty"`
	Perms             []string `json:"perms,omitempty"`
	Name              string   `json:"name,omitempty"`
	InternalImageName string   `json:"internalImageName,omitempty"`
	Description       string   `json:"description,omitempty"`
	Visibility        string   `json:"visibility,omitempty"`
	ImageType         string   `json:"imageType,omitempty"`
	OSName            string   `json:"osName,omitempty"`
	Tags              []string `json:"tags,omitempty"`
	Enabled           bool     `json:"enabled,omitempty"`
	SystemImage       bool     `json:"systemImage,omitempty"`
	NumOfNICs         int32    `json:"numOfNics,omitempty"`
	Count             int32    `json:"count,omitempty"`
}

func (s *Client) GetImages(tenantId int) ([]Image, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/images")
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

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/images/" + strconv.Itoa(imageId))

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

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(int(image.TenantId)) + "/images/" + image.Id)

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

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(int(image.TenantId)) + "/images/" + image.Id)

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
