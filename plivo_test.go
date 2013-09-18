package plivo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

var (
	// client is the API client being tested.
	client      *Client
	currTime    int64
	testAccount string
	config      map[string]string
	authID      string
	authToken   string
)

func loadJsonConfig(path string) (map[string]string, error) {
	contents, err := ioutil.ReadFile(path)
	var config map[string]string
	err = json.Unmarshal(contents, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func setup() {
	currTime = time.Now().Unix()
	testAccount = fmt.Sprintf("test_account_%d", currTime)
	config, err := loadJsonConfig("test_config.json")
	if err != nil {
		fmt.Println("Did not load config file")
	}
	authID = config["auth_id"]
	authToken = config["auth_token"]
}

func TestAccountGet(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	acc, _, err := client.Account.Get()
	if err != nil {
		t.Errorf("AccountGet failed: %v", err)
	}
	t.Logf("Account: %v\n", acc)
}

func TestAccountModify(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	acc := &Account{Name: "Test Name", City: "Test City", Address: "Test Address", AuthID: authID}
	acc, _, err := client.Account.Modify(acc)
	if err != nil {
		t.Errorf("AccountModify failed: %v", err)
	}
	t.Logf("Account: %v\n", acc)
}

func TestAccountCreateSubaccount(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	sacc := &Subaccount{Name: testAccount, Enabled: false}
	sacc, _, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("AccountCreateSubaccount failed: %v", err)
	}
	t.Logf("Account: %v\n", sacc)
}

func TestAccountModifySubaccount(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	sacc := &Subaccount{Name: testAccount + "_modify", Enabled: false}
	sacc, _, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("TestAccountModifySubaccount failed at account creation: %v", err)
	}
	sacc.Enabled = true
	sacc, _, err = client.Account.ModifySubaccount(sacc)
	if err != nil {
		t.Errorf("AccountModifySubaccount failed at account modification: %v", err)
	}
	t.Logf("Account: %v\n", sacc)
}

func TestAccountGetSubaccount(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	sacc := &Subaccount{Name: testAccount + "_get", Enabled: false}
	sacc, _, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("TestAccountGetSubaccount failed at account creation: %v", err)
	}

	sacc, _, err = client.Account.GetSubaccount(sacc.AuthID)
	if err != nil {
		t.Errorf("AccountGetSubaccount failed: %v", err)
	}
	t.Logf("Account: %v\n", sacc)
}

func TestAccountGetSubaccounts(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	for i := 0; i < 5; i++ {
		sacc := &Subaccount{Name: testAccount + fmt.Sprintf("_get_plural_%d", i), Enabled: false}
		sacc, _, err := client.Account.CreateSubaccount(sacc)
		if err != nil {
			t.Errorf("TestAccountGetSubaccount failed at account creation: %v", err)
		}
	}
	sacc, _, err := client.Account.GetSubaccounts(0, 0)
	if err != nil {
		t.Errorf("AccountGetSubaccount failed: %v", err)
	}
	t.Logf("Account: %v\n", sacc)
}

func TestAccountDeleteSubaccount(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	sacc := &Subaccount{Name: testAccount + "_delete", Enabled: false}
	sacc, _, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("TestAccountDeleteSubaccount failed at account creation: %v", err)
	}

	_, err = client.Account.DeleteSubaccount(sacc.AuthID)
	if err != nil {
		t.Errorf("AccountDeleteSubaccount failed: %v", err)
	}
	t.Logf("Account deleted")
}
