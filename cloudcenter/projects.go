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

type ProjectAPIResponse struct {
	Resource      *string   `json:"resource,omitempty"`
	Size          *int64    `json:"size,omitempty"`
	PageNumber    *int64    `json:"pageNumber,omitempty"`
	TotalElements *int64    `json:"totalElements,omitempty"`
	TotalPages    *int64    `json:"totalPages"`
	Projects      []Project `json:"projects,omitempty"`
}

type Project struct {
	Id             *string      `json:"id,omitempty"`
	Resource       *string      `json:"resource,omitempty"`
	Perms          *[]string    `json:"perms,omitempty"`
	Name           *string      `json:"name,omitempty"  validate:"nonzero"`
	Description    *string      `json:"description,omitempty"`
	ProjectOwnerId *int64       `json:"projectOwnerId,omitempty"`
	IsDraft        *bool        `json:"isDraft,omitempty"`
	TargetEndDate  *string      `json:"targetEndDate,omitempty"`
	NotifyUsers    *bool        `json:"notifyUsers,omitempty"`
	PlanType       *string      `json:"planType,omitempty"`
	Deleted        *bool        `json:"deleted,omitempty"`
	Quota          *Quota       `json:"quota,omitempty"`
	ProjectCost    *ProjectCost `json:"projectCost,omitempty"`
	Apps           *[]Apps      `json:"apps,omitempty"`
	Phases         *[]Phase     `json:"phases,omitempty"`
}

type Apps struct {
	Id       *string   `json:"id,omitempty"`
	Resource *string   `json:"resource,omitempty"`
	AppName  *string   `json:"appName,omitempty"`
	Perms    *[]string `json:"perms,omitempty"`
}

type ProjectCost struct {
	OriginalBalance  *float64 `json:"originalBalance,omitempty"`
	RemainingBalance *float64 `json:"remainingBalance,omitempty"`
	MeasurableUnit   *string  `json:"measurableUnit,omitempty"`
}

type Quota struct {
	Value          *float64 `json:"value,omitempty"`
	MeasurableUnit *string  `json:"measurableUnit,omitempty"`
	Type           *string  `json:"type,omitempty"`
}

func (s *Client) GetProjects() ([]Project, error) {

	var data ProjectAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/projects")

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

	projects := data.Projects
	return projects, nil
}

func (s *Client) GetProject(id int) (*Project, error) {

	var data Project

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + strconv.Itoa(id))
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

	project := &data
	return project, nil
}

func (s *Client) AddProject(project *Project) (*Project, error) {

	var data Project

	if errs := validator.Validate(project); errs != nil {
		return nil, errs
	}

	url := fmt.Sprintf(s.BaseURL + "/v1/projects")

	j, err := json.Marshal(project)

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

	project = &data

	return project, nil
}

func (s *Client) UpdateProject(project *Project) (*Project, error) {

	var data Project

	if errs := validator.Validate(project); errs != nil {
		return nil, errs
	}

	if nonzero(project.Id) {
		return nil, errors.New("Project.Id is missing")
	}

	projectId := *project.Id

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + projectId)

	j, err := json.Marshal(project)

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

	project = &data

	return project, nil
}

func (s *Client) DeleteProject(projectId int) error {

	url := fmt.Sprintf(s.BaseURL + "/v1/projects/" + strconv.Itoa(projectId))

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
