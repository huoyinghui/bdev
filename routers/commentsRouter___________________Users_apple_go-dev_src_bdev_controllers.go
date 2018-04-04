package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["bdev/controllers:MessageController"] = append(beego.GlobalControllerRouter["bdev/controllers:MessageController"],
		beego.ControllerComments{
			Method: "CreateMessage",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bdev/controllers:MessageController"] = append(beego.GlobalControllerRouter["bdev/controllers:MessageController"],
		beego.ControllerComments{
			Method: "DeleteMessage",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bdev/controllers:TokenController"] = append(beego.GlobalControllerRouter["bdev/controllers:TokenController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["bdev/controllers:ValidController"] = append(beego.GlobalControllerRouter["bdev/controllers:ValidController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
