package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarFormatoSalidas_20191218_170646 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarFormatoSalidas_20191218_170646{}
	m.Created = "20191218_170646"

	migration.Register("AgregarFormatoSalidas_20191218_170646", m)
}

// Run the migrations
func (m *AgregarFormatoSalidas_20191218_170646) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Salida', '{ \"funcionario\": \"number\", \"ubicacion\": \"number\" }', 'Formato para guardar una salida', 'SAL', 7.0, now(), now(), true);")

}

// Reverse the migrations
func (m *AgregarFormatoSalidas_20191218_170646) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'SAL';")

}
