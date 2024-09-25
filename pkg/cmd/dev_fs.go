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
	"github.com/analog-substance/carbon/deployments"
	"github.com/spf13/cobra"
	"log"
	"path"
)

// fsCmd represents the new command
var fsCmd = &cobra.Command{
	Use:   "fs",
	Short: "List fs contents",
	Long:  `for testing things`,
	Run: func(cmd *cobra.Command, args []string) {
		ListingDir(".")
	},
}

func init() {
	devCmd.AddCommand(fsCmd)
}

func ListingDir(dir string) {
	listing, err := deployments.Files.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range listing {
		fmt.Println(path.Join(dir, file.Name()))
		if file.IsDir() {
			ListingDir(path.Join(dir, file.Name()))
		}
	}
}