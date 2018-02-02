package main

import "github.com/cloudcenter/cloudcenter"
import "fmt"
import "strconv"

func example_add() {

	/*
		Define new cloudcenter client
	*/

	client := cloudcenter.NewClient("USERNAME", "API_KEY", "https://CLOUDCENTER.URL")

	/****************************************

				EXAMPLES - ADD

	****************************************/

	/*
			Create tenant


		newPreferences := cloudcenter.Preference{
			Name:  "PASSWORD_MIN_LENGTH",
			Value: "5",
		}

		preferencesArray := make([]cloudcenter.Preference, 1)

		preferencesArray[0] = newPreferences

		newTenant := cloudcenter.Tenant{
			UserId:                          "8",
			Name:                            "client",
			ShortName:                       "library",
			DomainName:                      "clientlibrary.cloudcenter.com",
			Phone:                           "1234567890",
			ExternalId:                      "",
			Url:                             "http://clientlibrary.cloudcenter.com",
			ContactEmail:                    "client@library.cloudcenter.com",
			LoginLogo:                       ".",
			HomePageLogo:                    ".",
			About:                           "clientlibrary tenant",
			TermsOfService:                  "termsofservice",
			PrivacyPolicy:                   "Privacy policy",
			EnablePurchaseOrder:             false,
			EnableEmailNotificationsToUsers: false,
			EnableMonthlyBilling:            false,
			DefaultChargeType:               "Hourly",
			Preferences:                     preferencesArray,
		}

		fmt.Println(client.AddTenant(&newTenant))

	*/

	/*
			Create user


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Create new user")
		fmt.Println("************************************************")
		fmt.Println()

		newUser := cloudcenter.User{

			FirstName:   "client",
			LastName:    "library",
			Password:    "myPassword",
			EmailAddr:   "clientlibrary@cloudcenter.com",
			CompanyName: "Company",
			PhoneNumber: "12345",
			ExternalId:  "23456",
			TenantId:    "1",
		}

		user, err := client.AddUser(&newUser)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("UserId: " + user.Id + ", Username: " + user.LastName + ", TenantId: " + user.TenantId)
		}

	*/

	/*
			Create bundle


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Create new bundle")
		fmt.Println("************************************************")
		fmt.Println()

		newBundle := cloudcenter.Bundle{

			TenantId:       "1",
			Name:           "clientlibraryBundle",
			Type:           "BUDGET_BUNDLE",
			Limit:          1,
			Price:          1,
			ExpirationDate: 1580679359000,
		}

		bundle, err := client.AddBundle(&newBundle)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bundle Id: " + bundle.Id + ", Disabled: " + strconv.FormatBool(bundle.Disabled))
		}

	*/

	/*
			Create contract



		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Create new contract")
		fmt.Println("************************************************")
		fmt.Println()

		newContract := cloudcenter.Contract{
			TenantId:     "1",
			Name:         "ClientLibrary contract",
			Length:       12,
			Terms:        "ClientLibrary contract terms",
			DiscountRate: 50,
		}

		contract, err := client.AddContract(&newContract)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Contract Id: " + contract.Id + ", Disabled: " + strconv.FormatBool(contract.Disabled))
		}

	*/

	/*
			Create plan


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Create new plan")
		fmt.Println("************************************************")
		fmt.Println()

		newPlan := cloudcenter.Plan{

			TenantId:        "1",
			Name:            "client library plan",
			Description:     "client library plan description",
			Type:            "UNLIMITED_PLAN",
			ShowOnlyToAdmin: false,
			Price:           5,
			OnetimeFee:      5,
			BillToVendor:    false,
		}

		plan, err := client.AddPlan(&newPlan)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Plan Id: " + plan.Id + ", Disabled: " + strconv.FormatBool(plan.Disabled))
		}

	*/
	/*
			Create activation profile



		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Create new activation profile")
		fmt.Println("************************************************")
		fmt.Println()

		activateRegionsArray := make([]cloudcenter.ActivateRegion, 2)

		newActivateRegions := cloudcenter.ActivateRegion{
			RegionId: "1",
		}

		activateRegionsArray[0] = newActivateRegions

		newActivateRegions = cloudcenter.ActivateRegion{
			RegionId: "2",
		}

		activateRegionsArray[1] = newActivateRegions

		newActivationProfile := cloudcenter.ActivationProfile{

			TenantId:        1,
			Name:            "Client Library activation profile",
			Description:     "Client Library activation profile description",
			PlanId:          "3",
			BundleId:        "1",
			ContractId:      "1",
			DepEnvId:        "1",
			ActivateRegions: activateRegionsArray,
		}

		activationProfile, err := client.AddActivationProfile(&newActivationProfile)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Activation Profile Id: " + activationProfile.Id + ", Description: " + activationProfile.Description)
		}
	*/
}
