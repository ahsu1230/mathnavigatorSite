package account1

import (
	"github.com/ahsu1230/mathnavigatorSite/clients/filler/utils"
	"log"
)

/*
This Account has the following features:
- One Guardian
- One Student
- Student enrolled into ap_calculus
- Account is fully paid
*/
func Fill() {
	log.Println("Fill account1")

	// ASSUMPTION! We are assuming this is the first account
	// TODO: When CreateAccount returns the accountId, use that instead.
	accountId := 1

	// Create Account with primary user
	accountJson := utils.CreateAccountJson(
		"joesmithdad@gmail.com",
		"asdf1234",
	)
	guardianJson := utils.CreateUserGuardianJson(
		accountId,
		"Joe",
		"",
		"Smith",
		"joesmithdad@gmail.com",
		"301-543-2412",
		"Father of Smith family",
	)
	utils.CreateAccount(
		accountJson,
		guardianJson,
	)

	// Add another user to account
	studentJson := utils.CreateUserStudentJson(
		accountId,
		"Jake",
		"",
		"Smith",
		"jakesmithstudent@gmail.com",
		"301-543-2424",
		"Son of Smith family",
		"Winston Churchill High School",
		2022,
	)
	utils.AddUser(studentJson)

	// Add two transactions to account
	utils.CreateTransaction(
		accountId,
		-100,
		"charge",
		"Enrolled into ap_calculus",
	)
	utils.CreateTransaction(
		accountId,
		100,
		"pay_paypal",
		"Paid for Jake's ap_calculus",
	)

	// Enroll student into classes
	// ASSUMPTION! We are assuming this is the first account
	// TODO: When CreateUser returns the userId, use that instead.
	userId := 2
	utils.CreateUserClass(
		accountId,
		userId,
		"ap_calculus_2021_summer_class1",
		0,
	)
}
