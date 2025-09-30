package comm

// BizConf 业务配置
var BizConf *BizConfig

type BizConfig struct {
	AllowOrigins []string `mapstructure:"allow_origins"` // 允许的跨域来源
}
