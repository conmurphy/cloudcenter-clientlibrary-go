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

//import "bytes"

type PhaseAPIResponse struct {
	Resource      *string `json:"resource,omitempty"`
	Size          *int64  `json:"size,omitempty"`
	PageNumber    *int64  `json:"pageNumber,omitempty"`
	TotalElements *int64  `json:"totalElements,omitempty"`
	TotalPages    *int64  `json:"totalPages"`
	Phases        []Phase `json:"phases,omitempty"`
}

type Phase struct {
	Id                     *string                `json:"id,omitempty"`
	ProjectId              *string                `json:"projectId,omitempty" validate:"nonzero"`
	Resource               *string                `json:"resource,omitempty"`
	Perms                  *[]string              `json:"perms,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	Order                  *float64               `json:"order,omitempty"`
	PhasePlan              *PhasePlan             `json:"phasePlan,omitempty"`
	PhaseBundles           *[]PhaseBundle         `json:"phaseBundles,omitempty"`
	PhaseCost              *PhaseCost             `json:"phaseCost,omitempty"`
	Deployments            *[]Deployment          `json:"deployments,omitempty"`
	DeploymentEnvironments *DeploymentEnvironment `json:"deploymentEnvironment,omitempty"`
}

type PhasePlan struct {
	Id       *string `json:"id,omitempty"`
	PlanName *string `json:"planName,omitempty"`
}

type PhaseBundle struct {
	Id    *string `json:"id,omitempty"`
	Name  *string `json:pName,omitempty"`
	Count *int64  `json:"count,omitempty"`
}

type PhaseCost struct {
	OriginalBalance  *float64 `json:"originalBalance,omitempty"`
	RemainingBalance *float64 `json:"remainingBalance,omitempty"`
	MeasurableUnit   *string  `json:"measurableUnit,omitempty"`
}

type DeploymentEnvironment struct {
	Id       *string   `json:"id,omitempty"`
	Resource *string   `json:"resource,omitempty"`
	Perms    *[]string `json:"perms,omitempty"`
	Name     *string   `json:"name,omitempty"`
}

type Deployment struct {
	Id                *string   `json:"id,omitempty"`
	Resource          *string   `json:"resource,omitempty"`
	Perms             *[]string `json:"perms,omitempty"`
	DeploymentName    *string   `json:"deploymentName,omitempty"`
	DeploymentOwnerId *string   `json:"deploymentOwnerId,omitempty"`
	DeploymentStatus  *string   `json:"deploymentStatus,omitempty"`
	AppName           *string   `json:"appName,omitempty"`
	AppVersion        *string   `json:"appVersion,omitempty"`
	AppLogoPath       *string   `json:"appLogoPath,omitempty"`
	SupportedActions  *[]string `json:"supportedActions,omitempty"`
}

func (s *Client) GetPhases(projectId int) ([]Phase, error) {

	var data PhaseAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + strconv.Itoa(projectId) + "/phases")

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

	phases := data.Phases
	return phases, nil
}

func (s *Client) GetPhase(projectId int, id int) (*Phase, error) {

	var data Phase

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + strconv.Itoa(projectId) + "/phases/" + strconv.Itoa(id))
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

	phase := &data
	return phase, nil
}

func (s *Client) AddPhase(phase *Phase) (*Phase, error) {

	var data Phase

	if errs := validator.Validate(phase); errs != nil {
		return nil, errs
	}

	phaseProjectID := *phase.ProjectId

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + phaseProjectID + "/phases")

	j, err := json.Marshal(phase)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	fmt.Println(string(bytes))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)

	if err != nil {
		return nil, err
	}

	phase = &data

	return phase, nil
}

func (s *Client) UpdatePhase(phase *Phase) (*Phase, error) {

	var data Phase

	if errs := validator.Validate(phase); errs != nil {
		return nil, errs
	}

	if nonzero(phase.Id) {
		return nil, errors.New("Phase.Id is missing")
	}

	phaseId := *phase.Id
	phaseProjectID := *phase.ProjectId

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + phaseProjectID + "/phases/" + phaseId)

	j, err := json.Marshal(phase)

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

	phase = &data

	return phase, nil
}

func (s *Client) DeletePhase(phaseProjectID int, phaseId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + strconv.Itoa(phaseProjectID) + "/phases/" + strconv.Itoa(phaseId))

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
