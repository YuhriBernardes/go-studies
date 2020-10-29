package studies

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "kfk",
		Short: "Kafka studies runner",
	}
)

type Study interface {
	Name() string
	Register()
	run(cmd *cobra.Command, args []string)
}

func Register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func Execute() error {
	rootCmd.PersistentFlags().String("bs", "localhost:9092", "Bootstrap servers comma separated")

	return rootCmd.Execute()
}
