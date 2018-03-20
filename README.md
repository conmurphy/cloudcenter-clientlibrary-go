# Cloudcenter Go Client Library

This is a Go Client Library used for accessing Cisco CloudCenter. 

It is currently a __Proof of Concept__ and has been developed and tested against Cisco CloudCenter 4.8.2 with Go version 1.9.3

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/overview.png)

Table of Contents
=================

  * [Cloudcenter Go Client Library](#cloudcenter-go-client-library)
      * [Quick Start](#quick-start)
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
	fmt.Println(”New user created. \n UserId: " + user.Id + ", User last name: " + user.LastName)
}
```

## Quick Start - Creation from JSON file

For some situations it may be easier to have the configuration represented as JSON rather than conifguring individually as per the two examples above. In this scenario you can either build the JSON file yourself or monitor the API POST call for the JSON data sent to CloudCenter. This can be achieved using the browsers built in developer tools.

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

### Actions

- [GetActions](#getactions)
- [GetAction](#getaction)

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

		fmt.Println("Id: " + action.Id + ", Name: " + action.Name)

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
	fmt.Println("Id: " + action.Id + ", Name: " + action.Name)
}
```

### ActivationProfiles

- [GetActivationProfiles](#getactivationprofiles)
- [GetActivationProfile](#getactivationprofile)

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

		fmt.Println("Id: " + activationProfile.Id + ", Name: " + activationProfile.Name)

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
	fmt.Println("Id: " + activationProfile.Id + ", Name: " + activationProfile.Name)
}
```

### AgingPolicies

- [GetAgingPolicies](#getagingpolicies)
- [GetAgingPolicy](#getagingpolicy)

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

		fmt.Println("Id: " + agingPolicy.Id + ", Name: " + agingPolicy.Name)

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
	fmt.Println("Id: " + agingPolicy.Id + ", Name: " + agingPolicy.Name)
}
```

### Apps

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
		fmt.Println("Id: " + app.Id + ", Name: " + app.Name)
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
	fmt.Println("Id: " + app.Id + ", Name: " + app.Name)
}
```


### Bundles

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
		fmt.Println("Id: " + bundle.Id + ", Name: " + bundle.Name)
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
	fmt.Println("Id: " + bundle.Id + ", Name: " + bundle.Name)
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
	fmt.Println("Id: " + bundle.Id + ", Name: " + bundle.Name)
}
```

### CloudAccounts

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
		fmt.Println("Id: " + cloudAccount.Id + ", Name: " + cloudAccount.DisplayName)
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
	fmt.Println("Id: " + cloudAccount.Id + ", Name: " + cloudAccount.DisplayName)
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
		fmt.Println("Id: " + cloudAccount.Id + ", Name: " + cloudAccount.AccountName + ", Name: " + cloudAccount.DisplayName)
	}
}
```

### CloudImageMapping

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
		fmt.Println("Id: " + cloudImage.Id + ", Resource: " + cloudImage.Resource)
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
	fmt.Println("Id: " + cloudImage.Id + ", Resource: " + cloudImage.Resource)
}
```

### CloudInstanceTypes

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
		fmt.Println("Id: " + cloudInstanceType.Id + ", Resource: " + cloudInstanceType.Resource)
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
	fmt.Println("Id: " + cloudInstanceType.Id + ", Resource: " + cloudInstanceType.Resource)
}
```

### CloudRegions

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
		fmt.Println("Id: " + cloudRegion.Id + ", Name: " + cloudRegion.RegionName)
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
	fmt.Println("Id: " + cloudRegion.Id + ", Name: " + cloudRegion.RegionName)
}
```

### CloudStorageTypes

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
		fmt.Println("Id: " + cloudStorageType.Id + ", Resource: " + cloudStorageType.Resource)
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
	fmt.Println("Id: " + cloudStorageType.Id + ", Name: " + cloudStorageType.Name)
}
```

### Clouds

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
		fmt.Println("Id: " + cloud.Id + ", Name: " + cloud.Name)
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
	fmt.Println("Id: " + cloud.Id + ", Name: " + cloud.Name)
}
```

