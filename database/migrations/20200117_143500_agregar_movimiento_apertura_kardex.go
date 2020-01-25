package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarMovimientoAperturaKardex_20200117_143500 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarMovimientoAperturaKardex_20200117_143500{}
	m.Created = "20200117_143500"

	migration.Register("AgregarMovimientoAperturaKardex_20200117_143500", m)
}

// Run the migrations
func (m *AgregarMovimientoAperturaKardex_20200117_143500) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (id, nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES (10,'Apertura de Kardex', '{ \"MetodoValoracion\": \"number\", \"CantidadMinima\": \"number\", \"CantidadMaxima\": \"number\" }', 'Formato para realizar la apertura de kardex de un elemento', 'AP_KDX', 10.0, now(), now(), true);")

}

// Reverse the migrations
func (m *AgregarMovimientoAperturaKardex_20200117_143500) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'AP_KDX';")

}
