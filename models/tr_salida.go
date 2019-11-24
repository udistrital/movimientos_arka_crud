package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TrSalida struct {
	Salida    *Movimiento
	Elementos []*ElementosMovimiento
}

// AddTransaccionProduccionAcademica Transacción para registrar toda la información de un grupo asociándolo a un catálogo
func AddTransaccionSalida(m *TrSalida) (err error) {
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

	var (
		formatoTipo      *FormatoTipoMovimiento = new(FormatoTipoMovimiento)
		estadoMovimiento *EstadoMovimiento      = new(EstadoMovimiento)
		elementoId       *Movimiento            = new(Movimiento)
		entrada          *Movimiento            = new(Movimiento)
	)

	m.Salida.Activo = true
	formatoTipo.Id = 7
	m.Salida.FormatoTipoMovimientoId = formatoTipo
	estadoMovimiento.Id = 3
	m.Salida.EstadoMovimientoId = estadoMovimiento

	if idSalida, err := o.Insert(m.Salida); err == nil {
		for _, elemento := range m.Elementos {
			elementoId.Id = int(idSalida)
			elemento.MovimientoId = elementoId
			if _, err := o.Insert(elemento); err != nil {
				panic(err.Error())
			}
		}
		entrada.Id = m.Salida.MovimientoPadreId.Id
		logs.Error(entrada)
		if o.Read(entrada) == nil {
			estadoMovimiento.Id = 4
			entrada.EstadoMovimientoId = estadoMovimiento
			if _, err := o.Update(entrada, "EstadoMovimientoId"); err != nil {
				panic(err.Error())
			}
		}
	} else {
		panic(err.Error())
	}

	return
}
