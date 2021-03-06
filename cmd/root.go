package cmd

import (
        "fmt"
        kitlog "github.com/go-kit/kit/log"
        "github.com/spf13/cobra"
        "github.com/spf13/viper"
        "log"
        "os"
)

func init() {Init()}

var (
        defaultConfig = viper.New()
        kitLog = kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
        Use:   "grpclab",
        Short: "100% of code generated by https://github.com/gofunct/cookiecutter-grpcgo",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func GetConfig() *viper.Viper {
        return defaultConfig
}

func GetKitLogger() kitlog.Logger {
        return kitLog
}

func Init() {
        {
                log.SetOutput(kitlog.NewStdlibAdapter(kitLog))
                log.Println("new json logger registered")
        }
        {
                defaultConfig = viper.New()
                defaultConfig.AutomaticEnv()
                defaultConfig.AddConfigPath(os.Getenv("$HOME")) // name of config file (without extension)
                defaultConfig.AddConfigPath(".")
                defaultConfig.SetEnvPrefix("grpclab")
                defaultConfig.SetDefault("viper_config_name", "config")
                defaultConfig.SetDefault("full_name", "Coleman Word")
                defaultConfig.SetDefault("github_username", "gofunct")
                defaultConfig.SetDefault("app_name", "grpclab")
                defaultConfig.SetDefault("project_short_description", "100% of code generated by https://github.com/gofunct/cookiecutter-grpcgo")
                defaultConfig.SetDefault("docker_hub_username", "colemanword")
                defaultConfig.SetDefault("docker_image", "alpine")
                defaultConfig.SetDefault("docker_build_image_version", "1.11")
                defaultConfig.SetDefault("json_logs", "y")
                defaultConfig.SetDefault("log_level", "debug")
                defaultConfig.SetDefault("grpc_port", ":9090")
                defaultConfig.SetDefault("http_port", ":8080")
        }

        // If a config file is found, read it in.
        if err := defaultConfig.ReadInConfig(); err != nil {
               log.Println("failed to read config file, writing defaults...")
                if err := defaultConfig.WriteConfigAs("config"+".yaml"); err != nil {
                        log.Fatal("failed to write config")
                        os.Exit(1)
                }

        } else {
                log.Print("Using config file-->", defaultConfig.ConfigFileUsed())
                if err := defaultConfig.WriteConfig(); err != nil {
                        log.Fatal("failed to write config file")
                        os.Exit(1)
                }
        }
}