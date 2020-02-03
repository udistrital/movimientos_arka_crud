package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjusteSolicitudBodega_20200203_140309 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjusteSolicitudBodega_20200203_140309{}
	m.Created = "20200203_140309"

	migration.Register("AjusteSolicitudBodega_20200203_140309", m)
}

// Run the migrations
func (m *AjusteSolicitudBodega_20200203_140309) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("UPDATE movimientos_arka.formato_tipo_movimiento SET formato = '{ \"Funcionario\": \"number\", \"Elementos\": [ { \"Ubicacion\": \"number\", \"ElementoActa\": \"number\", \"Cantidad\": \"number\", \"CantidadAprobada\": \"number\" } ] }' WHERE codigo_abreviacion = 'SOL_BOD';")

}

// Reverse the migrations
func (m *AjusteSolicitudBodega_20200203_140309) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("UPDATE movimientos_arka.formato_tipo_movimiento SET formato = '{ \"Funcionario\": \"number\", \"Elementos\": [ { \"Ubicacion\": \"number\", \"ElementoActa\": \"number\", \"Cantidad\": \"number\" } ] }' WHERE codigo_abreviacion = 'SOL_BOD';")

}
