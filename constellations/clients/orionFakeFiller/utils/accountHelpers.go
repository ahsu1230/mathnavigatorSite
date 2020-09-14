package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type AccountIdResp struct {
	AccountId uint `json:"accountId"`
}

func GetAccountIdFromBody(bytes []byte) (uint, error) {
	log.Printf("*** %s\n", string(bytes))
	var resp AccountIdResp
	if err := json.Unmarshal(bytes, &resp); err != nil {
		log.Printf("unexpected error: %v\n", err)
		return 0, err
	}
	return resp.AccountId, nil
}

func CreateAccount(accountJson, userJson string) (uint, error) {
	body := strings.NewReader(fmt.Sprintf(`{
		"account": %s,
		"user": %s
	}`, accountJson, userJson))
	log.Println("Creating account and primary user...")
	respBody := SendPostRequest("/api/accounts/create", body)
	id, _ := GetAccountIdFromBody(respBody)
	return id, nil
}

func CreateAccountJson(primaryEmail, password string) string {
	return fmt.Sprintf(`{
		"primaryEmail": "%s",
		"password": "%s"
	}`, primaryEmail, password)
}

func CreateUserStudentJson(
	accountId uint,
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
	accountId uint,
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
	accountId uint,
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
		"graduationYear": %d,
		"isAdminCreated": true
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

func AddUser(userJson string) (uint, error) {
	log.Println("Creating another user...")
	body := strings.NewReader(userJson)
	respBody := SendPostRequest("/api/users/create", body)
	id, _ := GetIdFromBody(respBody)
	return id, nil
}

func CreateTransaction(
	accountId uint,
	amount int,
	paymentType string,
	paymentNotes string) (uint, error) {
	transactionBody := strings.NewReader(fmt.Sprintf(`{
		"amount": %d,
		"type": "%s",
		"notes": "%s",
		"accountId": %d
		}`,
		amount,
		paymentType,
		paymentNotes,
		accountId,
	))
	log.Printf("Creating transaction for accountId %d, %s\n", accountId, paymentNotes)
	respBody := SendPostRequest("/api/transactions/create", transactionBody)
	id, _ := GetIdFromBody(respBody)
	return id, nil
}

func CreateUserClass(
	accountId uint,
	userId uint,
	classId string,
	userClassState int,
) (uint, error) {
	body := strings.NewReader(fmt.Sprintf(`{
		"userId": %d,
		"classId": "%s",
		"accountId": %d,
		"state": %d
	}`,
		userId,
		classId,
		accountId,
		userClassState,
	))
	log.Printf("Creating relation for user '%d' and class '%s'\n", userId, classId)
	respBody := SendPostRequest("/api/user-classes/create", body)
	id, _ := GetIdFromBody(respBody)
	return id, nil
}

func CreateUserAFH(
	accountId uint,
	userId uint,
	afhId uint,
) (uint, error) {
	body := strings.NewReader(fmt.Sprintf(`{
		"userId": %d,
		"afhId": %d,
		"accountId": %d
	}`,
		userId,
		afhId,
		accountId,
	))
	log.Printf("Creating relation for user '%d' and afh '%d'\n", userId, afhId)
	respBody := SendPostRequest("/api/user-afhs/create", body)
	id, _ := GetIdFromBody(respBody)
	return id, nil
}
