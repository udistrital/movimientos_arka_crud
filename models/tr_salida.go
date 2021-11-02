package models

import (
	"fmt"

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

		if _, err := o.QueryTable(new(ElementosMovimiento)).RelatedSel().Filter("Activo",true).Filter("MovimientoId__Id",id).All(&elementos); err != nil{
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

