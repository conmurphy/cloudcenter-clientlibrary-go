package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"

type CloudAPIResponse struct {
	Resource      string  `json:"resource,omitempty"`
	Size          int     `json:"size,omitempty"`
	PageNumber    int     `json:"pageNumber,omitempty"`
	TotalElements int     `json:"totalElements,omitempty"`
	TotalPages    int     `json:"totalPages,omitempty"`
	Clouds        []Cloud `json:"clouds,omitempty"`
}

type Cloud struct {
	Id          string   `json:"id,omitempty"`
	Resource    string   `json:"resource,omitempty"`
	Perms       []string `json:"perms,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	CloudFamily string   `json:"cloudFamily,omitempty"`
	PublicCloud bool     `json:"publicCloud,omitempty"`
	TenantId    string   `json:"tenantId,omitempty"`
	Detail      *Detail  `json:"detail,omitempty"`
}

type Detail struct {
	CloudAccounts []CloudAccount `json:"cloudAccounts"`
	CloudRegions  []CloudRegion  `json:"cloudRegions,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetail  string         `json:"statusDetail,omitempty"`
}

func (s *Client) GetClouds(tenantId int) ([]Cloud, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CloudAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	clouds := data.Clouds
	return clouds, nil
}

func (s *Client) GetCloud(tenantId int, cloudId int) (*Cloud, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId))

	var data Cloud

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

	cloud := &data

	return cloud, nil
}

func (s *Client) AddCloud(cloud *Cloud) (*Cloud, error) {

	var data Cloud

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloud.TenantId + "/clouds")

	j, err := json.Marshal(cloud)

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

	cloud = &data

	return cloud, nil
}

func (s *Client) UpdateCloud(cloud *Cloud) (*Cloud, error) {

	var data Cloud

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloud.TenantId + "/clouds/" + cloud.Id)

	j, err := json.Marshal(cloud)

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

	cloud = &data

	return cloud, nil
}

func (s *Client) DeleteCloud(tenantId int, cloudId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId))

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
