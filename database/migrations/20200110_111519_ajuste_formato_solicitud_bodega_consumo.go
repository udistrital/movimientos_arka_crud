package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AjusteFormatoSolicitudBodegaConsumo_20200110_111519 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AjusteFormatoSolicitudBodegaConsumo_20200110_111519{}
	m.Created = "20200110_111519"

	migration.Register("AjusteFormatoSolicitudBodegaConsumo_20200110_111519", m)
}

// Run the migrations
func (m *AjusteFormatoSolicitudBodegaConsumo_20200110_111519) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("UPDATE movimientos_arka.formato_tipo_movimiento SET formato = '{ \"Funcionario\": \"number\", \"Elementos\": [ { \"Ubicacion\": \"number\", \"ElementoActa\": \"number\", \"Cantidad\": \"number\" } ] }' WHERE codigo_abreviacion = 'SOL_BOD';")
}

// Reverse the migrations
func (m *AjusteFormatoSolicitudBodegaConsumo_20200110_111519) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("UPDATE movimientos_arka.formato_tipo_movimiento SET formato = '{ \"Elementos\": [ { \"Funcionario\": \"number\", \"Ubicacion\": \"number\", \"ElementoActa\": \"number\", \"Cantidad\": \"number\" } ] }' WHERE codigo_abreviacion = 'SOL_BOD';")

}
