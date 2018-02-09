package cloudcenter

import "fmt"
import "net/http"

import "encoding/json"

//import "strconv"
//import "bytes"

type OperationStatus struct {
	Id       string `json:"id,omitempty"`
	Status   string `json:"status,omitempty"`
	Resource string `json:"resource,omitempty"`
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
