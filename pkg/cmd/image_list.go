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
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// imageListCmd represents the image command
var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "list images",
	Long:  `list images`,
	Run: func(cmd *cobra.Command, args []string) {
		listBuilds, _ := cmd.Flags().GetBool("list-builds")
		if listBuilds {
			imagesBuilds, err := carbonObj.GetImageBuilds()
			if err != nil {
				log.Fatal(err)
			}
			for _, imageBuild := range imagesBuilds {
				fmt.Println(imageBuild)
			}
		} else {
			fmt.Println("Currently only listing image build is possible. to be implemented")
		}

	},
}

func init() {
	imageCmd.AddCommand(imageListCmd)
	imageListCmd.Flags().BoolP("builds", "b", false, "List build configs")
}
