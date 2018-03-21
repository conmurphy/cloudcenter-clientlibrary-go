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

type PlanAPIResponse struct {
	Resource      *string `json:"resource"`
	Size          *int64  `json:"size"`
	PageNumber    *int64  `json:"pageNumber"`
	TotalElements *int64  `json:"totalElements"`
	TotalPages    *int64  `json:"totalPages"`
	Plans         []Plan  `json:"plans"`
}

type Plan struct {
	Id                       *string   `json:"id,omitempty"`
	Resource                 *string   `json:"resource,omitempty"`
	Name                     *string   `json:"name,omitempty" validate:"nonzero"`
	Description              *string   `json:"description,omitempty"`
	Perms                    *[]string `json:"perms"`
	TenantId                 *string   `json:"tenantId,omitempty" validate:"nonzero"`
	Type                     *string   `json:"type,omitempty" validate:"nonzero"`
	MonthlyLimit             *int64    `json:"monthlyLimit,omitempty"`
	NodeHourIncrement        *float64  `json:"nodeHourIncrement,omitempty"`
	IncludedBundleId         *string   `json:"includedBundleId,omitempty"`
	Price                    *float64  `json:"price,omitempty"`
	OnetimeFee               *float64  `json:"onetimeFee,omitempty"`
	AnnualFee                *float64  `json:"annualFee,omitempty"`
	StorageRate              *float64  `json:"storageRate,omitempty"`
	HourlyRate               *float64  `json:"hourlyRate,omitempty"`
	OverageRate              *float64  `json:"overageRate,omitempty"`
	OverageLimit             *float64  `json:"overageLimit,omitempty"`
	RestrictedToAppStoreOnly *bool     `json:"restrictedToAppStoreOnly,omitempty"`
	BillToVendor             *bool     `json:"billToVendor,omitempty"`
	EnableRollover           *bool     `json:"enableRollover,omitempty"`
	Disabled                 *bool     `json:"disabled,omitempty"`
	ShowOnlyToAdmin          *bool     `json:"showOnlyToAdmin,omitempty"`
	NumberOfUsers            *int64    `json:"numberOfUsers,omitempty"`
	NumberOfProjects         *int64    `json:"numberOfProjects,omitempty"`
	PaymentProfileRequired   *bool     `json:"paymentProfileRequired,omitempty"`
}

func (s *Client) GetPlans(tenantId int) ([]Plan, error) {

	var data PlanAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/plans")

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

	plans := data.Plans
	return plans, nil
}

func (s *Client) GetPlan(tenantId int, planId int) (*Plan, error) {

	var data Plan

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/plans/" + strconv.Itoa(planId))
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

	plan := &data
	return plan, nil
}

func (s *Client) AddPlan(plan *Plan) (*Plan, error) {

	var data Plan

	if errs := validator.Validate(plan); errs != nil {
		return nil, errs
	}

	planTenantId := *plan.TenantId

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + planTenantId + "/plans")

	j, err := json.Marshal(plan)

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

	plan = &data

	return plan, nil
}

func (s *Client) UpdatePlan(plan *Plan) (*Plan, error) {

	var data Plan

	if errs := validator.Validate(plan); errs != nil {
		return nil, errs
	}

	if nonzero(plan.Id) {
		return nil, errors.New("Id is missing")
	}

	planTenantId := *plan.TenantId
	planId := *plan.Id

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + planTenantId + "/plans/" + planId)

	j, err := json.Marshal(plan)

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

	plan = &data

	return plan, nil
}

func (s *Client) DeletePlan(tenantId int, planId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/tenants/" + strconv.Itoa(tenantId) + "/plans/" + strconv.Itoa(planId))

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
