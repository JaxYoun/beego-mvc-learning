package controllers

func (c *MainController) Get() {
	c.Data["name"] = "yang"
	c.Data["id"] = "75"
	c.TplName = "test0.html"
}

func (c *MainController) Post() {
	c.Data["name"] = "jax"
	c.Data["id"] = "75"
	c.TplName = "test0.html"
}
