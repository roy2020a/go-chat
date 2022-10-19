package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type configuration struct {
	ServerInfo serverInfo
	RedisInfo  redisInfo
}

type serverInfo struct {
	Host string
}

type redisInfo struct {
	Host        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var Configuration = configuration{}

func Init() {
	//filePath := path.Join(os.Getenv("GOPATH"), "src/go-chat/config/config.json")
	filePath := path.Join(getCurrentAbPath(), "config.json")
	fmt.Println("123:", filePath)
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Printf("Open file error: %v\n", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
