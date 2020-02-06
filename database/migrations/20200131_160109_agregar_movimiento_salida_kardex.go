package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarMovimientoSalidaKardex_20200131_160109 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarMovimientoSalidaKardex_20200131_160109{}
	m.Created = "20200131_160109"

	migration.Register("AgregarMovimientoSalidaKardex_20200131_160109", m)
}

// Run the migrations
func (m *AgregarMovimientoSalidaKardex_20200131_160109) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (id, nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES (12,'Salida de Kardex', '{ }', 'Formato para realizar la salida de kardex de un elemento', 'SAL_KDX', 12.0, now(), now(), true);")

}

// Reverse the migrations
func (m *AgregarMovimientoSalidaKardex_20200131_160109) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'SAL_KDX';")

}
