package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "这是一个示例应用程序",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ResourceName)
	},
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加资源",
	Run: func(cmd *cobra.Command, args []string) {
		// 当 add 命令被执行时，打印出 ResourceName
		fmt.Println("Adding resource:", ResourceName)
	},
}

// ResourceName 用于存储资源名称
var ResourceName string

// init 函数用于初始化命令的标志
func init() {
	// 在根命令中定义 -n 标志
	rootCmd.PersistentFlags().StringVarP(&ResourceName, "name", "a", "test", "指定资源名称")
	// 将 add 命令添加到根命令
	rootCmd.AddCommand(addCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
