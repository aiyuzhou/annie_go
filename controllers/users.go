package controllers

import (
	"time"
  "annie_go/models"
  "github.com/astaxie/beego"
)

type UserController struct {
  BaseController
}

func (c *UserController) Register()  {
	//define a register form
  form := models.RegisterForm{}
	beego.Info(form) //can't get request form data
	//method:ParseForm parse form data(both the raw qurery from the URL and the request body)
	//if parse form error, return err
	if err := c.ParseForm(&form); err != nil {
    beego.Debug("ParseRegisterForm:", err)
    c.Data["json"] = models.NewErrorInfo(ErrInputData)
    c.ServeJSON()
    return
  }
	beego.Info(form) //can get request form data
  beego.Debug("ParseRegisterForm:", &form)

	//verify if the form data is valid, if not return ErrInputData
  if err := c.VerifyForm(&form); err != nil {
    beego.Debug("ValidRegisterForm:", err)
    c.Data["json"] = models.NewErrorInfo(ErrInputData)
    c.ServeJSON()
    return
  }
	// record user's registered date
  regDate := time.Now()
  user, err := models.NewUser(&form,regDate)
  if err != nil {
    beego.Error("NewUser:", err)
    c.Data["json"] = models.NewErrorInfo(ErrSystem)
    c.ServeJSON()
    return
  }
	// user is a set of json data, after encrypt password
  beego.Debug("NewUser:", user)

	// if code = 0, success; if code = -1, databese error, if code = -3, duplicate user
  if code, err := user.Insert(); err != nil {
    beego.Error("InsterUser:", err)
    if code == models.ErrDupRows {
      c.Data["json"] = models.NewErrorInfo(ErrDupUser)
    } else {
      c.Data["json"] = models.NewErrorInfo(ErrDatabase)
    }
    c.ServeJSON()
    return
  }

  go models.IncTotalUserCount(regDate)

  c.Data["json"] = models.NewNormalInfo("Succes")
  c.ServeJSON()
}
