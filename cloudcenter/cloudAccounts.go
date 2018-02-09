package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"

type CloudAccountAPIResponse struct {
	Resource      string         `json:"resource,omitempty"`
	Size          int            `json:"size,omitempty"`
	PageNumber    int            `json:"pageNumber,omitempty"`
	TotalElements int            `json:"totalElements,omitempty"`
	TotalPages    int            `json:"totalPages,omitempty"`
	CloudAccounts []CloudAccount `json:"cloudAccounts,omitempty"`
}

type CloudAccount struct {
	Id                 string            `json:"id,omitempty"`
	Resource           string            `json:"resource,omitempty"`
	Perms              []string          `json:"perms,omitempty"`
	DisplayName        string            `json:"displayName,omitempty"`
	CloudId            string            `json:"cloudId,omitempty"`
	UserId             string            `json:"userId,omitempty"`
	AccountId          string            `json:"accountId,omitempty"`
	AccountName        string            `json:"accountName,omitempty"`
	AccountPassword    string            `json:"accountPassword,omitempty"`
	AccountDescription string            `json:"accountDescription,omitempty"`
	ManageCost         bool              `json:"manageCost,omitempty"`
	PublicVisible      bool              `json:"publicVisible,omitempty"`
	AllowedUsers       []int             `json:"allowedUsers,omitempty"`
	AccessPermission   string            `json:"accessPermission,omitempty"`
	AccountProperties  []AccountProperty `json:"accountProperties,omitempty"`
	TenantId           string            `json:"tenantId,omitempty"`
}

type AccountProperty struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (s *Client) GetCloudAccounts(tenantId int, cloudId int) ([]CloudAccount, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/accounts/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data CloudAccountAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	cloudAccounts := data.CloudAccounts
	return cloudAccounts, nil
}

func (s *Client) GetCloudAccount(tenantId int, cloudId int, accountId int) (*CloudAccount, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/accounts/" + strconv.Itoa(accountId))

	var data CloudAccount

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

	cloudAccount := &data

	return cloudAccount, nil
}

func (s *Client) AddCloudAccount(cloudAccount *CloudAccount) (*CloudAccount, error) {

	var data CloudAccount

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudAccount.TenantId + "/clouds/" + cloudAccount.CloudId + "/accounts")

	j, err := json.Marshal(cloudAccount)

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

	cloudAccount = &data

	return cloudAccount, nil
}

func (s *Client) UpdateCloudAccount(cloudAccount *CloudAccount) (*CloudAccount, error) {

	var data CloudAccount

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + cloudAccount.TenantId + "/clouds/" + cloudAccount.CloudId + "/accounts/" + cloudAccount.Id)

	j, err := json.Marshal(cloudAccount)

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

	cloudAccount = &data

	return cloudAccount, nil
}

func (s *Client) DeleteCloudAccount(tenantId int, cloudId int, accountId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/clouds/" + strconv.Itoa(cloudId) + "/accounts/" + strconv.Itoa(accountId))

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
