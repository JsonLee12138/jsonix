package test

import (
	"json-server/test/controller"
	"json-server/test/service"
	"go.uber.org/dig"
)

// RegistertestModule 在 DI 容器中注册 test 模块
func TestModuleSetup(container *dig.Container) {
	container.Provide(service.NewtestService)
	container.Provide(controller.NewtestController)
}
