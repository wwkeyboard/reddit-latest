package latest

import (
	"fmt"

	"github.com/jzelinskie/geddit"

	"github.com/spf13/cobra"
)

// RunE is the entry poing for the application
func RunE(cmd *cobra.Command, args []string) error {
	fmt.Println("runing")

	session := geddit.NewSession("AaronsLonelyPostFinder")

	subOpts := geddit.ListingOptions{
		Limit: 10,
	}

	subs, err := session.SubredditSubmissions("photoclass2019", geddit.TopSubmissions, subOpts)
	if err != nil {
		return err
	}

	for _, s := range subs {
		fmt.Printf("-- %v, \n", s.Selftext)
	}

	return nil
}
