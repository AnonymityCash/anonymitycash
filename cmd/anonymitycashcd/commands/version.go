package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/anonymitycash/anonymitycash/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
