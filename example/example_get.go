package main

import "github.com/cloudcenter-clientlibrary-go/cloudcenter"
import "fmt"

func example_get() {

	/*
		Define new cloudcenter client
	*/

	client := cloudcenter.NewClient("USERNAME", "API_KEY", "https://CLOUDCENTER.URL")

	/****************************************

				EXAMPLES - GET

	****************************************/

	/*
			Retrieve all users


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all users")
		fmt.Println("************************************************")
		fmt.Println()

		users, err := client.GetUsers()

		if err != nil {
			fmt.Println(err)
		} else {
			for _, user := range users {

				fmt.Println("UserId: " + user.Id + ", Username: " + user.Username + ", TenantId: " + user.TenantId)

			}
		}

	*/

	/*
			Retrieve all details about a user


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details about a user ")
		fmt.Println("************************************************")
		fmt.Println()

		user, err := client.GetUser(1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("UserId: " + user.Id + ", Username: " + user.Username + ", TenantId: " + user.TenantId)
		}

	*/

	/*
			Retrieve all details about a user located by email address


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details about a user located by email address")
		fmt.Println("************************************************")
		fmt.Println()

		user, err = client.GetUserFromEmail("admin@clientlibrary.com")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("UserId: " + user.Id + ", Username: " + user.Username + ", TenantId: " + user.TenantId)
		}

	*/

	/*
			Retrieve all tenants


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all tenants ")
		fmt.Println("************************************************")
		fmt.Println()

		tenants, err := client.GetTenants()

		if err != nil {
			fmt.Println(err)
		} else {
			for _, tenant := range tenants {

				fmt.Println("TenantId: " + tenant.Id + ", Name: " + tenant.Name)

			}
		}

	*/

	/*
			Retrieve all details for a tenant


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for a tenant ")
		fmt.Println("************************************************")
		fmt.Println()

		tenant, err := client.GetTenant(2)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("TenantId: " + tenant.Id + ", Name: " + tenant.Name)
		}

	*/

	/*
			Retrieve all jobs


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all jobs")
		fmt.Println("************************************************")
		fmt.Println()

		jobs, err := client.GetJobs()

		if err != nil {
			fmt.Println(err)
		} else {
			for _, job := range jobs {

				fmt.Println("Id: " + job.Id + ", Name: " + job.Name)

			}
		}

	*/

	/*
			Retrieve all details for a job


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for a job")
		fmt.Println("************************************************")
		fmt.Println()

		job, err := client.GetJob(759)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Id: " + job.Id + ", Name: " + job.Name)
		}

	*/

	/*
			Retrieve all application profiles


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all application profiles ")
		fmt.Println("************************************************")
		fmt.Println()

		apps, err := client.GetApps()

		if err != nil {
			fmt.Println(err)
		} else {
			for _, app := range apps {

				fmt.Println("Id: " + app.Id + ", Name: " + app.Name)

			}
		}

	*/

	/*
			Retrieve all clouds


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all clouds ")
		fmt.Println("************************************************")
		fmt.Println()

		clouds, err := client.GetClouds(1)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, cloud := range clouds {

				fmt.Println("Id: " + cloud.Id + ", Name: " + cloud.Name)

			}
		}

	*/

	/*
			Retrieve all details about a cloud


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details about a cloud")
		fmt.Println("************************************************")
		fmt.Println()

		cloud, err := client.GetCloud(1, 1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Id: " + cloud.Id + ", Name: " + cloud.Name)
		}

	*/

	/*
			Retrieve all bundles


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all bundles ")
		fmt.Println("************************************************")
		fmt.Println()

		bundles, err := client.GetBundles(1)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, bundle := range bundles {

				fmt.Println("Id: " + bundle.Id + ", Name: " + bundle.Name)

			}
		}

	*/

	/*
			Retrieve all details for a bundle


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for a bundle")
		fmt.Println("************************************************")
		fmt.Println()

		bundle, err := client.GetBundle(1, 1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Id: " + bundle.Id + ", Name: " + bundle.Name)
		}

	*/

	/*
			Retrieve all Actions


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all actions ")
		fmt.Println("************************************************")
		fmt.Println()

		actions, err := client.GetActions()

		if err != nil {
			fmt.Println(err)
		} else {
			for _, action := range actions {

				fmt.Println("Id: " + action.Id + ", Name: " + action.Name)

			}
		}

	*/

	/*
			Retrieve all details for an action


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for an action")
		fmt.Println("************************************************")
		fmt.Println()

		action, err := client.GetAction(13)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Id: " + action.Id + ", Name: " + action.Name)
		}

	*/

	/*
			Retrieve all activation profiles


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all activation profiles")
		fmt.Println("************************************************")
		fmt.Println()

		activationProfiles, err := client.GetActivationProfiles(1)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, activationProfile := range activationProfiles {

				fmt.Println("Id: " + activationProfile.Id + ", Name: " + activationProfile.Name)

			}
		}

	*/

	/*
			Retrieve all details for an activation profile


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for an activation profile")
		fmt.Println("************************************************")
		fmt.Println()

		activationProfile, err := client.GetActivationProfile(1, 1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Id: " + activationProfile.Id + ", Name: " + activationProfile.Name)
		}

	*/

	/*
			Retrieve all plans


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all plans")
		fmt.Println("************************************************")
		fmt.Println()

		plans, err := client.GetPlans(1)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, plan := range plans {

				fmt.Println("Id: " + plan.Id + ", Name: " + plan.Name)

			}
		}

	*/

	/*
			Retrieve all details for a plan


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for a plan")
		fmt.Println("************************************************")
		fmt.Println()

		plan, err := client.GetPlan(1, 1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Id: " + plan.Id + ", Name: " + plan.Name)
		}

	*/

	/*
			Retrieve all contracts


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all contracts")
		fmt.Println("************************************************")
		fmt.Println()

		contracts, err := client.GetContracts(1)

		if err != nil {
			fmt.Println(err)
		} else {
			for _, contract := range contracts {

				fmt.Println("Id: " + contract.Id + ", Name: " + contract.Name)

			}
		}

	*/

	/*
			Retrieve all details for a contract


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for a contract")
		fmt.Println("************************************************")
		fmt.Println()

		contract, err := client.GetContract(1, 1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Id: " + contract.Id + ", Name: " + contract.Name)
		}

	*/

	/*
			Retrieve all environments


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all environments")
		fmt.Println("************************************************")
		fmt.Println()

		environments, err := client.GetEnvironments()

		if err != nil {
			fmt.Println(err)
		} else {

			for _, environment := range environments {

				fmt.Println("Id: " + environment.Id + ", Name: " + environment.Name)

			}
		}

	*/

	/*
			Retrieve all details for an environments


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Retrieve all details for an environments")
		fmt.Println("************************************************")
		fmt.Println()

		environment, err := client.GetEnvironment(1)

		if err != nil {
			fmt.Println(err)
		} else {

			fmt.Println("Id: " + environment.Id + ", Name: " + environment.Name)
		}

	*/

}
