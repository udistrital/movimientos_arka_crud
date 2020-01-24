package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
