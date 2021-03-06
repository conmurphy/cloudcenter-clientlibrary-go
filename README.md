# Cloudcenter Go Client Library

This is a Go Client Library used for accessing Cisco CloudCenter. 

It is currently a __Proof of Concept__ and has been developed and tested against Cisco CloudCenter 4.8.2 with Go version 1.9.3

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/overview.png)

Table of Contents
=================

  * [Cloudcenter Go Client Library](#cloudcenter-go-client-library)
      * [Quick Start](#quick-start)
      * [Quick Start - Creation from JSON file](#quick-start---creation-from-json-file)
      * [Helper Functions](#helper-functions)
         * [Without helper function](#without-helper-function)
         * [With helper function](#with-helper-function)
         * [Available Helper Functions](#available-helper-functions)
      * [Sync and Async](#sync-and-async)
      * [Reference](#reference)
         * [ActionPolicies](#actionpolicies)
         * [Actions](#actions)
         * [ActivationProfiles](#activationprofiles)
         * [AgingPolicies](#agingpolicies)
         * [Apps](#apps)
         * [Bundles](#bundles)
         * [CloudAccounts](#cloudaccounts)
         * [CloudImageMapping](#cloudimagemapping)
         * [CloudInstanceTypes](#cloudinstancetypes)
         * [CloudRegions](#cloudregions)
         * [CloudStorageTypes](#cloudstoragetypes)
         * [Clouds](#clouds)
         * [Contracts](#contracts)
         * [Environments](#environments)
         * [Groups](#groups)
         * [Images](#images)
         * [Jobs](#jobs)
         * [OperationStatus](#operationstatus)
         * [Phases](#phases)
         * [Plans](#plans)
         * [Projects](#projects)
         * [Roles](#roles)
         * [Services](#services)
         * [SuspensionPolicies](#suspensionpolicies)
         * [Tenants](#tenants)
         * [Users](#users)
         * [VirtualMachines](#virtualmachines)
      * [License](#license)

Created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc)

## Quick Start

```golang
package main
import "github.com/cloudcenter-clientlibrary-go/cloudcenter”

/*
	Define new cloudcenter client
*/

client := cloudcenter.NewClient("cliqradmin", ”myAPIKey", "https://ccm.cloudcenter-address.com")

/*
	Create user
*/

newUser := cloudcenter.User {
	FirstName:   cloudcenter.String("client"),
	LastName:    cloudcenter.String("library"),
	Password:    cloudcenter.String("myPassword"),
	EmailAddr:   cloudcenter.String("clientlibrary@cloudcenter-address.com"),
	CompanyName: cloudcenter.String("company"),
	TenantId:    cloudcenter.String("1"),
}

user, err := client.AddUser(&newUser)

if err != nil {
	fmt.Println(err)
} else {
	userId := *user.Id
	userEnabled := *user.Enabled
	fmt.Println(”New user created. \n UserId: " + user.Id + ", Enabled: " + strconv.FormatBool(userEnabled))
}
```

## Quick Start - Creation from JSON file

For some situations it may be easier to have the configuration represented as JSON rather than conifguring individually as per the  example above. In this scenario you can either build the JSON file yourself or monitor the API POST call for the JSON data sent to CloudCenter. This can be achieved using the browsers built in developer tools. See the following document for screenshots of how to find the POST call in the Chrome Developer Tools.

[Screenshots](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/README-DEVELOPER-TOOLS.md)


Example JSON File - newUser.json
```json
{
  "firstName": "Client",
  "lastName": "Library",
  "password": "myPassword",
  "emailAddr": "clientlibrary@cloudcenter-address.com",
  "companyName": "company",
  "tenantId": "1"
}
```

```golang
package main
import "github.com/cloudcenter-clientlibrary-go/cloudcenter”

/*
	Define new cloudcenter client
*/

client := cloudcenter.NewClient("cliqradmin", ”myAPIKey", "https://ccm.cloudcenter-address.com")

/*
	Create new user
*/

userJSONFile, err := os.Open("newUser.json")

if err != nil {
	fmt.Println(err)
}

bytes, _ := ioutil.ReadAll(userJSONFile)

var user *cloudcenter.User

json.Unmarshal(bytes, &user)

user, err = client.AddUser(user)

if err != nil {
	fmt.Println(err)
} else {
	userId := *user.Id
	userEnabled := *user.Enabled
	fmt.Println("Id: " + userId + ", Enabled: " + strconv.FormatBool(userEnabled))
}

defer userJSONFile.Close()
```

## Helper Functions

As per the following link, using the Marshal function from the encoding/json library treats false booleans as if they were nil values, and thus it omits them from the JSON response. To make a distinction between a non-existent boolean and false boolean we need to use a ```*bool``` in the struct. 

```golang
type User struct {
	Id                      *string `json:"id,omitempty"`
	FirstName               *string `json:"firstName,omitempty"`
	LastName                *string `json:"lastName,omitempty"`
	Password                *string `json:"password,omitempty"` 
	EmailAddr               *string `json:"emailAddr,omitempty"`
	Enabled                 *bool   `json:"enabled,omitempty"`
	TenantAdmin             *bool   `json:"tenantAdmin,omitempty"`
}
```
https://github.com/golang/go/issues/13284

Therefore in order to have a consistent experience all struct fields within this client library use pointers. This provides a way to differentiate between unset values, nil, and an intentional zero value, such as "", false, or 0. 

Helper functions have been created to simplify the creation of pointer types.

### Without helper function

```golang
firstName 	:= "client"
lastName 	:= "library"
password	:= "myPassword"
emailAddr	:= "clientlibrary@cloudcenter-address.com"
companyName	:= "company"
phoneNumber	:= "12345"
externalId	:= "23456"
tenantId	:= "1"


newUser := cloudcenter.User {
	FirstName:   &firstName,
	LastName:    &lastName,
	Password:    &password,
	EmailAddr:  &emailAddr,
	CompanyName: &companyName,
	PhoneNumber: &phoneNumber,
	ExternalId: &externalId,
	TenantId:    &tenantId,
}
```
### With helper function

```golang
newUser := cloudcenter.User {
	FirstName:   cloudcenter.String("client"),
	LastName:    cloudcenter.String("library"),
	Password:    cloudcenter.String("myPassword"),
	EmailAddr:   cloudcenter.String("clientlibrary@cloudcenter-address.com"),
	CompanyName: cloudcenter.String("company"),
	PhoneNumber: cloudcenter.String("12345"),
	ExternalId:  cloudcenter.String("23456"),
	TenantId:    cloudcenter.String("1"),
}
```

Reference: https://willnorris.com/2014/05/go-rest-apis-and-pointers

### Available Helper Functions

* cloudcenter.Bool()
* cloudcenter.Int()
* cloudcenter.Int64()
* cloudcenter.String()
* cloudcenter.Float32()
* cloudcenter.Float64()

## Sync and Async

* Synchronous APIs indicate that the program execution waits for a response to be returned by the API. The execution does not proceed until the call is completed. The real state of the API request is available in the response.

* Asynchronous APIs do not wait for the API call to complete. Program execution continues, and until the call completes,  you can issue GET requests to review the state after the submission, during the execution, and after the call completion

https://editor-docs.cloudcenter.cisco.com/display/40API/Synchronous+and+Asynchronous+APIs

Two options have been implemented in this library for each async API (AddCloudAccount used as example below):

*  __AddCloudAccountSync__: Client library will make an asynchronous call and wait until the task is complete. Once complete it will return either the newly created object or an error message.
*  __AddCloudAccountAsync__: Client library will make an asynchronous call and will return the operationStatus of the call. The client library user will be required to monitor the operation status and once successful retrieve the newly created object. 

## Reference

- [ActionPolicies](#actionpolicies)
- [Actions](#actions)
- [ActivationProfiles](#activationprofiles)
- [AgingPolicies](#agingpolicies)
- [Apps](#apps)
- [Bundles](#bundles)
- [Client](#client)
- [CloudAccounts](#cloudaccounts)
- [CloudImagemapping](#cloudimagemapping)
- [CloudInstancetypes](#cloudinstancetypes)
- [CloudRegions](#cloudregions)
- [CloudStoragetypes](#cloudstoragetypes)
- [Clouds](#clouds)
- [Contracts](#contracts)
- [Environments](#environments)
- [Groups](#groups)
- [Images](#images)
- [Jobs](#jobs)
- [OperationStatus](#operationstatus)
- [Phases](#phases)
- [Plans](#plans)
- [Projects](#projects)
- [Roles](#roles)
- [Services](#services)
- [SuspensionPolicies](#suspensionpolicies)
- [Tenants](#tenants)
- [Users](#users)
- [VirtualMachines](#virtualmachines)

### ActionPolicies

- [GetActionPolicies](#getactionpolicies)
- [GetActionPolicy](#getactionpolicy)
- [AddActionPolicy](#addactionpolicy)
- [UpdateActionPolicy](#updateactionpolicy)
- [DeleteActionPolicy](#deleteactionpolicy)

```go
type ActionPolicyAPIResponse struct {
	Resource       *string        
	Size           *int64         
	PageNumber     *int64         
	TotalElements  *int64         
	TotalPages     *int64         
	ActionPolicies []ActionPolicy 
}
```

```go
type ActionPolicy struct {
	Id          *string    
	Resource    *string    
	Perms       *[]string  
	Name        *string    
	Description *string    
	EntityType  *string    
	EventName   *string    
	Actions     *[]Actions 
	UserId      *string    
	Enabled     *bool      
	AutoEnable  *bool      
	ForceEnable *bool      
	Global      *bool      
}
```
#### GetActionPolicies

```go
func (s *Client) GetActionPolicies() ([]ActionPolicy, error)
```

##### Example

```go
actionPolicies, err := client.GetActionPolicies()

if err != nil {
	fmt.Println(err)
} else {
	for _, actionPolicy := range actionPolicies {

		fmt.Println("Id: " + actionPolicy.Id + ", Name: " + actionPolicy.Name)

	}
}
```

#### GetActionPolicy

```go
func (s *Client) GetActionPolicy(actionPolicyId int) (*ActionPolicy, error)
```

##### Example

```go
actionPolicy, err := client.GetActionPolicy(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Id: " + actionPolicy.Id + ", Name: " + actionPolicy.Name)
}
```

#### AddActionPolicy

```go
func (s *Client) AddActionPolicy(actionPolicy *ActionPolicy) (*ActionPolicy, error)
```

##### __Required Fields__
* Name
* EntityType
* EventName
* Actions


##### Example
```go
var actionInputs []cloudcenter.ActionInput

newActionInputs := cloudcenter.ActionInput{
	Name:  cloudcenter.String("toAddr"),
	Value: cloudcenter.String("%myEmail%"),
}

actionInputs = append(actionInputs, newActionInputs)

newActionInputs = cloudcenter.ActionInput{
	Name:  cloudcenter.String("bcc"),
	Value: cloudcenter.String(""),
}

actionInputs = append(actionInputs, newActionInputs)

newActionInputs = cloudcenter.ActionInput{
	Name:  cloudcenter.String("subject"),
	Value: cloudcenter.String("Deployment %jobName% for the application %appName% has reached maximum cluster size limit"),
}

actionInputs = append(actionInputs, newActionInputs)

newActionInputs = cloudcenter.ActionInput{
	Name:  cloudcenter.String("body"),
	Value: cloudcenter.String("Hello %firstName%\nDeployment %jobName% for the application %appName% has reached the maximum cluster size limit of %maxAppClusterSize% on %MaxClusterSizeReachedAt%. Click here %jobUrl% to view the status of the deployment.\nYour %vendorName% Team"),
}

actionInputs = append(actionInputs, newActionInputs)

var actions []cloudcenter.Actions

newAction := cloudcenter.Actions{
	ActionType:   cloudcenter.String("EMAIL"),
	ActionInputs: &actionInputs,
	}

actions = append(actions, newAction)

newActionPolicy := cloudcenter.ActionPolicy{
	Name:       cloudcenter.String("Client library action policy"),
	EntityType: cloudcenter.String("Application Deployment"),
	EventName:  cloudcenter.String("max_cluster_size_reached"),
	Actions:    &actions,
}

actionPolicy, err := client.AddActionPolicy(&newActionPolicy)

if err != nil {
	fmt.Println(err)
} else {
	actionPolicyId := *actionPolicy.Id
	actionPolicyName := *actionPolicy.Name
	fmt.Println("Id: " + actionPolicyId + ", Name: " + actionPolicyName)
}
```

#### UpdateActionPolicy

```go
func (s *Client) UpdateActionPolicy(actionPolicy *ActionPolicy) (*ActionPolicy, error)
```

##### __Required Fields__
* Id
* Name
* EntityType
* EventName
* Actions

##### Example
```go
var actionInputs []cloudcenter.ActionInput

newActionInputs := cloudcenter.ActionInput{
	Name:  cloudcenter.String("toAddr"),
	Value: cloudcenter.String("%myEmail%"),
}

actionInputs = append(actionInputs, newActionInputs)

newActionInputs = cloudcenter.ActionInput{
	Name:  cloudcenter.String("bcc"),
	Value: cloudcenter.String(""),
}

actionInputs = append(actionInputs, newActionInputs)

newActionInputs = cloudcenter.ActionInput{
	Name:  cloudcenter.String("subject"),
	Value: cloudcenter.String("Deployment %jobName% for the application %appName% has reached maximum cluster size limit"),
}

actionInputs = append(actionInputs, newActionInputs)

newActionInputs = cloudcenter.ActionInput{
	Name:  cloudcenter.String("body"),
	Value: cloudcenter.String("Hello %firstName%\nDeployment %jobName% for the application %appName% has reached the maximum cluster size limit of %maxAppClusterSize% on %MaxClusterSizeReachedAt%. Click here %jobUrl% to view the status of the deployment.\nYour %vendorName% Team"),
}

actionInputs = append(actionInputs, newActionInputs)

var actions []cloudcenter.Actions

newAction := cloudcenter.Actions{
	ActionType:   cloudcenter.String("EMAIL"),
	ActionInputs: &actionInputs,
	}

actions = append(actions, newAction)

newActionPolicy := cloudcenter.ActionPolicy{
	Id:       cloudcenter.String("2"),
	Name:       cloudcenter.String("Client library action policy"),
	EntityType: cloudcenter.String("Application Deployment"),
	EventName:  cloudcenter.String("max_cluster_size_reached"),
	Actions:    &actions,
}

actionPolicy, err := client.UpdateActionPolicy(&newActionPolicy)

if err != nil {
	fmt.Println(err)
} else {
	actionPolicyId := *actionPolicy.Id
	actionPolicyName := *actionPolicy.Name
	fmt.Println("Id: " + actionPolicyId + ", Name: " + actionPolicyName)
}
```

#### DeleteActionPolicy

```go
func (s *Client) DeleteActionPolicy(actionPolicyId int) error
```

##### Example
```go
err := client.DeleteActionPolicy(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Action policy deleted")
}
```

### Actions

- [GetActions](#getactions)
- [GetAction](#getaction)
- [AddAction](#addaction)
- [UpdateAction](#updateaction)
- [DeleteAction](#deleteaction)


```go
type ActionAPIResponse struct {
	Resource      *string  
	Size          *int64   
	PageNumber    *int64   
	TotalElements *int64   
	TotalPages    *int64   
	ActionJaxbs   []Action 
}
```

```go
type Action struct {
	Id                     *string                  
	Resource               *string                  
	Perms                  *[]string                
	Name                   *string                  
	Description            *string                  
	ActionType             *string                  
	LastUpdatedTime        *string                  
	TimeOut                *float64                 
	Enabled                *bool                    
	Encrypted              *bool                    
	Deleted                *bool                    
	SystemDefined          *bool                    
	BulkOperationSupported *bool                    
	IsAvailableToUser      *bool                    
	Owner                  *int64                   
	ActionParameters       *[]ActionParameter       
	ActionResourceMappings *[]ActionResourceMapping 
	ActionCustomParamSpecs *[]ActionCustomParamSpec 
}
```

```go
type ActionParameter struct {
	ParamName   *string 
	ParamValue  *string 
	CustomParam *bool   
	Required    *bool   
	Preference  *string 
}
```

```go
type ActionResourceMapping struct {
	Type                  *string                 
	ActionResourceFilters *[]ActionResourceFilter 
}
```

```go
type ActionCustomParamSpec struct {
	ParamName            *string              
	DisplayName          *string              
	HelpText             *string              
	Type                 *string              
	ValueList            *string              
	DefaultValue         *string              
	ConfirmValue         *string              
	PathSuffixValue      *string              
	UserVisible          *bool                
	UserEditable         *bool                
	SystemParam          *bool                
	ExampleValue         *string              
	DataUnit             *string              
	Optional             *bool                
	MultiselectSupported *bool                
	ValueConstraint      *ValueConstraint     
	Scope                *string              
	WebserviceListParams *WebserviceListParam 
	Preference           *string              
}
```

#### GetActions

```go
func (s *Client) GetActions() ([]Action, error)
```

##### Example

```go
actions, err := client.GetActions()

if err != nil {
	fmt.Println(err)
} else {
	for _, action := range actions {
		actionId := *action.Id
		actionName := *action.Name
		fmt.Println("Id: " + actionId + ", Name: " + actionName)
	}
}
```

#### GetAction

```go
func (s *Client) GetAction(id int) (*Action, error)
```

##### Example

```go
action, err := client.GetAction(1)

if err != nil {
	fmt.Println(err)
} else {
	actionId := *action.Id
	actionName := *action.Name
	fmt.Println("Id: " + actionId + ", Name: " + actionName)
}
```

#### AddAction

```go
func (s *Client) AddAction(action *Action) (*Action, error)
```

##### __Required Fields__
* Name
* ActionType 
* ActionParameters
* ActionResourceMappings


##### Example
```go
var actionParameters []cloudcenter.ActionParameter

newActionParameter := cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("bundlePath"),
	ParamValue:  cloudcenter.String("http://env.cliqrtech.com/Backup.zip"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("script"),
	ParamValue:  cloudcenter.String("backup.sh"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("downloadFromBundle"),
	ParamValue:  cloudcenter.String("true"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("executeOnContainer"),
	ParamValue:  cloudcenter.String(""),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("rebootInstance"),
	ParamValue:  cloudcenter.String("false"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("refreshInstanceInfo"),
	ParamValue:  cloudcenter.String("true"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

vmResource := cloudcenter.VmResource{
	Type:          cloudcenter.String("DEPLOYMENT_VM"),
	CloudRegions:  &[]string{"1"},
	CloudAccounts: &[]string{"1"},
}

var actionResourceFilters []cloudcenter.ActionResourceFilter

newActionResourceFilter := cloudcenter.ActionResourceFilter{
	VmResource: &vmResource,
}

actionResourceFilters = append(actionResourceFilters, newActionResourceFilter)

var actionResourceMappings []cloudcenter.ActionResourceMapping

newActionResourceMapping := cloudcenter.ActionResourceMapping{
	Type: cloudcenter.String("VIRTUAL_MACHINE"),
	ActionResourceFilters: &actionResourceFilters,
}

actionResourceMappings = append(actionResourceMappings, newActionResourceMapping)

newAction := cloudcenter.Action{
	Name:                   cloudcenter.String("ClientLibraryAction"),
	ActionType:             cloudcenter.String("EXECUTE_COMMAND"),
	ActionParameters:       &actionParameters,
	ActionResourceMappings: &actionResourceMappings,
}

action, err := client.AddAction(&newAction)

if err != nil {
	fmt.Println(err)
} else {
	actionId := *action.Id
	actionName := *action.Name
	fmt.Println("Id: " + actionId + ", Name: " + actionName)
}
```

#### UpdateAction

```go
func (s *Client) UpdateAction(action *Action) (*Action, error)
```

##### __Required Fields__
* Id
* Name
* ActionType 
* ActionParameters
* ActionResourceMappings

##### Example
```go
var actionParameters []cloudcenter.ActionParameter

newActionParameter := cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("bundlePath"),
	ParamValue:  cloudcenter.String("http://env.cliqrtech.com/Backup.zip"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("script"),
	ParamValue:  cloudcenter.String("backup.sh"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("downloadFromBundle"),
	ParamValue:  cloudcenter.String("true"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("executeOnContainer"),
	ParamValue:  cloudcenter.String(""),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("rebootInstance"),
	ParamValue:  cloudcenter.String("false"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

newActionParameter = cloudcenter.ActionParameter{
	ParamName:   cloudcenter.String("refreshInstanceInfo"),
	ParamValue:  cloudcenter.String("true"),
	CustomParam: cloudcenter.Bool(false),
	Required:    cloudcenter.Bool(true),
	Preference:  cloudcenter.String("VISIBLE_UNLOCKED"),
}

actionParameters = append(actionParameters, newActionParameter)

vmResource := cloudcenter.VmResource{
	Type:          cloudcenter.String("DEPLOYMENT_VM"),
	CloudRegions:  &[]string{"1"},
	CloudAccounts: &[]string{"1"},
}

var actionResourceFilters []cloudcenter.ActionResourceFilter

newActionResourceFilter := cloudcenter.ActionResourceFilter{
	VmResource: &vmResource,
}

actionResourceFilters = append(actionResourceFilters, newActionResourceFilter)

var actionResourceMappings []cloudcenter.ActionResourceMapping

newActionResourceMapping := cloudcenter.ActionResourceMapping{
	Type: cloudcenter.String("VIRTUAL_MACHINE"),
	ActionResourceFilters: &actionResourceFilters,
}

actionResourceMappings = append(actionResourceMappings, newActionResourceMapping)

newAction := cloudcenter.Action{
	Id:                   cloudcenter.String("14"),
	Name:                   cloudcenter.String("ClientLibraryAction"),
	ActionType:             cloudcenter.String("EXECUTE_COMMAND"),
	ActionParameters:       &actionParameters,
	ActionResourceMappings: &actionResourceMappings,
}

action, err := client.UpdateAction(&newAction)

if err != nil {
	fmt.Println(err)
} else {
	actionId := *action.Id
	actionName := *action.Name
	fmt.Println("Id: " + actionId + ", Name: " + actionName)
}
```

#### DeleteAction

```go
func (s *Client) DeleteAction(actionId int) error
```

##### Example
```go
err := client.DeleteAction(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Action deleted")
}
```

### ActivationProfiles

- [GetActivationProfiles](#getactivationprofiles)
- [GetActivationProfile](#getactivationprofile)
- [AddActivationProfile](#addactivationprofile)
- [UpdateActivationProfile](#updateactivationprofile)
- [DeleteActivationProfile](#deleteactivationprofile)

```go
type ActivationProfile struct {
	Id                  *string           
	Name                *string           
	Description         *string           
	Resource            *string           
	TenantId            *int64            
	PlanId              *string           
	BundleId            *string           
	ContractId          *string           
	DepEnvId            *string           
	ActivateRegions     *[]ActivateRegion 
	ImportApps          *[]string         
	AgreeToContract     *bool             
	SendActivationEmail *bool             
}
```

```go
type ActivateRegion struct {
	RegionId string 
}
```

#### GetActivationProfiles

```go
func (s *Client) GetActivationProfiles(tenantId int) ([]ActivationProfile, error)
```

##### Example

```go
activationProfiles, err := client.GetActivationProfiles(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, activationProfile := range activationProfiles {
		activationProfileId := *activationProfile.Id 
		activationProfileName := *activationProfile.Name
		fmt.Println("Id: " + activationProfileId + ", Name: " + activationProfileName)
	}
}
```

#### GetActivationProfile

```go
func (s *Client) GetActivationProfile(tenantId int, activationProfileId int) (*ActivationProfile, error)
```

##### Example

```go
activationProfile, err := client.GetActivationProfile(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	activationProfileId := *activationProfile.Id 
	activationProfileName := *activationProfile.Name
	fmt.Println("Id: " + activationProfileId + ", Name: " + activationProfileName)
}
```

#### AddActivationProfile

```go
func (s *Client) AddActivationProfile(activationProfile *ActivationProfile) (*ActivationProfile, error)
```

##### __Required Fields__
* TenantId
* Name

##### Example
```go
newActivationProfile := cloudcenter.ActivationProfile{
	TenantId: cloudcenter.Int64(1),
	Name:     cloudcenter.String("Client Library activation profile"),
	Description:     "Client Library activation profile description",
	PlanId:          "1",
	BundleId:        "1",
	ContractId:      "1",
	DepEnvId:        "1",
	ActivateRegions: activateRegions,
}

activationProfile, err := client.AddActivationProfile(&newActivationProfile)

if err != nil {
	fmt.Println(err)
} else {
	activationProfileId := *activationProfile.Id
	activationProfileName := *activationProfile.Name
	fmt.Println("Activation Profile Id: " + activationProfileId + ", Name: " + activationProfileName)
}
```

#### UpdateActivationProfile

```go
func (s *Client) UpdateActivationProfile(activationProfile *ActivationProfile) (*ActivationProfile, error)
```

##### __Required Fields__
* Id
* TenantId
* Name

##### Example
```go
newActivationProfile := cloudcenter.ActivationProfile{
	Id:       cloudcenter.String("1"),
	TenantId: cloudcenter.Int64(1),
	Name:     cloudcenter.String("Client Library activation profile"),
	Description:     "Client Library activation profile description",
	PlanId:          "1",
	BundleId:        "1",
	ContractId:      "1",
	DepEnvId:        "1",
	ActivateRegions: activateRegions,
}

activationProfile, err := client.UpdateActivationProfile(&newActivationProfile)

if err != nil {
	fmt.Println(err)
} else {
	activationProfileId := *activationProfile.Id
	activationProfileName := *activationProfile.Name
	fmt.Println("Activation Profile Id: " + activationProfileId + ", Name: " + activationProfileName)
}
```

#### DeleteActivationProfile

```go
func (s *Client) DeleteActivationProfile(tenantId int, activationProfileId int) error
```

##### Example
```go
err := client.DeleteActivationProfile(1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Activation profile deleted")
}
```

### AgingPolicies

- [GetAgingPolicies](#getagingpolicies)
- [GetAgingPolicy](#getagingpolicy)
- [AddAgingPolicy](#addagingpolicy)
- [UpdateAgingPolicy](#updateagingpolicy)
- [DeleteAgingPolicy](#deleteagingpolicy)

```go
type AgingPolicy struct {
	Id                             *string                
	Resource                       *string                
	Perms                          *[]string              
	Name                           *string                
	Description                    *string                
	Enabled                        *bool                  
	Type                           *string                
	Limit                          *Limit                 
	TerminateWhenPolicyEnds        *bool                  
	AllowGracePeriodForTermination *bool                  
	GraceLimit                     *GraceLimit            
	AllowPolicyExtension           *bool                  
	ExtensionLimit                 *ExtensionLimit        
	AllowGracePeriodNotification   *bool                  
	AllowPolicyExpiryNotification  *bool                  
	Notifications                  *[]Notification        
	IsPolicyActiveOnResources      *bool                  
	Created                        *float64               
	LastUpdated                    *float64               
	Resources                      *[]AgingPolicyResource 
	Priority                       *float64               
	OwnerId                        *int64                 
}
```

```go
type Notification struct {
	Template  *string     
	Type      *string     
	Enabled   *bool       
	Reminders *[]Reminder 
}
```

```go
type AgingPolicyResource struct {
	ResourceId                  *string  
	ResourceType                *string  
	AppliedDate                 *float64 
	ResourceStartTime           *float64 
	EstimatedPolicyEndTime      *float64 
	AllowedCost                 *float64 
	AccruedCost                 *float64 
	NumberOfExtensionsUsed      *int64   
	IsApprovalPending           *bool    
	IsPreviousExtensionDenied   *bool    
	IsPolicyReachingExpiry      *bool    
	IsPolicyReachingGraceExpiry *bool    
}
```

```go
type Reminder struct {
	Amount *float64 
	Unit   *string  
}
```

#### GetAgingPolicies

```go
func (s *Client) GetAgingPolicies() ([]AgingPolicy, error)
```

##### Example

```go
agingPolicies, err := client.GetAgingPolicies()

if err != nil {
	fmt.Println(err)
} else {
	for _, agingPolicy := range agingPolicies {
		agingPolicyId := *agingPolicy.Id 
		agingPolicyName := *agingPolicy.Name 
		fmt.Println("Id: " + agingPolicyId + ", Name: " + agingPolicyName)
	}
}
```

#### GetAgingPolicy

```go
func (s *Client) GetAgingPolicy(agingPolicyId int) (*AgingPolicy, error)
```

##### Example

```go
agingPolicy, err := client.GetAgingPolicies()

if err != nil {
	fmt.Println(err)
} else {
	agingPolicyId := *agingPolicy.Id 
	agingPolicyName := *agingPolicy.Name 
	fmt.Println("Id: " + agingPolicyId + ", Name: " + agingPolicyName)
}
```

#### AddAgingPolicy

```go
func (s *Client) AddAgingPolicy(agingPolicy *AgingPolicy) (*AgingPolicy, error)
```

##### __Required Fields__
* Name
* Enabled
* Type
* Limit



##### Example
```go
limit := cloudcenter.Limit{
	Unit:   cloudcenter.String("DAYS"),
	Amount: cloudcenter.Float64(1),
}

newAgingPolicy := cloudcenter.AgingPolicy{
	Name:    cloudcenter.String("Client Library AP"),
	Enabled: cloudcenter.Bool(true),
	Type:    cloudcenter.String("TIME"),
	Limit:   &limit,
}

agingPolicy, err := client.AddAgingPolicy(&newAgingPolicy)

if err != nil {
	fmt.Println(err)
} else {
	agingPolicyId := *agingPolicy.Id
	agingPolicyName := *agingPolicy.Name
	fmt.Println("Id: " + agingPolicyId + ", Name: " + agingPolicyName)
}
```

#### UpdateAgingPolicy

```go
func (s *Client) UpdateAgingPolicy(agingPolicy *AgingPolicy) (*AgingPolicy, error)
```

##### __Required Fields__
* Id
* Name
* Enabled
* Type
* Limit


##### Example
```go
limit := cloudcenter.Limit{
	Unit:   cloudcenter.String("DAYS"),
	Amount: cloudcenter.Float64(1),
}

newAgingPolicy := cloudcenter.AgingPolicy{
	Id:    cloudcenter.String("2"),
	Name:    cloudcenter.String("Client Library AP"),
	Enabled: cloudcenter.Bool(true),
	Type:    cloudcenter.String("TIME"),
	Limit:   &limit,
}

agingPolicy, err := client.UpdateAgingPolicy(&newAgingPolicy)

if err != nil {
	fmt.Println(err)
} else {
	agingPolicyId := *agingPolicy.Id
	agingPolicyName := *agingPolicy.Name
	fmt.Println("Id: " + agingPolicyId + ", Name: " + agingPolicyName)
}
```

#### DeleteAgingPolicy

```go
func (s *Client) DeleteAgingPolicy(agingPolicyId int) error
```

##### Example
```go
err := client.DeleteAgingPolicy(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Aging policy deleted")
}
```

### Apps

- [GetApps](#getapps)
- [GetApp](#getapp)
- [AddApp](#addapp)
- [UpdateApp](#updateapp)
- [DeleteApp](#deleteapp)

```go
type AppAPIResponse struct {
	Apps []App 
}
```

```go
type App struct {
	Id              *string           
	Resource        *string           
	Perms           *[]string         
	Name            *string           
	Description     *string           
	ServiceTierId   *string           
	Versions        *[]string         
	Version         *string           
	Executor        *string           
	Category        *string           
	ServiceTiers    *[]App            
	ProfileCategory *string           
	Service         *Service          
	Clusterable     *bool             
	HWProfile       *HWProfile        
	ParameterSpecs  *ParameterSpecs   
	Parameters      *Parameters       
	RevisionId      *int64            
	Metadatas       *[]Metadata       
	AppCategories   *[]AppCategory    
	LogoPath        *string           
	SupportedClouds *[]SupportedCloud 
}
```

```go
type HWProfile struct {
	MemorySize               *int64 
	NumOfCPUs                *int64 
	NetworkSpeed             *int64 
	NumOfNICs                *int64 
	LocalStorageCount        *int64 
	LocalStorageSize         *int64 
	CudaSupport              *bool  
	SSDSupport               *bool  
	SupportHardwareProvision *bool  
}
```

```go
type Metadata struct {
	Namespace *string 
	Name      *string 
	Value     *string 
	Editable  *bool   
	Required  *bool   
}
```

```go
type AppCategory struct {
	Id       *string   
	Resource *string   
	Perms    *[]string 
	Name     *string   
	Type     *string   
}
```

```go
type SupportedCloud struct {
	Id       *string 
	Resource *string 
}
```

```go
type ParameterSpecs struct {
	SystemParams *SystemParams 
	CustomParams *CustomParams 
	EnvVars      *EnvVar       
}
```

```go
type SystemParams struct {
	Params *[]Param 
	Size   *int64   
}
```

```go
type CustomParams struct {
	Params *[]Param 
	Size   *int64   
}
```

```go
type Param struct {
	ParamName            *string           
	DisplayName          *string           
	HelpText             *string           
	Type                 *string           
	ValueList            *string           
	DefaultValue         *string           
	ConfirmValue         *string           
	PathSuffixValue      *string           
	UserVisible          *bool             
	UserEditable         *bool             
	SystemParam          *bool             
	ExampleValue         *string           
	DataUnit             *string           
	Optional             *bool             
	MultiselectSupported *bool             
	ValueConstraint      *ValueConstraint  
	Scope                *string           
	WebserviceListParams *string           
	CollectionList       *[]CollectionList 
}
```

```go
type ValueConstraint struct {
	MinValue            *int64  
	MaxValue            *int64  
	MaxLength           *int64  
	Regex               *string 
	AllowSpaces         *bool   
	SizeValue           *int64  
	Step                *int64  
	CalloutWorkflowName *string 
}
```

```go
type CollectionList struct {
	ParamCollectionItem []ParamCollectionItem 
}
```

```go
type ParamCollectionItem struct {
	CollectionType         *string 
	CollectionName         *string 
	CollectionDisplayName  *string 
	CollectionValue        *string 
	CollectionDefaultValue *string 
	CollectionHelpText     *string 
	CollectionSampleText   *string 
	Optional               *bool   
}
```

#### GetApps

```go
func (s *Client) GetApps() ([]App, error)
```

##### Example

```go
apps, err := client.GetApps()

if err != nil {
	fmt.Println(err)
} else {
	for _, app := range apps {
		appId := *app.Id
		appName := *app.Name
		fmt.Println("Id: " + appId + ", Name: " + appName)
	}
}
```

#### GetApp

```go
func (s *Client) GetApp(appId int) (*App, error)
```

##### Example

```go
app, err := client.GetApp(760)

if err != nil {
	fmt.Println(err)
} else {
	appId := *app.Id
	appName := *app.Name
	fmt.Println("Id: " + appId + ", Name: " + appName)
}
```

#### ImportApp

```go
func (s *Client) ImportApp(filename string) error
```

##### __Required Fields__
* 


##### Example
```go

```

#### UpdateApp

```go
func (s *Client) UpdateApp(app *App) error
```

##### __Required Fields__
* Id
* Name
* Version - This should be an existing version


##### Example
```go
newApp := cloudcenter.App{
	Id:      cloudcenter.String("766"),
	Name:    cloudcenter.String("Ubuntu 14.04"),
	Version: cloudcenter.String("1"),
}

err := client.UpdateApp(&newApp)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("App updated")
}
```

#### DeleteApp

```go
func (s *Client) DeleteApp(appId int) error
```

##### Example
```go
err := client.DeleteApp(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("App deleted")
}
```

### Bundles

- [GetBundles](#getbundles)
- [GetBundle](#getbundle)
- [AddBundle](#addbundle)
- [UpdateBundle](#updatebundle)
- [DeleteBundle](#deletebundle)

```go
type BundleAPIResponse struct {
	Resource      *string  
	Size          *int64   
	PageNumber    *int64   
	TotalElements *int64   
	TotalPages    *int64   
	Bundles       []Bundle 
}
```

```go
type Bundle struct {
	Id               *string   
	Resource         *string   
	Perms            *[]string 
	Type             *string   
	Name             *string   
	Description      *string   
	Limit            *float64  
	Price            *float64  
	ExpirationDate   *float64  
	ExpirationMonths *int64    
	Disabled         *bool     
	ShowOnlyToAdmin  *bool     
	NumberOfUsers    *float64  
	TenantId         *string   
	PublishedAppIds  *[]string 
}
```

#### GetBundles

```go
func (s *Client) GetBundles(TenantId int) ([]Bundle, error)
```

##### Example

```go
bundles, err := client.GetBundles(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, bundle := range bundles {
		bundleId := *bundle.Id
		bundleName := *bundle.Name
		fmt.Println("Id: " + bundleId + ", Name: " + bundleName)
	}
}
```

#### GetBundle

```go
func (s *Client) GetBundle(TenantId int, BundleId int) (*Bundle, error)
```

##### Example

```go
bundle, err := client.GetBundle(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	bundleId := *bundle.Id
	bundleName := *bundle.Name
	fmt.Println("Id: " + bundleId + ", Name: " + bundleName)
}
```

#### GetBundleFromName

```go
func (s *Client) GetBundleFromName(TenantId int, BundleNameSearchString string) (*Bundle, error)
```

##### Example

```go
bundle, err := client.GetBundleFromName(1, "myBundle")

if err != nil {
	fmt.Println(err)
} else {
	bundleId := *bundle.Id
	bundleName := *bundle.Name
	fmt.Println("Id: " + bundleId + ", Name: " + bundleName)
}
```

#### AddBundle

```go
func (s *Client) AddBundle(bundle *Bundle) (*Bundle, error)
```

##### __Required Fields__
* TenantId
* Name
* Type
* ExpirationDate


##### Example
```go
newBundle := cloudcenter.Bundle{
	TenantId:       cloudcenter.String("1"),
	Name:           cloudcenter.String("clientlibraryBundle"),
	Type:           cloudcenter.String("BUDGET_BUNDLE"),
	ExpirationDate: cloudcenter.Float64(1580679359000),
}

bundle, err := client.AddBundle(&newBundle)

if err != nil {
	fmt.Println(err)
} else {
	bundleId := *bundle.Id
	bundleName := *bundle.Name
	fmt.Println("Id: " + bundleId + ", Name: " + bundleName)
}
```

#### UpdateBundle

```go
func (s *Client) UpdateBundle(bundle *Bundle) (*Bundle, error)
```

##### __Required Fields__
* Id
* TenantId
* Name
* Type
* ExpirationDate

##### Example
```go
newBundle := cloudcenter.Bundle{
	Id:       cloudcenter.String("1"),
	TenantId:       cloudcenter.String("1"),
	Name:           cloudcenter.String("clientlibraryBundle"),
	Type:           cloudcenter.String("BUDGET_BUNDLE"),
	ExpirationDate: cloudcenter.Float64(1580679359000),
}

bundle, err := client.UpdateBundle(&newBundle)

if err != nil {
	fmt.Println(err)
} else {
	bundleId := *bundle.Id
	bundleName := *bundle.Name
	fmt.Println("Id: " + bundleId + ", Name: " + bundleName)
}
```

#### DeleteBundle

```go
func (s *Client) DeleteBundle(tenantId int, bundleId int) error
```

##### Example
```go
err := client.DeleteBundle(1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Bundle deleted")
}
```

### CloudAccounts

- [GetCloudAccounts](#getcloudimageaccounts)
- [GetCloudAccount](#getcloudaccount)
- [AddCloudAccountAsync](#addcloudaccountasync)
- [AddCloudAccountSync](#addcloudaccountsync)
- [UpdateCloudAccount](#updatecloudaccount)
- [DeleteCloudAccount](#deletecloudaccount)

```go
type CloudAccountAPIResponse struct {
	Resource      *string        
	Size          *int64         
	PageNumber    *int64         
	TotalElements *int64         
	TotalPages    *int64         
	CloudAccounts []CloudAccount 
}
```

```go
type CloudAccount struct {
	Id                 *string            
	Resource           *string            
	Perms              *[]string          
	DisplayName        *string            
	CloudId            *string            
	UserId             *string            
	AccountId          *string            
	AccountName        *string            
	AccountPassword    *string            
	AccountDescription *string            
	ManageCost         *bool              
	PublicVisible      *bool              
	AllowedUsers       *[]int64           
	AccessPermission   *string            
	AccountProperties  *[]AccountProperty 
	TenantId           *string            
}
```

```go
type AccountProperty struct {
	Name  *string 
	Value *string 
}
```

#### GetCloudAccounts

```go
func (s *Client) GetCloudAccounts(tenantId int, cloudId int) ([]CloudAccount, error)
```

##### Example

```go
cloudAccounts, err := client.GetCloudAccounts(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	for _, cloudAccount := range cloudAccounts {
		cloudAccountId := *cloudAccount.Id
		cloudAccountDisplayName := *cloudAccount.DisplayName
		fmt.Println("Cloud Account Id: " + cloudAccountId + ", Name: " + cloudAccountDisplayName)
	}
}
```

#### GetCloudAccount

```go
func (s *Client) GetCloudAccount(tenantId int, cloudId int, accountId int) (*CloudAccount, error)
```

##### Example

```go
cloudAccount, err := client.GetCloudAccount(1, 1, 1)

if err != nil {
	fmt.Println(err)
} else {
	cloudAccountId := *cloudAccount.Id
	cloudAccountDisplayName := *cloudAccount.DisplayName
	fmt.Println("Cloud Account Id: " + cloudAccountId + ", Name: " + cloudAccountDisplayName)
}
```

#### GetCloudAccountByName

```go
func (s *Client) GetCloudAccountByName(tenantId int, cloudId int, displayName string) ([]CloudAccount, error)
```

##### Example

```go
cloudAccounts, err := client.GetCloudAccountByName(1, 1, "cloudAccountName")

if err != nil {
	fmt.Println(err)
} else {
	for _, cloudAccount := range cloudAccounts {
		cloudAccountId := *cloudAccount.Id
		cloudAccountDisplayName := *cloudAccount.DisplayName
		fmt.Println("Cloud Account Id: " + cloudAccountId + ", Name: " + cloudAccountDisplayName)
	}
}
```

#### AddCloudAccountAsync

```go
func (s *Client) AddCloudAccountAsync(cloudAccount *CloudAccount) (*OperationStatus, error)
```

##### __Required Fields__
* TenantId
* CloudId
* UserId
* AccountId
* AccountName
* DisplayName
* AccountPassword


##### Example
```go
newCloudAccount := cloudcenter.CloudAccount{
	TenantId:        cloudcenter.String("1"),
	CloudId:         cloudcenter.String("1"),
	UserId:          cloudcenter.String("2"),
	AccountId:       cloudcenter.String("myCloudAccountId"),
	AccountName:     cloudcenter.String("administrator@vsphere.local"),
	DisplayName:     cloudcenter.String("myCloudAccountName"),
	AccountPassword: cloudcenter.String("myPassword"),
}

operationStatus, err := client.AddCloudAccountAsync(&newCloudAccount)

if err != nil {
	fmt.Println(err)
} else {
	// Since this is an async call we will receive an operation status

	status := *operationStatus.Status
	operationStatusId := *operationStatus.OperationId

	fmt.Println("Operation Id: " + operationStatusId + ", Status: " + status)

	// We need to periodically check the status to find out if it is a success, failure, or still running

	for status == "RUNNING" {

		operationStatus, err := client.GetOperationStatus(operationStatusId)
		
		if err != nil {
			fmt.Println(err)
		}

		status = *operationStatus.Status

		// If it's still running it should have an operationId that we can use to update the status
		// If it's not running (i.e failed or success) it won't have an Id. We need this check
		// to ensure we don't have a "runtime error: invalid memory address or nil pointer dereference"
		if status == "RUNNING" {
			operationStatusId = *operationStatus.OperationId
		}

	}

	if status == "SUCCESS" {
		newCloudAccountDisplayName := *newCloudAccount.DisplayName
		cloudAccounts, err := client.GetCloudAccountByName(1, 1, newCloudAccountDisplayName)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, cloudAccount := range cloudAccounts {
				cloudAccountId := *cloudAccount.Id
				cloudAccountDisplayName := *cloudAccount.DisplayName
				fmt.Println("Cloud Account Id: " + cloudAccountId + ", Name: " + cloudAccountDisplayName)
			}
		}
	}

}
```
#### AddCloudAccountSync

```go
func (s *Client) AddCloudAccountSync(cloudAccount *CloudAccount) (*CloudAccount, error)
```

##### __Required Fields__
* TenantId
* CloudId
* UserId
* AccountId
* AccountName
* DisplayName
* AccountPassword

##### Example
```go
newCloudAccount := cloudcenter.CloudAccount{
	TenantId:        cloudcenter.String("1"),
	CloudId:         cloudcenter.String("1"),
	UserId:          cloudcenter.String("2"),
	AccountId:       cloudcenter.String("myCloudAccountId"),
	AccountName:     cloudcenter.String("administrator@vsphere.local"),
	DisplayName:     cloudcenter.String("myCloudAccountName"),
	AccountPassword: cloudcenter.String("myPassword"),
}

cloudAccount, err := client.AddCloudAccountSync(&newCloudAccount)

if err != nil {
	fmt.Println(err)
} else {
	cloudAccountId := *cloudAccount.Id
	cloudAccountDisplayName := *cloudAccount.DisplayName
	fmt.Println("Cloud Account Id: " + cloudAccountId + ", Name: " + cloudAccountDisplayName)
}
```

#### UpdateCloudAccountAsync

```go
func (s *Client) UpdateCloudAccountAsync(cloudAccount *CloudAccount) (*OperationStatus, error)
```

##### __Required Fields__
* Id
* TenantId
* CloudId
* UserId
* AccountId
* AccountName
* DisplayName
* AccountPassword

##### Example
```go
newCloudAccount := cloudcenter.CloudAccount{
	Id:        cloudcenter.String("19"),
	TenantId:        cloudcenter.String("1"),
	CloudId:         cloudcenter.String("1"),
	UserId:          cloudcenter.String("2"),
	AccountId:       cloudcenter.String("myCloudAccountId"),
	AccountName:     cloudcenter.String("administrator@vsphere.local"),
	DisplayName:     cloudcenter.String("myCloudAccountName"),
	AccountPassword: cloudcenter.String("myPassword"),
}

operationStatus, err := client.UpdateCloudAccountAsync(&newCloudAccount)

if err != nil {
	fmt.Println(err)
} else {
	// Since this is an async call we will receive an operation status

	status := *operationStatus.Status
	operationStatusId := *operationStatus.OperationId

	fmt.Println("Operation Id: " + operationStatusId + ", Status: " + status)

	// We need to periodically check the status to find out if it is a success, failure, or still running

	for status == "RUNNING" {

		operationStatus, err := client.GetOperationStatus(operationStatusId)
		
		if err != nil {
			fmt.Println(err)
		}

		status = *operationStatus.Status

		// If it's still running it should have an operationId that we can use to update the status
		// If it's not running (i.e failed or success) it won't have an Id. We need this check
		// to ensure we don't have a "runtime error: invalid memory address or nil pointer dereference"
		if status == "RUNNING" {
			operationStatusId = *operationStatus.OperationId
		}

	}

	if status == "SUCCESS" {
		newCloudAccountDisplayName := *newCloudAccount.DisplayName
		cloudAccounts, err := client.GetCloudAccountByName(1, 1, newCloudAccountDisplayName)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, cloudAccount := range cloudAccounts {
				cloudAccountId := *cloudAccount.Id
				cloudAccountDisplayName := *cloudAccount.DisplayName
				fmt.Println("Cloud Account Id: " + cloudAccountId + ", Name: " + cloudAccountDisplayName)
			}
		}
	}

}
```

#### UpdateCloudAccountSync

```go
func (s *Client) UpdateCloudAccountSync(cloudAccount *CloudAccount) (*CloudAccount, error)
```

##### __Required Fields__
* Id
* TenantId
* CloudId
* UserId
* AccountId
* AccountName
* DisplayName
* AccountPassword

##### Example
```go
newCloudAccount := cloudcenter.CloudAccount{
	Id: 		 cloudcenter.String("5"),
	TenantId:        cloudcenter.String("1"),
	CloudId:         cloudcenter.String("1"),
	UserId:          cloudcenter.String("2"),
	AccountId:       cloudcenter.String("myCloudAccountId"),
	AccountName:     cloudcenter.String("administrator@vsphere.local"),
	DisplayName:     cloudcenter.String("myCloudAccountName"),
	AccountPassword: cloudcenter.String("myPassword"),
}

cloudAccount, err := client.UpdateCloudAccountSync(&newCloudAccount)

if err != nil {
	fmt.Println(err)
} else {
	cloudAccountId := *cloudAccount.Id
	cloudAccountDisplayName := *cloudAccount.DisplayName
	fmt.Println("Cloud Account Id: " + cloudAccountId + ", Name: " + cloudAccountDisplayName)
}
```
```

#### DeleteCloudAccount

```go
func (s *Client) DeleteCloudAccount(tenantId int, cloudId int, accountId int) error
```

##### Example
```go
err := client.DeleteCloudAccount(1,1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Cloud account deleted")
}
```

### CloudImageMapping

- [GetCloudImageMappings](#getcloudimagemappings)
- [GetCloudImageMapping](#getcloudimagemapping)
- [AddCloudImageMapping](#addcloudimagemapping)
- [UpdateCloudImageMapping](#updatecloudimagemapping)
- [DeleteCloudImageMapping](#deletecloudimagemapping)

```go
type CloudImageMappingAPIResponse struct {
	Resource           *string             
	Size               *int64              
	PageNumber         *int64              
	TotalElements      *int64              
	TotalPages         *int64              
	CloudImageMappings []CloudImageMapping 
}
```

```go
type CloudImageMapping struct {
	Id                   *string     
	Resource             *string     
	Perms                *[]string   
	TenantId             *string     
	CloudId              *string     
	CloudRegionId        *string     
	RegionId             *string     
	CloudImageId         *string     
	CloudProviderImageId *string     
	LaunchUserName       *string     
	ImageId              *string     
	GrantAndRevoke       *bool       
	ImageCloudAccountId  *int64      
	Resources            *[]Resource 
	Mappings             *[]Mapping  
}
```

#### GetCloudImageMappings

```go
func (s *Client) GetCloudImageMappings(tenantId int, cloudId int, regionId int) ([]CloudImageMapping, error)
```

##### Example

```go
cloudImages, err := client.GetCloudImageMappings(1, 1, 1)

if err != nil {
	fmt.Println(err)
} else {
	for _, cloudImage := range cloudImages {
		cloudImageMappingId := *cloudImageMapping.Id
		cloudImageMappingResource := *cloudImageMapping.Resource
		fmt.Println("Id: " + cloudImageMappingId + ", Name: " + cloudImageMappingResource)
	}
}
```

#### GetCloudImageMapping

```go
func (s *Client) GetCloudImageMapping(tenantId int, cloudId int, regionId int, imageId int) (*CloudImageMapping, error)
```

##### Example

```go
cloudImage, err := client.GetCloudImageMapping(1, 1, 1,1)

if err != nil {
	fmt.Println(err)
} else {
	cloudImageMappingId := *cloudImageMapping.Id
	cloudImageMappingResource := *cloudImageMapping.Resource
	fmt.Println("Id: " + cloudImageMappingId + ", Name: " + cloudImageMappingResource)
}
```

#### AddCloudImageMapping

```go
func (s *Client) AddCloudImageMapping(cloudImage *CloudImageMapping) (*CloudImageMapping, error)
```

##### __Required Fields__
* TenantId
* CloudId
* RegionId
* CloudRegionId
* CloudProviderImageId
* ImageId


##### Example
```go
newCloudImageMapping := cloudcenter.CloudImageMapping{
	TenantId:             cloudcenter.String("1"),
	CloudId:              cloudcenter.String("1"),
	RegionId:             cloudcenter.String("1"),
	CloudRegionId:        cloudcenter.String("1"),
	CloudProviderImageId: cloudcenter.String("3a55133-749f-4cde-b80e-332781ae9b99"),
	ImageId:              cloudcenter.String("2"),
}

cloudImageMapping, err := client.AddCloudImageMapping(&newCloudImageMapping)

if err != nil {
	fmt.Println(err)
} else {
	cloudImageMappingId := *cloudImageMapping.Id
	cloudImageMappingResource := *cloudImageMapping.Resource
	fmt.Println("Id: " + cloudImageMappingId + ", Name: " + cloudImageMappingResource)
}
```

#### UpdateCloudImageMapping

```go
func (s *Client) UpdateCloudImageMapping(cloudImage *CloudImageMapping) (*CloudImageMapping, error)
```

##### __Required Fields__
* Id
* TenantId
* CloudId
* RegionId
* CloudRegionId
* CloudProviderImageId
* ImageId

##### Example
```go
newCloudImageMapping := cloudcenter.CloudImageMapping{
	Id:             cloudcenter.String("14"),
	TenantId:             cloudcenter.String("1"),
	CloudId:              cloudcenter.String("1"),
	RegionId:             cloudcenter.String("1"),
	CloudRegionId:        cloudcenter.String("1"),
	CloudProviderImageId: cloudcenter.String("3a55133-749f-4cde-b80e-332781ae9b99"),
	ImageId:              cloudcenter.String("2"),
}

cloudImageMapping, err := client.UpdateCloudImageMapping(&newCloudImageMapping)

if err != nil {
	fmt.Println(err)
} else {
	cloudImageMappingId := *cloudImageMapping.Id
	cloudImageMappingResource := *cloudImageMapping.Resource
	fmt.Println("Id: " + cloudImageMappingId + ", Name: " + cloudImageMappingResource)
}
```

#### DeleteCloudImageMapping

```go
func (s *Client) DeleteCloudImageMapping(tenantId int, cloudId int, regionId int, imageId int) error
```

##### Example
```go
err := client.DeleteCloudInstanceType(1,1,1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Cloud instance type deleted")
}
```

### CloudInstanceTypes

- [GetCloudInstanceTypes](#getcloudinstancetype)
- [GetCloudInstanceType](#getcloudinstancetype)
- [AddCloudInstanceType](#addcloudinstancetype)
- [UpdateCloudInstanceType](#updatecloudinstancetype)
- [DeleteCloudInstanceType](#deletecloudinstancetype)

```go
type CloudInstanceTypeAPIResponse struct {
	Resource           *string             
	Size               *int64              
	PageNumber         *int64              
	TotalElements      *int64              
	TotalPages         *int64              
	CloudInstanceTypes []CloudInstanceType 
}
```

```go
type CloudInstanceType struct {
	Id                        *string   
	Resource                  *string   
	Perms                     *[]string 
	Name                      *string   
	Description               *string   
	Type                      *string   
	TenantId                  *string   
	CloudId                   *string   
	RegionId                  *string   
	CostPerHour               *float64  
	MemorySize                *int64    
	NumOfCPUs                 *int64    
	NumOfNICs                 *int64    
	LocalStorageSize          *int64    
	SupportsSSD               *bool     
	SupportsCUDA              *bool     
	Supports32Bit             *bool     
	Supports64Bit             *bool     
	LocalStorageCount         *float64  
	SupportsHardwareProvision *bool     
}
```

#### GetCloudInstanceTypes

```go
func (s *Client) GetCloudInstanceTypes(tenantId int, cloudId int, regionId int) ([]CloudInstanceType, error)
```

##### Example

```go
cloudInstanceTypes, err := client.GetCloudInstanceTypes(1, 1, 1)

if err != nil {
	fmt.Println(err)
} else {
	for _, cloudInstanceType := range cloudInstanceTypes {
		cloudInstanceTypeId := *cloudInstanceType.Id
		cloudInstanceTypeName := *cloudInstanceType.Name
		fmt.Println("Id: " + cloudInstanceTypeId + ", Name: " + cloudInstanceTypeName)
	}
}
```

#### GetCloudInstanceType

```go
func (s *Client) GetCloudInstanceType(tenantId int, cloudId int, regionId int, instanceId int) (*CloudInstanceType, error)
```

##### Example

```go
cloudInstanceType, err := client.GetCloudInstanceType(1, 1, 1,1)

if err != nil {
	fmt.Println(err)
} else {
	cloudInstanceTypeId := *cloudInstanceType.Id
	cloudInstanceTypeName := *cloudInstanceType.Name
	fmt.Println("Id: " + cloudInstanceTypeId + ", Name: " + cloudInstanceTypeName)
}
```

#### AddCloudInstanceType

```go
func (s *Client) AddCloudInstanceType(cloudInstanceType *CloudInstanceType) (*CloudInstanceType, error)
```

##### __Required Fields__
* TenantId
* CloudId
* RegionId
* Name
* NumOfCPUs
* NumOfNICs
* MemorySize
* Supports32Bit
* Supports64Bit
* Type


##### Example
```go
newCloudInstanceType := cloudcenter.CloudInstanceType{
	TenantId:      cloudcenter.String("1"),
	CloudId:       cloudcenter.String("1"),
	RegionId:      cloudcenter.String("1"),
	Name:          cloudcenter.String("m1.medium.m1"),
	NumOfCPUs:     cloudcenter.Int64(1),
	NumOfNICs:     cloudcenter.Int64(2),
	MemorySize:    cloudcenter.Int64(1024),
	Supports32Bit: cloudcenter.Bool(true),
	Supports64Bit: cloudcenter.Bool(true),
	Type:          cloudcenter.String("m1.medium.db"),
}

cloudInstanceType, err := client.AddCloudInstanceType(&newCloudInstanceType)

if err != nil {
	fmt.Println(err)
} else {
	cloudInstanceTypeId := *cloudInstanceType.Id
	cloudInstanceTypeName := *cloudInstanceType.Name
	fmt.Println("Id: " + cloudInstanceTypeId + ", Name: " + cloudInstanceTypeName)
}
```

#### UpdateCloudInstanceType

```go
func (s *Client) UpdateCloudInstanceType(cloudInstanceType *CloudInstanceType) (*CloudInstanceType, error)
```

##### __Required Fields__
* Id
* TenantId
* CloudId
* RegionId
* Name
* NumOfCPUs
* NumOfNICs
* MemorySize
* Supports32Bit
* Supports64Bit
* Type

##### Example
```go
newCloudInstanceType := cloudcenter.CloudInstanceType{
	Id:      cloudcenter.String("5"),
	TenantId:      cloudcenter.String("1"),
	CloudId:       cloudcenter.String("1"),
	RegionId:      cloudcenter.String("1"),
	Name:          cloudcenter.String("m1.medium.m1"),
	NumOfCPUs:     cloudcenter.Int64(1),
	NumOfNICs:     cloudcenter.Int64(2),
	MemorySize:    cloudcenter.Int64(1024),
	Supports32Bit: cloudcenter.Bool(true),
	Supports64Bit: cloudcenter.Bool(true),
	Type:          cloudcenter.String("m1.medium.db"),
}

cloudInstanceType, err := client.UpdateCloudInstanceType(&newCloudInstanceType)

if err != nil {
	fmt.Println(err)
} else {
	cloudInstanceTypeId := *cloudInstanceType.Id
	cloudInstanceTypeName := *cloudInstanceType.Name
	fmt.Println("Id: " + cloudInstanceTypeId + ", Name: " + cloudInstanceTypeName)
}
```

#### DeleteCloudInstanceType

```go
func (s *Client) DeleteCloudInstanceType(tenantId int, cloudId int, regionId int, instanceId int) error
```

##### Example
```go
err := client.DeleteCloudInstanceType(1,1,1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Cloud instance type deleted")
}
```

### CloudRegions

- [GetCloudRegions](#getcloudregions)
- [GetCloudRegion](#getcloudregion)
- [AddCloudRegion](#addcloudregion)
- [UpdateCloudRegion](#updatecloudregion)
- [DeleteCloudRegion](#deletecloudregion)


```go
type CloudRegionAPIResponse struct {
	Resource      *string       
	Size          *int64        
	PageNumber    *int64        
	TotalElements *int64        
	TotalPages    *int64        
	CloudRegions  []CloudRegion 
}
```

```go
type CloudRegion struct {
	Id                     *string           
	TenantId               *string           
	CloudId                *string           
	CloudRegionId          *string           
	ImportRegion           *ImportRegion     
	Resource               *string           
	Perms                  *[]string         
	DisplayName            *string           
	RegionName             *string           
	Description            *string           
	Gateway                *Gateway          
	Storage                *Storage          
	Enabled                *bool             
	Activated              *bool             
	PublicCloud            *bool             
	NumUsers               *int32            
	Status                 *string           
	StatusDetail           *string           
	RegionProperties       *[]RegionProperty 
	ExternalBundleLocation *string           
	ExternalActions        *[]ExternalAction 
}
```

```go
type ImportRegion struct {
	Name        *string 
	DisplayName *string 
}
```

```go
type Gateway struct {
	Address        *string 
	DNSName        *string 
	Status         *string 
	CloudId        *string 
	CloudAccountId *string 
}
```

```go
type Storage struct {
	RegionId              *string                 
	CloudAccountId        *string                 
	Size                  *int64                  
	NumNodes              *int64                  
	CloudSpecificSettings *[]CloudSpecificSetting 
	Address               *string                 
}
```

```go
type CloudSpecificSetting struct {
	Name  *string 
	Value *string 
}
```

```go
type RegionProperty struct {
	Name  *string 
	Value *string 
}
```

```go
type ExternalAction struct {
	ActionName  *string 
	ActionType  *string 
	ActionValue *string 
}
```

#### GetCloudRegions

```go
func (s *Client) GetCloudRegions(tenantId int, cloudId int) ([]CloudRegion, error)

```

##### Example

```go
cloudRegions, err := client.GetCloudRegions(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	for _, cloudRegion := range cloudRegions {
		cloudRegionId := *cloudRegion.Id
		cloudRegionDisplayName := *cloudRegion.DisplayName
		fmt.Println("Cloud Region Id: " + cloudRegionId + ", DisplayName: " + cloudRegionDisplayName)
	}
}
```

#### GetCloudRegion

```go
func (s *Client) GetCloudRegion(tenantId int, cloudId int, regionId int) (*CloudRegion, error)
```

##### Example

```go
cloudRegion, err := client.GetCloudRegion(1, 1, 1)

if err != nil {
	fmt.Println(err)
} else {
	cloudRegionId := *cloudRegion.Id
	cloudRegionDisplayName := *cloudRegion.DisplayName
	fmt.Println("Cloud Region Id: " + cloudRegionId + ", DisplayName: " + cloudRegionDisplayName)
}
```

#### AddCloudRegion

```go
func (s *Client) AddCloudRegion(cloudRegion *CloudRegion) (*CloudRegion, error)
```

##### __Required Fields__
* TenantId
* CloudId
* RegionName
* DisplayName


##### Example
```go
newCloudRegion := cloudcenter.CloudRegion{
	TenantId:    cloudcenter.String("1"),
	CloudId:     cloudcenter.String("2"),
	RegionName:  cloudcenter.String("clientlibrary-west"),
	DisplayName: cloudcenter.String("Client Library West Region"),
}

cloudRegion, err := client.AddCloudRegion(&newCloudRegion)

if err != nil {
	fmt.Println(err)
} else {
	cloudRegionId := *cloudRegion.Id
	cloudRegionDisplayName := *cloudRegion.DisplayName
	fmt.Println("Cloud Region Id: " + cloudRegionId + ", DisplayName: " + cloudRegionDisplayName)
}
```

#### UpdateCloudRegion

```go
func (s *Client) UpdateCloudRegion(cloudRegion *CloudRegion) (*CloudRegion, error)
```

##### __Required Fields__
* Id
* TenantId
* CloudId
* DisplayName

##### Example
```go
newCloudRegion := cloudcenter.CloudRegion{
	Id:    cloudcenter.String("3"),
	TenantId:    cloudcenter.String("1"),
	CloudId:     cloudcenter.String("2"),
	DisplayName: cloudcenter.String("Client Library West Region"),
}

cloudRegion, err := client.UpdateCloudRegion(&newCloudRegion)

if err != nil {
	fmt.Println(err)
} else {
	cloudRegionId := *cloudRegion.Id
	cloudRegionDisplayName := *cloudRegion.DisplayName
	fmt.Println("Cloud Region Id: " + cloudRegionId + ", DisplayName: " + cloudRegionDisplayName)
}
```

#### DeleteCloudRegion

```go
func (s *Client) DeleteCloudRegion(tenantId int, cloudId int, cloudRegionId int) error
```

##### Example
```go
err := client.DeleteCloudRegion(1,1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Cloud region deleted")
}
```


### CloudStorageTypes

- [GetCloudStorageTypes](#getcloudstoragetypes)
- [GetCloudStorageType](#getcloudstoragetype)
- [AddCloudStorageType](#addcloudstoragetype)
- [UpdateCloudStorageType](#updatecloudstoragetype)
- [DeleteCloudStorageType](#deletecloudstoragetype)

```go
type CloudStorageTypeAPIResponse struct {
	Resource          *string            
	Size              *int64             
	PageNumber        *int64             
	TotalElements     *int64             
	TotalPages        *int64             
	CloudStorageTypes []CloudStorageType 
}
```

```go
type CloudStorageType struct {
	Id               *string  
	CloudId          *string  
	TenantId         *string  
	RegionId         *string  
	Resource         *string  
	Name             *string  
	Description      *string  
	Type             *string  
	CostPerMonth     *float64 
	MinVolumeSize    *int64   
	MaxVolumeSize    *int64   
	MaxIOPS          *int64   
	MaxThroughput    *int64   
	ProvisionedIOPS  *bool    
	IOPSCostPerMonth *float64 
}
```

#### GetCloudStorageTypes

```go
func (s *Client) GetCloudStorageTypes(tenantId int, cloudId int, regionId int) ([]CloudStorageType, error)
```

##### Example

```go
cloudStorageTypes, err := client.GetCloudStorageTypes(1, 1, 1)

if err != nil {
	fmt.Println(err)
} else {
	for _, cloudStorageType := range cloudStorageTypes {
		cloudStorageTypeId := *cloudStorageType.Id
		cloudStorageTypeResource := *cloudStorageType.Resource
		fmt.Println("Id: " + cloudStorageTypeId + ", Resource: " + cloudStorageTypeResource)
	}
}
```

#### GetCloudStorageType

```go
func (s *Client) GetCloudStorageType(tenantId int, cloudId int, regionId int, cloudStorageTypeId int) (*CloudStorageType, error)
```

##### Example

```go
cloudStorageType, err := client.GetCloudStorageType(1, 1, 1, 1)

if err != nil {
	fmt.Println(err)
} else {
	cloudStorageTypeId := *cloudStorageType.Id
	cloudStorageTypeResource := *cloudStorageType.Resource
	fmt.Println("Id: " + cloudStorageTypeId + ", Resource: " + cloudStorageTypeResource)
}
```

#### AddCloudStorageType

```go
func (s *Client) AddCloudStorageType(cloudStorageType *CloudStorageType) (*CloudStorageType, error)
```

##### __Required Fields__
* TenantId
* CloudId
* TegionId
* Name
* Type


##### Example
```go
newCloudStorageType := cloudcenter.CloudStorageType{
	TenantId: cloudcenter.String("1"),
	CloudId:  cloudcenter.String("1"),
	RegionId: cloudcenter.String("1"),
	Type:     cloudcenter.String("st3"),
	Name:     cloudcenter.String("Storage Type 01"),
}

cloudStorageType, err := client.AddCloudStorageType(&newCloudStorageType)

if err != nil {
	fmt.Println(err)
} else {
	cloudStorageTypeId := *cloudStorageType.Id
	cloudStorageTypeName := *cloudStorageType.Name
	fmt.Println("Id: " + cloudStorageTypeId + ", Name: " + cloudStorageTypeName)
}
```

#### UpdateCloudStorageType

```go
func (s *Client) UpdateCloudStorageType(cloudStorageType *CloudStorageType) (*CloudStorageType, error)
```

##### __Required Fields__
* Id
* TenantId
* CloudId
* TegionId
* Name
* Type

##### Example
```go
newCloudStorageType := cloudcenter.CloudStorageType{
	Id:       cloudcenter.String("1"),
	TenantId: cloudcenter.String("1"),
	CloudId:  cloudcenter.String("1"),
	RegionId: cloudcenter.String("1"),
	Type:     cloudcenter.String("st3"),
	Name:     cloudcenter.String("Updated storage Type 01"),
}

cloudStorageType, err := client.UpdateCloudStorageType(&newCloudStorageType)

if err != nil {
	fmt.Println(err)
} else {
	cloudStorageTypeId := *cloudStorageType.Id
	cloudStorageTypeName := *cloudStorageType.Name
	fmt.Println("Id: " + cloudStorageTypeId + ", Name: " + cloudStorageTypeName)
}
```

#### DeleteCloudStorageType

```go
func (s *Client) DeleteCloudStorageType(tenantId int, cloudId int, regionId int, cloudStorageTypeId int) error
```

##### Example
```go
err := client.DeleteCloudStorageType(1,1,1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Cloud storage type deleted")
}
```

### Clouds

- [GetClouds](#getclouds)
- [GetCloud](#getcloud)
- [AddCloud](#addcloud)
- [UpdateCloud](#updatecloud)
- [DeleteCloud](#deletecloud)

```go
type CloudAPIResponse struct {
	Resource      *string 
	Size          *int64  
	PageNumber    *int64  
	TotalElements *int64  
	TotalPages    *int64  
	Clouds        []Cloud 
}
```

```go
type Cloud struct {
	Id          *string   
	Resource    *string   
	Perms       *[]string 
	Name        *string   
	Description *string   
	CloudFamily *string   
	PublicCloud *bool     
	TenantId    *string   
	Detail      *Detail   
	CanDelete   *bool     
}
```

```go
type Detail struct {
	CloudAccounts *[]CloudAccount 
	CloudRegions  *[]CloudRegion  
	Status        *string         
	StatusDetail  *string         
}
```


#### GetClouds

```go
func (s *Client) GetClouds(tenantId int) ([]Cloud, error)
```

##### Example

```go
clouds, err := client.GetClouds(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, cloud := range clouds {
		cloudId := *cloud.Id
		cloudName := *cloud.Name
		fmt.Println("Cloud Id: " + cloudId + ", Name: " + cloudName)
	}
}
```

#### GetCloud

```go
func (s *Client) GetCloud(tenantId int, cloudId int) (*Cloud, error)
```

##### Example

```go
cloud, err := client.GetCloud(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	cloudId := *cloud.Id
	cloudName := *cloud.Name
	fmt.Println("Cloud Id: " + cloudId + ", Name: " + cloudName)
}
```

#### AddCloud

```go
func (s *Client) AddCloud(cloud *Cloud) (*Cloud, error)
```

##### __Required Fields__
* TenantId
* Name
* CloudFamily


##### Example
```go
newCloud := cloudcenter.Cloud{
	TenantId:    cloudcenter.String("1"),
	Name:        cloudcenter.String("ClientLibraryCloud"),
	CloudFamily: cloudcenter.String("Amazon"),
}

cloud, err := client.AddCloud(&newCloud)

if err != nil {
	fmt.Println(err)
} else {
	cloudId := *cloud.Id
	cloudName := *cloud.Name
	fmt.Println("Cloud Id: " + cloudId + ", Name: " + cloudName)
}
```

#### UpdateCloud

```go
func (s *Client) UpdateCloud(cloud *Cloud) (*Cloud, error)
```

##### __Required Fields__
* Id
* TenantId
* Name
* CloudFamily

##### Example
```go
newCloud := cloudcenter.Cloud{
	Id:          cloudcenter.String("3"),
	TenantId:    cloudcenter.String("1"),
	Name:        cloudcenter.String("ClientLibraryCloud"),
	CloudFamily: cloudcenter.String("Amazon"),
}

cloud, err := client.UpdateCloud(&newCloud)

if err != nil {
	fmt.Println(err)
} else {
	cloudId := *cloud.Id
	cloudName := *cloud.Name
	fmt.Println("Cloud Id: " + cloudId + ", Name: " + cloudName)
}
```

#### DeleteCloud

```go
func (s *Client) DeleteCloud(tenantId int, cloudId int) error
```

##### Example
```go
err := client.DeleteCloud(1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Cloud deleted")
}
```

### Contracts

- [GetContracts](#getcontracts)
- [GetContract](#getcontract)
- [AddContract](#addcontract)
- [UpdateContract](#updatecontract)
- [DeleteContract](#deletecontract)

```go
type ContractAPIResponse struct {
	Resource      *string    
	Size          *int64     
	PageNumber    *int64     
	TotalElements *int64     
	TotalPages    *int64     
	Contracts     []Contract 
}
```

```go
type Contract struct {
	Id              *string   
	Resource        *string   
	Name            *string   
	Description     *string   
	Perms           *[]string 
	TenantId        *string   
	Length          *int64    
	Terms           *string   
	DiscountRate    *float64  
	Disabled        *bool     
	ShowOnlyToAdmin *bool     
	NumberOfUsers   *int64    
}
```

#### GetContracts

```go
func (s *Client) GetContracts(tenantId int) ([]Contract, error)
```

##### Example

```go
contracts, err := client.GetContracts(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, contract := range contracts {
		fmt.Println("Id: " + contract.Id + ", Name: " + contract.Name)
	}
}
```

#### GetContract

```go
func (s *Client) GetContract(tenantId int, contractId int) (*Contract, error)
```

##### Example

```go
contract, err := client.GetContract(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Id: " + contract.Id + ", Name: " + contract.Name)
}
```

#### AddContract

```go
func (s *Client) AddContract(contract *Contract) (*Contract, error)
```

##### __Required Fields__
* TenantId
* Name
* Length
* Terms

##### Example
```go
newContract := cloudcenter.Contract{
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("ClientLibrary contract"),
	Length:   cloudcenter.Int64(12),
	Terms:    cloudcenter.String("ClientLibrary contract terms"),
}

contract, err := client.AddContract(&newContract)

if err != nil {
	fmt.Println(err)
} else {
	contractId := *contract.Id
	contractDisabled := *contract.Disabled
	fmt.Println("Contract Id: " + contractId + ", Disabled: " + strconv.FormatBool(contractDisabled))
}
```

#### UpdateContract

```go
func (s *Client) UpdateContract(contract *Contract) (*Contract, error)
```

##### __Required Fields__
* Id
* TenantId
* Name
* Length
* Terms

##### Example
```go
newContract := cloudcenter.Contract{
	Id:       cloudcenter.String("2"),
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("ClientLibrary contract"),
	Length:   cloudcenter.Int64(12),
	Terms:    cloudcenter.String("ClientLibrary contract terms"),
}

contract, err := client.UpdateContract(&newContract)

if err != nil {
	fmt.Println(err)
} else {
	contractId := *contract.Id
	contractDisabled := *contract.Disabled
	fmt.Println("Contract Id: " + contractId + ", Disabled: " + strconv.FormatBool(contractDisabled))
}
```

#### DeleteContract

```go
func (s *Client) DeleteContract(tenantId int, contractId int) error
```

##### Example
```go
err := client.DeleteContract(1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Contract deleted")
}
```

### Environments

- [GetEnvironments](#getenvironment)
- [GetEnvironment](#getenvironment)
- [AddEnvironment](#addenvironment)
- [UpdateEnvironment](#updateenvironment)
- [DeleteEnvironment](#deleteenvironment)

```go
type EnvironmentAPIResponse struct {
	Resource      *string       
	Size          *int64        
	PageNumber    *int64        
	TotalElements *int64        
	TotalPages    *int64        
	Environments  []Environment 
}
```

```go
type Environment struct {
	Id                 *string            
	Resource           *string            
	Name               *string            
	Perms              *[]string          
	Description        *string            
	AllowedClouds      *string            
	DefaultSettings    *string            
	RequiresApproval   *bool              
	AssociatedClouds   *[]AssociatedCloud 
	TotalDeployments   *int64             
	ExtensionId        *string            
	CostDetails        *CostDetail        
	NetworkTypes       *[]NetworkType     
	NetworkTypeEnabled *bool              
	RestrictedUser     *bool              
	DefaultRegionId    *string            
	Owner              *int64             
}
```

```go
type AssociatedCloud struct {
	RegionId                 *string                    
	RegionName               *string                    
	RegionDisplayName        *string                    
	CloudFamily              *string                    
	CloudId                  *string                    
	CloudAccountId           *string                    
	CloudName                *string                    
	CloudAccountName         *string                    
	CloudAssociationDefaults *[]CloudAssociationDefault 
	Default                  *bool                      
}
```

```go
type CostDetail struct {
	TotalCloudCost *float64 
	TotalAppCost   *float64 
	TotalJobsCost  *float64 
}
```

```go
type NetworkType struct {
	Name                  *string  
	Description           *string  
	NumberOfNetworkMapped *float64 
}
```

```go
type CloudAssociationDefault struct {
	Name  *string 
	Value *string 
}
```

#### GetEnvironments

```go
func (s *Client) GetEnvironments() ([]Environment, error)
```

##### Example

```go
environments, err := client.GetEnvironments()

if err != nil {
	fmt.Println(err)
} else {
	for _, environment := range environments {
		environmentId := *environment.Id
		environmentName := *environment.Name
		fmt.Println("Environment Id: " + environmentId + ", Name: " + environmentName)
	}
}
```

#### GetEnvironment

```go
func (s *Client) GetEnvironment(id int) (*Environment, error)
```

##### Example

```go
environment, err := client.GetEnvironment(1)

if err != nil {
	fmt.Println(err)
} else {
	environmentId := *environment.Id
	environmentName := *environment.Name
	fmt.Println("Environment Id: " + environmentId + ", Name: " + environmentName)
}
```

#### AddEnvironment

```go
func (s *Client) AddEnvironment(environment *Environment) (*Environment, error)
```

##### __Required Fields__
* Name

##### Example
```go
newEnvironment := cloudcenter.Environment{
	Name: cloudcenter.String("Client Library environment"),
}

environment, err := client.AddEnvironment(&newEnvironment)

if err != nil {
	fmt.Println(err)
} else {
	environmentId := *environment.Id
	environmentResource := *environment.Resource
	fmt.Println("Environment Id: " + environmentId + ", Resource: " + environmentResource)
}
```

#### UpdateEnvironment

```go
func (s *Client) UpdateEnvironment(environment *Environment) (*Environment, error)
```

##### __Required Fields__
* Id
* Name

##### Example
```go
newEnvironment := cloudcenter.Environment{
	Id: cloudcenter.String("5"),
	Name: cloudcenter.String("Client Library environment"),
}

environment, err := client.UpdateEnvironment(&newEnvironment)

if err != nil {
	fmt.Println(err)
} else {
	environmentId := *environment.Id
	environmentResource := *environment.Resource
	fmt.Println("Environment Id: " + environmentId + ", Resource: " + environmentResource)
}
```
#### DeleteEnvironment

```go
func (s *Client) DeleteEnvironment(environmentId int) error
```

##### Example
```go
err := client.DeleteGroup(4)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Environment deleted")
}
```

### Groups

- [GetGroups](#getgroups)
- [GetGroup](#getgroup)
- [AddGroup](#addgroup)
- [UpdateGroup](#updategroup)
- [DeleteGroup](#deletegroup)

```go
type GroupAPIResponse struct {
	Resource      *string 
	Size          *int    
	PageNumber    *int    
	TotalElements *int    
	TotalPages    *int    
	Groups        []Group 
}
```

```go
type Group struct {
	Id           *string   
	Resource     *string   
	Perms        *[]string 
	Name         *string   
	Description  *string   
	TenantId     *string   
	Users        *[]User   
	Roles        *[]Role   
	Created      *int      
	LastUpdated  *int      
	CreatedBySso *bool     
}
```

#### GetGroups

```go
func (s *Client) GetGroups(tenantId int) ([]Group, error)
```

##### Example

```go
groups, err := client.GetGroups(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, group := range groups {
		groupId := *group.Id
		groupName := *group.Name
		fmt.Println("Group Id: " + groupId + ", Name: " + groupName)
	}
}
```

#### GetGroup

```go
func (s *Client) GetGroup(tenantId int, groupId int) (*Group, error)
```

##### Example

```go
group, err := client.GetGroup(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	groupId := *group.Id
	groupName := *group.Name
	fmt.Println("Group Id: " + groupId + ", Name: " + groupName)
}
```

#### AddGroup

```go
func (s *Client) AddGroup(group *Group) (*Group, error)
```

##### __Required Fields__
* TenantId 
* Name

##### Example
```go
newGroup := cloudcenter.Group{
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("New Client Library group"),
}

group, err := client.AddGroup(&newGroup)

if err != nil {
	fmt.Println(err)
} else {
	groupId := *group.Id
	groupName := *group.Name
	fmt.Println("Group Id: " + groupId + ", Name: " + groupName)
}
```
#### UpdateGroup

```go
func (s *Client) UpdateGroup(group *Group) (*Group, error)
```

##### __Required Fields__
* Id
* TenantId 
* Name

##### Example
```go
newGroup := cloudcenter.Group{
	Id:       cloudcenter.String("4"),
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("New Client Library group"),
}

group, err := client.UpdateGroup(&newGroup)

if err != nil {
	fmt.Println(err)
} else {
	groupId := *group.Id
	groupName := *group.Name
	fmt.Println("Group Id: " + groupId + ", Name: " + groupName)
}

```
#### DeleteGroup

```go
func (s *Client) DeleteGroup(tenantId int, groupId int) error
```

##### Example
```go
err := client.DeleteGroup(1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Group deleted")
}
```

### Images

- [GetImages](#getimages)
- [GetImage](#getimage)
- [AddImage](#addimage)
- [UpdateImage](#updateimage)
- [DeleteImage](#deleteimage)

```go
type ImageAPIResponse struct {
	Resource      *string 
	Size          *int64  
	PageNumber    *int64  
	TotalElements *int64  
	TotalPages    *int64  
	Images        []Image 
}
```

```go
type Image struct {
	Id                *string       
	TenantId          *int64        
	Resource          *string       
	Perms             *[]string     
	Name              *string       
	InternalImageName *string       
	Description       *string       
	Visibility        *string       
	ImageType         *string       
	OSName            *string       
	Tags              *[]string     
	Enabled           *bool         
	SystemImage       *bool         
	NumOfNICs         *int64        
	AttachCount       *int64        
	Details           *ImageDetails 
}
```

```go
type ImageDetails struct {
	Count       *int64        
	CloudImages *[]CloudImage 
}
```

```go
type CloudImage struct {
	Id                   *string     
	Resource             *string     
	Perms                *[]string   
	RegionId             *string     
	CloudProviderImageId *string     
	LaunchUserName       *string     
	ImageId              *string     
	GrantAndRevoke       *bool       
	ImageCloudAccountId  *int64      
	Resources            *[]Resource 
	Mappings             *[]Mapping  
}
```

```go
type Resource struct {
	Name  *string 
	Value *string 
}
```

```go
type Mapping struct {
	Id                           *string            
	CloudInstanceType            *CloudInstanceType 
	CostOverride                 *float64           
	CloudProviderImageIdOverride *string            
}
```

```go
type CloudInstanceType struct {
	Id                        *string   
	Resource                  *string   
	Perms                     *[]string 
	Name                      *string   
	Description               *string   
	Type                      *string   
	TenantId                  *string   
	CloudId                   *string   
	RegionId                  *string   
	CostPerHour               *float64  
	MemorySize                *int64    
	NumOfCPUs                 *int64    
	NumOfNICs                 *int64    
	LocalStorageSize          *int64    
	SupportsSSD               *bool     
	SupportsCUDA              *bool     
	Supports32Bit             *bool     
	Supports64Bit             *bool     
	LocalStorageCount         *float64  
	SupportsHardwareProvision *bool     
}
```

#### GetImages

```go
func (s *Client) GetImages(tenantId int) ([]Image, error)
```

##### Example

```go
images, err := client.GetImages(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, image := range images {
		imageId := *image.Id
		imageResource := *image.Resource
		fmt.Println("Id: " + imageId + ", Resource: " + imageResource)
	}
}
```

#### GetImage

```go
func (s *Client) GetImage(tenantId int, imageId int) (*Image, error)
```

##### Example

```go
image, err := client.GetImage(1, 2)

if err != nil {
	fmt.Println(err)
} else {
	imageId := *image.Id
	imageResource := *image.Resource
	fmt.Println("Id: " + imageId + ", Resource: " + imageResource)
}
```

#### AddImage

```go
func (s *Client) AddImage(image *Image) (*Image, error)
```

##### __Required Fields__
* TenantId 
* Name
* ImageType
* OSName

##### Example
```go
newImage := cloudcenter.Image{
	TenantId:  cloudcenter.Int64(1),
	Name:      cloudcenter.String("Ubuntu 14.4"),
	ImageType: cloudcenter.String("CLOUD_WORKER"),
	OSName:    cloudcenter.String("LINUX"),
}

image, err := client.AddImage(&newImage)

if err != nil {
	fmt.Println(err)
} else {
	imageId := *image.Id
	imageResource := *image.Resource
	fmt.Println("Id: " + imageId + ", Resource: " + imageResource)
}
```
#### UpdateImage

```go
func (s *Client) UpdateImage(image *Image) (*Image, error)
```

##### __Required Fields__
* Id
* TenantId 
* Name
* ImageType
* OSName

##### Example
```go
newImage := cloudcenter.Image{
	Id:      cloudcenter.String("12"),
	TenantId:  cloudcenter.Int64(1),
	Name:      cloudcenter.String("Updated Ubuntu 14.4"),
	ImageType: cloudcenter.String("CLOUD_WORKER"),
	OSName:    cloudcenter.String("LINUX"),
}

image, err := client.UpdateImage(&newImage)

if err != nil {
	fmt.Println(err)
} else {
	imageId := *image.Id
	imageResource := *image.Resource
	fmt.Println("Id: " + imageId + ", Resource: " + imageResource)
}
```
#### DeleteImage

```go
func (s *Client) DeleteImage(tenantId int, imageId int) error
```

##### Example
```go
err := client.DeleteImage(1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Image deleted")
}
```

### Jobs
### OperationStatus

```go
type OperationStatus struct {
	OperationId          *string                
	Id                   *string                
	Status               *string                
	Resource             *string                
	Msg                  *string                
	Progress             *int64                 
	AdditionalParameters *[]AdditionalParameter 
}
```

```go
type AdditionalParameter struct {
	Name  *string 
	Value *string 
}
```

### Phases

- [GetPhases](#getphases)
- [GetPhase](#getphase)
- [AddPhase](#addphase)
- [UpdatePhase](#updatephase)
- [DeletePhase](#deletephase)

```go
type PhaseAPIResponse struct {
	Resource      *string 
	Size          *int64  
	PageNumber    *int64  
	TotalElements *int64  
	TotalPages    *int64  
	Phases        []Phase 
}
```

```go
type Phase struct {
	Id                     *string                
	ProjectId              *string                
	Resource               *string                
	Perms                  *[]string              
	Name                   *string                
	Order                  *float64               
	PhasePlan              *PhasePlan             
	PhaseBundles           *[]PhaseBundle         
	PhaseCost              *PhaseCost             
	Deployments            *[]Deployment          
	DeploymentEnvironments *DeploymentEnvironment 
}
```

```go
type PhasePlan struct {
	Id       *string 
	PlanName *string 
}
```

```go
type PhaseBundle struct {
	Id    *string 
	Name  *string 
	Count *int64  
}
```

```go
type PhaseCost struct {
	OriginalBalance  *float64 
	RemainingBalance *float64 
	MeasurableUnit   *string  
}
```

#### GetPhases

```go
func (s *Client) GetPhases(projectId int) ([]Phase, error)
```

##### Example

```go
phases, err := client.GetPhases(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, phase := range phases {
		phaseId := *phase.Id
		phaseName := *phase.Name
		fmt.Println("Id: " + phaseId + ", Name: " + phaseName)
	}
}
```

#### GetPhase

```go
func (s *Client) GetPhase(projectId int, id int) (*Phase, error)
```

##### Example

```go
phase, err := client.GetPhase(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	phaseId := *phase.Id
	phaseName := *phase.Name
	fmt.Println("Id: " + phaseId + ", Name: " + phaseName)
}
```

#### AddPhase

```go
func (s *Client) AddPhase(phase *Phase) (*Phase, error)
```

##### __Required Fields__
* ProjectId
* Name


##### Example
```go

```
#### UpdatePhase

```go
func (s *Client) UpdatePhase(phase *Phase) (*Phase, error)
```

##### __Required Fields__
* Id
* ProjectId
* Name


##### Example
```go

```
#### DeletePhase

```go
func (s *Client) DeletePhase(phaseProjectID int, phaseId int) error
```

##### Example
```go
err := client.DeletePhase(1,1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Phase deleted")
}
```

### Plans

- [GetPlans](#getplans)
- [GetPlan](#getplan)
- [AddPlan](#addplan)
- [UpdatePlan](#updateplan)
- [DeletePlan](#deleteplan)

```go
type PlanAPIResponse struct {
	Resource      *string 
	Size          *int64  
	PageNumber    *int64  
	TotalElements *int64  
	TotalPages    *int64  
	Plans         []Plan  
}
```

```go
type Plan struct {
	Id                       *string   
	Resource                 *string   
	Name                     *string   
	Description              *string   
	Perms                    *[]string 
	TenantId                 *string   
	Type                     *string   
	MonthlyLimit             *int64    
	NodeHourIncrement        *float64  
	IncludedBundleId         *string   
	Price                    *float64  
	OnetimeFee               *float64  
	AnnualFee                *float64  
	StorageRate              *float64  
	HourlyRate               *float64  
	OverageRate              *float64  
	OverageLimit             *float64  
	RestrictedToAppStoreOnly *bool     
	BillToVendor             *bool     
	EnableRollover           *bool     
	Disabled                 *bool     
	ShowOnlyToAdmin          *bool     
	NumberOfUsers            *int64    
	NumberOfProjects         *int64    
	PaymentProfileRequired   *bool     
}
```

#### GetPlans

```go
func (s *Client) GetPlans(tenantId int) ([]Plan, error)
```

##### Example

```go
plans, err := client.GetPlans(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, plan := range plans {
		planId := *plan.Id
		planName := *plan.Name
		fmt.Println("Id: " + planId + ", Name: " + planName)
	}
}
```

#### GetPlan

```go
func (s *Client) GetPlan(tenantId int, planId int) (*Plan, error)
```

##### Example

```go
plan, err := client.GetPlan(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	planId := *plan.Id
	planName := *plan.Name
	fmt.Println("Id: " + planId + ", Name: " + planName)
}
```

#### AddPlan

```go
func (s *Client) AddPlan(plan *Plan) (*Plan, error)
```

##### __Required Fields__
* TenantId
* Name
* Type

##### Example
```go
newPlan := cloudcenter.Plan{
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("client library plan"),
	Type: cloudcenter.String("UNLIMITED_PLAN"),
}

plan, err := client.AddPlan(&newPlan)

if err != nil {
	fmt.Println(err)
} else {
	planId := *plan.Id
	planDisabled := *plan.Disabled
	fmt.Println("Plan Id: " + planId + ", Disabled: " + strconv.FormatBool(planDisabled))
}
```
#### UpdatePlan

```go
func (s *Client) UpdatePlan(plan *Plan) (*Plan, error)
```

##### __Required Fields__
* Id
* TenantId
* Name
* Type

##### Example
```go
newPlan := cloudcenter.Plan{
	Id: cloudcenter.String("2"),
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("client library plan"),
	Type: cloudcenter.String("UNLIMITED_PLAN"),
}

plan, err := client.UpdatePlan(&newPlan)

if err != nil {
	fmt.Println(err)
} else {
	planId := *plan.Id
	planDisabled := *plan.Disabled
	fmt.Println("Plan Id: " + planId + ", Disabled: " + strconv.FormatBool(planDisabled))
}
```
#### DeletePlan

```go
func (s *Client) DeletePlan(tenantId int, planId int) error
```

##### Example
```go
err := client.DeletePlan(1, 12)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Plan deleted")
}
```


### Projects

- [GetProjects](#getprojects)
- [GetProject](#getproject)
- [AddProject](#addproject)
- [UpdateProject](#updateproject)
- [DeleteProject](#deleteproject)

```go
type ProjectAPIResponse struct {
	Resource      *string   
	Size          *int64    
	PageNumber    *int64    
	TotalElements *int64    
	TotalPages    *int64    
	Projects      []Project 
}
```

```go
type Project struct {
	Id             *string      
	Resource       *string      
	Perms          *[]string    
	Name           *string      
	Description    *string      
	ProjectOwnerId *int64       
	IsDraft        *bool        
	TargetEndDate  *string      
	NotifyUsers    *bool        
	PlanType       *string      
	Deleted        *bool        
	Quota          *Quota       
	ProjectCost    *ProjectCost 
	Apps           *[]Apps      
	Phases         *[]Phase     
}
```

```go
type ProjectCost struct {
	OriginalBalance  *float64 
	RemainingBalance *float64 
	MeasurableUnit   *string  
}
```

#### GetProjects

```go
func (s *Client) GetProjects() ([]Project, error)
```

##### Example

```go
projects, err := client.GetProjects()

if err != nil {
	fmt.Println(err)
} else {
	for _, project := range projects {
		projectName := *project.Name
		projectId := *project.Id
		fmt.Println("Project Id: " + projectId + ", Name: " + projectName)
	}
}
```

#### GetProject

```go
func (s *Client) GetProject(id int) (*Project, error)
```

##### Example

```go
project, err := client.GetProject(1)

if err != nil {
	fmt.Println(err)
} else {
	projectName := *project.Name
	projectId := *project.Id
	fmt.Println("Project Id: " + projectId + ", Name: " + projectName)
}
```

#### AddProject

```go
func (s *Client) AddProject(project *Project) (*Project, error)
```

##### __Required Fields__
* Name

##### Example
```go
newProject := cloudcenter.Project{
	Name: cloudcenter.String("ClientLibrary project"),
}

project, err := client.AddProject(&newProject)

if err != nil {
	fmt.Println(err)
} else {
	projectName := *project.Name
	projectId := *project.Id
	fmt.Println("Project Id: " + projectId + ", Name: " + projectName)
}
```
#### UpdateProject

```go
func (s *Client) UpdateProject(project *Project) (*Project, error)
```

##### __Required Fields__
* Id
* Name

##### Example
```go
newProject := cloudcenter.Project{
	Id: cloudcenter.String("1"),
	Name: cloudcenter.String("ClientLibrary project"),
}

project, err := client.UpdateProject(&newProject)

if err != nil {
	fmt.Println(err)
} else {
	projectName := *project.Name
	projectId := *project.Id
	fmt.Println("Project Id: " + projectId + ", Name: " + projectName)
}
```
#### DeleteProject

```go
func (s *Client) DeleteProject(projectId int) error
```

##### Example
```go
err := client.DeleteProject(12)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Project deleted")
}
```

### Roles

- [GetRoles](#getroles)
- [GetRole](#getrole)
- [AddRole](#addrole)
- [UpdateRole](#updaterole)
- [DeleteRole](#deleterole)

```go
type RoleAPIResponse struct {
	Resource      *string 
	Size          *int64  
	PageNumber    *int64  
	TotalElements *int64  
	TotalPages    *int64  
	Roles         []Role  
}
```

```go
type Role struct {
	Id          *string       
	Resource    *string       
	Perms       *[]string     
	Name        *string       
	Description *string       
	TenantId    *string       
	ObjectPerms *[]ObjectPerm 
	Users       *[]User       
	Groups      *[]Group      
	OobRole     *bool         
	LastUpdated *int64        
	Created     *int64        
}
```

```go
type ObjectPerm struct {
	ObjectType *string   
	Perms      *[]string 
}
```

#### GetRoles

```go
func (s *Client) GetRoles(tenantId int) ([]Role, error)
```

##### Example

```go
roles, err := client.GetRoles(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, role := range roles {
		roleName := *role.Name
		roleId := *role.Id
		fmt.Println("Role Id: " + roleId + ", Name: " + roleName)
	}
}
```

#### GetRole

```go
func (s *Client) GetRole(tenantId int, roleId int) (*Role, error)
```

##### Example

```go
role, err := client.GetRole(1, 1)

if err != nil {
	fmt.Println(err)
} else {
	roleName := *role.Name
	roleId := *role.Id
	fmt.Println("Role Id: " + roleId + ", Name: " + roleName)
}
```

#### AddRole

```go
func (s *Client) AddRole(role *Role) (*Role, error)
```

##### __Required Fields__
* TenantId
* Name

##### Example
```go
newRole := cloudcenter.Role{
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("ClientLibrary Role"),
}

role, err := client.AddRole(&newRole)

if err != nil {
	fmt.Println(err)
} else {
	roleName := *role.Name
	roleId := *role.Id
	fmt.Println("Role Id: " + roleId + ", Name: " + roleName)
}
```

#### UpdateRole

```go
func (s *Client) UpdateRole(role *Role) (*Role, error)
```

##### __Required Fields__
* Id
* TenantId
* Name

##### Example
```go
newRole := cloudcenter.Role{
	Id: cloudcenter.String("24"),
	TenantId: cloudcenter.String("1"),
	Name:     cloudcenter.String("ClientLibrary Role"),
}

role, err := client.UpdateRole(&newRole)

if err != nil {
	fmt.Println(err)
} else {
	roleName := *role.Name
	roleId := *role.Id
	fmt.Println("Role Id: " + roleId + ", Name: " + roleName)
}
```

#### DeleteRole

```go
func (s *Client) DeleteRole(tenantId int, roleId int) error
```
##### Example
```go
err := client.DeleteRole(1, 12)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Role deleted")
}
```

### Services

- [GetServices](#getservices)
- [GetService](#getservice)
- [AddService](#addservice)
- [UpdateService](#updateservice)
- [DeleteService](#deleteservice)

```go
type ServiceAPIResponse struct {
	Resource      *string   
	Size          *int64    
	PageNumber    *int64    
	TotalElements *int64    
	TotalPages    *int64    
	Services      []Service 
}
```

```go
type Service struct {
	Id                     *string              
	OwnerUserId            *string              
	TenantId               *string              
	ParentService          *bool                
	ParentServiceId        *string              
	Resource               *string              
	Perms                  *[]string            
	Name                   *string              
	DisplayName            *string              
	LogoPath               *string              
	Description            *string              
	DefaultImageId         *int64               
	ServiceType            *string              
	SystemService          *bool                
	ExternalService        *bool                
	Visible                *bool                
	ExternalBundleLocation *string              
	BundleLocation         *string              
	CostPerHour            *float64             
	OwnerId                *string              
	ServiceActions         *[]ServiceAction     
	ServicePorts           *[]ServicePort       
	ServiceParamSpecs      *[]ServiceParamSpec  
	EgressRestrictions     *[]EgressRestriction 
	Images                 *[]Image             
	Repositories           *[]Repository        
	ChildServices          *[]Service           
	ExternalActions        *[]ExternalAction    
}
```

```go
type ServiceAction struct {
	ActionName  *string 
	ActionType  *string 
	ActionValue *string 
}
```

```go
type ServicePort struct {
	Protocol *string 
	FromPort *string 
	ToPort   *string 
	CloudId  *string 
}
```

```go
type ServiceParamSpec struct {
	ParamName            *string                
	DisplayName          *string                
	HelpText             *string                
	Type                 *string                
	ValueList            *string                
	WebserviceListParams *[]WebserviceListParam 
	DefaultValue         *string                
	UserVisible          *bool                  
	UserEditable         *bool                  
	SystemParam          *bool                  
	ExampleValue         *string                
	Optional             *bool                  
	ValueConstraint      *ValueConstraint       
}
```

```go
type EgressRestriction struct {
	EgressServiceName *string 
}
```

```go
type WebserviceListParam struct {
	URL           *string 
	Protocol      *string 
	Username      *string 
	Password      *string 
	RequestType   *string 
	ContentType   *string 
	CommandParams *string 
	RequestBody   *string 
	ResultString  *string 
}
```

```go
type Repository struct {
	Id          *string   
	Resource    *string   
	Perms       *[]string 
	Hostname    *string   
	DisplayName *string   
	Protocol    *string   
	Description *string   
	Port        *int64    
}
```

```go
type ExternalAction struct {
	ActionName  *string 
	ActionType  *string 
	ActionValue *string 
}
```

#### GetServices

```go
func (s *Client) GetServices(tenantId int) ([]Service, error)
```

##### Example

```go
services, err := client.GetServices(1)

if err != nil {
	fmt.Println(err)
} else {
	for _, service := range services {
		serviceId := *service.Id
		serviceName := *service.Name
		fmt.Println("Id: " + serviceId + ", Name: " + serviceName)
	}
}
```

#### GetService

```go
func (s *Client) GetService(tenantId int, serviceId int) (*Service, error)
```

##### Example

```go
service, err := client.GetService(1, 2)

if err != nil {
	fmt.Println(err)
} else {
	serviceId := *service.Id
	serviceName := *service.Name
	fmt.Println("Id: " + serviceId + ", Name: " + serviceName)
}
```

#### AddService

```go
func (s *Client) AddService(service *Service) (*Service, error)
```

##### __Required Fields__
* TenantId
* Name
* DisplayName
* Images


##### Example
```go
var images []cloudcenter.Image

newImage := cloudcenter.Image{
	Id: cloudcenter.String("1"),
}

images = append(images, newImage)

newService := cloudcenter.Service{
	TenantId:    cloudcenter.String("1"),
	Name:        cloudcenter.String("ClientLibraryService"),
	DisplayName: cloudcenter.String("Client Library Service"),
	Images:      &images,
}

service, err := client.AddService(&newService)

if err != nil {
	fmt.Println(err)
} else {
	serviceId := *service.Id
	serviceName := *service.Name
	fmt.Println("Id: " + serviceId + ", Name: " + serviceName)
}
	
```

#### UpdateService

```go
func (s *Client) UpdateService(service *Service) (*Service, error)
```

##### __Required Fields__
* Id
* TenantId
* Name
* DisplayName
* Images

##### Example
```go
var images []cloudcenter.Image

newImage := cloudcenter.Image{
	Id: cloudcenter.String("1"),
}

images = append(images, newImage)

newService := cloudcenter.Service{
	Id:          cloudcenter.String("66"),
	TenantId:    cloudcenter.String("1"),
	Name:        cloudcenter.String("ClientLibraryService"),
	DisplayName: cloudcenter.String("Client Library Service"),
	Images:      &images,
}

service, err := client.UpdateService(&newService)

if err != nil {
	fmt.Println(err)
} else {
	serviceId := *service.Id
	serviceName := *service.Name
	fmt.Println("Id: " + serviceId + ", Name: " + serviceName)
}
```


#### DeleteService

```go
func (s *Client) DeleteService(tenantId int, serviceId int) error
```
##### Example
```go
err := client.DeleteService(1, 66)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Service deleted")
}
```

### SuspensionPolicies

- [GetSuspensionPolicies](#getsuspensionpolicies)
- [GetSuspensionPolicy](#getsuspensionpolicy)
- [AddSuspensionPolicy](#addsuspensionpolicy)
- [UpdateSuspensionPolicy](#updatesuspensionpolicy)
- [DeleteSuspensionPolicy](#deletesuspensionpolicy)

```go
type SuspensionPolicyAPIResponse struct {
	Resource           *string            
	Size               *int64             
	PageNumber         *int64             
	TotalElements      *int64             
	TotalPages         *int64             
	SuspensionPolicies []SuspensionPolicy 
}
```

```go
type SuspensionPolicy struct {
	Id                        *string           
	Resource                  *string           
	Perms                     *[]string         
	Name                      *string           
	Description               *string           
	Enabled                   *bool             
	Schedules                 *[]Schedule       
	BlockoutPeriods           *[]BlockoutPeriod 
	IsPolicyActiveOnResources *bool             
	ResourcesMaps             *[]ResourcesMap   
	Priority                  *float64          
	Created                   *float64          
	LastUpdated               *float64          
	OwnerId                   *int64            
}
```

```go
type Schedule struct {
	Type      *string   
	Days      *[]string 
	StartTime *string   
	EndTime   *string   
	Repeats   *string   
}
```

```go
type BlockoutPeriod struct {
	StartDate *float64 
	EndDate   *float64 
}
```

```go
type ResourcesMap struct {
	ResourceId                  *string  
	ResourceType                *string  
	AppliedDate                 *float64 
	ResourceStartTime           *float64 
	EstimatedPolicyEndTime      *float64 
	AllowedCost                 *float64 
	AccruedCost                 *float64 
	NumberOfExtensionsUsed      *int64   
	IsApprovalPending           *bool    
	IsPreviousExtensionDenied   *bool    
	IsPolicyReachingExpiry      *bool    
	IsPolicyReachingGraceExpiry *bool    
}
```

#### GetSuspensionPolicies

```go
func (s *Client) GetSuspensionPolicy(suspensionPolicyId int) (*SuspensionPolicy, error)
```

##### Example

```go
suspensionPolicies, err := client.GetSuspensionPolicies()

if err != nil {
	fmt.Println(err)
} else {
	for _, suspensionPolicy := range suspensionPolicies {
		suspensionPolicyId := *suspensionPolicy.Id
		suspensionPolicyName := *suspensionPolicy.Name
		fmt.Println("Id: " + suspensionPolicyId + ", Name: " + suspensionPolicyName)
	}
}
```

#### GetSuspensionPolicy

```go
func (s *Client) GetSuspensionPolicies() ([]SuspensionPolicy, error)
```

##### Example

```go
suspensionPolicy, err := client.GetSuspensionPolicy(2)

if err != nil {
	fmt.Println(err)
} else {
	suspensionPolicyId := *suspensionPolicy.Id
	suspensionPolicyName := *suspensionPolicy.Name
	fmt.Println("Id: " + suspensionPolicyId + ", Name: " + suspensionPolicyName)
}
```

#### AddSuspensionPolicy

```go
func (s *Client) AddSuspensionPolicy(suspensionPolicy *SuspensionPolicy) (*SuspensionPolicy, error)
```

##### __Required Fields__
* Name
* Enabled
* Schedules 

##### Example
```go
var schedules []cloudcenter.Schedule

newSchedule := cloudcenter.Schedule{
	Type:      cloudcenter.String("1"),
	StartTime: cloudcenter.String("12:00"),
	EndTime:   cloudcenter.String("15:00"),
	Repeats:   cloudcenter.String("1"),
	Days:      &[]string{"WED"},
}

schedules = append(schedules, newSchedule)

newSuspensionPolicy := cloudcenter.SuspensionPolicy{
	Name:      cloudcenter.String("ClientLibrarySuspenion"),
	Enabled:   cloudcenter.Bool(false),
	Schedules: &schedules,
}

suspensionPolicy, err := client.AddSuspensionPolicy(&newSuspensionPolicy)

if err != nil {
	fmt.Println(err)
} else {
	suspensionPolicyId := *suspensionPolicy.Id
	suspensionPolicyName := *suspensionPolicy.Name
	fmt.Println("Id: " + suspensionPolicyId + ", Name: " + suspensionPolicyName)
}
```

#### UpdateSuspensionPolicy

```go
func (s *Client) UpdateSuspensionPolicy(suspensionPolicy *SuspensionPolicy) (*SuspensionPolicy, error)
```

##### __Required Fields__
* Id
* Name
* Enabled
* Schedules 

##### Example
```go
var schedules []cloudcenter.Schedule

newSchedule := cloudcenter.Schedule{
	Type:      cloudcenter.String("1"),
	StartTime: cloudcenter.String("12:00"),
	EndTime:   cloudcenter.String("15:00"),
	Repeats:   cloudcenter.String("1"),
	Days:      &[]string{"WED"},
}

schedules = append(schedules, newSchedule)

newSuspensionPolicy := cloudcenter.SuspensionPolicy{
	Id:        cloudcenter.String("3"),
	Name:      cloudcenter.String("ClientLibrarySuspenion"),
	Enabled:   cloudcenter.Bool(false),
	Schedules: &schedules,
}

suspensionPolicy, err := client.UpdateSuspensionPolicy(&newSuspensionPolicy)

if err != nil {
	fmt.Println(err)
} else {
	suspensionPolicyId := *suspensionPolicy.Id
	suspensionPolicyName := *suspensionPolicy.Name
	fmt.Println("Id: " + suspensionPolicyId + ", Name: " + suspensionPolicyName)
}
```

#### DeleteSuspensionPolicy

```go
func (s *Client) DeleteSuspensionPolicy(suspensionPolicyId int) error
```

##### Example
```go
err := client.DeleteSuspensionPolicy(1)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Suspension policy deleted")
}
```

### Tenants

- [GetTenants](#gettenants)
- [GetTenant](#gettenant)
- [AddTenant](#addtenant)
- [UpdateTenant](#updatetenant)
- [DeleteTenantAsync](#deletetenantasync)
- [DeleteTenantSync](#deletetenantsync)

```go
type TenantAPIResponse struct {
	Resource      *string  
	Size          *int     
	PageNumber    *int     
	TotalElements *int     
	TotalPages    *int     
	Tenants       []Tenant 
}
```

```go
type Tenant struct {
	Id                              *string       
	Resource                        *string       
	Name                            *string       
	Url                             *string       
	About                           *string       
	ContactEmail                    *string       
	Phone                           *string       
	UserId                          *string       
	TermsOfService                  *string       
	PrivacyPolicy                   *string       
	RevShareRate                    *float64      
	CcTransactionFeeRate            *float64      
	MinAppFeeRate                   *float64      
	EnableConsolidatedBilling       *bool         
	ShortName                       *string       
	EnablePurchaseOrder             *bool         
	EnableEmailNotificationsToUsers *bool         
	ParentTenantId                  *int64        
	ExternalId                      *string       
	DefaultActivationProfileId      *string       
	EnableMonthlyBilling            *bool         
	DefaultChargeType               *string       
	LoginLogo                       *string       
	HomePageLogo                    *string       
	DomainName                      *string       
	SkipDefaultUserSecurityGroup    *bool         
	DisableAllEmailNotification     *bool         
	TrademarkURL                    *string       
	Deleted                         *bool         
	Preferences                     *[]Preference 
}
```

```go
type Preference struct {
	Name  *string 
	Value *string 
}
```

#### GetTenants

```go
func (s *Client) GetTenants() ([]Tenant, error)
```

##### Example

```go
tenants, err := client.GetTenants()

if err != nil {
	fmt.Println(err)
} else {
	for _, tenant := range tenants {
		tenantId := *tenant.Id
		tenantName := *tenant.Name
		fmt.Println("Id: " + tenantId + ", Name: " + tenantName)
	}
}
```

#### GetTenant

```go
func (s *Client) GetTenant(id int) (*Tenant, error)
```

##### Example

```go
tenant, err := client.GetTenant(1)

if err != nil {
	fmt.Println(err)
} else {
	tenantId := *tenant.Id
	tenantName := *tenant.Name
	fmt.Println("Id: " + tenantId + ", Name: " + tenantName)
}
```

#### AddTenant

```go
func (s *Client) AddTenant(tenant *Tenant) error
```

##### __Required Fields__
* Name
* ShortName
* UserId

##### Example

```golang
newTenant := cloudcenter.Tenant{
	UserId:                           cloudcenter.String("5",
	Name:                             cloudcenter.String("client-library-tenant"),
	ShortName:                        cloudcenter.String("client-library-tenant"),
	DomainName:                       cloudcenter.String("clientlibrary.cloudcenter.com"),
	Phone:                            cloudcenter.String("1234567890"),
	Url:                              cloudcenter.String("http://clientlibrary.cloudcenter.com"),
	ContactEmail:                     cloudcenter.String("poweruser@dcloud.cisco.com"),
	About:                            cloudcenter.String("clientlibrary tenant"),
	EnablePurchaseOrder:              cloudcenter.Bool(false),
	EnableEmailNotificationsToUsers:  cloudcenter.Bool(false),
	EnableMonthlyBilling:             cloudcenter.Bool(false),
	DefaultChargeType:                cloudcenter.String("Hourly)",
}

fmt.Println(client.AddTenant(&newTenant))
```

#### UpdateTenant

```go
func (s *Client) UpdateTenant(tenant *Tenant) (*Tenant, error)
```

##### __Required Fields__
* Id 
* UserId 
* Name 
* ShortName 

##### Example

```golang
newTenant := cloudcenter.Tenant{
	Id:                              cloudcenter.String("2"),
	UserId:                          cloudcenter.String("5"),
	Name:                            cloudcenter.String("client-library-tenant"),
	ShortName:                       cloudcenter.String("client-library-tenant"),
	DomainName:                      cloudcenter.String("clientlibrary.cloudcenter.com"),
	Phone:                           cloudcenter.String("1234567890"),
	Url:                             cloudcenter.String("http://clientlibrary.cloudcenter.com"),
	ContactEmail:                    cloudcenter.String("poweruser@dcloud.cisco.com"),
	About:                           cloudcenter.String("clientlibrary tenant"),
	EnablePurchaseOrder:             cloudcenter.Bool(false),
	EnableEmailNotificationsToUsers: cloudcenter.Bool(false),
	EnableMonthlyBilling:            cloudcenter.Bool(false),
	DefaultChargeType:               cloudcenter.String("Hourly"),
}

fmt.Println(client.UpdateTenant(&newTenant))
```

#### DeleteTenantAsync

```go
func (s *Client) DeleteTenantAsync(tenantId int) (*OperationStatus, error)
```
##### Example

```go
operationStatus, err := client.DeleteTenantAsync(3)

if err != nil {
	fmt.Println(err)
} else {

	if err != nil {
		fmt.Println(err)
	} else {
		operationStatusId := *operationStatus.Id
		operationStatusStatus := *operationStatus.Status
		operationStatusMsg := *operationStatus.Msg
		fmt.Println("Operation Status: " + operationStatusId + ", Status: " + operationStatusStatus + ", Message: " + operationStatusMsg)

		for operationStatusStatus == "RUNNING" {
			operationStatusStatus = *operationStatus.Status
			operationStatusId := *operationStatus.Id
			operationStatus, err = client.GetOperationStatus(operationStatusId)
		}

		if operationStatusStatus == "SUCCESS" {
			fmt.Println("Tenant deleted")
		}
	}
}
```

#### DeleteTenantSync

```go
func (s *Client) DeleteTenantSync(tenantId int) error
```
##### Example

```go
err := client.DeleteTenant(6)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("Tenant deleted")
}
```

### Users

- [GetUsers](#getusers)
- [GetUser](#getuser)
- [GetUserFromEmail](#getuserfromemail)
- [AddUser](#adduser)
- [UpdateUser](#updateuser)
- [DeleteUser](#deleteuser)

```go
type UserAPIResponse struct {
	Resource      string 
	Size          int    
	PageNumber    int    
	TotalElements int    
	TotalPages    int    
	Users         []User 
}
```

```go
type User struct {
	Id                      *string 
	Resource                *string 
	Username                *string  
	Password                *string  
	Enabled                 *bool   
	Type                    *string 
	FirstName               *string 
	LastName                *string 
	CompanyName             *string 
	EmailAddr               *string 
	EmailVerified           *bool   
	PhoneNumber             *string 
	ExternalId              *string 
	AccessKeys              *string 
	DisableReason           *string 
	AccountSource           *string 
	Status                  *string 
	Detail                  *string 
	ActivationData          *string 
	Created                 *int64  
	LastUpdated             *int64  
	CoAdmin                 *bool   
	TenantAdmin             *bool   
	ActivationProfileId     *string 
	HasSubscriptionPlanType *bool   
	TenantId                *string 
}
```

#### GetUsers

```go
func (s *Client) GetUsers() ([]User, error)
```

##### Example

```golang
users, err := client.GetUsers()

if err != nil {
	fmt.Println(err)
} else {
	for _, user := range users {
		userId := *user.Id
		username := *user.Username
		userTenantId := *user.TenantId
		fmt.Println("UserId: " + userId + ", Username: " + username + ", TenantId: " + userTenantId)
	}
}
```

#### GetUser

```go
func (s *Client) GetUser(id int) (*User, error)
```

##### Example

```golang
user, err := client.GetUser(1)

if err != nil {
	fmt.Println(err)
} else {
	userId := *user.Id
	username := *user.Username
	userTenantId := *user.TenantId
	fmt.Println("UserId: " + userId + ", Username: " + username + ", TenantId: " + userTenantId)
}
```

#### GetUserFromEmail

```go
func (s *Client) GetUserFromEmail(emailToSearch string) (*User, error)
```

##### Example

```golang
user, err = client.GetUserFromEmail("admin@cliqrtech.com")

if err != nil {
	fmt.Println(err)
} else {
	userId := *user.Id
	username := *user.Username
	userTenantId := *user.TenantId
	fmt.Println("UserId: " + userId + ", Username: " + username + ", TenantId: " + userTenantId)
}
```

#### AddUser

```go
func (s *Client) AddUser(user *User) (*User, error)
```

##### __Required Fields__
* EmailAddr
* TenantId
* Password
* ActivateRegions (When creating user with activation)
* ActivationProfileId (required if using activation profile method)

##### Example

```golang
newUser := cloudcenter.User {
	FirstName:   cloudcenter.String("client"),
	LastName:    cloudcenter.String("library"),
	Password:    cloudcenter.String("myPassword"),
	EmailAddr:   cloudcenter.String("clientlibrary@cloudcenter-address.com"),
	CompanyName: cloudcenter.String("company"),
	PhoneNumber: cloudcenter.String("12345"),
	ExternalId:  cloudcenter.String("23456"),
	TenantId:    cloudcenter.String("1"),
}

user, err := client.AddUser(&newUser)

if err != nil {
	fmt.Println(err)
} else {
	userId := *user.Id
	username:= *user.Username
	fmt.Println("UserId: " + userId + ", Username: " + username)
}
```

#### UpdateUser

```go
func (s *Client) UpdateUser(user *User) (*User, error)
```

##### __Required Fields__
* Id 
* TenantId 
* AccountSource
* Username 
* Type 
* EmailAddr


##### Example

```golang
newUser := cloudcenter.User{
	Id:        cloudcenter.String("2"),
	TenantId:  cloudcenter.String("1"),
	Username:  cloudcenter.String("cliqradmin"),
	Type:      cloudcenter.String("TENANT"),
	EmailAddr: cloudcenter.String("admin@cliqrtech.com"),
	AccountSource: cloudcenter.String("AdminCreated"),
}

user, err := client.UpdateUser(&newUser)

if err != nil {
	fmt.Println(err)
} else {
	userId := *user.Id
	username:= *user.Username
	fmt.Println("UserId: " + userId + ", Username: " + username)
}
```

#### DeleteUser

```go
func (s *Client) DeleteUser(userId int) error
```
##### Example
```go
err := client.DeleteUser(6)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("User deleted")
}
```

### VirtualMachines

- [GetVirtualMachines](#getvirtualmachines)
- [GetVirtualMachine](#getvirtualmachine)
- [GetVirtualMachineCostSummary](#getvirtualmachinecostsummary)
- [AddVirtualMachine](#addvirtualmachine)
- [UpdateVirtualMachine](#updatevirtualmachine)
- [DeleteVirtualMachine](#deletevirtualmachine)

```go
type VirtualMachineAPIResponse struct {
	CostSummary *CostSummary 
	Details     *Details     
	Filters     *Filters     
}
```

```go
type CostSummary struct {
	TotalNumberOfVMs          *int64   
	TotalNumberOfRunningVMs   *int64   
	TotalCloudCost            *float64 
	EstimatedMonthlyCloudCost *float64 
	TotalNodeHours            *float64 
}
```

```go
type Details struct {
	Resource              *string                 
	Size                  *int64                  
	PageNumber            *int64                  
	TotalElements         *int64                  
	TotalPages            *int64                  
	VirtualMachineDetails []VirtualMachineDetails 
}
```

```go
type VirtualMachineDetails struct {
	Id                        *string           
	Resource                  *string           
	Perms                     *[]string         
	Name                      *string           
	Description               *string           
	Enabled                   *bool             
	Schedules                 *[]Schedule       
	BlockoutPeriods           *[]BlockoutPeriod 
	IsPolicyActiveOnResources *bool             
	ResourcesMaps             *[]ResourcesMap   
	Priority                  *float64          
	Created                   *float64          
	LastUpdated               *float64          
	NodeId                    *string           
	HostName                  *string           
	NodeStartTime             *int64            
	NodeEndTime               *int64            
	NumberOfCPUs              *int64            
	MemorySize                *int64            
	StorageSize               *int64            
	OSName                    *string           
	CostPerHour               *float64          
	Status                    *string           
	ReviewNodeStatus          *bool             
	CloudFamily               *string           
	CloudId                   *string           
	CloudName                 *string           
	CloudAccountId            *string           
	CloudAccountName          *string           
	RegionId                  *string           
	RegionName                *string           
	RegionDisplayName         *string           
	TenantId                  *string           
	UserId                    *string           
	FirstName                 *string           
	LastName                  *string           
	Email                     *string           
	InstanceTypeId            *string           
	InstanceTypeName          *string           
	InstanceCost              *float64          
	NICs                      *[]NIC            
	Metatdata                 *[]Metadata       
	NodeProperties            *[]NodeProperty   
	JobStartTime              *float64          
	CloudNameAndAccountName   *string           
	AgentVersion              *string           
	JobId                     *string           
	JobName                   *string           
	JobEndTime                *float64          
	ParentJobId               *string           
	ParentJobName             *string           
	ParentJobStatus           *string           
	BenchmarkId               *int64            
	DeploymentEnvironmentId   *string           
	DeploymentEnvironmentName *string           
	AppId                     *string           
	AppName                   *string           
	AppVersion                *string           
	AppLogoPath               *string           
	ServiceId                 *string           
	ServiceName               *string           
	Tags                      *[]string         
	PublicIpAddresses         *string           
	PrivateIpAddresses        *string           
	CloudCost                 *float64          
	NodeHours                 *float64          
	UserFavorite              *bool             
	RecordTimestamp           *float64          
	ImageId                   *string           
	TerminateProtection       *bool             
	ImportedTime              *float64          
	Running                   *bool             
	RunTime                   *float64          
	AgingPolicy               *string           
	ScalingPolicy             *string           
	SecurityProfile           *string           
	Storage                   *string           
	StorageIP                 *string           
	NodeStatus                *string           
	Type                      *string           
	Actions                   *[]Action         
}
```

```go
type Filters struct {
	CloudFamilies          *[]Filter 
	Groups                 *[]Filter 
	Regions                *[]Filter 
	CloudAccounts          *[]Filter 
	UserNames              *[]Filter 
	OSNames                *[]Filter 
	MemorySizes            *[]Filter 
	CPUses                 *[]Filter 
	StorageSizes           *[]Filter 
	AppNames               *[]Filter 
	ServiceNames           *[]Filter 
	DeploymentEnvironments *[]Filter 
	Tags                   *[]Filter 
	Statuses               *[]Filter 
	NodeStatuses           *[]Filter 
	ParentJobStatuses      *[]Filter 
	CloudAndAccountNames   *[]Filter 
}
```

```go
type Filter struct {
	DisplayName *string  
	Field       *string  
	Value       *float64 
}
```

```go
type NIC struct {
	Name             *string 
	PublicIpAddress  *string 
	PrivateIpAddress *string 
	Index            *int64  
}
```

```go
type NodeProperty struct {
	Name  *string  
	Value *float64 
}
```

#### GetVirtualMachines

```go
func (s *Client) GetVirtualMachines() ([]VirtualMachineDetails, error)
```

##### Example

```golang
vms, err := client.GetVirtualMachines()

if err != nil {
	fmt.Println(err)
} else {
	for _, vm := range vms {
		vmId := *vm.Id
		vmName := *vm.Name
		fmt.Println("Id: " + vmId + ", Name: " + vmName)
	}
}
```

#### GetVirtualMachine

```go
func (s *Client) GetVirtualMachine(virtualMachineId int) (*VirtualMachineDetails, error)
```

##### Example

```golang
vm, err := client.GetVirtualMachine(14)

if err != nil {
	fmt.Println(err)
} else {
	vmId := *vm.Id
	vmName := *vm.Name
	fmt.Println("Id: " + vmId + ", Name: " + vmName)
}
```

#### GetVirtualMachineCostSummary

```go
func (s *Client) GetVirtualMachineCostSummary() (*CostSummary, error)
```

##### Example

```golang
costSummary, err := client.GetVirtualMachineCostSummary()

if err != nil {
	fmt.Println(err)
} else {
	totalNumberOfVMs := int(*costSummary.TotalNumberOfVMs)
	fmt.Println("TotalNumberOfVMs: " + strconv.Itoa(totalNumberOfVMs))
}
```

## License

This project is licensed to you under the terms of the [Cisco Sample
Code License](./LICENSE).
