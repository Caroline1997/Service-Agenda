package cmd

import (
    "fmt"
    "github.com/Caroline1997/Service-Agenda/cli/service"
    "github.com/spf13/cobra"
    "log"
    "os"
)

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
    Use:   "queryMeeting -s [StartTime] -e [EndTime]",
    Short: "To query the meeting during [StartTime] and [EndTime]",
    Long: `usage of using this command is to query the meeting during [StartTime] and [EndTime]`,
    Run: func(cmd *cobra.Command, args []string) {
        startTime, _ := cmd.Flags().GetString("StartTime")
        endTime, _ := cmd.Flags().GetString("EndTime")
        meetings, err := service.Query_meeting(startTime, endTime)
        if err != nil {
            fmt.Println(err)
        }
        if len(meetings) == 0 {
            fmt.Println("There is no meeting!")
            return
        }
        for temp, m := range meetings {
				    fmt.Println("Title: %s", m.Title)
				    fmt.Println("Sponsor: %s", m.Sponsor)
            fmt.Println("Participators:")
            for _, participator := range m.Participators {
                fmt.Println("%s", participator)
            }
            fmt.Println("StartTime: %s", m.StartTime)
				    fmt.Println("EndTime: %s", m.EndTime)
        }
    },
}

func init() {
    RootCmd.AddCommand(queryMeetingCmd)
    queryMeetingCmd.Flags().StringP("StartTime", "s", "", "startTime of the meeting")
    queryMeetingCmd.Flags().StringP("EndTime", "e", "", "endTime of the meeting")
}
