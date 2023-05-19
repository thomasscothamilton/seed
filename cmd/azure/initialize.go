package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	AzureInitializeCmd = &cobra.Command{
		Use: "initialize --billing-account-name <billing-account-name> --billing-profile-name <billing-profile-name> --invoice-section-name <invoice-section-name>",
		Run: func(cmd *cobra.Command, args []string) {
			// create a new azure seed configuration

			if _, err := os.Stat("seed.yaml"); err == nil {
				fmt.Printf("file seed.yaml already exists")
			} else if errors.Is(err, os.ErrNotExist) {
				yamlData, err := yaml.Marshal(&azureSeed)
				if err != nil {
					panic(err)
				}
				err = ioutil.WriteFile("seed.yaml", yamlData, 0)
				if err != nil {
					panic("Unable to write data into the file")
				}

			} else {
				// do nothing
			}

			// create a new cdktf configuration
			if _, err := os.Stat("cdktf.json"); err == nil {
				// do nothing
			} else if errors.Is(err, os.ErrNotExist) {
				yamlData, err := yaml.Marshal(&azureSeed)
				if err != nil {
					panic(err)
				}
				err = ioutil.WriteFile("cdktf.json", yamlData, 0)
				if err != nil {
					panic("Unable to write data into the file")
				}
			} else {
				// do nothing
			}
		},
	}
)

func init() {
	AzureCmd.AddCommand(AzureInitializeCmd)
}
