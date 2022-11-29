package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type cConfig struct {
	KeySecrets     g.MapStrStr `json:"keySecrets"`
	AllowAnonymous bool        `json:"allowAnonymous"`
	BaseUrl        string      `json:"base-url"`
}

var (
	Config = cConfig{}
)

func init() {
	ctx := gctx.GetInitCtx()
	Config.KeySecrets = g.Cfg().MustGetWithEnv(ctx, "oss.keySecrets").MapStrStr()
	Config.AllowAnonymous = g.Cfg().MustGetWithEnv(ctx, "oss.allowAnonymous").Bool()
	Config.BaseUrl = g.Cfg().MustGetWithEnv(ctx, "oss.baseUrl").String()

	// g.Log().Info(ctx, "Config.KeySecrets", Config.KeySecrets)
	// g.Log().Info(ctx, "Config.AllowAnonymous", Config.AllowAnonymous)
	g.Dump(Config)
}
