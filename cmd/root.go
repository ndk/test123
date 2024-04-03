package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ndk/test123/internal"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use: "app",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	rootCmd.AddCommand(&cobra.Command{
		Use:   "run",
		Short: "Run application",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Run(cmd.Context())
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
