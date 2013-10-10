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
	FromNumber  string
	ToNumber    string
	AnswerURL   string
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
	testAccount = fmt.Sprintf("test_acc_%d", currTime)
	config, err := loadJsonConfig("test_config.json")
	if err != nil {
		fmt.Println("Did not load config file")
	}
	authID = config["auth_id"]
	authToken = config["auth_token"]

	FromNumber = config["from_number"]
	ToNumber = config["to_number"]
	AnswerURL = config["answer_url"]
}

func TestAccountGet(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	acc, _, err := client.Account.Get()
	if err != nil {
		t.Errorf("AccountGet failed: %v", err)
	} else {
		t.Logf("Account: %v\n", acc)
	}
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
	_, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("AccountCreateSubaccount failed: %v", err)
	} else {
		t.Logf("Account: %v\n", sacc)
	}
}

func TestAccountModifySubaccount(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	sacc := &Subaccount{Name: testAccount + "_mod", Enabled: false}
	 _, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("TestAccountModifySubaccount failed at account creation: %v", err)
	}
	sacc.Enabled = true
	sacc, _, err = client.Account.ModifySubaccount(sacc)
	if err != nil {
		t.Errorf("AccountModifySubaccount failed at account modification: %v", err)
	} else {
		t.Logf("Account: %v\n", sacc)
	}
}

func TestAccountGetSubaccount(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	sacc := &Subaccount{Name: testAccount + "_get", Enabled: false}
	_, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("TestAccountGetSubaccount failed at account creation: %v", err)
	}

	sacc, _, err = client.Account.GetSubaccount(sacc.AuthID)
	if err != nil {
		t.Errorf("AccountGetSubaccount failed: %v", err)
	} else {
		t.Logf("Account: %v\n", sacc)
	}
}

func TestAccountGetSubaccounts(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	for i := 0; i < 2; i++ {
		sacc := &Subaccount{Name: testAccount + fmt.Sprintf("_get_mult_%d", i), Enabled: false}
		_, err := client.Account.CreateSubaccount(sacc)
		if err != nil {
			t.Errorf("TestAccountGetSubaccounts failed at account creation: %v", err)
		}
	}
	sacc, _, err := client.Account.GetSubaccounts(0, 0)
	if err != nil {
		t.Errorf("AccountGetSubaccounts failed: %v", err)
	} else {
		t.Logf("Account: %v\n", sacc)
	}
}

func TestAccountDeleteSubaccount(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	sacc := &Subaccount{Name: testAccount + "_del", Enabled: false}
	_, err := client.Account.CreateSubaccount(sacc)
	if err != nil {
		t.Errorf("TestAccountDeleteSubaccount failed at account creation: %v", err)
	}

	_, err = client.Account.DeleteSubaccount(sacc.AuthID)
	if err != nil {
		t.Errorf("AccountDeleteSubaccount failed: %v", err)
	} else {
		t.Logf("Account deleted")
	}
}

func TestApplicationCreate(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	app := &Application{AnswerURL: AnswerURL, AppName: "Test App (Create)"}
	app, _, err := client.Application.Create(app)
	if err != nil {
		t.Errorf("ApplicationCreate failed: %v", err)
	}
	t.Logf("Application: %v\n", app)
}

func TestApplicationGetApplications(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	apps, _, err := client.Application.GetApplications(0, 0)
	if err != nil {
		t.Errorf("ApplicationGetApplications failed: %v", err)
	} else {
		t.Logf("Account: %v\n", apps)
	}
}

func TestApplicationGet(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	app := &Application{AnswerURL: "http://example.com/answer/", AppName: "Test App (Get)"}
	app, _, err := client.Application.Create(app)
	if err != nil {
		t.Errorf("ApplicationGet failed at application creation: %v", err)
	}

	app, _, err = client.Application.Get(app.AppID)
	if err != nil {
		t.Errorf("ApplicationGet failed: %v", err)
	} else {
		t.Logf("Account: %v\n", app)
	}
}

func TestApplicationDelete(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)

	app := &Application{AnswerURL: AnswerURL, AppName: "Test App (Delete)"}
	app, _, err := client.Application.Create(app)
	if err != nil {
		t.Errorf("ApplicationDelete failed at application creation: %v", err)
	}

	_, err = client.Application.Delete(app.AppID)
	if err != nil {
		t.Errorf("ApplicationDelete failed: %v", err)
	} else {
		t.Logf("Application deleted")
	}
}

func TestCallMake(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	cp := &CallMakeParams{From: FromNumber, To: ToNumber, AnswerURL: AnswerURL}
	_, err := client.Call.Make(cp)
	if err != nil {
		t.Errorf("CallMake failed: %v", err)
	}
	t.Logf("Call Parameters: %v\n", cp)
}

func TestCallGetAll(t *testing.T) {
	setup()
	client = NewClient(authID, authToken)
	cp := &CallGetAllParams{}
	calls, _, err := client.Call.GetAll(cp)
	if err != nil {
		t.Errorf("CallGetAll failed: %v", err)
	}
	t.Logf("Calls: %v\n", calls)
}

// TODO: Iterate through calls returned by GetAll to fix TestCallGet
// func TestCallGet(t *testing.T) {
// 	setup()
// 	client = NewClient(authID, authToken)
// 	c := &Call{CallUUID: ""}
// 	call, _, err := client.Call.Get(c.CallUUID)
// 	if err != nil {
// 		t.Errorf("CallGet failed: %v", err)
// 	}
// 	t.Logf("Call: %v\n", call)
// }

// TODO: Getting Live Calls isn't quite straightforward to test, but unlikely to be broken.
// func TestCallGetAllLive(t *testing.T) {
// 	setup()
// 	client = NewClient(authID, authToken)
// 	calls, _, err := client.Call.GetAllLive()
// 	if err != nil {
// 		t.Errorf("CallGetAllLive failed: %v", err)
// 	}
// 	t.Logf("Calls: %v\n", calls)
// }
