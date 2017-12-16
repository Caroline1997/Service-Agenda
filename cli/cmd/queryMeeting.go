package cmd

import (
    "Service-Agenda/cli/service"
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
        service.Query_meeting(startTime, endTime)
    },
}

func init() {
    RootCmd.AddCommand(queryMeetingCmd)
    queryMeetingCmd.Flags().StringP("StartTime", "s", "", "startTime of the meeting")
    queryMeetingCmd.Flags().StringP("EndTime", "e", "", "endTime of the meeting")
}
