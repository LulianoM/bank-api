package controllers

type ControllersContainer struct{}

func NewControllerContainer() *ControllersContainer {
	controllersContainer := &ControllersContainer{}
	return controllersContainer
}
