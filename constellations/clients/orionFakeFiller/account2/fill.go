package account2

import (
	"github.com/ahsu1230/mathnavigatorSite/constellations/clients/orionFakeFiller/utils"
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
func Fill() {
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
	utils.AddUser(guardianJson)

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
	utils.AddUser(studentJson)
	studentJson = utils.CreateUserStudentJson(
		accountId,
		"Marcus",
		"",
		"Chang",
		"marcuschangstudent@gmail.com",
		"301-555-4447",
		"Son of Chang family",
		"Thomas Wootton High School",
		2022,
	)
	utils.AddUser(studentJson)

	// Add two transactions to account
	utils.CreateTransaction(
		accountId,
		-400,
		"charge",
		"Michelle Chang enrolled into AP Calculus",
	)
	utils.CreateTransaction(
		accountId,
		-400,
		"charge",
		"Michelle Chang enrolled into AP Java",
	)
	utils.CreateTransaction(
		accountId,
		-400,
		"charge",
		"Michael Chang enrolled into AP Java",
	)
	utils.CreateTransaction(
		accountId,
		800,
		"pay_cash",
		"Cash payment for Michelle Chang",
	)

	// Enroll student into classes
	// ASSUMPTION! We are assuming this is the first account
	// TODO: When CreateUser returns the userId, use that instead.
	// For Michelle
	userId := 5
	utils.CreateUserClass(
		accountId,
		userId,
		"ap_calculus_2021_summer_class1",
		0,
	)
	utils.CreateUserClass(
		accountId,
		userId,
		"ap_java_2021_summer_class1",
		0,
	)

	// For Michael
	userId = 6
	utils.CreateUserClass(
		accountId,
		userId,
		"ap_java_2021_summer_class1",
		0,
	)
}
