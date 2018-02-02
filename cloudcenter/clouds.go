package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"

type CloudAPIResponse struct {
	Resource      string  `json:"resource"`
	Size          int     `json:"size"`
	PageNumber    int     `json:"pageNumber"`
	TotalElements int     `json:"totalElements"`
	TotalPages    int     `json:"totalPages"`
	Clouds        []Cloud `json:"clouds"`
}

type Cloud struct {
	Id          string   `json:"id"`
	Resource    string   `json:"resource"`
	Perms       []string `json:"perms"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CloudFamily string   `json:"cloudFamily"`
	PublicCloud bool     `json:"publicCloud"`
	TenantId    string   `json:"tenantId"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CloudInstanceType struct {
	Id                       string  `json:"id"`
	Resource                 string  `json:"resource"`
	InstanceType             string  `json:"instanceType"`
	Name                     string  `json:"name"`
	Description              string  `json:"description"`
	CostPerHour              float32 `json:"costPerHour"`
	MemorySize               int     `json:"memorySize"`
	NumOfCPUs                int     `json:"numOfCPUs"`
	NumOfNICs                int     `json:"numOfNICs"`
	LocalStorageCount        int     `json:"localStorageCount"`
	LocalStorageSize         int     `json:"localStorageSize"`
	CudaSupport              bool    `json:"cudaSupport"`
	SsdSupport               bool    `json:"ssdSupport"`
	Support32Bit             bool    `json:"support32Bit"`
	Support64Bit             bool    `json:"support64Bit"`
	Dummy                    bool    `json:"dummy"`
	Deleted                  bool    `json:"deleted"`
	Mutability               string  `json:"mutability"`
	SupportHardwareProvision bool    `json:"supportHardwareProvision"`
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
