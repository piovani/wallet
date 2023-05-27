package cmd

import (
	"log"

	"github.com/piovani/wallet/infra/config"
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Use:     "wallet",
		Version: "1.0.0",
	}

	cmd.AddCommand(
		// HTTP
		Rest,
		// CRON
		CurrentDollar,
		// TEST
		Test,
	)

	CheckFatal(cmd.Execute())
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitConfig() {
	CheckFatal(config.InitConfig())
}
