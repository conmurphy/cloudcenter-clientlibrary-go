package cloudcenter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	validator "gopkg.in/validator.v2"
)

//UserAPIResponse
type UserAPIResponse struct {
	Resource      string `json:"resource,omitempty"`
	Size          int    `json:"size,omitempty"`
	PageNumber    int    `json:"pageNumber,omitempty"`
	TotalElements int    `json:"totalElements,omitempty"`
	TotalPages    int    `json:"totalPages,omitempty"`
	Users         []User `json:"users,omitempty"`
}

type User struct {
	Id                      *string `json:"id,omitempty"`
	Resource                *string `json:"resource,omitempty"`
	Username                *string `json:"username,omitempty"`
	Password                *string `json:"password,omitempty" `
	Enabled                 *bool   `json:"enabled,omitempty"`
	Type                    *string `json:"type,omitempty"`
	FirstName               *string `json:"firstName,omitempty"`
	LastName                *string `json:"lastName,omitempty"`
	CompanyName             *string `json:"companyName,omitempty"`
	EmailAddr               *string `json:"emailAddr,omitempty" validate:"nonzero"`
	EmailVerified           *bool   `json:"emailVerified,omitempty"`
	PhoneNumber             *string `json:"phoneNumber,omitempty"`
	ExternalId              *string `json:"externalId,omitempty"`
	AccessKeys              *string `json:"accessKeys,omitempty"`
	DisableReason           *string `json:"disableReason,omitempty"`
	AccountSource           *string `json:"accountSource,omitempty"`
	Status                  *string `json:"status,omitempty"`
	Detail                  *string `json:"detail,omitempty"`
	ActivationData          *string `json:"activationData,omitempty"`
	Created                 *int64  `json:"created,omitempty"`
	LastUpdated             *int64  `json:"lastUpdated,omitempty"`
	CoAdmin                 *bool   `json:"coAdmin,omitempty"`
	TenantAdmin             *bool   `json:"tenantAdmin,omitempty"`
	ActivationProfileId     *string `json:"activationProfileId,omitempty"`
	HasSubscriptionPlanType *bool   `json:"hasSubscriptionPlanType,omitempty"`
	TenantId                *string `json:"tenantId,omitempty" validate:"nonzero"`
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

func (s *Client) GetUserFromEmail(emailToSearch string) (*User, error) {

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

	for _, user := range users {

		email := *user.EmailAddr

		if emailToSearch == email {
			return &user, nil
		}
	}

	return nil, errors.New("USER NOT FOUND")
}

func (s *Client) AddUser(user *User) (*User, error) {

	var data User

	if errs := validator.Validate(user); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/v1/users")

	j, err := json.Marshal(user)

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

	user = &data

	return user, nil
}

func (s *Client) UpdateUser(user *User) (*User, error) {

	var data User

	if errs := validator.Validate(user); errs != nil {
		return nil, errs
	}

	if nonzero(user.Id) {
		return nil, errors.New("User.Id is missing")
	}

	if nonzero(user.Username) {
		return nil, errors.New("User.Username is missing")
	}

	if nonzero(user.Type) {
		return nil, errors.New("User.Type is missing")
	}

	if nonzero(user.AccountSource) {
		return nil, errors.New("User.AccountSource is missing")
	}

	userId := *user.Id

	url := fmt.Sprintf(s.BaseURL + "/v1/users/" + userId)

	j, err := json.Marshal(user)

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

	user = &data

	return user, nil
}

func (s *Client) DeleteUser(userId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/users/" + strconv.Itoa(userId))

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

func (s *Client) DeleteUserByEmail(emailToSearch string) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/users")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return err
	}
	var data UserAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	users := data.Users

	for _, user := range users {

		email := *user.EmailAddr
		if email == emailToSearch {

			userId := *user.Id

			url := fmt.Sprintf(s.BaseURL + "/v1/users/" + userId)

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
	}

	return errors.New("USER NOT FOUND")
}
