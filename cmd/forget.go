package cmd

import (
	"github.com/cupcakearmy/autorestic/internal"
	"github.com/cupcakearmy/autorestic/internal/lock"
	"github.com/spf13/cobra"
)

var forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget and optionally prune snapshots according the specified policies",
	Run: func(cmd *cobra.Command, args []string) {
		err := lock.Lock()
		CheckErr(err)
		defer lock.Unlock()

		CheckErr(internal.CheckConfig())

		selected, err := internal.GetAllOrSelected(cmd, false)
		CheckErr(err)
		prune, _ := cmd.Flags().GetBool("prune")
		dry, _ := cmd.Flags().GetBool("dry-run")
		for _, name := range selected {
			location, _ := internal.GetLocation(name)
			err := location.Forget(prune, dry)
			CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(forgetCmd)
	internal.AddFlagsToCommand(forgetCmd, false)
	forgetCmd.Flags().Bool("prune", false, "Also prune repository")
	forgetCmd.Flags().Bool("dry-run", false, "Do not write changes, show what would be affected")
}
