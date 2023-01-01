package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Tag = "develop"
var Version = "develop"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s - Manga volumes downloading tool\n", color.YellowString("Manga Downloader"))
		fmt.Printf("All Rights Reserved © 2023 %s\n", color.BlackString("Òscar Casajuana Alonso"))
		fmt.Printf("Version: %s\n", color.MagentaString("%s (%s)", Version, Tag))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
