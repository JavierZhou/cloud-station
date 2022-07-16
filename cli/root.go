package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// ossProvider string
	version     bool
)

var (
	RootCmd = &cobra.Command{
		Use:     "cloud-station-cli",
		Long:    "cloud-station-cli 云中转站CLI",
		Short:   "cloud-station-cli 云中转站CLI",
		Example: "cloud-station-cli cmds",
		RunE: func(cmd *cobra.Command, args []string) error {
			if version {
				fmt.Println("cloud-station v0.0.1")
			}
			return nil
		},
	}
)

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&version, "version", "v", false, "cloud station 版本信息")

}
