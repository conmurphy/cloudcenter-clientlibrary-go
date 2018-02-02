package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"

type BundleAPIResponse struct {
	Resource      string   `json:"resource"`
	Size          int      `json:"size"`
	PageNumber    int      `json:"pageNumber"`
	TotalElements int      `json:"totalElements"`
	TotalPages    int      `json:"totalPages"`
	Bundles       []Bundle `json:"bundles"`
}

type Bundle struct {
	Id               string   `json:"id,omitempty"`
	Resource         string   `json:"resource,omitempty"`
	Perms            []string `json:"perms,omitempty"`
	Type             string   `json:"type,omitempty"`
	Name             string   `json:"name,omitempty"`
	Description      string   `json:"description,omitempty"`
	Limit            float32  `json:"limit,omitempty"`
	Price            float32  `json:"price,omitempty"`
	ExpirationDate   int64    `json:"expirationDate,omitempty"`
	ExpirationMonths int64    `json:"expirationMonths,omitempty"`
	Disabled         bool     `json:"disabled,omitempty"`
	ShowOnlyToAdmin  bool     `json:"showOnlyToAdmin,omitempty"`
	NumberOfUsers    float32  `json:"numberOfUsers,omitempty"`
	TenantId         string   `json:"tenantId,omitempty"`
	PublishedAppIds  []int    `json:"publishedAppIds,omitempty"`
}

func (s *Client) GetBundles(TenantId int) ([]Bundle, error) {

	var data BundleAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(TenantId) + "/bundles")

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

	bundles := data.Bundles
	return bundles, nil
}

func (s *Client) GetBundle(TenantId int, BundleId int) (*Bundle, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(TenantId) + "/bundles/" + strconv.Itoa(BundleId))

	var data Bundle

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

	bundle := &data

	return bundle, nil
}

func (s *Client) AddBundle(bundle *Bundle) (*Bundle, error) {

	var data Bundle

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + bundle.TenantId + "/bundles")

	j, err := json.Marshal(bundle)

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

	bundle = &data

	return bundle, nil
}

func (s *Client) UpdateBundle(bundle *Bundle) (*Bundle, error) {

	var data Bundle

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + bundle.TenantId + "/bundles/" + bundle.Id)

	j, err := json.Marshal(bundle)

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

	bundle = &data

	return bundle, nil
}

func (s *Client) DeleteBundle(tenantId int, bundleId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/bundles/" + strconv.Itoa(bundleId))

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
