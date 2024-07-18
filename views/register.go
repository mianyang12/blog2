package views

import (
	"goBlog/common"
	"goBlog/config"
	"net/http"
)

func (*HTMLApi) Register(w http.ResponseWriter, r *http.Request) {
	register := common.Template.Login

	register.WriteData(w, config.Cfg.Viewer)
}

