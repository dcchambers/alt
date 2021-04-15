package cmd

import (
	"encoding/json"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

type Programs struct {
	Programs []Program `json:"programs"`
}

type Program struct {
	Name         string   `json:"name"`
	FullName     string   `json:"fullName"`
	Description  string   `json:"description"`
	URL          string   `json:"url"`
	Alternatives []string `json:"alternatives"`
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alt",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	//Args: cobra.ArbitraryArgs,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		programs := readJson("programs.json")
		//fmt.Println(programs)
		for i := 0; i < len(programs.Programs); i++ {
			if programs.Programs[i].Name == args[0] {
				fmt.Println("Name: " + programs.Programs[i].Name)
				fmt.Println("Full Name: " + programs.Programs[i].FullName)
				fmt.Println("Description: " + programs.Programs[i].Description)
				fmt.Println("URL: " + programs.Programs[i].URL)
				fmt.Printf("Alternatives: ")
				for j := 0; j < len(programs.Programs[i].Alternatives); j++ {
					fmt.Printf(programs.Programs[i].Alternatives[j] + " ")
				}
				fmt.Println()
			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.alt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".alt" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".alt")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func readJson(fileName string) Programs {
	// Open jsonFile
	jsonFile, err := os.Open(fileName)
	// handle error os.Open may return
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//initialize programs array
	var programs Programs

	// unmarshal the byteArray from jsonFile
	json.Unmarshal(byteValue, &programs)
	return programs

	// we iterate through every program within our programs array and
	// print out some attributes
	/*
			for i := 0; i < len(programs.Programs); i++ {
	        fmt.Println("Name: " + programs.Programs[i].Name)
	        fmt.Println("Full Name: " + programs.Programs[i].FullName)
	        fmt.Println("Description: " + programs.Programs[i].Description)
	        fmt.Println("URL: " + programs.Programs[i].URL)
	        for j :=0; j<len(programs.Programs[i].Alternatives); j++ {
	            fmt.Println("Alternatives: " + programs.Programs[i].Alternatives[j])
	        }
	        fmt.Println("---")
	    }
	*/
}
