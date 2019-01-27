/*Copyright (c) 2019 Cisco and/or its affiliates.

This software is licensed to you under the terms of the Cisco Sample
Code License, Version 1.0 (the "License"). You may obtain a copy of the
License at

               https://developer.cisco.com/docs/licenses

All use of the material herein must be in accordance with the terms of
the License. All rights not expressly granted by the License are
reserved. Unless required by applicable law or agreed to separately in
writing, software distributed under the License is distributed on an "AS
IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
or implied.
*/

package cloudcenter

import "fmt"
import "net/http"

import "encoding/json"

//import "strconv"
//import "bytes"

type OperationStatus struct {
	OperationId          *string                `json:"operationId,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Status               *string                `json:"status,omitempty"`
	Resource             *string                `json:"resource,omitempty"`
	Msg                  *string                `json:"msg,omitempty"`
	Progress             *int64                 `json:"progress,omitempty"`
	AdditionalParameters *[]AdditionalParameter `json:"additionalParameters,omitempty"`
}

type AdditionalParameter struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
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
