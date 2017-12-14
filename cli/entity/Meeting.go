package entity

import(
    "fmt"
    "net/http"
)

type Meeting struct {
    Sponsor string `json:"sponsor"`
    Participators []string `json:"participators"`
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    Title string `json:"title"`
}

// for create meeting
func CreateMeeting(meeting *Meeting, err error) {
   	var code int
    var responseBody struct {
        Message string `json:"message"`
    }
    code, err = request("POST", "/api/meetings", meeting, &responseBody)
    if err != nil {
        return
    }
    err = fmt.Errorf("%s", responseBody.Message)
    return
}

// for qurey meeting
func QueryMeeting(startTime string, endTime string, err error) []Meeting {
    var code int
    var responseBody []Meeting
    code, err = request("GET", "/api/meetings"+"?startTime="+startTime+"&endTime="+endTime, nil, &responseBody)
    if err != nil {
        return nil
    }
    // 200
    if code == http.StatusOK {
        return responseBody
    }
    // 401
    if code == http.StatusUnauthorized {
        return fmt.Errorf("Please log in !")
    }
    return responseBody
}
