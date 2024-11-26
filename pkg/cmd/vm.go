package cmd

import (
	"fmt"
	"github.com/analog-substance/carbon/pkg/types"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
	"os"
	"strings"
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

func vmTable(vms []types.VM, privateIPs bool) {

	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
	typeColors := map[string]lipgloss.Color{
		"Running":    lipgloss.Color("#D7FF87"),
		"QEMU":       lipgloss.Color("#FDFF90"),
		"AWS":        lipgloss.Color("#FF875F"),
		"VirtualBox": lipgloss.Color("#FF87D7"),
	}
	dimTypeColors := map[string]lipgloss.Color{
		"Running":    lipgloss.Color("#97AD64"),
		"QEMU":       lipgloss.Color("#FCFF5F"),
		"AWS":        lipgloss.Color("#C77252"),
		"VirtualBox": lipgloss.Color("#C97AB2"),
	}

	headers := []string{"Name", "IP", "State", "Up Time", "Type", "Environment", "Profile", "Provider"}

	CapitalizeHeaders := func(data []string) []string {
		for i := range data {
			data[i] = strings.ToUpper(data[i])
		}
		return data
	}

	data := [][]string{}
	for _, vm := range vms {
		var name string
		if vm.ID() == vm.Name() {
			name = vm.Name()
		} else {
			name = fmt.Sprintf("%s (%s)", vm.Name(), vm.ID())
		}

		ipAddrs := vm.IPAddress()
		if privateIPs {
			ipAddrs = vm.PrivateIPAddress()
		}

		data = append(data, []string{
			name,
			ipAddrs,
			vm.State(),
			vm.UpTime().String(),
			vm.Type(),
			vm.Environment().Name(),
			vm.Profile().Name(),
			vm.Provider().Name(),
		})
	}

	ct := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers(CapitalizeHeaders(headers)...).
		Rows(data...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}

			even := row%2 == 0

			switch col {
			case 2, 7: // Type 1 + 2
				c := typeColors
				if even {
					c = dimTypeColors
				}

				color := c[fmt.Sprint(data[row][col])]
				return baseStyle.Foreground(color)
			}

			if even {
				return baseStyle.Foreground(lipgloss.Color("245"))
			}
			return baseStyle.Foreground(lipgloss.Color("252"))
		})
	fmt.Println(ct)

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
