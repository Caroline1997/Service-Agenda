package service

import (
    "fmt"
    "entity"
    "time"
)

// log's command

// Login POST /v1/user/login
func Login(name string, password string) (err error) {
    err = entity.Check_Login()
    if err != nil {
        return
    }
    if Check_Login() == true {
        var currentUser string
        currentUser, err = entity.GetCurrentUser()
        err = fmt.Errorf("There exist someone logged as '%s', please logout!", currentUser)
        return
    }
    _, err = entity.Login(username, password)
    if err != nil {
        return
    }
    return
}

// Logout POST /v1/user/logout
func Logout() (err error) {
    err = entity.Check_Login()
    if err != nil {
        return
    }
    if Check_Login() == true {
        err = entity.Logout()
        if err != nil {
            return
        }
        //entity.CurSessionModel.SetCurOpenid("")
    }
    return
}

// User's command

func check_user_valid(user *entity.User) error {
    if len(user.Username) == 0 {
        return fmt.Errorf("Please enter your username!")
    }
    if len(user.Password) == 0 {
        return fmt.Errorf("Please enter your password!")
    }
    return nil
}

// Register POST /v1/users
func Register(username string, password string, email string, phone string) (err error) {
    User1 := &entity.User {
        Username : username,
        Password : password,
        Email : email,
        Phone : phone,
    }
    err = check_user_valid(User1)
    if err != nil {
        return
    }
    err = entity.CreateUser(User1)
    return
}

// DELETE DELETE /v1/user/account
func Delete_user(username string, password string) (err error) {
    err = Check_Login()
    if err != nil {
        return err
    }
    if Check_Login() == true {
        err = entity.DeleteUser()
        if err != nil {
            return err
        }
    }
    return
}

// Meeting's command

func check_meeting_valid(meeting *entity.Meeting) error {
    if len(meeting.Title) == 0 {
        return fmt.Errorf("Please enter your title!")
    }
    if len(meeting.Participators) == 0 {
        return fmt.Errorf("Please enter at least one participator!")
    }
    return nil
}

// Create Meeting POST /v1/meetings
func Create_meeting(title string, participators []string, startTime string, endTime string) (err error) {
    err = Check_Login()
    if err != nil {
        return err
    }
    if Check_Login() == true {
        sponsor, _ := entity.GetCurrentUser()
        if err != nil {
            return
        }
        Meeting1 := &entity.Meeting{
            Sponsor : sponsor,
            Participators : participators,
            StartTime : startTime,
            EndTime : endTime,
            Title : title,
      	}
        err = check_meeting_valid(Meeting1)
        if err != nil {
            return
        }
        entity.CreateMeeting(Meeting1, err)
        return
    }
}

// Query Meetings GET /v1/meetings{?startDate,endDate}
func Query_meeting(startTime string, endTime string) ([]entity.Meeting, error) {
    err = Check_Login()
    if err != nil {
        return nil, err
    }
    if Check_Login() == true {
        // check start time
        if len(startTime) == 0 {
            return nil, fmt.Errorf("Please enter startTime!")
        }
        // 2006-01-02 15:04:05 is time format
        _, err := time.Parse("2006-01-02 15:04:05", startTime)
        if err != nil {
            return nil, fmt.Errorf("Please enter correct startTime format!")
        }
        // check end time
        if len(endTime) == 0 {
            return nil, fmt.Errorf("Please enter endTime!")
        }
        _, err = time.Parse("2006-01-02 15:04:05", endTime)
        if err != nil {
            return nil, fmt.Errorf("Please enter correct endTime format!")
        }
        // compare
        if startTime > endTime {
            return nil, fmt.Errorf("startTime should not larger than endTime!")
        }
        meetings := entity.QueryMeeting(startTime, endTime, err)
        return meetings, nil
    }
}
