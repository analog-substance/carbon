package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// vmVNCCmd represents the config command
var vmVNCCmd = &cobra.Command{
	Use:   "vnc",
	Short: "vnc to a vm",
	Long:  `vnc to a vm`,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		vms := getVMsFromArgs(cmd, args)
		if len(vms) > 1 {
			fmt.Println("Too many vms specified.")
		} else if len(vms) == 1 {
			err := vms[0].StartVNC(user)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Println("VM not found")
		}
	},
}

func init() {
	vmCmd.AddCommand(vmVNCCmd)
}
