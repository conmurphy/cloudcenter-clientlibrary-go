# Cloudcenter Go Client Library

This is a Go Client Library used for accessing Cisco CloudCenter. 

It is currently a __Proof of Concept__ and has been developed and tested against Cisco CloudCenter 4.8.2 with Go version 1.9.3

![alt tag](https://github.com/conmurphy/cloudcenter-clientlibrary-go/blob/master/images/overview.png)

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



### Bundles
### Client
### CloudAccounts
### CloudImageMapping
### CloudInstanceTypes
### CloudRegions
### CloudStorageTypes
### Clouds
### Contracts
### Environments
### Groups
### Images
### Jobs
### OperationStatus
### Phases
### Plans
### Projects
### Roles
### Services

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

### SuspensionPolicies
### Tenants

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

### Users

- [GetUsers](#getusers)
- [GetUser](#getuser)
- [GetUserByEmail](#getuserbyemail)
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

/*
	Create user
*/

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

#### DeleteUser

```go
func (s *Client) DeleteUser(userId int) error
```
##### Example
```go
fmt.Println()
/*
	Delete user
*/

err := client.DeleteUser(6)

if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("User deleted")
}
```
### VirtualMachines


WARNING:

These scripts are meant for educational/proof of concept purposes only. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
