package cmd

import (
	"log"

	"github.com/ccsunnyfd/practice/word/internal/word"
	"github.com/spf13/cobra"
)

//
const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
)

var str string
var mode int8

// wordCmdCmd represents the wordCmd command
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  `支持多种单词格式转换`,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpperCase(str)
		case ModeLower:
			content = word.ToLowerCase(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		default:
			log.Fatalf("help word")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init() {
	rootCmd.AddCommand(wordCmd)
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
