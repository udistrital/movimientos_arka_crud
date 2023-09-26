package models

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrSalida struct {
	Salida    *Movimiento
	Elementos []*ElementosMovimiento
}
type SalidaGeneral struct {
	Salidas []TrSalida
}

// AddTransaccionSalida Transacción para registrar todas las salidas asociadas a una entrada
func AddTransaccionSalida(n *SalidaGeneral) (err error) {

	o := orm.NewOrm()

	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	err = o.Begin()
	if err != nil {
		return
	}

	var estado EstadoMovimiento
	err = o.QueryTable(new(EstadoMovimiento)).RelatedSel().Filter("Nombre", "Entrada Con Salida").One(&estado)
	if err != nil {
		panic(err)
	}

	n.Salidas[0].Salida.MovimientoPadreId.EstadoMovimientoId = &estado
	if _, err = o.Update(n.Salidas[0].Salida.MovimientoPadreId, "EstadoMovimientoId"); err != nil {
		panic(err)
	}

	for _, m := range n.Salidas {
		idSalida, err := o.Insert(m.Salida)
		if err != nil {
			panic(err)
		}

		mov := Movimiento{Id: int(idSalida)}
		for _, elemento := range m.Elementos {
			elemento.MovimientoId = &mov
			_, err = o.Insert(elemento)
			if err != nil {
				panic(err)
			}
		}
	}

	return
}

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de un grupo asociándolo a un catálogo
func GetTransaccionSalida(id int) (salida map[string]interface{}, err error) {
	o := orm.NewOrm()

	var elementos []interface{}
	var movimiento Movimiento
	salida = map[string]interface{}{}

	_, err = o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", id).All(&movimiento)
	if err != nil || movimiento.Id == 0 {
		return
	}

	query := map[string]string{"MovimientoId__Id": strconv.Itoa(id)}
	fields := []string{"Id", "ElementoActaId", "Unidad", "ValorUnitario", "ValorTotal", "SaldoCantidad", "SaldoValor", "VidaUtil", "ValorResidual"}
	elementos, err = GetAllElementosMovimiento(query, fields, nil, nil, 0, -1)
	if err != nil {
		return
	}

	salida = map[string]interface{}{
		"Salida":    movimiento,
		"Elementos": elementos,
	}

	return
}

// PutTransaccionSalida Transacción para registrar todas las salidas asociadas a una entrada
func PutTransaccionSalida(n *SalidaGeneral) (err error) {

	o := orm.NewOrm()

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	err = o.Begin()
	if err != nil {
		return
	}

	for _, m := range n.Salidas {
		// Se actualiza la salida con el id y consecutivo original
		if m.Salida.Id > 0 {
			_, err = o.Update(m.Salida)
			if err != nil {
				panic(err)
			}

			for _, elemento := range m.Elementos {
				_, err := o.Update(elemento, "VidaUtil", "ValorResidual")
				if err != nil {
					panic(err)
				}
			}
		} else {
			// Las demás salidas se insertan como un movimiento adicional y este Id se asigna a los elementos
			idSalida, err := o.Insert(m.Salida)
			if err != nil {
				panic(err)
			}

			mov := Movimiento{Id: int(idSalida)}
			for _, elemento := range m.Elementos {
				elemento.MovimientoId = &mov
				_, err = o.Update(elemento, "MovimientoId", "VidaUtil", "ValorResidual")
				if err != nil {
					panic(err)
				}
			}
		}
	}

	return
}
