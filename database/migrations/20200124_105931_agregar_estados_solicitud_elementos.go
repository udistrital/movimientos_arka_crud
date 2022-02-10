package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarEstadosSolicitudElementos_20200124_105931 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarEstadosSolicitudElementos_20200124_105931{}
	m.Created = "20200124_105931"

	migration.Register("AgregarEstadosSolicitudElementos_20200124_105931", m)
}

// Run the migrations
func (m *AgregarEstadosSolicitudElementos_20200124_105931) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.estado_movimiento(nombre, activo, fecha_creacion, fecha_modificacion, descripcion) VALUES ('Solicitud Pendiente', TRUE, now(), now(), 'Estado para marcar una solicitud pendiente');")
	m.SQL("INSERT INTO movimientos_arka.estado_movimiento(nombre, activo, fecha_creacion, fecha_modificacion, descripcion) VALUES ('Solicitud Aprobada', TRUE, now(), now(), 'Estado para marcar una solicitud Aprobada');")
	m.SQL("INSERT INTO movimientos_arka.estado_movimiento(nombre, activo, fecha_creacion, fecha_modificacion, descripcion) VALUES ('Solicitud Parcial', TRUE, now(), now(), 'Estado para marcar una solicitud Aprobada Parcialmente');")
	m.SQL("INSERT INTO movimientos_arka.estado_movimiento(nombre, activo, fecha_creacion, fecha_modificacion, descripcion) VALUES ('Solicitud Rechazada', TRUE, now(), now(), 'Estado para marcar una solicitud Aprobada');")
}

// Reverse the migrations
func (m *AgregarEstadosSolicitudElementos_20200124_105931) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE id = 'Solicitud Pendiente';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE id = 'Solicitud Aprobada';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE id = 'Solicitud Parcial';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE id = 'Solicitud Rechazada';")
}
