package config

const (
	CookieAge      = 60 * 60 * 24 * 30 // cookie存活时长
	CookiePath     = "/"               // cookie路径
	CookieDomain   = "127.0.0.1"       // cookie作用域
	CookieSecure   = false             // 为true时表示只能通过https访问
	CookieHttpOnly = true              // 为true时不能被js获取
)
