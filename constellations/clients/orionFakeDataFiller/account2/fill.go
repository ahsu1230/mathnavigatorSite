package account2

import (
	"github.com/ahsu1230/mathnavigatorSite/clients/filler/utils"
	"log"
)

/*
This Account has the following features:
- Two Guardians (one is primary account holder)
- Two Students
- studentA enrolled into ap_calculus & ap_java
- studentB enrolled only in ap_java
- Account is still pending payment
*/
func Fill(hostAddress string) {
	log.Println("Fill account2")

	// ASSUMPTION! We are assuming this is the second account
	// TODO: When CreateAccount returns the accountId, use that instead.
	accountId := 2

	// Create Account with primary user
	accountJson := utils.CreateAccountJson(
		"marychangmom@gmail.com",
		"asdf1234",
	)
	guardianJson := utils.CreateUserGuardianJson(
		accountId,
		"Mary",
		"Mei-Li",
		"Chang",
		"marychangmom@gmail.com",
		"301-555-4444",
		"Mother of Chang family",
	)
	utils.CreateAccount(
		hostAddress,
		accountJson,
		guardianJson,
	)

	// Add another guardian to account
	guardianJson = utils.CreateUserGuardianJson(
		accountId,
		"Michael",
		"",
		"Chang",
		"michaelchangdad@gmail.com",
		"301-555-4445",
		"Father of Chang family",
	)
	utils.AddUser(hostAddress, guardianJson)

	// Add two students to account
	studentJson := utils.CreateUserStudentJson(
		accountId,
		"Michelle",
		"",
		"Chang",
		"michellechangstudent@gmail.com",
		"301-555-4446",
		"Daughter of Chang family",
		"Thomas Wootton High School",
		2024,
	)
	utils.AddUser(hostAddress, studentJson)
	studentJson = utils.CreateUserStudentJson(
		accountId,
		"Michael",
		"",
		"Chang",
		"michaelchangstudent@gmail.com",
		"301-555-4447",
		"Son of Chang family",
		"Thomas Wootton High School",
		2022,
	)
	utils.AddUser(hostAddress, studentJson)

	// Add two transactions to account
	utils.CreateTransaction(
		hostAddress,
		accountId,
		-400,
		"charge",
		"Michelle enrolled into AP Calculus",
	)
	utils.CreateTransaction(
		hostAddress,
		accountId,
		-400,
		"charge",
		"Michelle enrolled into AP Java",
	)
	utils.CreateTransaction(
		hostAddress,
		accountId,
		-400,
		"charge",
		"Michael enrolled into AP Java",
	)
	utils.CreateTransaction(
		hostAddress,
		accountId,
		800,
		"pay_cash",
		"Cash payment for Michelle",
	)
}
