package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

type accionPersistenciaElementoSalida int

const (
	insertarElementoSalida accionPersistenciaElementoSalida = iota
	reutilizarElementoSalida
)

// AddTransaccionSalida Transacción para registrar todas las salidas asociadas a una entrada
func AddTransaccionSalida(n *SalidaGeneral) (err error) {

	o := orm.NewOrm()

	defer func() {
		r := recover()
		if r != nil {
			err = r.(error)
			o.Rollback()
			logs.Error(r)
		} else if err != nil {
			o.Rollback()
		} else {
			o.Commit()
		}
	}()

	err = o.Begin()
	if err != nil {
		return
	}

	if err = validarTransaccionSalida(n); err != nil {
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
		now := time.Now()
		m.Salida.FechaCreacion = now
		m.Salida.FechaModificacion = now
		idSalida, err := o.Insert(m.Salida)
		if err != nil {
			panic(err)
		}

		mov := Movimiento{Id: int(idSalida)}
		for _, elemento := range m.Elementos {
			if err = persistirElementoSalida(o, &mov, elemento, now); err != nil {
				panic(err)
			}
		}
	}

	return
}

func validarTransaccionSalida(n *SalidaGeneral) error {
	if n == nil || len(n.Salidas) == 0 {
		return fmt.Errorf("no se recibieron salidas para registrar")
	}

	if n.Salidas[0].Salida == nil || n.Salidas[0].Salida.MovimientoPadreId == nil || n.Salidas[0].Salida.MovimientoPadreId.Id <= 0 {
		return fmt.Errorf("la salida debe incluir un movimiento padre valido")
	}

	return nil
}

func persistirElementoSalida(o orm.Ormer, mov *Movimiento, elemento *ElementosMovimiento, now time.Time) error {
	if elemento == nil {
		return fmt.Errorf("se recibio un elemento de salida vacio")
	}

	elemento.MovimientoId = mov
	elemento.FechaModificacion = now

	existente, err := buscarElementoMovimientoPorActa(o, elemento.ElementoActaId)
	if err != nil {
		return err
	}

	accion, err := resolverAccionElementoSalida(existente)
	if err != nil {
		return err
	}

	switch accion {
	case reutilizarElementoSalida:
		elemento.Id = existente.Id
		elemento.ElementoActaId = existente.ElementoActaId
		elemento.FechaCreacion = existente.FechaCreacion
		_, err = o.Update(
			elemento,
			"MovimientoId",
			"Unidad",
			"ValorUnitario",
			"ValorTotal",
			"SaldoCantidad",
			"SaldoValor",
			"VidaUtil",
			"ValorResidual",
			"Activo",
			"FechaModificacion",
		)
		return err
	default:
		elemento.FechaCreacion = now
		_, err = o.Insert(elemento)
		return err
	}
}

func buscarElementoMovimientoPorActa(o orm.Ormer, elementoActaID *int) (*ElementosMovimiento, error) {
	if elementoActaID == nil || *elementoActaID <= 0 {
		return nil, nil
	}

	var elemento ElementosMovimiento
	err := o.QueryTable(new(ElementosMovimiento)).RelatedSel().Filter("ElementoActaId", *elementoActaID).One(&elemento)
	if err == orm.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if elemento.MovimientoId != nil && elemento.MovimientoId.Id > 0 {
		movimiento := Movimiento{Id: elemento.MovimientoId.Id}
		if err = o.QueryTable(new(Movimiento)).RelatedSel().Filter("Id", elemento.MovimientoId.Id).One(&movimiento); err != nil {
			return nil, err
		}
		elemento.MovimientoId = &movimiento
	}

	return &elemento, nil
}

func resolverAccionElementoSalida(existente *ElementosMovimiento) (accionPersistenciaElementoSalida, error) {
	if existente == nil {
		return insertarElementoSalida, nil
	}

	if existente.MovimientoId == nil {
		return reutilizarElementoSalida, nil
	}

	estado := ""
	movimientoID := existente.MovimientoId.Id
	if existente.MovimientoId.EstadoMovimientoId != nil {
		estado = existente.MovimientoId.EstadoMovimientoId.Nombre
	}

	if estadoSalidaPermiteReuso(estado) {
		return reutilizarElementoSalida, nil
	}

	return insertarElementoSalida, fmt.Errorf(
		"el elemento de acta %d ya tiene una salida asociada (movimiento %d, estado %q) y no puede registrarse nuevamente",
		valorElementoActa(existente.ElementoActaId),
		movimientoID,
		estado,
	)
}

func estadoSalidaPermiteReuso(nombre string) bool {
	estado := strings.ToLower(strings.TrimSpace(nombre))
	return estado == "salida rechazada" || estado == "salida anulada"
}

func valorElementoActa(elementoActaID *int) int {
	if elementoActaID == nil {
		return 0
	}
	return *elementoActaID
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
			err = r.(error)
			o.Rollback()
			logs.Error(r)
		} else if err != nil {
			o.Rollback()
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
			now := time.Now()
			m.Salida.FechaCreacion = now
			m.Salida.FechaModificacion = now
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
