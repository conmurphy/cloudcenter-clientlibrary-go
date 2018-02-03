package main

//import "github.com/cloudcenter-clientlibrary-go/cloudcenter"

//import "fmt"
//import "strconv"

func example_update() {

	/*
		Define new cloudcenter client
	*/

	//client := cloudcenter.NewClient("USERNAME", "API_KEY", "https://CLOUDCENTER.URL")

	/****************************************

				EXAMPLES - UPDATE

	****************************************/

	/*
			Update user


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Update user")
		fmt.Println("************************************************")
		fmt.Println()

		newUser := cloudcenter.User{
			Id:        "27",
			FirstName: "Client",
			LastName:  "Library",
			Password:    "myPassword",
			EmailAddr:     "admin@clientlibrary.com",
			CompanyName:   "Company",
			PhoneNumber:   "12345",
			ExternalId:    "23456",
			TenantId:      "1",
			AccountSource: "AdminCreated",
			Type:          "STANDARD",
			Username:      "clientlibrary_1",
		}

		user, err := client.UpdateUser(&newUser)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("UserId: " + user.Id + ", Username: " + user.LastName + ", TenantId: " + user.TenantId)
		}

	*/

	/*
			Update plan


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Update plan")
		fmt.Println("************************************************")
		fmt.Println()

		newPlan := cloudcenter.Plan{

			Id:              "3",
			TenantId:        "1",
			Name:            "Client Library plan",
			Description:     "Client Library  plan description updated",
			Type:            "UNLIMITED_PLAN",
			ShowOnlyToAdmin: false,
			Price:           5,
			OnetimeFee:      5,
			BillToVendor:    false,
		}

		plan, err := client.UpdatePlan(&newPlan)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Plan Id: " + plan.Id + ", Disabled: " + strconv.FormatBool(plan.Disabled))
		}

	*/

	/*
			Update contract


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Update contract")
		fmt.Println("************************************************")
		fmt.Println()

		newContract := cloudcenter.Contract{
			Id:           "2",
			TenantId:     "1",
			Name:         "Client Library contract",
			Length:       12,
			Terms:        "Client Library  contract terms updated",
			DiscountRate: 50,
		}



		contract, err := client.UpdateContract(&newContract)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Contract Id: " + contract.Id + ", Disabled: " + strconv.FormatBool(contract.Disabled))
		}

	*/

	/*
			Update bundle


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Update bundle")
		fmt.Println("************************************************")
		fmt.Println()

		newBundle := cloudcenter.Bundle{

			Id:             "3",
			TenantId:       "1",
			Name:           "Client Library Bundle",
			Type:           "BUDGET_BUNDLE",
			Limit:          1,
			Price:          1,
			ExpirationDate: 2580679359000,
		}

		bundle, err := client.UpdateBundle(&newBundle)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bundle Id: " + bundle.Id + ", Disabled: " + strconv.FormatBool(bundle.Disabled))
		}

	*/

	/*
			Update activation profile



		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Update activation profile")
		fmt.Println("************************************************")
		fmt.Println()

		var activateRegions []cloudcenter.ActivateRegion

		newActivateRegion := cloudcenter.ActivateRegion{
			RegionId: "1",
		}

		activateRegions = append(activateRegions, newActivateRegion)

		newActivationProfile := cloudcenter.ActivationProfile{

			TenantId:        1,
			Id:              "1",
			Name:            "Client Library activation profile",
			Description:     "Client Library activation profile description updated",
			PlanId:          "2",
			BundleId:        "1",
			ContractId:      "1",
			DepEnvId:        "1",
			ActivateRegions: activateRegions,
		}

		activationProfile, err := client.UpdateActivationProfile(&newActivationProfile)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Activation Profile Id: " + activationProfile.Id + ", Description: " + activationProfile.Description)
		}

	*/
}
