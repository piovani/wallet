package cmd

import (
	"fmt"

	"github.com/piovani/wallet/infra/notification"
	"github.com/spf13/cobra"
)

var (
	Test = &cobra.Command{
		Use:     "test",
		Short:   "Start listen http type rest",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			InitConfig()

			not := notification.NewNotification()
			msg := not.GetMsgEmail()

			if err := not.SendEmail(msg); err != nil {
				fmt.Println("DEU RUIM", err)
			}
		},
	}
)
