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
	Command() *cobra.Command
	run(cmd *cobra.Command, args []string)
}

func AddCommand(study Study) {
	rootCmd.AddCommand(study.Command())
}

func AddCommands(studies []Study) {
	for _, study := range studies {
		AddCommand(study)
	}
}

func Execute() error {
	rootCmd.PersistentFlags().String("bs", "localhost:9092", "Bootstrap servers comma separated")

	return rootCmd.Execute()
}
