package controllers

type HelpController struct {
	BaseController
}

func (this *HelpController) Index() {

	this.Data["pageTitle"] = "使用帮助"
	this.display()
}

func (this *HelpController) About() {

	this.Data["pageTitle"] = "关于系统"
	this.display()
}
