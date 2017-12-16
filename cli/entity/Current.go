package entity

import (
    "fmt"
    "net/http"
)


// for log in
func Login(username string, password string) (err error) {
    requestBody := struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }{username, password}
    var responseBody struct {
        //Openid string `json:"openid"`
        Message string `json:"message"`
    }
    var code int
    code, err = request("POST", "/api/user/login", &requestBody, &responseBody)
    if err != nil {
        return
    }
    // 200
	  if code == http.StatusOK {
        return
    }
    err = fmt.Errorf("%s", responseBody.Message)
    return
}

// for log out
func Logout() (err error) {
    var code int
    code, err = request("POST", "/api/user/logout", nil, nil)
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

// check if there exist current user
func Check_Login() (flag bool, err error) {
    var responseBody struct {
        Username string `json:"username"`
    }
    var code int
    code, err = request("GET", "/api/user/login", nil, &responseBody)
    if err != nil {
        return
    }
    flag = false
    // 200
    if code == http.StatusOK {
        flag = true
        return flag
    }
    return flag
}

// get current username
func GetCurrentUser() (username string, err error) {
    var responseBody struct {
        Username string `json:"username"`
    }
    var code int
    code, err = request("GET", "/api/user/login", nil, &responseBody)
    if err != nil {
        return
    }
    // 200
    if code == http.StatusOK {
        username = responseBody.Username
    }
    // 401
    if code == http.StatusUnauthorized {
        return
    }
    return
}
