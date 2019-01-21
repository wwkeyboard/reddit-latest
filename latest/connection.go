package latest

import (
	"fmt"

	"github.com/jzelinskie/geddit"

	"github.com/spf13/cobra"
)

// RunE is the entry poing for the application
func RunE(cmd *cobra.Command, args []string) error {
	session := geddit.NewSession("AaronsLonelyPostFinder")

	subOpts := geddit.ListingOptions{
		Limit: 10,
	}

	topics, err := session.SubredditSubmissions("photoclass2019", geddit.NewSubmissions, subOpts)
	if err != nil {
		return err
	}

	for _, t := range topics {
		fmt.Printf("%+v\n", t.Title)
	}

	return nil
}
