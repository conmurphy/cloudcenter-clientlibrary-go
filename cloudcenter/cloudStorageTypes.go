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

type CloudStorageTypeAPIResponse struct {
	Resource          *string            `json:"resource,omitempty"`
	Size              *int64             `json:"size,omitempty"`
	PageNumber        *int64             `json:"pageNumber,omitempty"`
	TotalElements     *int64             `json:"totalElements,omitempty"`
	TotalPages        *int64             `json:"totalPages,omitempty"`
	CloudStorageTypes []CloudStorageType `json:"cloudStorageTypes,omitempty"`
}

type CloudStorageType struct {
	Id               *string  `json:"id,omitempty"`
	CloudId          *string  `json:"cloudId,omitempty" validate:"nonzero"`
	TenantId         *string  `json:"tenantId,omitempty" validate:"nonzero"`
	RegionId         *string  `json:"regionId,omitempty" validate:"nonzero"`
	Resource         *string  `json:"resource,omitempty"`
	Name             *string  `json:"name,omitempty" validate:"nonzero"`
	Description      *string  `json:"description,omitempty"`
	Type             *string  `json:"type,omitempty" validate:"nonzero"`
	CostPerMonth     *float64 `json:"costPerMonth,omitempty"`
	MinVolumeSize    *int64   `json:"minVolumeSize,omitempty"`
	MaxVolumeSize    *int64   `json:"maxVolumeSize,omitempty"`
	MaxIOPS          *int64   `json:"maxIOPS,omitempty"`
	MaxThroughput    *int64   `json:"maxThroughput,omitempty"`
	ProvisionedIOPS  *bool    `json:"provisionedIOPS,omitempty"`
	IOPSCostPerMonth *float64 `json:"iopsCostPerMonth,omitempty"`
}

func (s *Client) GetCloudStorageTypes(tenantId int, cloudId int, regionId int) ([]CloudStorageType, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/storageTypes")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CloudStorageTypeAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	cloudStorageType := data.CloudStorageTypes
	return cloudStorageType, nil
}

func (s *Client) GetCloudStorageType(tenantId int, cloudId int, regionId int, cloudStorageTypeId int) (*CloudStorageType, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/storageTypes/" + strconv.Itoa(cloudStorageTypeId))

	var data CloudStorageType

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

	cloudStorageType := &data

	return cloudStorageType, nil
}

func (s *Client) AddCloudStorageType(cloudStorageType *CloudStorageType) (*CloudStorageType, error) {

	var data CloudStorageType

	if errs := validator.Validate(cloudStorageType); errs != nil {
		return nil, errs
	}

	cloudStorageTypeTenantId := *cloudStorageType.TenantId
	cloudStorageTypeCloudId := *cloudStorageType.CloudId
	cloudStorageTypeRegionId := *cloudStorageType.RegionId
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudStorageTypeTenantId + "/clouds/" + cloudStorageTypeCloudId + "/regions/" + cloudStorageTypeRegionId + "/storageTypes")

	j, err := json.Marshal(cloudStorageType)

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

	cloudStorageType = &data

	return cloudStorageType, nil

}

func (s *Client) UpdateCloudStorageType(cloudStorageType *CloudStorageType) (*CloudStorageType, error) {

	var data CloudStorageType

	if errs := validator.Validate(cloudStorageType); errs != nil {
		return nil, errs
	}

	if nonzero(cloudStorageType.Id) {
		return nil, errors.New("CloudStorageType.Id is missing")
	}

	cloudStorageTypeTenantId := *cloudStorageType.TenantId
	cloudStorageTypeCloudId := *cloudStorageType.CloudId
	cloudStorageTypeId := *cloudStorageType.Id
	cloudStorageTypeRegionId := *cloudStorageType.RegionId

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudStorageTypeTenantId + "/clouds/" + cloudStorageTypeCloudId + "/regions/" + cloudStorageTypeRegionId + "/storageTypes/" + cloudStorageTypeId)

	j, err := json.Marshal(cloudStorageType)

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

	cloudStorageType = &data

	return cloudStorageType, nil
}

func (s *Client) DeleteCloudStorageType(tenantId int, cloudId int, regionId int, cloudStorageTypeId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/storageTypes/" + strconv.Itoa(cloudStorageTypeId))

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
