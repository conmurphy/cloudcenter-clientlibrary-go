package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"
import "errors"

type BundleAPIResponse struct {
	Resource      *string  `json:"resource"`
	Size          *int64   `json:"size"`
	PageNumber    *int64   `json:"pageNumber"`
	TotalElements *int64   `json:"totalElements"`
	TotalPages    *int64   `json:"totalPages"`
	Bundles       []Bundle `json:"bundles"`
}

type Bundle struct {
	Id               *string   `json:"id,omitempty"`
	Resource         *string   `json:"resource,omitempty"`
	Perms            *[]string `json:"perms,omitempty"`
	Type             *string   `json:"type,omitempty"`
	Name             *string   `json:"name,omitempty"`
	Description      *string   `json:"description,omitempty"`
	Limit            *float64  `json:"limit,omitempty"`
	Price            *float64  `json:"price,omitempty"`
	ExpirationDate   *float64  `json:"expirationDate,omitempty"`
	ExpirationMonths *int64    `json:"expirationMonths,omitempty"`
	Disabled         *bool     `json:"disabled,omitempty"`
	ShowOnlyToAdmin  *bool     `json:"showOnlyToAdmin,omitempty"`
	NumberOfUsers    *float64  `json:"numberOfUsers,omitempty"`
	TenantId         *string   `json:"tenantId,omitempty"`
	PublishedAppIds  []string  `json:"publishedAppIds,omitempty"`
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

func (s *Client) GetBundleFromName(TenantId int, BundleNameSearchString string) (*Bundle, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(TenantId) + "/bundles")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data BundleAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	bundles := data.Bundles

	for _, bundle := range bundles {
		bundleName := *bundle.Name
		if bundleName == BundleNameSearchString {

			return &bundle, nil
		}
	}

	return nil, errors.New("BUNDLE NOT FOUND")

}

func (s *Client) AddBundle(bundle *Bundle) (*Bundle, error) {

	var data Bundle

	bundleTenantId := *bundle.TenantId
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + bundleTenantId + "/bundles")

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

	bundleTenantId := *bundle.TenantId
	bundleId := *bundle.Id

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + bundleTenantId + "/bundles/" + bundleId)

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
