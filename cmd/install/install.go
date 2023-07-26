// Package install /*
package install

import (
	"github.com/williamotieno/butler/services"

	"github.com/spf13/cobra"
)

// ButlerInstallCmd represents the installation command
var ButlerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Subcommand to install tools",
	Long: `Specifically install a tool or a variety of tools. For example:

butler install redis
butler install docker nvm redis
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			switch arg {
			case "test":
				services.InstallTest()
			case "docker":
				services.InstallDocker()
			case "redis":
				services.InstallRedis()
			case "nvm":
				services.InstallNVM()
			case "homebrew":
				services.InstallHomebrew()
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
