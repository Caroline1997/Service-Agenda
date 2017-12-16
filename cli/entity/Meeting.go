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
func CreateMeeting(meeting *Meeting) (err error) {
   	var code int
    var requestBody struct {
        Title string `json:"title"`
        Participators []string `json:"participators"`
        StartTime string `json:"startTime"`
        EndTime string `json:"endTime"`
    }{meeting.Title, meeting.Participators, meeting.StartTime, meeting.EndTime}
    var responseBody struct {
        Message string `json:"message"`
    }
    code, err = request("POST", "/api/meetings", &requestBody, &responseBody)
    if err != nil {
        return err
    }
    err = fmt.Errorf("%s", responseBody.Message)
    return
}

// for qurey meeting
func QueryMeeting(startTime string, endTime string) (responseBody []Meeting, err error) {
    var code int
    code, err = request("GET", "/api/meetings"+"?startTime="+startTime+"&endTime="+endTime, nil, &responseBody)
    if err != nil {
        return err
    }
    // 200
    if code == http.StatusOK {
        return responseBody
    }
    // 401
    if code == http.StatusUnauthorized {
        return fmt.Errorf("Please log in !")
    }
    return
}