### Contracts

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

### Environments

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
		fmt.Println("Id: " + environment.Id + ", Name: " + environment.Name)
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
	fmt.Println("Id: " + environment.Id + ", Name: " + environment.Name)
}
```

### Groups

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
		fmt.Println("Id: " + group.Id + ", Name: " + group.Name)
		for _, user := range group.Users {
			fmt.Println("Id: " + user.Id + ", Name: " + user.Username)
		}
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
	fmt.Println("Id: " + group.Id + ", Name: " + group.Name)
	for _, user := range group.Users {
		fmt.Println("Id: " + user.Id + ", Name: " + user.Username)
	}
}
```

### Images

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
		fmt.Println("Id: " + image.Id + ", Resource: " + image.Resource)
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
	fmt.Println("Id: " + image.Id + ", Name: " + image.Name)
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
		fmt.Println("Id: " + phase.Id + ", Name: " + phase.Name)
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
	fmt.Println("Id: " + phase.Id + ", Name: " + phase.Name)
}
```


### Plans

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
		fmt.Println("Id: " + plan.Id + ", Name: " + plan.Name)
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
	fmt.Println("Id: " + plan.Id + ", Name: " + plan.Name)
}
```

### Projects

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
		fmt.Println("Id: " + project.Id + ", Name: " + project.Name)
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
	fmt.Println("Id: " + project.Id + ", Name: " + project.Name)
}
```


### Roles

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
		fmt.Println("Id: " + role.Id + ", Name: " + role.Name)
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
	fmt.Println("Id: " + role.Id + ", Name: " + role.Name)
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
		fmt.Println("Id: " + service.Id + ", DisplayName: " + service.DisplayName)
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
	fmt.Println("Id: " + service.Id + ", Name: " + service.Name)
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
		fmt.Println("Id: " + tenant.Id + ", Name: " + tenant.Name)
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
	fmt.Println("Id: " + tenant.Id + ", Name: " + tenant.Name)
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
	UserId:                          "5",
	Name:                            "client-library-tenant",
	ShortName:                       "client-library-tenant",
	DomainName:                      "clientlibrary.cloudcenter.com",
	Phone:                           "1234567890",
	Url:                             "http://clientlibrary.cloudcenter.com",
	ContactEmail:                    "poweruser@dcloud.cisco.com",
	About:                           "clientlibrary tenant",
	EnablePurchaseOrder:             false,
	EnableEmailNotificationsToUsers: false,
	EnableMonthlyBilling:            false,
	DefaultChargeType:               "Hourly",
}

fmt.Println(client.AddTenant(&newTenant))
```

#### UpdateTenant

```go
func (s *Client) UpdateTenant(tenant *Tenant) (*Tenant, error)
```

##### __Required Fields__
* Id 
  * (Value of field should not be changed)
* UserId 
  * (Value of field should not be changed)
* Name 
  * (Value of field should not be changed)
* ShortName 
  * (Value of field should not be changed)

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
	fmt.Println(”New user created. \n UserId: " + user.Id + ", User last name: " + user.LastName)
}
```

#### UpdateUser

```go
func (s *Client) UpdateUser(user *User) (*User, error)
```

##### __Required Fields__
* Id 
  * (Value of field should not be changed)
* TenantId 
  * (Value of field should not be changed)
* Username 
  * (Value of field should not be changed)
* Type 
  * (Value of field should not be changed)
* EmailAddr


##### Example

```golang

newUser := cloudcenter.User{
	Id:        cloudcenter.String("2"),
	TenantId:  cloudcenter.String("1"),
	Username:  cloudcenter.String("cliqradmin"),
	Type:      cloudcenter.String("TENANT"),
	EmailAddr: cloudcenter.String("admin@cliqrtech.com"),
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


DISCLAIMER:

These scripts are meant for educational/proof of concept purposes only. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
