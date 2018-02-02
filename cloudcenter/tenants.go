package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
import "bytes"

type TenantAPIResponse struct {
	Resource      string   `json:"resource"`
	Size          int      `json:"size"`
	PageNumber    int      `json:"pageNumber"`
	TotalElements int      `json:"totalElements"`
	TotalPages    int      `json:"totalPages"`
	Tenants       []Tenant `json:"tenants"`
}

type Tenant struct {
	Id                              string  `json:"id"`
	Resource                        string  `json:"resource"`
	Name                            string  `json:"name"`
	Url                             string  `json:"url"`
	About                           string  `json:"about"`
	ContactEmail                    string  `json:"contactEmail"`
	Phone                           string  `json:"phone"`
	UserId                          string  `json:"userId"`
	TermsOfService                  string  `json:"termsOfService"`
	PrivacyPolicy                   string  `json:"privacyPolicy"`
	RevShareRate                    float32 `json:"revShareRate"`
	CcTransactionFeeRate            float32 `json:"ccTransactionFeeRate"`
	MinAppFeeRate                   float32 `json:"minAppFeeRate"`
	EnableConsolidatedBilling       bool    `json:"enableConsolidatedBilling"`
	ShortName                       string  `json:"shortName"`
	EnablePurchaseOrder             bool    `json:"enablePurchaseOrder"`
	EnableEmailNotificationsToUsers bool    `json:"enableEmailNotificationsToUsers"`
	ParentTenantId                  int     `json:"parentTenantId"`
	ExternalId                      string  `json:"externalId"`
	DefaultActivationProfileId      string  `json:"defaultActivationProfileId"`
	EnableMonthlyBilling            bool    `json:"enableMonthlyBilling"`
	DefaultChargeType               string  `json:"defaultChargeType"`
	LoginLogo                       string  `json:"loginLogo"`
	HomePageLogo                    string  `json:"homePageLogo"`
	DomainName                      string  `json:"domainName"`
	//ActivationCodes                 string `json:"activationCodes"`
	//FirewallProfiles                string `json:"firewallProfiles"`
	SkipDefaultUserSecurityGroup bool         `json:"skipDefaultUserSecurityGroup"`
	DisableAllEmailNotification  bool         `json:"disableAllEmailNotification"`
	TrademarkURL                 string       `json:"trademarkURL"`
	Deleted                      bool         `json:"deleted"`
	Preferences                  []Preference `json:"preferences"`
}

type Preference struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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
