package cmd

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"strconv"
	"time"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "根据url截取图片并保存",
	Run: func(cmd *cobra.Command, args []string) {
		basePath := Identify()

		urlStr := cmd.Flags().Lookup("url").Value.String()
		filename := cmd.Flags().Lookup("filename").Value.String()
		element := cmd.Flags().Lookup("element").Value.String()
		width, _ := strconv.Atoi(cmd.Flags().Lookup("width").Value.String())
		height, _ := strconv.Atoi(cmd.Flags().Lookup("height").Value.String())

		if isValidUrl(urlStr) == false {
			println("url 不合法")
			os.Exit(0)
		}

		filename = fmt.Sprintf("%s/%s", basePath, filename)
		if len(element) < 1 {
			page := rod.New().MustConnect().MustPage(urlStr).MustWaitLoad()
			page.MustScreenshot(filename)
		} else {
			page := rod.New().MustConnect().MustPage(urlStr).MustWaitLoad().MustSetViewport(width, height, 0 ,false).MustElement(element)
			page.MustScreenshot(filename)
		}
	},
}

// Identify 判断文件夹是否存在
func Identify() string {
	basePath, _ := os.Getwd()
	date := time.Now().Format("2006/01/02")
	path := fmt.Sprintf("%s/%s", basePath, date)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(date, os.ModePerm)
	}

	return path
}

// 验证url是否合法
func isValidUrl(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}

	u, err := url.Parse(urlStr)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func init() {
	f := execCmd.Flags()
	f.StringP("url", "u", "https://baidu.com", "想要截图的URL")
	f.StringP("filename", "f", "", "保存图片的名称")
	f.StringP("element", "e", "", "页面元素 (default '#s_lg_img')")
	f.StringP("width", "w", "1200", "页面宽度")
	f.StringP("height", "g", "800", "页面高度")
	_ = execCmd.MarkFlagRequired("url")
	_ = execCmd.MarkFlagRequired("filename")
	rootCmd.AddCommand(execCmd)
}
