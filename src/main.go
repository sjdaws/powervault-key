package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

func main() {
	var array Array
	var featureEnableIdentifier string
	var output string

	var command = &cobra.Command{
		Long: `powervault-key is a proof of concept for reverse engineers Dell PowerVault premium feature keys
		  
	   Based on the work of https://github.com/ReverseAllTheThings/PowervaultKeygen`,
		Run: func(cmd *cobra.Command, args []string) {
			if output == "<feature-enable-identifier>.key" {
				output = fmt.Sprintf("%s.key", featureEnableIdentifier)
			}

			err := generate(cmd.Flags().Lookup("array").Value.String(), featureEnableIdentifier, output)
			if err != nil {
				log.Fatalf("Unable to create licence file: %v", err)
			}

			fmt.Printf("Licence file successfully generated and saved to %s\n", output)
		},
		Short: "powervault-key - a reverse engineering tool for Dell PowerVault premium feature keys",
		Use:   "powervault-key",
	}

	// Get list of supported arrays
	var supportedArrays []string
	for _, value := range arrayIds {
		supportedArrays = append(supportedArrays, value...)
	}

	command.Flags().VarP(enumflag.NewWithoutDefault(&array, "array", arrayIds, enumflag.EnumCaseInsensitive), "array", "a", fmt.Sprintf("supported arrays: %s", strings.Join(supportedArrays, ", ")))
	command.MarkFlagRequired("array")
	command.Flags().StringVarP(&featureEnableIdentifier, "feature-enable-identifier", "f", "", "the feature enable identifier from your MD array")
	command.MarkFlagRequired("feature-enable-identifier")
	command.Flags().StringVarP(&output, "output", "o", "<feature-enable-identifier>.key", "the full path and filename to save the licence to")

	_ = command.Execute()
}
