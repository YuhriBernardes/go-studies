package sample

import "github.com/spf13/cobra"

var (
	samples = []Sample{
		&CountDown{},
	}
)

type Sample interface {
	Init()
	run(*cobra.Command, []string)
	Command() *cobra.Command
}

func LoadSamples(rootCmd *cobra.Command) {

	for _, sample := range samples {
		sample.Init()
		rootCmd.AddCommand(sample.Command())
	}

}
