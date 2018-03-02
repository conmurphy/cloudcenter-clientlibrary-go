package cloudcenter

import "fmt"
import "net/http"

import "encoding/json"

//import "strconv"
//import "bytes"

type OperationStatus struct {
	OperationId          string                `json:"operationId,omitempty"`
	Id                   string                `json:"id,omitempty"`
	Status               string                `json:"status,omitempty"`
	Resource             string                `json:"resource,omitempty"`
	Msg                  string                `json:"msg,omitempty"`
	Progress             int64                 `json:"progress,omitempty"`
	AdditionalParameters []AdditionalParameter `json:"additionalParameters,omitempty"`
}

type AdditionalParameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (s *Client) GetOperationStatus(operationId string) (*OperationStatus, error) {

	url := fmt.Sprintf(s.BaseURL + "/v1/operationStatus/" + operationId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data OperationStatus

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	operationStatus := data
	return &operationStatus, nil
}
