package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type IndexReq struct {
	g.Meta         `path:"/" method:"post"`
	OSSAccessKeyId string            `json:"OSSAccessKeyId"`
	Policy         string            `json:"policy"`
	Signature      string            `json:"Signature"`
	Key            string            `json:"key"`
	File           *ghttp.UploadFile `json:"file" type:"file"`
}
type IndexRes struct {
	// 上传成功后的文件访问路径
	Url string `json:"url,omitempty"`
}
