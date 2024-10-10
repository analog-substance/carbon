package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/charmbracelet/huh"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
)

// vmCmd represents the config command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Manage and interact with VMs.",
	Long:  `Manage and interact with VMs.`,
}

func init() {
	CarbonCmd.AddCommand(vmCmd)

	vmCmd.PersistentFlags().StringP("name", "n", "", "Name of the VM.")
	vmCmd.PersistentFlags().StringP("id", "i", "", "ID of machine to start.")
	vmCmd.PersistentFlags().StringP("user", "u", "ubuntu", "SSH Username.")
	vmCmd.PersistentFlags().StringSlice("host", []string{}, "Hostname or IP Address.")

	err := vmCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getVMNames()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}

	err = vmCmd.RegisterFlagCompletionFunc("id", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		names := getVMIDs()
		return names, cobra.ShellCompDirectiveDefault
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getVMsFromArgs(cmd *cobra.Command, args []string) []types.VM {
	hosts, _ := cmd.Flags().GetStringSlice("host")
	id, _ := cmd.Flags().GetString("id")
	name, _ := cmd.Flags().GetString("name")

	if len(hosts) > 0 {
		return carbonObj.VMsFromHosts(hosts)
	}

	if name != "" {
		return carbonObj.FindVMByName(name)
	} else if id != "" {
		return carbonObj.FindVMByID(id)
	}
	return carbonObj.GetVMs()
}

func AskIfSure(msg string) bool {
	var confirm bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(msg).
				//Description(getLatestFile(outputDir)).
				Affirmative("Yes!").
				Negative("No.").
				Value(&confirm),
		),
	)

	err := form.Run()
	if err != nil {
		return false
	}

	return confirm

}

func vmTable(vms []types.VM) {
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

func getVMNames() (vmNames []string) {
	vms := carbonObj.GetVMs()
	for _, vm := range vms {
		vmNames = append(vmNames, vm.Name())
	}
	return vmNames
}

func getVMIDs() (vmIDs []string) {
	vms := carbonObj.GetVMs()
	for _, vm := range vms {
		vmIDs = append(vmIDs, vm.ID())
	}
	return vmIDs
}
