package gclplugin

import "github.com/golangci/plugin-module-register/register"

func init() {
	register.Plugin("logmsg", New)
}
