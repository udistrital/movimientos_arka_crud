// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"

	"github.com/udistrital/movimientos_arka_crud/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/formato_tipo_movimiento",
			beego.NSInclude(
				&controllers.FormatoTipoMovimientoController{},
			),
		),

		beego.NSNamespace("/movimiento",
			beego.NSInclude(
				&controllers.MovimientoController{},
			),
		),

		beego.NSNamespace("/estado_movimiento",
			beego.NSInclude(
				&controllers.EstadoMovimientoController{},
			),
		),

		beego.NSNamespace("/soporte_movimiento",
			beego.NSInclude(
				&controllers.SoporteMovimientoController{},
			),
		),

		beego.NSNamespace("/elementos_movimiento",
			beego.NSInclude(
				&controllers.ElementosMovimientoController{},
			),
		),

		beego.NSNamespace("/tr_salida",
			beego.NSInclude(
				&controllers.TrSalidaController{},
			),
		),
		beego.NSNamespace("/tr_kardex",
			beego.NSInclude(
				&controllers.TrkardexController{},
			),
		),
		beego.NSNamespace("/bajas",
			beego.NSInclude(
				&controllers.BajasController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
