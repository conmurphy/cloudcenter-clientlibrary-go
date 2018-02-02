package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"
import "bytes"

type ActivationProfileAPIResponse struct {
	Resource           string              `json:"resource"`
	Size               int                 `json:"size"`
	PageNumber         int                 `json:"pageNumber"`
	TotalElements      int                 `json:"totalElements"`
	TotalPages         int                 `json:"totalPages"`
	ActivationProfiles []ActivationProfile `json:"activationProfiles"`
}

type ActivationProfile struct {
	Id                  string           `json:"id,omitempty"`
	Name                string           `json:"name,omitempty"`
	Description         string           `json:"description,omitempty"`
	TenantId            int              `json:"tenantId,omitempty"`
	PlanId              string           `json:"planId,omitempty"`
	BundleId            string           `json:"bundleId,omitempty"`
	ContractId          string           `json:"contractId,omitempty"`
	DepEnvId            string           `json:"depEnvId,omitempty"`
	ActivateRegions     []ActivateRegion `json:"activateRegions,omitempty"`
	ImportApps          []string         `json:"importApps,omitempty"`
	AgreeToContract     bool             `json:"agreeToContract,omitempty"`
	SendActivationEmail bool             `json:"sendActivationEmail,omitempty"`
}

type ActivateRegion struct {
	RegionId string `json:"regionId"`
}

func (s *Client) GetActivationProfiles(tenantId int) ([]ActivationProfile, error) {

	var data ActivationProfileAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/activationProfiles")

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

	activationProfile := data.ActivationProfiles
	return activationProfile, nil
}

func (s *Client) GetActivationProfile(tenantId int, activationProfileId int) (*ActivationProfile, error) {

	var data ActivationProfile

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/activationProfiles/" + strconv.Itoa(activationProfileId))
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

	activationProfile := &data
	return activationProfile, nil
}

func (s *Client) AddActivationProfile(activationProfile *ActivationProfile) (*ActivationProfile, error) {

	var data ActivationProfile

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(activationProfile.TenantId) + "/activationProfiles")

	j, err := json.Marshal(activationProfile)

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

	activationProfile = &data

	return activationProfile, nil
}

func (s *Client) UpdateActivationProfile(activationProfile *ActivationProfile) (*ActivationProfile, error) {

	var data ActivationProfile

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(activationProfile.TenantId) + "/activationProfiles/" + activationProfile.Id)

	j, err := json.Marshal(activationProfile)

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

	activationProfile = &data

	return activationProfile, nil
}

func (s *Client) DeleteActivationProfile(tenantId int, activationProfileId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/activationProfiles/" + strconv.Itoa(activationProfileId))

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
