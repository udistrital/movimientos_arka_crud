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

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de un grupo asociándolo a un catálogo
func AddTransaccionSalida(n *SalidaGeneral) (err error) {
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


	for _, m := range n.Salidas {


		fmt.Println(m.Salida.Detalle)
		if idSalida, err := o.Insert(m.Salida); err == nil {
			fmt.Println(idSalida)
			mov := Movimiento{Id : int(idSalida)}
			for _, elemento := range m.Elementos {
				elemento.MovimientoId = &mov
				fmt.Println("elemento" ,elemento)
				if _, err := o.Insert(elemento); err != nil {
					panic(err.Error())
				}
			}
			fmt.Println("ok2")

			if m.Salida.MovimientoPadreId != nil {
				entrada := Movimiento{Id : m.Salida.MovimientoPadreId.Id}
				if err := o.Read(entrada) ; err == nil {
					entrada.EstadoMovimientoId.Id = 4
					if _, err := o.Update(entrada, "EstadoMovimientoId"); err != nil {
						panic(err.Error())
					}
				}
			}
			
		} else {
			panic(err.Error())
		}

	}

	return
}

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de un grupo asociándolo a un catálogo
func GetTransaccionSalida(id int) (Salida []map[string]interface{}, err error) {
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
					"MovimientoId":      	elemento.MovimientoId,
				})
			}
			Salida = append(Salida, map[string]interface{}{
				"Salida":		v,
				"Elementos":	Elementos,
			})
			return Salida, nil
		}
	} else {
		return nil, err
	}
}

