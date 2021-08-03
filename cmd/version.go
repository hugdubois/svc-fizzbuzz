package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command.
var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Return service version",
		Long: `Use this command

Example :
  $ svc-fizzbuzz version

`,
		Run: version,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// version is the core function of this command.
func version(cmd *cobra.Command, args []string) {
	fmt.Printf("%s version %s - %s", svc.Name, svc.Version, svcName)
}
