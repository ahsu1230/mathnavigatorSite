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
func Fill(afhId1, afhId2 uint) {
	log.Println("Fill account2")

	// Create Account with primary user
	accountJson := utils.CreateAccountJson(
		"marychangmom@gmail.com",
		"asdf1234",
	)
	guardianJson := utils.CreateUserGuardianJson(
		0, // will be filled automatically by endpoint
		"Mary",
		"Mei-Li",
		"Chang",
		"marychangmom@gmail.com",
		"301-555-4444",
		"Mother of Chang family",
	)
	accountId, _ := utils.CreateAccount(
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
	userId2, _ := utils.AddUser(studentJson)

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
	userId3, _ := utils.AddUser(studentJson)

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
	// For Michelle
	utils.CreateUserClass(
		accountId,
		userId2,
		"ap_calculus_2021_summer_class1",
		0,
	)
	utils.CreateUserClass(
		accountId,
		userId2,
		"ap_java_2021_summer_class1",
		0,
	)

	// For Michael
	utils.CreateUserClass(
		accountId,
		userId3,
		"ap_java_2021_summer_class1",
		0,
	)

	// Both will attend the first two AP Java AFH sessions
	utils.CreateUserAFH(
		accountId,
		userId2,
		afhId1,
	)
	utils.CreateUserAFH(
		accountId,
		userId2,
		afhId2,
	)
	utils.CreateUserAFH(
		accountId,
		userId3,
		afhId1,
	)
	utils.CreateUserAFH(
		accountId,
		userId3,
		afhId2,
	)
}
