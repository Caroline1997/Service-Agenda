package entity

import (
    "fmt"
    "net/http"
)

type User struct{
    Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
    Phone string `json:"phonenumber"`
}

// for register
func CreateUser(user *User) (err error) {
	  var code int
    var responseBody struct {
        Message string `json:"message"`
    }
    code, err = request("POST", "/api/users", user, &responseBody)
    if err != nil {
        return 
    }
    // 201
    if code == http.StatusCreated {
        return
    }
    err = fmt.Errorf("%s", responseBody.Message)
    return
}

// for delete user
func DeleteUser() (err error) {
    var code int
    code, err = request("DELETE", "/api/user/account", nil, nil)
    if err != nil {
        return
    }
    // 200
    if code == http.StatusOK {
        return
    }
    err = fmt.Errorf("%d", code)
    return
}

func ListAllUsers() (user []User, err error) {
    var code int
    code, err = request("GET", "/api/users", nil, &user)
    if err != nil {
        return
    }
    // 200
    if code == http.StatusOK {
        return user
    }
    err = fmt.Errorf("%d", code)
    return
}
