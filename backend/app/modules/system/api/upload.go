package api

import (
	"nft_object/library/response"
	"nft_object/library/upload"
	"nft_object/statusCode"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	Upload = uploadApi{}
)

type uploadApi struct {
}

// 单图片上传
func (c *uploadApi) UpImg(r *ghttp.Request) {

	files := r.GetUploadFile("upload-file")
	// *ghttp.UploadFile
	names, err := upload.UpImg(files)
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	saveUrl := g.Config().GetString("upload.saveUrl", "")
	// 小程序访问地址
	names.SaveUrl = saveUrl + names.FileUrl
	// 返回的地址  应该是 hostorigin := r.GetHeader("Origin")
	Url := g.Config().GetString("upload.url", "sytest.woa.com")
	names.FileUrl = Url + names.FileUrl
	response.Json(r, statusCode.SUCCESS, "success", names)
}

// 多图片上传
func (c *uploadApi) UpImgs(r *ghttp.Request) {

	file := r.GetUploadFile("upload-file")
	// *ghttp.UploadFile
	file.Filename = "text.xlxs"
	names, err := file.Save("/tmp/")
	if err != nil {
		response.Json(r, statusCode.ERROR, err.Error())
	}
	response.Json(r, statusCode.SUCCESS, "success", names)
}
