package cloudcenter

import "fmt"
import "net/http"
import "strconv"

import "encoding/json"
import "errors"
import "bytes"
import "time"

type JobAPIResponse struct {
	Resource      *string `json:"resource"`
	Size          *int    `json:"size"`
	PageNumber    *int    `json:"pageNumber"`
	TotalElements *int    `json:"totalElements"`
	TotalPages    *int    `json:"totalPages"`
	Jobs          []Job   `json:"jobs"`
}

type Job struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
	//Perms                  []string              `json:"perms,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	Description            *string                `json:"description,omitempty"`
	Status                 *string                `json:"status,omitempty"`
	JobStatusMessage       *string                `json:"jobStatusMessage,omitempty"`
	Favorite               *bool                  `json:"favorite,omitempty"`
	ApprovalRequest        *ApprovalRequest       `json:"approvalRequest,omitempty"`
	ApprovalRequestAction  *string                `json:"approvalRequestAction,omitempty"`
	ApprovalRequestStatus  *string                `json:"approvalRequestStatus,omitempty"`
	StartTime              *string                `json:"startTime,omitempty"`
	EndTime                *string                `json:"endTime,omitempty"`
	FavoriteCreationTime   *string                `json:"favouriteCreationTime,omitempty"`
	CloudFamily            *string                `json:"cloudFamily,omitempty"`
	AgentUpgradeInProgress *bool                  `json:"agentUpgradeInProgress,omitempty"`
	DeploymentEnvironment  *DeploymentEnvironment `json:"deploymentEnvironment,omitempty"`
	Application            *Application           `json:"application,omitempty"`
	DeploymentEntity       *DeploymentEntity      `json:"deploymentEntity,omitempty"`
	Actions                *[]string              `json:"actions,omitempty"`
	TerminateProtection    *bool                  `json:"terminateProtection,omitempty"`
	Hidden                 *bool                  `json:"hidden,omitempty"`
	Benchmark              *bool                  `json:"benchmark,omitempty"`
	Owner                  *bool                  `json:"owner,omitempty"`
	OwnerEmailAddress      *string                `json:"ownerEmailAddress,omitempty"`
	PolicyIds              *[]string              `json:"policyIds,omitempty"`
	AppId                  *string                `json:"appId,omitempty"`
	EnvironmentId          *string                `json:"environmentId,omitempty,omitempty"`
	AppName                *string                `json:"appName,omitempty"`
	AppVersion             *string                `json:"appVersion,omitempty"`
	KeepExistingDeployment *bool                  `json:"keepExistingDeployment,omitempty"`
	TagIds                 *[]float64             `json:"tagIds,omitempty"`
	Tags                   *[]Tag                 `json:"tags,omitempty"`
	Parameters             *Parameter             `json:"parameters,omitempty"`
	Jobs                   *[]Jobs                `json:"jobs,omitempty"`
	CloudProperties        *[]CloudProperty       `json:"cloudProperties,omitempty"`
	ChildJobs              *[]ChildJob            `json:"childJobs,omitempty"`
	Metadata               *[]Metadata            `json:"metadata,omitempty"`
	LastUpdatedTime        *string                `json:"lastUpdatedTime,omitempty"`
	Scalable               *bool                  `json:"scalable,omitempty"`
	WindowsJob             *bool                  `json:"windowsJob,omitempty"`
	AccessLink             *string                `json:"accessLink,omitempty"`
	ParentJob              *ParentJob             `json:"parentJob,omitempty"`
	VirtualMachines        *[]VirtualMachine      `json:"virtualMachines,omitempty"`
	SecurityProfiles       *[]SecurityProfile     `json:"securityProfiles,omitempty"`
	BareMetalMachines      *[]BareMetalMachine    `json:"bareMetalMachines,omitempty"`
	TotalCost              *float64               `json:"totalCost,omitempty"`
	NodeHours              *float64               `json:"nodeHours,omitempty"`
}

