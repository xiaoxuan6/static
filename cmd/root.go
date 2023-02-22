package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "care-screenshot",
	Short: "一款命令行辅助截图工具",
}

func Execute()  {
	cobra.CheckErr(rootCmd.Execute())
}
