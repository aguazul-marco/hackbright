package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/aguazul-marco/hackbright/capstone/cli/client"
	"github.com/spf13/cobra"
)

var DailyFlag bool
var WeeklyFlag bool
var MonthlyFlag bool
var CreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"create", "c"},
	Short:   `add a goal. use the flags to specify daily, weekly or monthly goal.`,
	Long:    `add a goal. use the flags to specify daily, weekly or monthly goal. If no flag is use the default is daily goal. Use the -h for help when needed.`,
	Example: "create [flag] <user_id> run a mile",
	Run: func(cmd *cobra.Command, args []string) {

		switch {
		case DailyFlag:
			goal := "dailygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			arg := args[1:]
			dis := strings.Join(arg, " ")

			client.CreateGoal(id, dis, goal)
		case WeeklyFlag:
			goal := "weeklygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			arg := args[1:]
			dis := strings.Join(arg, " ")

			client.CreateGoal(id, dis, goal)
		case MonthlyFlag:
			goal := "monthlygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			arg := args[1:]
			dis := strings.Join(arg, " ")

			client.CreateGoal(id, dis, goal)
		default:
			goal := "dailygoal"
			client := client.NewClient()
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatalf("most use a numeric value: %v", err)
			}
			arg := args[1:]
			dis := strings.Join(arg, " ")

			client.CreateGoal(id, dis, goal)
		}

	},
}

func init() {
	CreateCmd.Flags().BoolVarP(&DailyFlag, "daily", "d", false, "add a daily goal")
	CreateCmd.Flags().BoolVarP(&WeeklyFlag, "weekly", "w", false, "add a weekly goal")
	CreateCmd.Flags().BoolVarP(&MonthlyFlag, "monthly", "m", false, "add a monthly goal")
}
