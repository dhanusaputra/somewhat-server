package env

import "github.com/dhanusaputra/somewhat-server/util/envutil"

type authMiddlewareConf struct {
	AuthEnable              bool
	AuthMethodBlacklist     map[string]bool
	AuthRequestURIBlacklist map[string]bool
}

const (
	defaultAuthEnable = true
)

var (
	// AuthMiddlewareConf ...
	AuthMiddlewareConf         = &authMiddlewareConf{}
	defaultAuthMethodBlacklist = map[string]bool{
		"GET": true,
	}
	defaultAuthRequestURIBlacklist = map[string]bool{
		"/v1/login": true,
	}
)

// Init ...
func Init() {
	AuthMiddlewareConf = &authMiddlewareConf{
		AuthEnable:              envutil.GetEnvAsBool("AUTH_ENABLE", defaultAuthEnable),
		AuthMethodBlacklist:     envutil.GetEnvAsMapBool("AUTH_METHOD_BLACKLIST", defaultAuthMethodBlacklist, ","),
		AuthRequestURIBlacklist: envutil.GetEnvAsMapBool("AUTH_REQUESTURI_BLACKLIST", defaultAuthRequestURIBlacklist, ","),
	}
}
