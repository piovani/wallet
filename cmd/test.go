package cmd

import (
	"fmt"
	"time"

	"github.com/piovani/wallet/app/shared"
	"github.com/spf13/cobra"
)

var (
	Test = &cobra.Command{
		Use:     "test",
		Short:   "Start listen http type rest",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig()

			checkDayIsHoliday := shared.NewCheckDayIsHoliday()
			res := checkDayIsHoliday.Execute(time.Now())

			fmt.Println(res)
		},
	}
)
