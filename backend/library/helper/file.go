package helper

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/genv"
)

// LoadConfig 加载config目录的配置文件
func LoadConfig(filename, index string) interface{} {

	filenameall := path.Base(filename)
	filesuffix := path.Ext(filenameall)
	basename := strings.TrimSuffix(filenameall, filesuffix)

	env := genv.Get("env", "pro")
	cfg := g.Cfg(basename)
	cfgDir, err := cfg.GetFilePath()
	if err != nil {
		cfgDir = "./config"
	}
	// 根据env参数判断配置文件是否存在
	cfg_path := filepath.Dir(cfgDir) + "/" + env + "/" + filename

	// 根据环境变量读取配置文件
	_, err = os.Stat(cfg_path)
	if err != nil && os.IsNotExist(err) {
		cfg.SetFileName(filename)
	} else {
		cfg.SetFileName(env + "/" + filename)
	}
	if index != "" {
		menus := cfg.Get(index)
		return menus
	}

	arr := cfg.Array()
	if len(arr) > 0 {
		return arr[0]
	}
	return []string{}
}

// 获取文件内容的类型os.open()后调用
func GetFileContentType(out *os.File) (string, error) {
	// 只需前512 个字节即可
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	t := http.DetectContentType(buffer)
	return t, nil
}

// 递归获取目录树
func ReadDirTree(dirPath string, isDir bool) ([]map[string]interface{}, error) {
	treeData := []map[string]interface{}{}
	flist, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return treeData, err
	}
	for _, f := range flist {
		if f.IsDir() {
			tmp := map[string]interface{}{
				"label":   f.Name(),
				"value":   dirPath + "/" + f.Name(),
				"size":    f.Size(),
				"modtime": f.ModTime(),
				"isDir":   f.IsDir(),
			}
			child, _ := ReadDirTree(dirPath+"/"+f.Name(), isDir)
			if len(child) > 0 {
				tmp["children"] = child
			}
			treeData = append(treeData, tmp)
		} else if !isDir {
			treeData = append(treeData, map[string]interface{}{
				"label":   f.Name(),
				"value":   dirPath + "/" + f.Name(),
				"size":    f.Size(),
				"modtime": f.ModTime(),
				"isDir":   f.IsDir(),
			})
		}
	}
	return treeData, nil
}

// 获取目录下所有文件
func ReadDirFiles(dirPath string) ([]map[string]interface{}, error) {
	treeData := []map[string]interface{}{}
	flist, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return treeData, err
	}
	for _, f := range flist {
		if !f.IsDir() {
			treeData = append(treeData, map[string]interface{}{
				"label":   f.Name(),
				"value":   dirPath + "/" + f.Name(),
				"size":    f.Size(),
				"modtime": f.ModTime(),
				"isDir":   f.IsDir(),
			})
		}
	}
	return treeData, nil
}

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		// return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { // if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

// Copy 复制文件
func Copy(src, dst string) (int64, error) {

	// dst 要过虑非正常路径，如&之类的字符
	filterStr := []string{"..", "&", ":", ";", "|", "$", "%", "?", "\r", "\n", "`", ","}
	for _, str := range filterStr {
		if strings.Contains(dst, str) {
			return 0, errors.New("目标路径存在特殊字符: " + str)
		}
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// IsDir 判断文件夹是否存在
func IsDir(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !f.IsDir() {
		return errors.New("不是文件夹")
	}
	return nil
}

// IsFile 判断文件是否存在
func IsFile(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return errors.New("不是文件类型")
	}
	return nil
}

// FileExist 判断文件是否存在
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

// Mkdir 创建目录
func Mkdir(path string, perm os.FileMode) error {
	if path == "" {
		return errors.New("path empty")
	}
	// recordID 要过虑非正常路径，如&之类的字符
	filterStr := []string{"..", "&", ":", ";", "|", "$", "%", "?", "\r", "\n", "`", ",", " "}
	for _, str := range filterStr {
		if strings.Contains(path, str) {
			return errors.New("目标路径存在特殊字符: " + str)
		}
	}
	err := os.MkdirAll(path, perm)
	return err
}
