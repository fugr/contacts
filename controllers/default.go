package controllers

import (
	"contacts/models"
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	op := this.GetString("op")

	if op == "modify" {
		fmt.Println("modify")
		id, err := this.GetInt("id")
		if err != nil {
			fmt.Println("get wrong id", err)
		}
		contact := models.Get(id)
		this.Data["Id"] = contact.Id
		this.Data["Name"] = contact.Name
		this.Data["Telephone"] = contact.Telephone
		this.Data["Address"] = contact.Address
		this.TplNames = "modify.html"
		return
	}

	this.Data["Contacts"] = models.GetAll()
	this.TplNames = "views.html"
	return
}

func (this *MainController) Post() {
	contacts := models.GetAll()
	contact := models.Contact{}
	var err error
	contact.Id, err = this.GetInt("Id")
	if err != nil {
		fmt.Println("get wrong id", err)
		this.Ctx.Output.Body([]byte("学号错误！"))
		return
	}
	contact.Name = this.GetString("Name")
	contact.Telephone = this.GetString("Telephone")
	contact.Address = this.GetString("Address")

	contacts.Post(&contact)

	this.Data["Contacts"] = models.GetAll()
	this.TplNames = "views.html"
}

func (this *MainController) Download() {
}

func (this *MainController) Delete() {
	id, err := this.GetInt("id")
	if err != nil {
		fmt.Println("get wrong id", err)
	}
	contacts := models.GetAll()
	contacts.Delete(id)
	this.Redirect("/", 302)
	return
}
