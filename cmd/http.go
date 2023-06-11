package cmd

import (
	"github.com/piovani/wallet/ui/rest"
	"github.com/spf13/cobra"
)

var (
	Rest = &cobra.Command{
		Use:     "rest",
		Short:   "Start listen http type rest",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			CheckFatal(rest.NewRest().Start())
		},
	}
)
