package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func f20220307RenameFile() {
	root := `C:\Users\jxdxm\Desktop\新建文件夹`
	var walkFunc = func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if strings.HasSuffix(path, ".livp") {
				log.Println("rename file: ", path)
				//time.Sleep(time.Millisecond * 500)
				return os.Rename(path, path+".rar")
			}
		}
		return err
	}
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("end")
	}
}

func f20220307DeleteRepeatFile() {
	//root := `F:\生活照片`
	root := `E:\生活照片-备份\来自：iPhone`
	var walkFunc = func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if !validPath(path, info.Name()) {
				log.Println("delete file: ", path)
				//time.Sleep(time.Millisecond * 300)
				return os.Remove(path)
			}
		}
		return err
	}
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("end")
	}
}

func validPath(path, name string) bool {
	if !strings.Contains(path, "照片") {
		panic("保护,只处理“照片”文件夹")
	}
	if strings.Contains(path, ".downloading") {
		log.Println("downloading", path)
		return true
	}
	ext := strings.ToLower(filepath.Ext(path))
	if (strings.Contains(path, "(") && strings.Contains(path, ")")) ||
		ext == ".go" || ext == ".md" {
		return false
	}
	var rs = []rune(name)
	for _, r := range rs {
		if r > 255 {
			return false // 有非ascii字符
		}
	}
	return true
}
