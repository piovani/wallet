package cmd

import (
	"fmt"

	"github.com/piovani/wallet/app/usecase"
	"github.com/spf13/cobra"
)

var (
	CurrentDollar = &cobra.Command{
		Use:     "current-dollar",
		Short:   "collects the current value of the dollar in reais",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig()

			currentDollar := usecase.NewCurrentDollar()
			fmt.Println(currentDollar.Execute())
		},
	}
)
