package cmd

import (
	"log"
	"strconv"

	"github.com/aguazul-marco/hackbright/capstone/cli/client"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"update", "u"},
	Short: `update goal status, "true" if it's complete or "false" if it's incomplete. Default is false. 
If no flag is use the default is daily goal. Use the -h for help when needed.`,
	Example: "update [flag] <user_id> <goal_id> true",
	Run: func(cmd *cobra.Command, args []string) {

		switch {
		case DailyFlag:
			goal := "dailygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			goalID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			status, _ := strconv.ParseBool(args[2])

			client.UpdateGoals(id, goalID, status, goal)
		case WeeklyFlag:
			goal := "weeklygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			goalID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			status, _ := strconv.ParseBool(args[2])

			client.UpdateGoals(id, goalID, status, goal)
		case MonthlyFlag:
			goal := "monthlygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			goalID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			status, _ := strconv.ParseBool(args[2])

			client.UpdateGoals(id, goalID, status, goal)
		default:
			goal := "dailygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			goalID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			status, _ := strconv.ParseBool(args[2])

			client.UpdateGoals(id, goalID, status, goal)
		}

	},
}

func init() {
	UpdateCmd.Flags().BoolVarP(&DailyFlag, "daily", "d", false, "update daily goal")
	UpdateCmd.Flags().BoolVarP(&WeeklyFlag, "weekly", "w", false, "update weekly goal")
	UpdateCmd.Flags().BoolVarP(&MonthlyFlag, "monthly", "m", false, "update monthly goal")
}
