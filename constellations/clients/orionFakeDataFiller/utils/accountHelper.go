package utils

import (
	"fmt"
	"log"
	"strings"
)

func CreateAccount(hostAddress, accountJson, userJson string) error {
	body := strings.NewReader(fmt.Sprintf(`{
		"account": %s,
		"user": %s
	}`, accountJson, userJson))
	log.Println("Creating account and primary user...")
	SendPostRequest(hostAddress+"/api/accounts/create", body)
	return nil
}

func CreateAccountJson(primaryEmail, password string) string {
	return fmt.Sprintf(`{
		"primaryEmail": "%s",
		"password": "%s"
	}`, primaryEmail, password)
}

func CreateUserStudentJson(
	accountId int,
	firstName, middleName, lastName, email, phone string,
	notes string,
	school string,
	graduationYear int) string {
	return createUserJson(
		accountId,
		firstName,
		middleName,
		lastName,
		email,
		phone,
		false,
		school,
		graduationYear,
		notes,
	)
}

func CreateUserGuardianJson(
	accountId int,
	firstName, middleName, lastName, email, phone string,
	notes string) string {
	return createUserJson(
		accountId,
		firstName,
		middleName,
		lastName,
		email,
		phone,
		true,
		"",
		0,
		notes,
	)
}

func createUserJson(
	accountId int,
	firstName, middleName, lastName, email, phone string,
	isGuardian bool,
	school string,
	graduationYear int,
	notes string) string {
	return fmt.Sprintf(`{
		"accountId": %d,
		"firstName": "%s",
		"middleName": "%s",
		"lastName": "%s",
		"email": "%s",
		"phone": "%s",
		"notes": "%s",
		"isGuardian": %t,
		"school": "%s",
		"graduationYear": %d
	}`,
		accountId,
		firstName,
		middleName,
		lastName,
		email,
		phone,
		notes,
		isGuardian,
		school,
		graduationYear,
	)
}

func AddUser(hostAddress string, userJson string) error {
	log.Println("Creating another user...")
	body := strings.NewReader(userJson)
	SendPostRequest(hostAddress+"/api/users/create", body)
	return nil
}

func CreateTransaction(hostAddress string, accountId int, amount int, paymentType string, paymentNotes string) error {
	transactionBody := strings.NewReader(fmt.Sprintf(`{
		"amount": %d,
		"paymentType": "%s",
		"paymentNotes": "%s",
		"accountId": %d
		}`,
		amount,
		paymentType,
		paymentNotes,
		accountId,
	))
	log.Printf("Creating transaction for accountId %d, %s\n", accountId, paymentNotes)
	SendPostRequest(hostAddress+"/api/transactions/create", transactionBody)
	return nil
}
