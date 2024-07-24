package tools

import (
	"time"
)


type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"sina": {
		AuthToken: "pass787B",
		Username: "sina",
	},
	"test": {
		AuthToken: "test",
		Username: "test",
	},
}

var mockBalanceDetails = map[string]BalanceDetails{
	"sina": {
		Balance: 1000,
		Username: "sina",
	},
	"test": {
		Balance: 2528,
		Username: "test",
	},	
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserBalance(username string) *BalanceDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = BalanceDetails{}
	clientData, ok := mockBalanceDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}