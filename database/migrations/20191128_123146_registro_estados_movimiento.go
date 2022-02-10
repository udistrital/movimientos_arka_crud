package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type RegistroEstadosMovimiento_20191128_123146 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RegistroEstadosMovimiento_20191128_123146{}
	m.Created = "20191128_123146"

	migration.Register("RegistroEstadosMovimiento_20191128_123146", m)
}

// Run the migrations
func (m *RegistroEstadosMovimiento_20191128_123146) Up() {
	m.SQL("INSERT INTO movimientos_arka.estado_movimiento(nombre, activo, fecha_creacion, fecha_modificacion, descripcion) VALUES ('Entrada Aceptada', TRUE, now(), now(), 'Formato para marcar una entrada como aceptada/aprobada');")
	m.SQL("INSERT INTO movimientos_arka.estado_movimiento(nombre, activo, fecha_creacion, fecha_modificacion, descripcion) VALUES ('Salida Aceptada', TRUE, now(), now(), 'Formato para marcar una salida como aceptada/aprobada');")
	m.SQL("INSERT INTO movimientos_arka.estado_movimiento(nombre, activo, fecha_creacion, fecha_modificacion, descripcion) VALUES ('Entrada Con Salida', TRUE, now(), now(), 'Formato para marcar una entrada la cual ya tuvo salida');")
}

// Reverse the migrations
func (m *RegistroEstadosMovimiento_20191128_123146) Down() {
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE nombre = 'Entrada Aceptada';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE nombre = 'Salida Aceptada';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE nombre = 'Entrada Con Salida';")
}
