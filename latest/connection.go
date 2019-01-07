package latest

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RunE is the entry poing for the application
func RunE(cmd *cobra.Command, args []string) error {
	fmt.Println("runing")
	return nil
}
