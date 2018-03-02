package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"
import "errors"

type TenantAPIResponse struct {
	Resource      *string  `json:"resource"`
	Size          *int     `json:"size"`
	PageNumber    *int     `json:"pageNumber"`
	TotalElements *int     `json:"totalElements"`
	TotalPages    *int     `json:"totalPages"`
	Tenants       []Tenant `json:"tenants"`
}

type Tenant struct {
	Id                              *string  `json:"id"`
	Resource                        *string  `json:"resource,omitempty"`
	Name                            *string  `json:"name,omitempty"`
	Url                             *string  `json:"url,omitempty"`
	About                           *string  `json:"about,omitempty"`
	ContactEmail                    *string  `json:"contactEmail,omitempty"`
	Phone                           *string  `json:"phone,omitempty"`
	UserId                          *string  `json:"userId,omitempty"`
	TermsOfService                  *string  `json:"termsOfService,omitempty"`
	PrivacyPolicy                   *string  `json:"privacyPolicy,omitempty"`
	RevShareRate                    *float32 `json:"revShareRate,omitempty"`
	CcTransactionFeeRate            *float32 `json:"ccTransactionFeeRate,omitempty"`
	MinAppFeeRate                   *float32 `json:"minAppFeeRate,omitempty"`
	EnableConsolidatedBilling       *bool    `json:"enableConsolidatedBilling,omitempty"`
	ShortName                       *string  `json:"shortName,omitempty"`
	EnablePurchaseOrder             *bool    `json:"enablePurchaseOrder,omitempty"`
	EnableEmailNotificationsToUsers *bool    `json:"enableEmailNotificationsToUsers,omitempty"`
	ParentTenantId                  *int     `json:"parentTenantId,omitempty"`
	ExternalId                      *string  `json:"externalId,omitempty"`
	DefaultActivationProfileId      *string  `json:"defaultActivationProfileId,omitempty"`
	EnableMonthlyBilling            *bool    `json:"enableMonthlyBilling,omitempty"`
	DefaultChargeType               *string  `json:"defaultChargeType,omitempty"`
	LoginLogo                       *string  `json:"loginLogo,omitempty"`
	HomePageLogo                    *string  `json:"homePageLogo,omitempty"`
	DomainName                      *string  `json:"domainName,omitempty"`
	//ActivationCodes                 string `json:"activationCodes"`
	//FirewallProfiles                string `json:"firewallProfiles"`
	SkipDefaultUserSecurityGroup *bool        `json:"skipDefaultUserSecurityGroup,omitempty"`
	DisableAllEmailNotification  *bool        `json:"disableAllEmailNotification,omitempty"`
	TrademarkURL                 *string      `json:"trademarkURL,omitempty"`
	Deleted                      *bool        `json:"deleted,omitempty"`
	Preferences                  []Preference `json:"preferences,omitempty"`
}

type Preference struct {
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

func (s *Client) GetTenants() ([]Tenant, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data TenantAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	tenants := data.Tenants
	return tenants, nil
}

func (s *Client) GetTenant(id int) (*Tenant, error) {

	var data Tenant

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(id))
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

	tenant := &data
	return tenant, nil
}

func (s *Client) AddTenant(tenant *Tenant) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants")

	j, err := json.Marshal(tenant)

	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)

	return err
}

func (s *Client) UpdateTenant(tenant *Tenant) (*Tenant, error) {

	var data Tenant

	tenantId := *tenant.Id
	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + tenantId)

	j, err := json.Marshal(tenant)

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

	tenant = &data

	return tenant, nil
}

func (s *Client) DeleteTenantSync(tenantId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)

	if err != nil {

		byt := []byte(err.Error())

		var dat map[string]interface{}

		if err := json.Unmarshal(byt, &dat); err != nil {
			return err
		}

		msg := dat["msg"].(string)

		if msg == "Delete tenant request accepted" {
			return errors.New("Delete tenant request accepted. The tenant deletion is successful only when the following conditions are completed: \n\n - All the running jobs must be terminated for all users – users cannot be deleted before the jobs are terminated.\n\n - All users in the tenant are deleted \n\n - All the sub tenants under the tenant must be deleted prior to issuing this API call. If any sub-tenant is not deleted, then a validation message states that you do this first.\n ")
		}

		return err
	}

	return nil
}

func (s *Client) DeleteTenantAsync(tenantId int) (*OperationStatus, error) {

	var data OperationStatus

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	body, err := s.doRequest(req)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
