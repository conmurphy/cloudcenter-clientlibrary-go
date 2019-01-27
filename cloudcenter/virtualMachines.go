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
import "strconv"
import "encoding/json"

//import "bytes"

type VirtualMachineAPIResponse struct {
	CostSummary *CostSummary `json:"costSummary,omitempty"`
	Details     *Details     `json:"details,omitempty"`
	Filters     *Filters     `json:"filters,omitempty"`
}

type CostSummary struct {
	TotalNumberOfVMs          *int64   `json:"totalNumberOfVMs,omitempty"`
	TotalNumberOfRunningVMs   *int64   `json:"totalNumberOfRunningVMs,omitempty"`
	TotalCloudCost            *float64 `json:"totalCloudCost,omitempty"`
	EstimatedMonthlyCloudCost *float64 `json:"estimatedMonthlyCloudCost,omitempty"`
	TotalNodeHours            *float64 `json:"totalNodeHours,omitempty"`
}

type Details struct {
	Resource              *string                 `json:"resource,omitempty"`
	Size                  *int64                  `json:"size,omitempty"`
	PageNumber            *int64                  `json:"pageNumber,omitempty"`
	TotalElements         *int64                  `json:"totalElements,omitempty"`
	TotalPages            *int64                  `json:"totalPages"`
	VirtualMachineDetails []VirtualMachineDetails `json:"virtualMachineDetails,omitempty"`
}

type VirtualMachineDetails struct {
	Id                        *string           `json:"id,omitempty"`
	Resource                  *string           `json:"resource,omitempty"`
	Perms                     *[]string         `json:"perms,omitempty"`
	Name                      *string           `json:"name,omitempty"`
	Description               *string           `json:"description,omitempty"`
	Enabled                   *bool             `json:"enabled,omitempty"`
	Schedules                 *[]Schedule       `json:"schedules,omitempty"`
	BlockoutPeriods           *[]BlockoutPeriod `json:"blockoutPeriods,omitempty"`
	IsPolicyActiveOnResources *bool             `json:"isPolicyActiveOnResources,omitempty"`
	ResourcesMaps             *[]ResourcesMap   `json:"resourcesMaps,omitempty"`
	Priority                  *float64          `json:"priority,omitempty"`
	Created                   *float64          `json:"created,omitempty"`
	LastUpdated               *float64          `json:"lastUpdated,omitempty"`
	NodeId                    *string           `json:"nodeId,omitempty"`
	HostName                  *string           `json:"hostName,omitempty"`
	NodeStartTime             *int64            `json:"nodeStartTime,omitempty"`
	NodeEndTime               *int64            `json:"nodeEndTime,omitempty"`
	NumberOfCPUs              *int64            `json:"numberOfCpus,omitempty"`
	MemorySize                *int64            `json:"memorySize,omitempty"`
	StorageSize               *int64            `json:"storageSize,omitempty"`
	OSName                    *string           `json:"osName,omitempty"`
	CostPerHour               *float64          `json:"costPerHour,omitempty"`
	Status                    *string           `json:"status,omitempty"`
	ReviewNodeStatus          *bool             `json:"reviewNodeStatus,omitempty"`
	CloudFamily               *string           `json:"cloudFamily,omitempty"`
	CloudId                   *string           `json:"cloudId,omitempty"`
	CloudName                 *string           `json:"cloudName,omitempty"`
	CloudAccountId            *string           `json:"cloudAccountId,omitempty"`
	CloudAccountName          *string           `json:"cloudAccountName,omitempty"`
	RegionId                  *string           `json:"regionId,omitempty"`
	RegionName                *string           `json:"regionName,omitempty"`
	RegionDisplayName         *string           `json:"regionDisplayName,omitempty"`
	TenantId                  *string           `json:"tenantId,omitempty"`
	UserId                    *string           `json:"userId,omitempty"`
	FirstName                 *string           `json:"firstName,omitempty"`
	LastName                  *string           `json:"lastName,omitempty"`
	Email                     *string           `json:"email,omitempty"`
	InstanceTypeId            *string           `json:"instanceTypeId,omitempty"`
	InstanceTypeName          *string           `json:"instanceTypeName,omitempty"`
	InstanceCost              *float64          `json:"instanceCost,omitempty"`
	NICs                      *[]NIC            `json:"nics,omitempty"`
	Metatdata                 *[]Metadata       `json:"metadata,omitempty"`
	NodeProperties            *[]NodeProperty   `json:"nodeProperties,omitempty"`
	JobStartTime              *float64          `json:"jobStartTime,omitempty"`
	CloudNameAndAccountName   *string           `json:"cloudNameAndAccountName,omitempty"`
	AgentVersion              *string           `json:"agentVersion,omitempty"`
	JobId                     *string           `json:"jobId,omitempty"`
	JobName                   *string           `json:"jobName,omitempty"`
	JobEndTime                *float64          `json:"jobEndTime,omitempty"`
	ParentJobId               *string           `json:"parentJobId,omitempty"`
	ParentJobName             *string           `json:"parentJobName,omitempty"`
	ParentJobStatus           *string           `json:"parentJobStatus,omitempty"`
	BenchmarkId               *int64            `json:"benchmarkId,omitempty"`
	DeploymentEnvironmentId   *string           `json:"deploymentEnvironmentId,omitempty"`
	DeploymentEnvironmentName *string           `json:"deploymentEnvironmentName,omitempty"`
	AppId                     *string           `json:"appId,omitempty"`
	AppName                   *string           `json:"appName,omitempty"`
	AppVersion                *string           `json:"appVersion,omitempty"`
	AppLogoPath               *string           `json:"appLogoPath,omitempty"`
	ServiceId                 *string           `json:"serviceId,omitempty"`
	ServiceName               *string           `json:"serviceName,omitempty"`
	Tags                      *[]string         `json:"tags,omitempty"`
	PublicIpAddresses         *string           `json:"publicIpAddresses,omitempty"`
	PrivateIpAddresses        *string           `json:"privateIpAddresses,omitempty"`
	CloudCost                 *float64          `json:"cloudCost,omitempty"`
	NodeHours                 *float64          `json:"nodeHours,omitempty"`
	UserFavorite              *bool             `json:"userFavorite,omitempty"`
	RecordTimestamp           *float64          `json:"recordTimestamp,omitempty"`
	ImageId                   *string           `json:"imageId,omitempty"`
	TerminateProtection       *bool             `json:"terminateProtection,omitempty"`
	ImportedTime              *float64          `json:"importedTime,omitempty"`
	Running                   *bool             `json:"running,omitempty"`
	RunTime                   *float64          `json:"runTime,omitempty"`
	AgingPolicy               *string           `json:"agingPolicy,omitempty"`
	ScalingPolicy             *string           `json:"scalingPolicy,omitempty"`
	SecurityProfile           *string           `json:"securityProfile,omitempty"`
	Storage                   *string           `json:"storage,omitempty"`
	StorageIP                 *string           `json:"storageIp,omitempty"`
	NodeStatus                *string           `json:"nodeStatus,omitempty"`
	Type                      *string           `json:"type,omitempty"`
	Actions                   *[]Action         `json:"actions,omitempty"`
}

