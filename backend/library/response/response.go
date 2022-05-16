package response

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"nft_object/library/auth"
	"nft_object/library/helper"
	"nft_object/statusCode"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gogf/gf/crypto/gmd5"

	"github.com/gogf/gf/text/gstr"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// Json 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:F
// code:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func Json(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	list := []interface{}(nil)

	if len(data) > 1 {
		list = data[1:]
	}

	json := g.Map{
		"code": code,
		"msg":  msg,
		"data": responseData,
		"list": list,
	}

	// 防止xss攻击（让浏览器解析 javascript 代码，而不会是 html 输出）
	r.Response.Header().Set("Content-Type", "text/javascript")

	r.Response.WriteJson(json)
	r.Exit()
}

// WriteHtmlFile 输出html文件
func WriteHtmlFile(fileName string) {
	//1.判断静态文件是否存在
	if helper.FileExist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println("移除文件失败")
		}
	}
	//2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()
	//template.Execute(file, &product)
}

// DownloadFile 下载文件
func DownloadFile(r *ghttp.Request, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		Json(r, statusCode.ERROR, "9.138.64.22")
	}
	defer file.Close()
	fileName := path.Base(filePath)

	// 防止中文乱码
	fileName = url.PathEscape(fileName)
	r.Response.Header().Add("Content-Type", "application/octet-stream")
	r.Response.Header().Add("content-disposition", "attachment; filename=\""+fileName+"\"")
	r.Response.Header().Add("FileName", fileName)
	r.Response.Header().Add("Access-Control-Expose-Headers", "FileName")
	_, err = io.Copy(r.Response.Writer, file)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, err.Error())
	}
}

// OutputFile 输出文件流(音频、视频、图片、pdf)，传递路径读取内容输出
func OutputFile(r *ghttp.Request, filePath string) {

	f, err := os.Open(filePath)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, "Error: 文件不存在")
	}
	defer f.Close()

	fileType, err := helper.GetFileContentType(f)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, err.Error())
	}

	r.Response.ClearBuffer()

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, err.Error())
	}
	fileInfo, _ := f.Stat()
	length := fileInfo.Size()

	// Range参数处理，主要兼容safari浏览器识别音频文件流问题
	rangeStr := r.Header.Get("Range")
	r.Response.Header().Set("Content-Length", gconv.String(length))
	start := "0"
	end := gconv.String(length - 1)
	if rangeStr != "" {
		rangeStr = strings.ReplaceAll(rangeStr, "bytes=", "")
		rangeArr := gstr.Split(rangeStr, "-")
		if rangeArr[0] != "" {
			start = rangeArr[0]
		}

		if len(rangeArr) == 2 && rangeArr[1] != "" {
			end = rangeArr[1]
		}
		if end == "1" {
			// safari浏览器传1时情况处理
			r.Response.Header().Set("Content-Length", "2")
		}
	}
	rangeStr = start + "-" + end + "/" + gconv.String(length)
	r.Response.Header().Set("Content-Range", "bytes "+rangeStr)
	etag, _ := gmd5.EncryptString(filePath)
	r.Response.Header().Set("ETag", etag)
	r.Response.Header().Set("Content-Type", fileType)
	r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Pragma", "private")
	r.Response.Header().Set("Content-Transfer-Encoding", "binary")
	r.Response.Header().Set("Cache-Control", "must-revalidate")
	r.Response.Write(content)
}

// OutputFileByByte 输出文件流(音频、视频、图片、pdf)，传递内容输出
func OutputFileByByte(r *ghttp.Request, data []byte) {

	fileType := http.DetectContentType(data)

	content := string(data)
	r.Response.ClearBuffer()

	r.Response.Header().Add("Content-Type", fileType)
	r.Response.Header().Add("Content-Length", strconv.Itoa(len(content)))
	r.Response.Header().Add("Pragma", "private")
	r.Response.Header().Add("Content-Transfer-Encoding", "binary")
	r.Response.Header().Add("Cache-Control", "must-revalidate")
	r.Response.Write(content)
}

// CheckAuthSession check系统登录(tof接口)
func CheckAuthSession(r *ghttp.Request) {
	// 本地测试环境
	v, ok := r.Header["Auth-Sign"]
	if !ok || len(v) == 0 {
		Json(r, statusCode.ERROR_NO_LOGIN, "Not logged in")
		r.Exit()
	}

	admin_info, errorCode, err := auth.NewAuth().CheckSign(v[0])
	if err != nil {
		Json(r, errorCode, err.Error())
		r.Exit()
	}

	if g.IsEmpty(admin_info) {
		Json(r, statusCode.ERROR_NO_LOGIN, "Not logged in")
		r.Exit()
	}
	// 超级管理员账户设置
	superAdmin := gconv.SliceStr(g.Config().GetArray("auth.superAdmin", []string{"fourteen"}))
	login_name := admin_info.LoginName
	if gstr.InArray(superAdmin, gconv.String(login_name)) {
		admin_info.IsAdmin = 1
	}

	r.SetCtxVar(statusCode.SESSION_ADMIN_INFO, admin_info)
	// 设置名称快捷获取
	r.SetCtxVar(statusCode.SESSION_CACHE_ADMIN_NAME, login_name)
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}
