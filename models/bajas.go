package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrRevisionBaja struct {
	Bajas         []int
	Aprobacion    bool
	Observaciones string
}

type FormatoBaja struct {
	Consecutivo    string
	Elementos      []int
	FechaRevisionA string
	FechaRevisionC string
	Funcionario    int
	Revisor        int
}

// PostRevisionComite hace la actualización de los movimientos de acuerdo a la revisión
func PostRevisionComite(n *TrRevisionBaja) (ids []int, err error) {
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

		if !n.Aprobacion {
			v.Observacion += n.Observaciones
		} else {
			var detalle FormatoBaja
			if err := json.Unmarshal([]byte(v.Detalle), &detalle); err != nil {
				panic(err)
			}
			detalle.FechaRevisionC = time.Now().UTC().String()
			if detalle_, err := json.Marshal(detalle); err != nil {
				panic(err)
			} else {
				v.Detalle = string(detalle_[:])
			}
		}

		if _, err = o.Update(&v); err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}
