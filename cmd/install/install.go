// Package install /*
package install

import (
	"github.com/williamotieno/butler/services"

	"github.com/spf13/cobra"
)

// ButlerInstallCmd represents the install command
var ButlerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Subcommand to install tools",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

butler install redis`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			switch arg {
			case "docker":
				services.InstallDocker()
			case "test":
				services.InstallTest()
			}
		}
	},
}

func init() {
	//rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
