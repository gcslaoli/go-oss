package controller

import (
	"github.com/hjmcloud/go-oss/internal/config"
	"github.com/hjmcloud/go-oss/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/hjmcloud/go-oss/api/v1"
)

var (
	Index = cIndex{}
)

type cIndex struct {
}

// Index is the default controller method for the controller.
func (c *cIndex) Index(ctx g.Ctx, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	// g.Dump(req)
	// policy, err := gbase64.DecodeToString(req.Policy)
	// g.Dump(policy)
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "文件为空")
	}
	// 如果允许匿名上传且没有传入AccessKeyId，则调用匿名上传
	if config.Config.AllowAnonymous && req.OSSAccessKeyId == "" {
		return service.Index.UploadAnonymous(ctx, req)
	}
	// 检验参数是否为空
	if req.OSSAccessKeyId == "" {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "AccessKeyId为空")
	}
	if req.Signature == "" {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "Signature为空")
	}
	if req.Policy == "" {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "Policy为空")
	}
	if req.Key == "" {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "Key为空")
	}
	return service.Index.Upload(ctx, req)
}
