//package test

package test

import (
	"TheLabSystem/Dao/UserDao"
	"TheLabSystem/Types/User"
	"fmt"
	"sync"
)

// tested successfully
func testInsertUser(wt *sync.WaitGroup) {
	var user = User.User{
		UserType:    1,
		Username:    "OceanCT",
		DisplayName: "OceanCT",
		Password:    "123455",
	}
	err := UserDao.InsertUser(user)
	if err != nil {
		fmt.Println(err)
	}
	wt.Done()
}

// tested successfully
func testDeleteUser() {
	var user = User.User{
		UserID: 2,
	}
	err := UserDao.DeleteUser(user)
	if err != nil {
		fmt.Println(err)
	}
}

// tested successfully
func testFindUser() {
	fmt.Println(UserDao.FindUserByUsername("OceanCT"))
	fmt.Println(UserDao.FindUserByID(3))
}

//func test UpdateUser()
func testUpdateUser() {
	user := User.User{UserID: 3, UserType: 2, Username: "sdafs"}
	fmt.Println(UserDao.UpdateUser(user))
}
func main() {
	//testUpdateUser()
	//testFindUser()
	testDeleteUser()
	//waiter := &sync.WaitGroup{}
	//waiter.Add(90)
	//for i := 1; i <= 90; i++ {
	//	go testInsertUser(waiter)
	//}
	//waiter.Wait()
	//users, err := UserDao.FindUserByOffset(1, 100)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	for key := range users {
	//		fmt.Println(users[key])
	//	}
	//}
}
