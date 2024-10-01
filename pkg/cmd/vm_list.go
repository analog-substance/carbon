package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
)

// vmList represents the config command
var vmList = &cobra.Command{
	Use:   "list",
	Short: "List VMs",
	Long:  `List VMs`,
	Run: func(cmd *cobra.Command, args []string) {

		vms := getVMsFromArgs(cmd, args)
		if jsonOutput {
			out, err := json.MarshalIndent(vms, "", "  ")
			if err != nil {
				log.Error("error marshalling JSON", err)
			}
			fmt.Println(string(out))
		} else {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Name", "IP", "State", "Up Time", "Type", "Environment", "Profile", "Provider"})

			for _, vm := range vms {
				var name string
				if vm.ID() == vm.Name() {
					name = vm.Name()
				} else {
					name = fmt.Sprintf("%s (%s)", vm.Name(), vm.ID())
				}
				t.AppendRows([]table.Row{
					{
						name,
						vm.IPAddress(),
						vm.State(),
						vm.UpTime(),
						vm.Type(),
						vm.Environment().Name(),
						vm.Profile().Name(),
						vm.Provider().Name(),
					},
				})
			}

			t.Render()

		}
	},
}

func init() {
	vmCmd.AddCommand(vmList)
}
