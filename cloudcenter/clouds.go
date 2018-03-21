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

type CloudAPIResponse struct {
	Resource      *string `json:"resource,omitempty"`
	Size          *int64  `json:"size,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	TotalElements *int64  `json:"totalElements,omitempty"`
	TotalPages    *int64  `json:"totalPages,omitempty"`
	Clouds        []Cloud `json:"clouds,omitempty"`
}

type Cloud struct {
	Id          *string   `json:"id,omitempty"`
	Resource    *string   `json:"resource,omitempty"`
	Perms       *[]string `json:"perms,omitempty"`
	Name        *string   `json:"name,omitempty" validate:"nonzero"`
	Description *string   `json:"description,omitempty"`
	CloudFamily *string   `json:"cloudFamily,omitempty" validate:"nonzero"`
	PublicCloud *bool     `json:"publicCloud,omitempty"`
	TenantId    *string   `json:"tenantId,omitempty" validate:"nonzero"`
	Detail      *Detail   `json:"detail,omitempty"`
	CanDelete   *bool     `json:"canDelete,omitempty"`
}

type Detail struct {
	CloudAccounts *[]CloudAccount `json:"cloudAccounts"`
	CloudRegions  *[]CloudRegion  `json:"cloudRegions,omitempty"`
	Status        *string         `json:"status,omitempty"`
	StatusDetail  *string         `json:"statusDetail,omitempty"`
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

	if errs := validator.Validate(cloud); errs != nil {
		return nil, errs
	}

	cloudTenantId := *cloud.TenantId

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudTenantId + "/clouds")

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

	if errs := validator.Validate(cloud); errs != nil {
		return nil, errs
	}

	cloudTenantId := *cloud.TenantId

	if nonzero(cloud.Id) {
		return nil, errors.New("Cloud.Id is missing")
	}

	cloudId := *cloud.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudTenantId + "/clouds/" + cloudId)

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
