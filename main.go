package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringP("dir", "d", "", "指定文件服务器的根目录")
	rootCmd.Flags().StringP("host", "", "0.0.0.0", "指定服务器的IP")
	rootCmd.Flags().StringP("port", "p", "18020", "指定服务器的端口")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var verbose = `文件服务器已启动:
  监听IP: %s
  本机IP: %s
  端口: %s
  目录: %s

使用：
  查看文件列表:
  curl "http://%s:%s/"

  下载文件: curl -o 文件名 主机:端口/文件名 
  curl -o %s %s:%s/%s %s 
			
`
var rootCmd = &cobra.Command{
	Use:   "fileserver",
	Short: "文件服务器",
	Long:  `文件服务器`,
	Run: func(cmd *cobra.Command, args []string) {
		file := "text.txt"
		if len(args) == 1 {
			file = args[0]

		}

		host := FlagsGetString(cmd, "host")
		ip := GetIp()

		port := FlagsGetString(cmd, "port")
		dir := FlagsGetString(cmd, "dir")

		if dir == "" {
			current, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				fmt.Println("获取不到当前执行路径！Error:", err)
				cmd.Help()
				return
			}
			dir = current
		}

		// 输出提示
		otherCmd := ""
		if !strings.Contains(file, ".") || strings.HasSuffix(file, ".sh") {
			otherCmd = "&& chmod +x " + file
		}

		fmt.Printf(verbose, host, ip, port, dir, ip, port, file, ip, port, file, otherCmd)

		http.Handle("/", http.FileServer(http.Dir(dir)))
		err := http.ListenAndServe(host+":"+port, nil)
		if err != nil {
			fmt.Println("服务器启动失败！ Error: ", err.Error())
			os.Exit(1)
		}

	},
}

func FlagsGetString(cmd *cobra.Command, flag string) string {
	flagVal, err := cmd.Flags().GetString(flag)
	if err != nil {
		fmt.Printf("指定%s参数时出错 Error: %s\n", flag, err.Error())
		os.Exit(1)
	}

	return flagVal
}

func GetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}