type DeploymentEnvironment struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type Application struct {
	Id       *string `json:"id,omitempty"`
	Version  *string `json:"version,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type DeploymentEntity struct {
	Type *string `json:"type,omitempty"`
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Metadata struct {
	Namespace *string `json:"namespace,omitempty"`
	Name      *string `json:"name,omitempty"`
	Value     *string `json:"value,omitempty"`
	Editable  *bool   `json:"editable,omitempty"`
	Required  *bool   `json:"required,omitempty"`
}

type Jobs struct {
	TierId     *string   `json:"tierId,omitempty"`
	NodeIds    *string   `json:"nodeIds,omitempty"`
	Parameters Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	CloudParams CloudParam `json:"cloudParams,omitempty"`
	AppParams   []AppParam `json:"appParams,omitempty"`
	EnvParams   []EnvParam `json:"envParams,omitempty"`
}

type CloudParam struct {
	Cloud    *string `json:"cloud,omitempty"`
	Instance *string `json:"instance,omitempty"`
	//Storage         Storage         `json:"storage,omitempty"`
	RootVolumeSize  *string         `json:"rootVolumeSize,omitempty"`
	CloudProperties []CloudProperty `json:"cloudProperties,omitempty"`
}

type CloudProperty struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type AppParam struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type EnvParam struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type SecurityProfile struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type ChildJob struct {
	Id           *string      `json:"id,omitempty"`
	Resource     *string      `json:"resource,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

type Dependency struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type ParentJob struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type ApprovalRequest struct {
	Id                           *string                      `json:"id,omitempty"`
	Resource                     *string                      `json:"resource,omitempty"`
	CurrentDeploymentEnvironment CurrentDeploymentEnvironment `json:"currentDeploymentEnvironment,omitempty"`
	FromDeploymentEnvironment    FromDeploymentEnvironment    `json:"fromDeploymentEnvironment,omitempty"`
	InitiatingUser               InitiatingUser               `json:"initiatingUser,omitempty"`
	ApprovingUser                ApprovingUser                `json:"approvingUser,omitempty"`
	RequestTime                  *string                      `json:"requestTime,omitempty"`
	ApprovalTime                 *string                      `json:"approvalTime,omitempty"`
	Status                       *string                      `json:"status,omitempty"`
	Message                      *string                      `json:"message,omitempty"`
}

type CurrentDeploymentEnvironment struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type FromDeploymentEnvironment struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type InitiatingUser struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type ApprovingUser struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type Tag struct {
	Id       *string `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
}

type VirtualMachine struct {
	Id                    *string                `json:"id,omitempty"`
	VirtualMachineId      *string                `json:"virtualMachineId,omitempty"`
	PublicIpAddr          *string                `json:"publicIpAddr,omitempty"`
	PrivateIpAddr         *string                `json:"privateIpAddr,omitempty"`
	HostName              *string                `json:"hostName,omitempty"`
	Zone                  *string                `json:"zone,omitempty"`
	Status                *string                `json:"status,omitempty"`
	StartTime             *string                `json:"startTime,omitempty"`
	EndTime               *string                `json:"endTime,omitempty"`
	NodeNetworkInterfaces []NodeNetworkInterface `json:"nodeNetworkInterfaces,omitempty"`
	CostDetails           CostDetails            `json:"costDetails",omitempty"`
	TaskDetails           []TaskDetails          `json:"costDetails",omitempty"`
	AdditionalInfo        []AdditionalInfo       `json:"additionalInfo",omitempty"`
}

type NodeNetworkInterface struct {
	PublicIpAddr       *string `json:"publicIpAddr,omitempty"`
	PrivateIpAddr      *string `json:"privateIpAddr,omitempty"`
	NetworkDisplayName *string `json:"networkDisplayName,omitempty"`
	InterfaceIndex     *int64  `json:"interfaceIndex,omitempty"`
}

type CostDetails struct {
	NodeId                 *string  `json:"nodeId,omitempty"`
	NodeHour               *float64 `json:"nodeHour,omitempty"`
	MgmtBillNodeHour       *float64 `json:"mgmtBillNodeHour,omitempty"`
	CloudBillNodeTime      *float64 `json:"cloudBillNodeTime,omitempty"`
	CloudCostBillStartTime *string  `json:"cloudCostBillStartTime,omitempty"`
	CloudCostBillEndTime   *string  `json:"cloudCostBillEndTime,omitempty"`
	RecordTimestamp        *string  `json:"recordTimestamp,omitempty"`
	TotalCloudCost         *float64 `json:"totalCloudCost,omitempty"`
	TotalAppCost           *float64 `json:"totalAppCost,omitempty"`
	TotalJobsCost          *float64 `json:"totalJobsCost,omitempty"`
}

type TaskDetails struct {
	Id              *string  `json:"id,omitempty"`
	TaskId          *string  `json:"taskId,omitempty"`
	TaskName        *string  `json:"taskName,omitempty"`
	Status          *string  `json:"status,omitempty"`
	LastUpdatedTime *float64 `json:"lastUpdatedTime,omitempty"`
	NodeId          *string  `json:"nodeId,omitempty"`
	Msg             *string  `json:"msg,omitempty"`
}
type AdditionalInfo struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type BareMetalMachine struct {
	PublicIp    *string     `json:"publicIp,omitempty"`
	PrivateIp   *string     `json:"privateIp,omitempty"`
	HostName    *string     `json:"hostName,omitempty"`
	Status      *string     `json:"status,omitempty"`
	StartTime   *string     `json:"startTime,omitempty"`
	EndTime     *string     `json:"endTime,omitempty"`
	CostDetails CostDetails `json:"costDetails",omitempty"`
}

func (s *Client) GetJobs() ([]Job, error) {

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data JobAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	jobs := data.Jobs
	return jobs, nil
}

func (s *Client) GetJob(id int) (*Job, error) {

	var data Job

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs/" + strconv.Itoa(id))
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

	job := &data
	return job, nil
}

func (s *Client) GetJobByName(name string) ([]Job, error) {

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs?search=[deploymentEntity.name,eq," + name + "]")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data JobAPIResponse

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	jobs := data.Jobs
	return jobs, nil
}

func (s *Client) AddJobSync(job *Job, retrySeconds int) (*Job, error) {

	var data Job

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs")

	j, err := json.Marshal(job)

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
	} else {

		err = json.Unmarshal(bytes, &data)

		job := data

		tmpJobId := *job.Id

		jobId, err := strconv.Atoi(tmpJobId)

		if err != nil {
			return nil, err
		}

		tmpJob, err := s.GetJob(jobId)

		jobStatus := *tmpJob.Status

		if err != nil {
			return nil, err
		} else {

			for jobStatus == "JobStarting" || jobStatus == "JobSubmitted" || jobStatus == "JobInProgress" || jobStatus == "JobResuming" {

				tmpJob, err := s.GetJob(jobId)
				jobStatus = *tmpJob.Status

				if err != nil {
					return nil, err
				}

				time.Sleep(time.Duration(retrySeconds) * time.Second)

			}
		}

		if jobStatus != "JobCanceled" && jobStatus != "JobCancelling" && jobStatus != "JobError" && jobStatus != "JobStoppingError" && jobStatus != "JobRejected" && jobStatus != "JobSuspending" && jobStatus != "JobSuspended" {

			job, err := s.GetJob(jobId)

			if err != nil {
				return nil, err
			} else {

				return job, nil
			}
		} else {

			return nil, errors.New("Job deployment failed")

		}
	}

	return nil, errors.New("Job deployment failed")

}

