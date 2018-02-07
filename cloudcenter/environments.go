package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"
import "bytes"

type EnvironmentAPIResponse struct {
	Resource      string        `json:"resource"`
	Size          int           `json:"size"`
	PageNumber    int           `json:"pageNumber"`
	TotalElements int           `json:"totalElements"`
	TotalPages    int           `json:"totalPages"`
	Environments  []Environment `json:"deploymentEnvironments"`
}

type Environment struct {
	Id               string            `json:"id"`
	Resource         string            `json:"resource"`
	Name             string            `json:"name"`
	Perms            []string          `json:"perms"`
	Description      string            `json:"description"`
	AllowedClouds    string            `json:"allowedClouds"`
	DefaultSettings  string            `json:"defaultSettings"`
	RequiresApproval bool              `json:"requiresApproval"`
	AssociatedClouds []AssociatedCloud `json:"associatedClouds"`
	TotalDeployments int32             `json:"totalDeployments"`
	ExtensionId      string            `json:"extensionId"`
	CostDetails      CostDetail        `json:"costDetails"`
}

type CostDetail struct {
	TotalCloudCost float32 `json:"totalCloudCost"`
	TotalAppCost   float32 `json:"totalAppCost"`
	TotalJobsCost  float32 `json:"totalJobsCost"`
}

type AssociatedCloud struct {
	RegionId                 string                    `json:"regionId"`
	RegionName               string                    `json:"regionName"`
	RegionDisplayName        string                    `json:"regionDisplayName"`
	CloudFamily              string                    `json:"cloudFamily"`
	CloudId                  string                    `json:"cloudId"`
	CloudAccountId           string                    `json:"cloudAccountId"`
	CloudName                string                    `json:"cloudName"`
	CloudAccountName         string                    `json:"cloudAccountName"`
	CloudAssociationDefaults []CloudAssociationDefault `json:"cloudAssociationDefaults"`
	Default                  bool                      `json:"default"`
}

type CloudAssociationDefault struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (s *Client) GetEnvironments() ([]Environment, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/environments")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data EnvironmentAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	environment := data.Environments
	return environment, nil
}

func (s *Client) GetEnvironment(id int) (*Environment, error) {

	var data Environment

	url := fmt.Sprintf(s.BaseURL + "/v1/environments/" + strconv.Itoa(id))
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

	environment := &data
	return environment, nil
}

func (s *Client) AddEnvironment(environment *Environment) (*Environment, error) {

	var data Environment

	url := fmt.Sprintf(s.BaseURL + "/v1/environments")

	j, err := json.Marshal(environment)

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

	environment = &data

	return environment, nil
}

func (s *Client) UpdateEnvironment(environment *Environment) (*Environment, error) {

	var data Environment

	url := fmt.Sprintf(s.BaseURL + "/v1/environments/" + environment.Id)

	j, err := json.Marshal(environment)

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

	environment = &data

	return environment, nil
}

func (s *Client) DeleteEnvironment(environmentId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/environments/" + strconv.Itoa(environmentId))

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
