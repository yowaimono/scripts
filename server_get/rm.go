package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/urfave/cli/v2"
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
	app := &cli.App{
		Name:  "app",
		Usage: "A simple tool to upload and download files using scp",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./cnf.toml",
				Usage:   "Path to the configuration file",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "Download a file from a server",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return fmt.Errorf("Usage: %s get <server> <remote_file> [local_file]", c.App.Name)
					}

					serverName := c.Args().Get(0)
					remoteFile := c.Args().Get(1)
					localFile := filepath.Base(remoteFile)
					if c.NArg() > 2 {
						localFile = c.Args().Get(2)
					}

					config, err := loadConfig(c.String("config"))
					if err != nil {
						return fmt.Errorf("Failed to load config: %v", err)
					}

					server, ok := config.Servers[serverName]
					if !ok {
						return fmt.Errorf("Server %s not found in config", serverName)
					}

					return downloadFile(server, localFile, remoteFile)
				},
			},
			{
				Name:  "up",
				Usage: "Upload a file to a server",
				Action: func(c *cli.Context) error {
					if c.NArg() < 2 {
						return fmt.Errorf("Usage: %s up <server> <local_file> [remote_file]", c.App.Name)
					}

					serverName := c.Args().Get(0)
					localFile := c.Args().Get(1)
					remoteFile := filepath.Base(localFile)
					if c.NArg() > 2 {
						remoteFile = c.Args().Get(2)
					}

					config, err := loadConfig(c.String("config"))
					if err != nil {
						return fmt.Errorf("Failed to load config: %v", err)
					}

					server, ok := config.Servers[serverName]
					if !ok {
						return fmt.Errorf("Server %s not found in config", serverName)
					}

					return uploadFile(server, localFile, remoteFile)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
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
	// 确保远程文件路径不包含 Windows 路径分隔符
	remoteFile = strings.ReplaceAll(remoteFile, "\\", "/")

	// 确保本地文件路径在 Windows 系统上正确处理
	localFile = filepath.FromSlash(localFile)

	content := fmt.Sprintf("%s@%s:%s", server.Name, server.Addr, remoteFile)

	fmt.Println(content)
	cmd := exec.Command("C:\\Program Files\\Git\\usr\\bin\\scp.exe", content, localFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func uploadFile(server Server, localFile, remoteFile string) error {
	// 确保远程文件路径不包含 Windows 路径分隔符
	remoteFile = strings.ReplaceAll(remoteFile, "\\", "/")

	// 确保本地文件路径在 Windows 系统上正确处理
	localFile = filepath.FromSlash(localFile)

	cmd := exec.Command("C:\\Program Files\\Git\\usr\\bin\\scp.exe", localFile, fmt.Sprintf("%s@%s:%s", server.Name, server.Addr, remoteFile))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}