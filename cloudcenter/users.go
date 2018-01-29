package cloudcenter

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"

type UserAPIResponse struct {
	Resource      string `json:"resource"`
	Size          int    `json:"size"`
	PageNumber    int    `json:"pageNumber"`
	TotalElements int    `json:"totalElements"`
	TotalPages    int    `json:"totalPages"`
	Users         []User `json:"users"`
}

type User struct {
	Id                      string `json:"id"`
	Resource                string `json:"resource"`
	Username                string `json:"username"`
	Password                string `json:"password"`
	Enabled                 bool   `json:"enabled"`
	Type                    string `json:"type,omitempty"`
	FirstName               string `json:"firstName,omitempty"`
	LastName                string `json:"lastName,omitempty"`
	CompanyName             string `json:"companyName,omitempty"`
	EmailAddr               string `json:"emailAddr,omitempty"`
	EmailVerified           bool   `json:"emailVerified,omitempty"`
	PhoneNumber             string `json:"phoneNumber,omitempty"`
	ExternalId              string `json:"externalId,omitempty"`
	AccessKeys              string `json:"accessKeys,omitempty"`
	DisableReason           string `json:"disableReason,omitempty"`
	AccountSource           string `json:"accountSource,omitempty"`
	Status                  string `json:"status,omitempty"`
	Detail                  string `json:"detail,omitempty"`
	ActivationData          string `json:"activationData,omitempty"`
	Created                 int64  `json:"created,omitempty"`
	LastUpdated             int64  `json:"lastUpdated,omitempty"`
	CoAdmin                 bool   `json:"coAdmin,omitempty"`
	TenantAdmin             bool   `json:"tenantAdmin,omitempty"`
	ActivationProfileId     string `json:"activationProfileId,omitempty"`
	HasSubscriptionPlanType bool   `json:"hasSubscriptionPlanType,omitempty"`
}

func (s *Client) GetUsers() ([]User, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/users")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data UserAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	users := data.Users
	return users, nil
}

func (s *Client) GetUser(id int) (*User, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/users/" + strconv.Itoa(id))

	var data User

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

	user := &data

	return user, nil
}
