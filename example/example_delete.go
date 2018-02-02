package main

import "github.com/cloudcenter-clientlibrary-go/cloudcenter"
import "fmt"

func example_delete() {

	/*
		Define new cloudcenter client
	*/

	client := cloudcenter.NewClient("USERNAME", "API_KEY", "https://CLOUDCENTER.URL")

	/****************************************

				EXAMPLES - DELETE

	****************************************/

	/*
			Delete user


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Delete user")
		fmt.Println("************************************************")
		fmt.Println()

		err := client.DeleteUser(6)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("User deleted")
		}
	*/
	/*
			Delete user based on email


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Delete user based on email")
		fmt.Println("************************************************")
		fmt.Println()

		err := client.DeleteUserByEmail("email@clientlibrary.com")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("User deleted")
		}

	*/

	/*
			Delete activated user


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Delete activated user")
		fmt.Println("************************************************")
		fmt.Println()

		err = client.DeleteUser(1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("User deleted")
		}

	*/

	/*
			Delete activation profile


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Delete activation profile")
		fmt.Println("************************************************")
		fmt.Println()

		err = client.DeleteActivationProfile(1, 2)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Activation profile deleted")
		}

	*/

	/*
			Delete plan


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Delete plan")
		fmt.Println("************************************************")
		fmt.Println()

		err = client.DeletePlan(1, 1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Plan deleted")
		}

	*/

	/*
			Delete bundle


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Delete bundle")
		fmt.Println("************************************************")
		fmt.Println()

		err = client.DeleteBundle(1, 1)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bundle deleted")
		}

	*/

	/*
			Delete contract


		fmt.Println()
		fmt.Println("************************************************")
		fmt.Println("Delete contract")
		fmt.Println("************************************************")
		fmt.Println()

		err = client.DeleteContract(1, 2)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Contract deleted")
		}

	*/
}
