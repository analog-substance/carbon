/*
Copyright © 2024 defektive

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// imageBootstrapCmd represents the image command
var imageBootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "create image build configs",
	Long:  `create image build configs`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		osDir, _ := cmd.Flags().GetString("template")
		serviceProvider, _ := cmd.Flags().GetString("service")
		err := carbonObj.CreateImageBuild(name, osDir, serviceProvider)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	imageCmd.AddCommand(imageBootstrapCmd)
	imageBootstrapCmd.Flags().StringP("name", "n", "", "Name of image build")
	imageBootstrapCmd.Flags().StringP("template", "t", "ubuntu-24.04", "Template to use")
	imageBootstrapCmd.Flags().StringP("service", "s", "", "Service provider (aws, virtualbox, qemu, multipass)")
}
