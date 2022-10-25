package service

import (
	"crypto/hmac"
	"crypto/sha1"

	"github.com/hjmcloud/go-oss/internal/config"

	v1 "github.com/hjmcloud/go-oss/api/v1"

	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sIndex struct {
}

// Upload 兼容OSS上传
func (s *sIndex) Upload(ctx g.Ctx, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	var (
		accessKeySecret string
		policy          = req.Policy
		basePath        = gfile.Join(gfile.SelfDir(), "upload")
	)
	if secret, ok := config.Config.KeySecrets[req.OSSAccessKeyId]; !ok {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "AccessKeyId不存在")
	} else {
		accessKeySecret = secret
	}
	// 解析policy
	policyBytes, err := gbase64.DecodeString(policy)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "Policy解析失败")
	}
	policyMap := gconv.MapStrStr(gjson.New(policyBytes))
	g.Dump(policyMap)
	// 检验policy是否过期
	expiration := policyMap["expiration"]
	expirationTime, err := gtime.StrToTime(expiration)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "Policy过期时间解析失败")
	}
	if expirationTime.Before(gtime.Now()) {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "Policy已过期")
	}
	// hmac-sha1 校验
	h := hmac.New(sha1.New, []byte(accessKeySecret))
	h.Write([]byte(policy))
	signature := gbase64.EncodeToString(h.Sum(nil))
	if signature != req.Signature {
		return nil, gerror.New("签名错误")
	}

	filePath := gfile.Join(basePath, req.Key)
	// 获取文件绝对路径
	fileAbs := gfile.Abs(filePath)
	fileDir := gfile.Dir(fileAbs)
	// fileAbs 应当以 basePath 开头，防止越权访问
	if gstr.Pos(fileAbs, basePath) != 0 {
		return nil, gerror.New("文件路径错误")
	}
	if gfile.Exists(filePath) {
		return nil, gerror.New("文件已存在")
	}
	// 保存文件
	req.File.Filename = gfile.Basename(filePath)
	_, err = req.File.Save(fileDir)
	res = &v1.IndexRes{
		Url: config.Config.BaseUrl + "/" + req.Key,
	}
	return
}

// UploadAnonymous 匿名上传
func (s *sIndex) UploadAnonymous(ctx g.Ctx, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	// 获取当前日期
	date := gtime.Now().Format("Ymd")
	basePath := gfile.Join(gfile.SelfDir(), "upload", date)
	// 保存文件
	filename, err := req.File.Save(basePath, true)
	res = &v1.IndexRes{
		Url: config.Config.BaseUrl + "/" + date + "/" + filename,
	}
	return
}

var (
	Index = sIndex{}
)
