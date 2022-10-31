package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:BajasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:BajasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:DepreciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:DepreciacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:DepreciacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:DepreciacionController"],
        beego.ControllerComments{
            Method: "GetCorte",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "GetByFuncionario",
            Router: "/funcionario/:funcionarioId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:ElementosMovimientoController"],
        beego.ControllerComments{
            Method: "GetHistorial",
            Router: "/historial/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:EstadoMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:FormatoTipoMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetAllBajasByTerceroId",
            Router: "/baja/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetMovimientoByActa",
            Router: "/entrada/:acta_recibido_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "GetAllTrasladoByTerceroId",
            Router: "/traslado/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:NovedadElementoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:SoporteMovimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrSalidaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"],
        beego.ControllerComments{
            Method: "PostRechazarSolicitud",
            Router: "/rechazar_solicitud/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"] = append(beego.GlobalControllerRouter["github.com/udistrital/movimientos_arka_crud/controllers:TrkardexController"],
        beego.ControllerComments{
            Method: "PostRespuestaSolicitud",
            Router: "/responder_solicitud/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
