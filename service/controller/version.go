package controller

import (
	"V2RayA/global"
	"V2RayA/model/v2ray"
	"V2RayA/service"
	"V2RayA/tools"
	"github.com/gin-gonic/gin"
)

func GetVersion(ctx *gin.Context) {
	err := v2ray.CheckTProxySupported()
	var transparentProxyValid string
	if err == nil {
		transparentProxyValid = "yes"
	} else {
		transparentProxyValid = err.Error()
	}
	tools.ResponseSuccess(ctx, gin.H{
		"version":          global.Version,
		"dockerMode":       global.ServiceControlMode == global.DockerMode,
		"foundNew":         global.FoundNew,
		"remoteVersion":    global.RemoteVersion,
		"serviceValid":     v2ray.IsV2rayServiceValid(),
		"transparentValid": transparentProxyValid,
	})
}

func GetRemoteGFWListVersion(ctx *gin.Context) {
	c, err := tools.GetHttpClientAutomatically()
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	t, err := service.GetRemoteGFWListUpdateTime(c)
	if err != nil {
		tools.ResponseError(ctx, err)
		return
	}
	tools.ResponseSuccess(ctx, gin.H{"remoteGFWListVersion": t.Format("2006-01-02")})
}
