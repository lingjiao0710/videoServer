package dbops

import (
	"testing"

)


func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestUserWorkFlow(t *testing.T){
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T){
	err := AddUserCredential("admin", "123")
	if err != nil {
		t.Errorf("Error of AddUser :%v", err)
	}
}

func testGetUser(t *testing.T){
	pwd, err := GetUserCredential("admin")
	if pwd!= "123"||  err != nil {
		t.Errorf("Error of GetUser :%v", err) 
	}
}

func testDeleteUser(t *testing.T){
	err := DeleteUser("admin", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser :%v", err) 
	}
}

func testRegetUser(t *testing.T){
	pwd, err := GetUserCredential("admin")
	if err != nil {
		t.Errorf("Error of RegetUser :%v", err) 
	}

	//删除之后应当返回空值
	if pwd != ""{
		t.Errorf("Deleting user test failed")
	}
}
