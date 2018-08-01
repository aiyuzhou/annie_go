package controllers

type ErrorController struct {
  BaseController
}

func (c *ErrorController) RetError(e *ControllerError) {
  c.Data["json"] = e
  c.ServeJSON()
}

func (c *ErrorController) Error404() {
  c.RetError(err404)
}
