package xlog

type (
	ZapCFG struct {
		Development       bool `json:"development"` // 是否开发环境
		Debug             bool `json:"debug"`       // 是否 debug
		Sample            bool `json:"sample"`      // 是否采样，默认zap是采样的，在生产环境设置为false，关闭采样
		CallerSkip        int  `json:"caller_skip"` // callerSkip 打印文件和行号
		DisableStackTrace bool `json:"disable_stack_trace"`

		Fields     *Fields     `json:"fields"`     // 携带一些自定义信息
		Lumberjack *Lumberjack `json:"lumberjack"` // 日志分割 options
	}
	Fields struct {
		App         string                 `json:"app"`
		ExtraFields map[string]interface{} `json:"extra_fields"`
		// add more info
	}
	Lumberjack struct {
		LogPath    string `json:"log_path"`    // 日志文件路径，默认 os.TempDir()
		MaxSize    int    `json:"max_size"`    // 日志保存大小，默认 100 MB
		MaxBackups int    `json:"max_backups"` // 日志备份数
		MaxAge     int    `json:"max_age"`     // 最长保存天数
		Compress   bool   `json:"compress"`    // 是否压缩，默认不压缩
	}
)
