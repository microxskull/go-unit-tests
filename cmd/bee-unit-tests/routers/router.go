// Package routers definition.
// @APIVersion 1.0.0.
// @Title beego Test API.
// @Description beego has a very cool tools to autogenerate documents for your API.
// @Contact astaxie@gmail.com.
// @TermsOfServiceUrl http://beego.me/.
// @License Apache 2.0.
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html.
package routers

import (
	"github.com/microxskull/go-unit-tests/cmd/bee-unit-tests/controllers"

	"github.com/astaxie/beego"
)

func init() {
	objectCtrl := &controllers.ObjectController{}
	{
		beego.Router("/v1/object", objectCtrl, "post:Post")
		beego.Router("/v1/object/:objectId", objectCtrl, "get:Get")
		beego.Router("/v1/object", objectCtrl, "get:GetAll")
		beego.Router("/v1/object/:objectId", objectCtrl, "put:Put")
		beego.Router("/v1/object/:objectId", objectCtrl, "delete:Delete")
	}

	userCtrl := &controllers.UserController{}
	{
		beego.Router("/v1/user", userCtrl, "post:Post")
		beego.Router("/v1/user/:uid", userCtrl, "get:Get")
		beego.Router("/v1/user", userCtrl, "get:GetAll")
		beego.Router("/v1/user/:uid", userCtrl, "put:Put")
		beego.Router("/v1/user/:uid", userCtrl, "delete:Delete")
		beego.Router("/v1/user/login", userCtrl, "get:Login")
		beego.Router("/v1/user/logout", userCtrl, "get:Logout")
	}
}