type Filters struct {
	CloudFamilies          *[]Filter `json:"cloudFamilies,omitempty"`
	Groups                 *[]Filter `json:"groups,omitempty"`
	Regions                *[]Filter `json:"regions,omitempty"`
	CloudAccounts          *[]Filter `json:"cloudAccounts,omitempty"`
	UserNames              *[]Filter `json:"userNames,omitempty"`
	OSNames                *[]Filter `json:"osNames,omitempty"`
	MemorySizes            *[]Filter `json:"memorySizes,omitempty"`
	CPUses                 *[]Filter `json:"cpuses,omitempty"`
	StorageSizes           *[]Filter `json:"storageSizes,omitempty"`
	AppNames               *[]Filter `json:"appNames,omitempty"`
	ServiceNames           *[]Filter `json:"serviceNames,omitempty"`
	DeploymentEnvironments *[]Filter `json:"deploymentEnvironments,omitempty"`
	Tags                   *[]Filter `json:"tags,omitempty"`
	Statuses               *[]Filter `json:"statuses,omitempty"`
	NodeStatuses           *[]Filter `json:"nodeStatuses,omitempty"`
	ParentJobStatuses      *[]Filter `json:"parentJobStatuses,omitempty"`
	CloudAndAccountNames   *[]Filter `json:"cloudAndAccountNames,omitempty"`
}

type Filter struct {
	DisplayName *string  `json:"displayName,omitempty"`
	Field       *string  `json:"field,omitempty"`
	Value       *float64 `json:"value,omitempty"`
}

type NIC struct {
	Name             *string `json:"name,omitempty"`
	PublicIpAddress  *string `json:"publicIpAddress,omitempty"`
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	Index            *int64  `json:"index,omitempty"`
}

type NodeProperty struct {
	Name  *string  `json:"name,omitempty"`
	Value *float64 `json:"value,omitempty"`
}

func (s *Client) GetVirtualMachines() ([]VirtualMachineDetails, error) {

	var data VirtualMachineAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/virtualMachines")

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

	virtualMachine := data.Details.VirtualMachineDetails

	return virtualMachine, nil
}

func (s *Client) GetVirtualMachine(virtualMachineId int) (*VirtualMachineDetails, error) {

	var data VirtualMachineDetails

	url := fmt.Sprintf(s.BaseURL + "/v1/virtualMachines/" + strconv.Itoa(virtualMachineId))
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

	virtualMachine := &data
	return virtualMachine, nil
}

func (s *Client) GetVirtualMachineCostSummary() (*CostSummary, error) {

	var data VirtualMachineAPIResponse

	url := fmt.Sprintf(s.BaseURL + "/v1/virtualMachines")

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

	virtualMachineAPIResponse := &data
	costSummary := virtualMachineAPIResponse.CostSummary
	return costSummary, nil
}
