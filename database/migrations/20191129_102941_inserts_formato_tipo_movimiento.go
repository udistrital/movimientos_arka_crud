package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertsFormatoTipoMovimiento_20191129_102941 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsFormatoTipoMovimiento_20191129_102941{}
	m.Created = "20191129_102941"

	migration.Register("InsertsFormatoTipoMovimiento_20191129_102941", m)
}

// Run the migrations
func (m *InsertsFormatoTipoMovimiento_20191129_102941) Up() {
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Adquisición', '{\"acta_recibido_id\": \"number\", \"consecutivo\":\"string\",  \"documento_contable_id\":\"number\",\"contrato_id\":\"number\", \"vigencia_contrato\":\"number\",\"importacion\":\"boolean\"}', 'Formato para guardar una entrada por adquisición', 'EA', 1.0, now(), now(), true);")
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Elaboración Propia', '{\"acta_recibido_id\": \"number\", \"consecutivo\": \"string\", \"documento_contable_id\": \"number\", \"vigencia_ordenador\": \"string\", \"ordenador_gasto_id\": \"number\", \"solicitante_id\": \"number\" }', 'Formato para guardar una entrada por elaboración propia', 'EEP', 2.0, now(), now(), true);")
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Donación', '{ \"acta_recibido_id\": \"number\", \"contrato_id\": \"number\", \"vigencia_contrato\": \"string\", \"consecutivo\": \"string\", \"documento_contable_id\": \"number\", \"vigencia_solicitante\": \"string\", \"ordenador_gasto_id\": \"number\" }', 'Formato para guardar una entrada por donación', 'ED', 3.0, now(), now(), true);")
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Sobrante', '{ \"acta_recibido_id\": \"number\", \"consecutivo\": \"string\", \"documento_contable_id\": \"number\", \"vigencia_ordenador\": \"string\", \"ordenador_gasto_id\": \"number\" }', 'Formato para guardar una entrada por sobrante de inventario', 'ESI', 4.0, now(), now(), true);")
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Terceros', '{ \"acta_recibido_id\": \"number\", \"consecutivo\": \"string\", \"documento_contable_id\": \"number\", \"contrato_id\": \"number\", \"vigencia_contrato\": \"number\", \"tercero_id\": \"number\" }', 'Formato para guardar una entrada por terceros', 'ET', 5.0, now(), now(), true);")
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES ('Reposición', '{ \"acta_recibido_id\": \"number\", \"consecutivo\": \"string\", \"documento_contable_id\": \"number\", \"placa_id\": \"string\", \"encargado_id\": \"number\" }', 'Formato para guardar una entrada por reposición', 'EPR', 6.0, now(), now(), true);")

}

// Reverse the migrations
func (m *InsertsFormatoTipoMovimiento_20191129_102941) Down() {
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE codigo_abreviacion = 'EA';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE codigo_abreviacion = 'EEP';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE codigo_abreviacion = 'ED';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE codigo_abreviacion = 'ESI';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE codigo_abreviacion = 'ET';")
	m.SQL("DELETE FROM movimientos_arka.estado_movimiento WHERE codigo_abreviacion = 'EPR';")
}
