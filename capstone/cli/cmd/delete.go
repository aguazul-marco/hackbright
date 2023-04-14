package cmd

import (
	"log"
	"strconv"

	"github.com/aguazul-marco/hackbright/capstone/cli/client"
	"github.com/spf13/cobra"
)

const (
	dailyGoal   = "dailygoal"
	weeklyGoal  = "weeklygoal"
	monthlyGoal = "monthlygoal"
)

var GoalFlag bool
var DeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"delete", "d"},
	Short:   "delete account or goals. Use the flags to specify",
	Run: func(cmd *cobra.Command, args []string) {

		if !GoalFlag {
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			client.DeleteData(id, 0, "")
		} else {
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}

			gID, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			switch {
			case DailyFlag:
				client.DeleteData(id, gID, dailyGoal)
			case WeeklyFlag:
				client.DeleteData(id, gID, weeklyGoal)
			case MonthlyFlag:
				client.DeleteData(id, gID, monthlyGoal)
			}
		}

	},
}

func init() {
	DeleteCmd.Flags().BoolVarP(&GoalFlag, "goal", "g", false, "delete a goal")
	DeleteCmd.Flags().BoolVarP(&DailyFlag, "daily", "d", false, "used in conjunction with -goal to specific what type of goal to delete")
	DeleteCmd.Flags().BoolVarP(&WeeklyFlag, "weekly", "w", false, "used in conjunction with -goal to specific what type of goal to delete")
	DeleteCmd.Flags().BoolVarP(&MonthlyFlag, "monthly", "m", false, "used in conjunction with -goal to specific what type of goal to delete")
}
