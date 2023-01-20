package models

import (
	"fmt"

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

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de un grupo asociándolo a un catálogo
func GetTransaccionKardex(id int) (Salida map[string]interface{}, err error) {
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

	
	var elementos []ElementosMovimiento
	var Elementos []map[string]interface{}

	v := &Movimiento{Id: id}
	if err = o.Read(v); err == nil {

		if _, err := o.QueryTable(new(ElementosMovimiento)).RelatedSel().Filter("Activo",true).Filter("ElementoCatalogoId__Id",id).All(&elementos); err != nil{
			panic(err.Error())
		} else {

			for _, elemento := range elementos {
				Elementos = append(Elementos, map[string]interface{}{
					"Id":					elemento.Id,                
					"ElementoActaId":    	elemento.ElementoActaId,
					"Unidad":            	elemento.Unidad,
					"ValorUnitario":     	elemento.ValorUnitario,
					"ValorTotal":        	elemento.ValorTotal,
					"SaldoCantidad":     	elemento.SaldoCantidad,
					"SaldoValor":        	elemento.SaldoValor,
					"Activo":            	elemento.Activo,
					"FechaCreacion":     	elemento.FechaCreacion,
					"FechaModificacion": 	elemento.FechaModificacion,
					// "MovimientoId":      	elemento.MovimientoId,
				})
			}
			Salida = map[string]interface{}{
				"Salida": v,
				"Elementos": Elementos,
			}
			
			return Salida, nil
		}
	} else {
		return nil, err
	}
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

func GetAllAperturas(conSaldo bool) (ids []ElementosMovimiento, err error) {
	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	var ids_ []int
	query :=
		`
	WITH aperturas AS (
		SELECT DISTINCT ON (1)
			em.elemento_catalogo_id
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
			em.fecha_creacion,
			em.*
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

	SELECT id FROM saldo
	`

	if conSaldo {
		query +=
			`
		WHERE saldo_cantidad > 0;
		`
	}

	if _, err = o.Raw(query).QueryRows(&ids_); err != nil {
		return
	}

	_, err = o.QueryTable(new(ElementosMovimiento)).RelatedSel().Filter("Id__in", ids_).All(&ids)
	return
}