func (s *Client) AddJobAsync(job *Job) (*Job, error) {

	var data Job

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs")

	j, err := json.Marshal(job)

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
	} else {

		err = json.Unmarshal(bytes, &data)

		if err != nil {
			return nil, err
		} else {
			return &data, nil
		}
	}

	return nil, errors.New("Job deployment failed")
}

func (s *Client) UpdateJobSync(job *Job, retrySeconds int) (*Job, error) {

	var data Job

	jobId := *job.Id

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs/" + jobId)

	j, err := json.Marshal(job)

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
	} else {

		err = json.Unmarshal(bytes, &data)

		job := data

		tmpJobId := *job.Id

		jobId, err := strconv.Atoi(tmpJobId)

		if err != nil {
			return nil, err
		}

		tmpJob, err := s.GetJob(jobId)
		jobStatus := *tmpJob.Status

		if err != nil {
			return nil, err
		} else {
			for jobStatus == "JobStarting" || jobStatus == "JobSubmitted" || jobStatus == "JobInProgress" || jobStatus == "JobResuming" {

				tmpJob, err := s.GetJob(jobId)
				jobStatus = *tmpJob.Status

				if err != nil {
					return nil, err
				}

				time.Sleep(time.Duration(retrySeconds) * time.Second)

			}
		}

		if jobStatus != "JobCanceled" && jobStatus != "JobCancelling" && jobStatus != "JobError" && jobStatus != "JobStoppingError" && jobStatus != "JobRejected" && jobStatus != "JobSuspending" && jobStatus != "JobSuspended" {

			job, err := s.GetJob(jobId)

			if err != nil {
				return nil, err
			} else {

				return job, nil
			}
		} else {

			return nil, errors.New("Job update failed")

		}
	}

	return nil, errors.New("Job update failed")

}

func (s *Client) UpdateJobAsync(job *Job) (*Job, error) {

	var data Job

	jobId := *job.Id

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs/" + jobId)

	j, err := json.Marshal(job)

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
	} else {

		err = json.Unmarshal(bytes, &data)

		if err != nil {
			return nil, err
		} else {
			return &data, nil
		}
	}

	return nil, errors.New("Job update failed")
}

func (s *Client) DeleteJobSync(jobId int) error {

	var operationStatus OperationStatus

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs/" + strconv.Itoa(jobId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	bytes, err := s.doRequest(req)

	err = json.Unmarshal(bytes, &operationStatus)

	status := &operationStatus

	operationId := *status.Id

	if err != nil {

		return err
	} else {

		status, err = s.GetOperationStatus(operationId)

		currentStatus := *status.Status
		for currentStatus == "RUNNING" {

			status, err = s.GetOperationStatus(operationId)
			currentStatus = *status.Status

		}

		if currentStatus == "SUCCESS" {
			return nil
		} else {
			return errors.New("Job deletion failed")
		}

	}

	return errors.New("Job deletion failed")

}

func (s *Client) DeleteJobAsync(jobId int) (*OperationStatus, error) {

	var operationStatus OperationStatus

	url := fmt.Sprintf(s.BaseURL + "/v2/jobs/" + strconv.Itoa(jobId))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)

	err = json.Unmarshal(bytes, &operationStatus)

	if err != nil {
		return nil, err
	}

	status := &operationStatus

	return status, nil
}
