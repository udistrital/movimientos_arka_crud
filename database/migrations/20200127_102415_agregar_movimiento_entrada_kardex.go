package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarMovimientoEntradaKardex_20200127_102415 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarMovimientoEntradaKardex_20200127_102415{}
	m.Created = "20200127_102415"

	migration.Register("AgregarMovimientoEntradaKardex_20200127_102415", m)
}

// Run the migrations
func (m *AgregarMovimientoEntradaKardex_20200127_102415) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (id, nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES (11,'Entrada de Kardex', '{ }', 'Formato para realizar la entrada de kardex de un elemento', 'ENT_KDX', 11.0, now(), now(), true);")

}

// Reverse the migrations
func (m *AgregarMovimientoEntradaKardex_20200127_102415) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ENT_KDX';")
}
