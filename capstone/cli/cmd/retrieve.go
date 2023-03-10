package cmd

import (
	"log"
	"strconv"

	"github.com/aguazul-marco/hackbright/capstone/cli/client"
	"github.com/spf13/cobra"
)

var CompleteFlag bool
var IncompleteFlag bool
var RetrieveCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"get", "g"},
	Short: `get a list of all goals. Use the flags to specific the type of goal, as well as incomplete and complete goals. 
No flag will retreive all goals. Use the -h for help when needed.`,
	Args:    cobra.ExactArgs(1),
	Example: "get [flag] [flag] <user_id>",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case IncompleteFlag:
			switch {
			case DailyFlag:
				flag := "uncompleted"
				goal := "dailygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			case WeeklyFlag:
				flag := "uncompleted"
				goal := "weeklygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			case MonthlyFlag:
				flag := "uncompleted"
				goal := "monthlygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			}
		case CompleteFlag:
			switch {
			case DailyFlag:
				flag := "completed"
				goal := "dailygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			case WeeklyFlag:
				flag := "completed"
				goal := "weeklygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			case MonthlyFlag:
				flag := "completed"
				goal := "monthlygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			}
		default:
			switch {
			case DailyFlag:
				flag := ""
				goal := "dailygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			case WeeklyFlag:
				flag := ""
				goal := "weeklygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			case MonthlyFlag:
				flag := ""
				goal := "monthlygoals"
				client := client.NewClient()
				id, err := strconv.ParseInt(args[0], 10, 64)
				if err != nil {
					log.Fatalf("most use a numeric value: %v", err)
				}

				client.GetGoals(id, goal, flag)
			}
		}

	},
}

func init() {
	RetrieveCmd.Flags().BoolVarP(&DailyFlag, "daily", "d", false, "add a daily goal")
	RetrieveCmd.Flags().BoolVarP(&WeeklyFlag, "weekly", "w", false, "add a weekly goal")
	RetrieveCmd.Flags().BoolVarP(&MonthlyFlag, "monthly", "m", false, "add a monthly goal")
	RetrieveCmd.Flags().BoolVarP(&CompleteFlag, "complete", "c", false, "get all completed goals")
	RetrieveCmd.Flags().BoolVarP(&IncompleteFlag, "incomplete", "i", false, "get all incomplete goals")
}
