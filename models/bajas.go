package models

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrRevisionBaja struct {
	Aprobacion     bool
	Bajas          []int
	DependenciaId  int
	FechaRevisionC string
	RazonRechazo   string
	Resolucion     string
}

type FormatoBaja struct {
	Consecutivo    string
	ConsecutivoId  int
	DependenciaId  int
	Elementos      []int
	FechaRevisionA string
	FechaRevisionC string
	Funcionario    int
	Revisor        int
	RazonRechazo   string
	Resolucion     string
}

// PostRevisionComite hace la actualización de los movimientos de acuerdo a la revisión
func PostRevisionComite(n *TrRevisionBaja) (ids []int, err error) {
	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		logs.Error(err)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			logs.Error(r)
		}
		o.Commit()
	}()

	var estado EstadoMovimiento
	var estadoString string
	if n.Aprobacion {
		estadoString = "Baja Aprobada"
	} else {
		estadoString = "Baja Rechazada"
	}

	if _, err = o.QueryTable(new(EstadoMovimiento)).RelatedSel().Filter("Nombre", estadoString).All(&estado); err != nil {
		panic(err.Error())
	}

	for _, id := range n.Bajas {
		v := Movimiento{Id: id}

		if err = o.Read(&v); err != nil {
			panic(err)
		}

		v.EstadoMovimientoId = &estado
		var detalle FormatoBaja
		if err := json.Unmarshal([]byte(v.Detalle), &detalle); err != nil {
			panic(err)
		}

		if !n.Aprobacion {
			detalle.RazonRechazo += n.RazonRechazo
		} else {
			detalle.FechaRevisionC = n.FechaRevisionC
			detalle.Resolucion = n.Resolucion
			detalle.DependenciaId = n.DependenciaId

			for _, el := range detalle.Elementos {
				novedad := NovedadElemento{
					MovimientoId:         &Movimiento{Id: id},
					ElementoMovimientoId: &ElementosMovimiento{Id: el},
					Activo:               true,
				}

				if _, err = o.Insert(&novedad); err != nil {
					panic(err.Error())
				}

			}

		}

		if detalle_, err := json.Marshal(detalle); err != nil {
			panic(err)
		} else {
			v.Detalle = string(detalle_[:])
		}

		if _, err = o.Update(&v); err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}
