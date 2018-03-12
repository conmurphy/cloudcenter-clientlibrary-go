package cloudcenter

import "fmt"
import "net/http"
import "strconv"
import "encoding/json"
import "bytes"

type ContractAPIResponse struct {
	Resource      *string    `json:"resource"`
	Size          *int64     `json:"size"`
	PageNumber    *int64     `json:"pageNumber"`
	TotalElements *int64     `json:"totalElements"`
	TotalPages    *int64     `json:"totalPages"`
	Contracts     []Contract `json:"contracts"`
}

type Contract struct {
	Id              *string   `json:"id,omitempty"`
	Resource        *string   `json:"resource,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Description     *string   `json:"description,omitempty"`
	Perms           *[]string `json:"perms"`
	TenantId        *string   `json:"tenantId,omitempty"`
	Length          *int64    `json:"length,omitempty"`
	Terms           *string   `json:"terms,omitempty"`
	DiscountRate    *float64  `json:"discountRate,omitempty"`
	Disabled        *bool     `json:"disabled,omitempty"`
	ShowOnlyToAdmin *bool     `json:"showOnlyToAdmin,omitempty"`
	NumberOfUsers   *int64    `json:"numberOfUsers,omitempty"`
}

func (s *Client) GetContracts(tenantId int) ([]Contract, error) {

	var data ContractAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/contracts")

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

	contracts := data.Contracts
	return contracts, nil
}

func (s *Client) GetContract(tenantId int, contractId int) (*Contract, error) {

	var data Contract

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/contracts/" + strconv.Itoa(contractId))
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

	contract := &data
	return contract, nil
}

func (s *Client) AddContract(contract *Contract) (*Contract, error) {

	var data Contract

	contractTenantId := *contract.TenantId
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + contractTenantId + "/contracts")

	j, err := json.Marshal(contract)

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

	contract = &data

	return contract, nil
}

func (s *Client) UpdateContract(contract *Contract) (*Contract, error) {

	var data Contract

	contractTenantId := *contract.TenantId
	contractId := *contract.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + contractTenantId + "/contracts/" + contractId)

	j, err := json.Marshal(contract)

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

	contract = &data

	return contract, nil
}

func (s *Client) DeleteContract(tenantId int, contractId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/contracts/" + strconv.Itoa(contractId))

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
