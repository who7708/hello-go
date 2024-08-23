package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	jarURL         = "http://192.168.100.47:19090/cyy-acs-fe.jar"
	tempFileName   = "downloaded-jar-temp.jar"
	targetFileName = "cyy-acs-fe.jar"
)

func main() {
	// 下载 JAR 包
	err := downloadJar(jarURL)
	if err != nil {
		log.Fatalf("下载 JAR 包失败：%v", err)
	}

	// 停止旧的服务（假设通过进程名查找并杀死）
	stopOldService()

	// 移动临时文件到目标文件名
	err = moveTempToTarget()
	if err != nil {
		log.Fatalf("移动文件失败：%v", err)
	}

	// 启动新的 JAR 服务
	startNewService()
}

func downloadJar(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tempFile, err := os.Create(tempFileName)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	return err
}

func stopOldService() {
	// 尝试多次停止旧服务，确保完全停止
	for i := 0; i < 5; i++ {
		cmd := exec.Command("pkill", "-f", "java -jar "+targetFileName)
		err := cmd.Run()
		if err != nil {
			log.Printf("第 %d 次尝试停止旧服务失败：%v", i+1, err)
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
}

func moveTempToTarget() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	tempPath := filepath.Join(dir, tempFileName)
	targetPath := filepath.Join(dir, targetFileName)

	// 删除旧的目标文件（如果存在）
	if _, err := os.Stat(targetPath); err == nil {
		err = os.Remove(targetPath)
		if err != nil {
			return err
		}
	}

	return os.Rename(tempPath, targetPath)
}

func startNewService() {
	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("获取当前目录失败：%v", err)
		return
	}

	// 查找 JAR 包
	var jarPath string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("读取目录失败：%v", err)
		return
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jar") {
			jarPath = filepath.Join(dir, file.Name())
			break
		}
	}

	if jarPath == "" {
		log.Println("未找到 JAR 包")
		return
	}

	// 启动 JAR 服务
	cmd := exec.Command("java", "-jar", jarPath)
	err = cmd.Start()
	if err != nil {
		log.Printf("启动新服务失败：%v", err)
		return
	}
	log.Println("新服务已启动")
}
