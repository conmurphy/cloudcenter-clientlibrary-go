package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"

type CloudInstanceTypeAPIResponse struct {
	Resource           string              `json:"resource,omitempty"`
	Size               int                 `json:"size,omitempty"`
	PageNumber         int                 `json:"pageNumber,omitempty"`
	TotalElements      int                 `json:"totalElements,omitempty"`
	TotalPages         int                 `json:"totalPages,omitempty"`
	CloudInstanceTypes []CloudInstanceType `json:"cloudInstanceTypes,omitempty"`
}

type CloudInstanceType struct {
	Id                        string   `json:"id,omitempty"`
	Resource                  string   `json:"resource,omitempty"`
	Perms                     []string `json:"perms,omitempty"`
	Name                      string   `json:"name,omitempty"`
	Description               string   `json:"description,omitempty"`
	Type                      string   `json:"type,omitempty"`
	TenantId                  string   `json:"tenantId,omitempty"`
	CloudId                   string   `json:"cloudId,omitempty"`
	RegionId                  string   `json:"regionId,omitempty"`
	CostPerHour               float64  `json:"costPerHour,omitempty"`
	MemorySize                int32    `json:"memorySize,omitempty"`
	NumOfCPUs                 int32    `json:"numOfCpus,omitempty"`
	NumOfNICs                 int32    `json:"numOfNics,omitempty"`
	LocalStorageSize          int32    `json:"localStorageSize,omitempty"`
	SupportsSSD               bool     `json:"supportsSSD,omitempty"`
	Supports32Bit             bool     `json:"supports32Bit,omitempty"`
	Supports64Bit             bool     `json:"supports64Bit,omitempty"`
	LocalStorageCount         float64  `json:"localStorageCount,omitempty"`
	SupportsHardwareProvision bool     `json:"supportsHardwareProvision,omitempty"`
}

func (s *Client) GetCloudInstanceTypes(tenantId int, cloudId int, regionId int) ([]CloudInstanceType, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/instanceTypes/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CloudInstanceTypeAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	cloudInstanceType := data.CloudInstanceTypes
	return cloudInstanceType, nil
}

func (s *Client) GetCloudInstanceType(tenantId int, cloudId int, regionId int, instanceId int) (*CloudInstanceType, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/instanceTypes/" + strconv.Itoa(instanceId))

	var data CloudInstanceType

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

	cloudInstanceType := &data

	return cloudInstanceType, nil
}

func (s *Client) AddCloudInstanceType(cloudInstanceType *CloudInstanceType) (*CloudInstanceType, error) {

	var data CloudInstanceType

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudInstanceType.TenantId + "/clouds/" + cloudInstanceType.CloudId + "/regions/" + cloudInstanceType.RegionId + "/instanceTypes/")

	j, err := json.Marshal(cloudInstanceType)

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

	cloudInstanceType = &data

	return cloudInstanceType, nil
}

func (s *Client) UpdateCloudInstanceType(cloudInstanceType *CloudInstanceType) (*CloudInstanceType, error) {

	var data CloudInstanceType

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudInstanceType.TenantId + "/clouds/" + cloudInstanceType.CloudId + "/regions/" + cloudInstanceType.RegionId + "/instanceTypes/" + cloudInstanceType.Id)

	j, err := json.Marshal(cloudInstanceType)

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

	cloudInstanceType = &data

	return cloudInstanceType, nil
}

func (s *Client) DeleteCloudInstanceType(tenantId int, cloudId int, regionId int, instanceId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/instanceTypes/" + strconv.Itoa(instanceId))

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

func (s *Client) SyncCloudInstanceTypes(tenantId int, cloudId int, regionId int) ([]CloudInstanceType, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/regions/" + strconv.Itoa(regionId) + "/syncInstanceTypes/")
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CloudInstanceTypeAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	cloudInstanceType := data.CloudInstanceTypes
	return cloudInstanceType, nil
}
