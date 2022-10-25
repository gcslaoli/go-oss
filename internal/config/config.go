package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type cConfig struct {
	KeySecrets     g.MapStrStr `json:"key-secrets"`
	AllowAnonymous bool        `json:"allow-anonymous"`
	BaseUrl        string      `json:"base-url"`
}

var (
	Config = cConfig{}
)

func init() {
	ctx := gctx.GetInitCtx()
	Config.KeySecrets = g.Cfg().MustGetWithEnv(ctx, "oss.key-secrets").MapStrStr()
	Config.AllowAnonymous = g.Cfg().MustGetWithEnv(ctx, "oss.allow-anonymous").Bool()
	Config.BaseUrl = g.Cfg().MustGetWithEnv(ctx, "oss.base-url").String()

	// g.Log().Info(ctx, "Config.KeySecrets", Config.KeySecrets)
	// g.Log().Info(ctx, "Config.AllowAnonymous", Config.AllowAnonymous)
	g.Dump(Config)
}
