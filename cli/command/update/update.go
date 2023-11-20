package update

import (
	"github.com/TheGameProfi/proxmox-api-go/cli"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "With this command you can update existing items within proxmox",
}

func init() {
	cli.RootCmd.AddCommand(updateCmd)
}