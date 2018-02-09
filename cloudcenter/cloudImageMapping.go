package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"

type CloudImageMappingAPIResponse struct {
	Resource           string              `json:"resource,omitempty"`
	Size               int                 `json:"size,omitempty"`
	PageNumber         int                 `json:"pageNumber,omitempty"`
	TotalElements      int                 `json:"totalElements,omitempty"`
	TotalPages         int                 `json:"totalPages,omitempty"`
	CloudImageMappings []CloudImageMapping `json:"cloudImages,omitempty"`
}

type CloudImageMapping struct {
	Id                   string     `json:"id,omitempty"`
	Resource             string     `json:"resource,omitempty"`
	Perms                []string   `json:"perms,omitempty"`
	TenantId             string     `json:"tenantId,omitempty"`
	CloudId              string     `json:"cloudId,omitempty"`
	CloudRegionId        string     `json:"cloudRegionId,omitempty"`
	RegionId             string     `json:"regionId,omitempty"`
	CloudImageId         string     `json:"cloudImageId,omitempty"`
	CloudProviderImageId string     `json:"cloudProviderImageId,omitempty"`
	LaunchUserName       string     `json:"launchUserName,omitempty"`
	ImageId              string     `json:"imageId,omitempty"`
	GrantAndRevoke       bool       `json:"grantAndRevoke,omitempty"`
	ImageCloudAccountId  int        `json:"imageCloudAccountId,omitempty"`
	Resources            []Resource `json:"resources,omitempty"`
	Mappings             []Mapping  `json:"mappings,omitempty"`
}

type Resource struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Mapping struct {
	Id                           string            `json:"id,omitempty"`
	CloudInstanceType            CloudInstanceType `json:"cloudInstanceType,omitempty"`
	CostOverride                 float64           `json:"costOverride,omitempty"`
	CloudProviderImageIdOverride string            `json:"CloudProviderImageIdOverride,omitempty"`
}

func (s *Client) GetCloudImageMappings(tenantId int, cloudId int, regionId int) ([]CloudImageMapping, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/images/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CloudImageMappingAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	cloudImage := data.CloudImageMappings
	return cloudImage, nil
}

func (s *Client) GetCloudImageMapping(tenantId int, cloudId int, regionId int, imageId int) (*CloudImageMapping, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/images/" + strconv.Itoa(imageId))

	var data CloudImageMapping

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

	cloudImage := &data

	return cloudImage, nil
}

func (s *Client) AddCloudImageMapping(cloudImage *CloudImageMapping) (*CloudImageMapping, error) {

	var data CloudImageMapping

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudImage.TenantId + "/clouds/" + cloudImage.CloudId + "/regions/" + cloudImage.RegionId + "/images/")

	j, err := json.Marshal(cloudImage)

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

	cloudImage = &data

	return cloudImage, nil
}

func (s *Client) UpdateCloudImageMapping(cloudImage *CloudImageMapping) (*CloudImageMapping, error) {

	var data CloudImageMapping

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudImage.TenantId + "/clouds/" + cloudImage.CloudId + "/regions/" + cloudImage.RegionId + "/images/" + cloudImage.Id)

	j, err := json.Marshal(cloudImage)

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

	cloudImage = &data

	return cloudImage, nil
}

func (s *Client) DeleteCloudImageMapping(tenantId int, cloudId int, regionId int, imageId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/images/" + strconv.Itoa(imageId))

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
