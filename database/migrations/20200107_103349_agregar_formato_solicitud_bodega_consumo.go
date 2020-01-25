package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarFormatoSolicitudBodegaConsumo_20200107_103349 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarFormatoSolicitudBodegaConsumo_20200107_103349{}
	m.Created = "20200107_103349"

	migration.Register("AgregarFormatoSolicitudBodegaConsumo_20200107_103349", m)
}

// Run the migrations
func (m *AgregarFormatoSolicitudBodegaConsumo_20200107_103349) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Solicitud Elementos Bodega de Consumo', '{ \"Elementos\": [ { \"Funcionario\": \"number\", \"Ubicacion\": \"number\", \"ElementoActa\": \"number\", \"Cantidad\": \"number\" } ] }', 'Formato para guardar una solicitud de Bodega de Consumo', 'SOL_BOD', 8.0, now(), now(), true);")

}

// Reverse the migrations
func (m *AgregarFormatoSolicitudBodegaConsumo_20200107_103349) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'SOL_BOD';")

}
