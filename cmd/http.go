package cmd

import (
	"fmt"

	"github.com/piovani/wallet/infra/config"
	"github.com/spf13/cobra"
)

var (
	Rest = &cobra.Command{
		Use:     "rest",
		Short:   "Start listen http type rest",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig()
			fmt.Println(config.Env.ApiRestPort)
		},
	}
)
