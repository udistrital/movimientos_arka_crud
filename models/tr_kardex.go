package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrKardex struct {
	Kardex    *Movimiento
	Elementos []*ElementosMovimiento
}
type KardexGeneral struct {
	Movimiento []*TrKardex
}

type Apertura struct {
	CantidadMinima     int
	CantidadMaxima     int
	ElementoCatalogoId int
	FechaCreacion      time.Time
	MetodoValoracion   int
	SaldoCantidad      float64
	SaldoValor         float64
	Unidad             float64
	ValorUnitario      float64
	ValorTotal         float64
}

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de un grupo asociándolo a un catálogo
func AddTransaccionKardex(n *KardexGeneral) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()
		fmt.Println("ok")


	for _, m := range n.Movimiento {


		fmt.Println(m.Kardex.Detalle)
		if idSalida, err := o.Insert(m.Kardex); err == nil {
			fmt.Println(idSalida)
			mov := Movimiento{Id : int(idSalida)}
			for _, elemento := range m.Elementos {
				elemento.MovimientoId = &mov
				var elemento__ ElementosMovimiento
				id__ := elemento.Id
				elemento.Id = 0;
				fmt.Println("elemento" ,elemento)

				if _, err := o.Insert(elemento); err == nil {
					if _, err := o.QueryTable(new(ElementosMovimiento)).RelatedSel().Filter("Id",id__).All(&elemento__) ; err == nil {
						fmt.Println(elemento__)
						elemento__.Activo = false
						if _, err := o.Update(&elemento__, "Activo"); err != nil {
							panic(err.Error())
						}
					}

				} else {
					panic(err.Error())
				}
			}
			fmt.Println("ok2")
			
		} else {
			panic(err.Error())
		}

	}

	return
}

func ResponderSolicitud(Solicitud *KardexGeneral) (err error) {

	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()
		fmt.Println("ok")

	if err := AddTransaccionKardex(Solicitud); err == nil {
		id := Solicitud.Movimiento[0].Kardex.MovimientoPadreId
		
		var elemento__ Movimiento
		fmt.Println("asdkjsdfhsdlfkghsldfkjghlsdkf")
		fmt.Println(id)
		fmt.Println(id.EstadoMovimientoId)
		fmt.Println(id)
		if _, err := o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id",id.Id).All(&elemento__) ; err == nil {
			fmt.Println(elemento__)
			
			elemento__.EstadoMovimientoId = id.EstadoMovimientoId
			elemento__.Detalle = id.Detalle
			if _, err := o.Update(&elemento__, "EstadoMovimientoId", "Detalle"); err != nil {
				panic(err.Error())
			}
		}

	}
	return

}

func RechazarSolicitud(id *Movimiento) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()
		fmt.Println("ok")

		var elemento__ Movimiento
	if _, err := o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id",id.Id).All(&elemento__) ; err == nil {
		fmt.Println(elemento__)
		
		elemento__.EstadoMovimientoId = id.EstadoMovimientoId
		elemento__.Detalle = id.Detalle
		if _, err := o.Update(&elemento__, "EstadoMovimientoId", "Detalle"); err != nil {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	return
}

func GetAllAperturas(conSaldo bool) (aperturas []Apertura, err error) {

	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	query :=
		`
	WITH aperturas AS (
		SELECT DISTINCT ON (1)
			em.elemento_catalogo_id,
			m.detalle
		FROM
			movimientos_arka.elementos_movimiento em,
			movimientos_arka.movimiento m,
			movimientos_arka.formato_tipo_movimiento tm
		WHERE
				tm.codigo_abreviacion = 'AP_KDX'
			AND m.formato_tipo_movimiento_id = tm.id
			AND em.movimiento_id = m.id
	), saldo AS (
		SELECT DISTINCT ON (1)
			em.elemento_catalogo_id,
			m.fecha_creacion,
			em.saldo_cantidad,
			em.saldo_valor,
			em.unidad,
			em.valor_total,
			em.valor_unitario,
			aperturas.detalle
		FROM
			movimientos_arka.elementos_movimiento em,
			movimientos_arka.movimiento m,
			movimientos_arka.formato_tipo_movimiento tm,
			aperturas
		WHERE
				tm.codigo_abreviacion IN ('AP_KDX', 'ENT_KDX', 'SAL_KDX')
			AND m.formato_tipo_movimiento_id = tm.id
			AND em.movimiento_id = m.id
			AND aperturas.elemento_catalogo_id = em.elemento_catalogo_id
		ORDER BY 1, 2 DESC
	)

	SELECT *,
		detalle ->> 'Metodo_Valoracion' metodo_valoracion,
		detalle ->> 'Cantidad_Minima' cantidad_minima,
		detalle ->> 'Cantidad_Maxima' cantidad_maxima
	FROM saldo
	`

	if conSaldo {
		query +=
			`
		WHERE saldo_cantidad > 0;
		`
	}

	_, err = o.Raw(query).QueryRows(&aperturas)
	return
}
