package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Servers map[string]Server `toml:"servers"`
}

type Server struct {
	Name string `toml:"name"`
	Addr string `toml:"addr"`
	Pass string `toml:"pass"`
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <get|up> <server> [local_file] [remote_file]", os.Args[0])
	}

	action := os.Args[1]
	serverName := os.Args[2]
	file := ""
	path := "."

	if len(os.Args) > 3 {
		file = os.Args[3]
	}

	if len(os.Args) > 4 {
		path = os.Args[4]
	}

	// 将 :: 替换为斜杠
	file = strings.ReplaceAll(file, ":", "/")
	path = strings.ReplaceAll(path, ":", "/")

	fmt.Println("remotefile: ", file)
	fmt.Println("dir is: ", path)

	// 读取配置文件
	config, err := loadConfig("./cnf.toml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server, ok := config.Servers[serverName]
	if !ok {
		log.Fatalf("Server %s not found in config", serverName)
	}

	switch action {
	case "get":
		err = downloadFile(server, path, file)
	case "up":
		err = uploadFile(server, file, path)
	default:
		log.Fatalf("Unknown action: %s", action)
	}

	if err != nil {
		log.Fatalf("Failed to %s file: %v", action, err)
	}
}

func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func downloadFile(server Server, localFile, remoteFile string) error {
	if remoteFile == "" {
		fmt.Println("error remote file is null")
		return fmt.Errorf("error remote file is null")
	}

	// 确保本地文件路径在 Windows 系统上正确处理
	localFile = filepath.FromSlash(localFile)

	content := fmt.Sprintf("%s@%s:%s", server.Name, server.Addr, remoteFile)

	fmt.Println(content)
	cmd := exec.Command("scp", content, localFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func uploadFile(server Server, localFile, remoteFile string) error {
	if remoteFile == "" {
		remoteFile = filepath.Base(localFile)
	}

	// 确保本地文件路径在 Windows 系统上正确处理
	localFile = filepath.FromSlash(localFile)

	cmd := exec.Command("scp", localFile, fmt.Sprintf("%s@%s:%s", server.Name, server.Addr, remoteFile))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
