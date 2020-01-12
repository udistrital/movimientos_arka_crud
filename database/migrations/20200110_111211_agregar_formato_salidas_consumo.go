package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarFormatoSalidasConsumo_20200110_111211 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarFormatoSalidasConsumo_20200110_111211{}
	m.Created = "20200110_111211"

	migration.Register("AgregarFormatoSalidasConsumo_20200110_111211", m)
}

// Run the migrations
func (m *AgregarFormatoSalidasConsumo_20200110_111211) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (id, nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES (9, 'Salida de Consumo', '{ \"ubicacion\": \"number\" }', 'Formato para guardar una salida para la bodega de consumo', 'SAL_CONS', 9.0, now(), now(), true);")

}

// Reverse the migrations
func (m *AgregarFormatoSalidasConsumo_20200110_111211) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE codigo_abreviacion = 'SAL_CONS';")

}
