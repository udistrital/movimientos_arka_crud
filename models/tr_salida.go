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
	Salidas []*TrSalida
}

// AddTransaccionSalida Transacción para registrar todas las salidas asociadas a una entrada
func AddTransaccionSalida(n *SalidaGeneral) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()
	if err != nil {
		logs.Error(err)
		return
	}

	var estado EstadoMovimiento
	if _, err = o.QueryTable(new(EstadoMovimiento)).RelatedSel().Filter("Nombre", "Entrada Con Salida").All(&estado); err != nil {
		panic(err.Error())
	}
	n.Salidas[0].Salida.MovimientoPadreId.EstadoMovimientoId = &estado
	if _, err = o.Update(n.Salidas[0].Salida.MovimientoPadreId, "EstadoMovimientoId"); err != nil {
		panic(err.Error())
	}

	for _, m := range n.Salidas {
		if idSalida, err := o.Insert(m.Salida); err == nil {
			mov := Movimiento{Id: int(idSalida)}
			for _, elemento := range m.Elementos {
				elemento.MovimientoId = &mov
				if _, err := o.Insert(elemento); err != nil {
					panic(err.Error())
				}
			}
		} else {
			panic(err.Error())
		}
	}

	return
}

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de un grupo asociándolo a un catálogo
func GetTransaccionSalida(id int) (Salida map[string]interface{}, err error) {
	o := orm.NewOrm()
	err = o.Begin()

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	if err != nil {
		logs.Error(err)
		return
	}

	var elementos []interface{}
	var movimiento Movimiento

	if _, err := o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", id).All(&movimiento); err != nil {
		panic(err.Error())
	}

	query := map[string]string{"MovimientoId__Id": strconv.Itoa(id)}
	fields := []string{"Id", "ElementoActaId", "Unidad", "ValorUnitario", "ValorTotal", "SaldoCantidad", "SaldoValor", "VidaUtil", "ValorResidual"}
	if elementos, err = GetAllElementosMovimiento(query, fields, nil, nil, 0, -1); err != nil {
		panic(err.Error())
	}

	Salida = map[string]interface{}{
		"Salida":    movimiento,
		"Elementos": elementos,
	}

	return Salida, nil
}

// PutTransaccionSalida Transacción para registrar todas las salidas asociadas a una entrada
func PutTransaccionSalida(n *SalidaGeneral) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		} else {
			o.Commit()
		}
	}()

	if err != nil {
		logs.Error(err)
		return
	}

	for _, m := range n.Salidas {
		// Se actualiza la salida con el id y consecutivo original
		if m.Salida.Id > 0 {
			if _, err = o.Update(m.Salida); err != nil {
				panic(err.Error())
			}
			for _, elemento := range m.Elementos {
				if _, err := o.Update(elemento, "VidaUtil", "ValorResidual"); err != nil {
					panic(err.Error())
				}
			}
		} else {
			// Las demás salidas se insertan como un movimiento adicional y este Id se asigna a los elementos
			if idSalida, err := o.Insert(m.Salida); err == nil {
				mov := Movimiento{Id: int(idSalida)}
				for _, elemento := range m.Elementos {
					elemento.MovimientoId = &mov
					if _, err := o.Update(elemento, "MovimientoId", "VidaUtil", "ValorResidual"); err != nil {
						panic(err.Error())
					}
				}
			} else {
				panic(err.Error())
			}
		}
	}

	return
}
