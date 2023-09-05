package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarSolicitudBajas_20200302_135939 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarSolicitudBajas_20200302_135939{}
	m.Created = "20200302_135939"

	migration.Register("AgregarSolicitudBajas_20200302_135939", m)
}

// Run the migrations
func (m *AgregarSolicitudBajas_20200302_135939) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("INSERT INTO movimientos_arka.formato_tipo_movimiento (id, nombre, formato, descripcion, codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo) VALUES (13,'Solicitud de Bajas', '{ \"Funcionario\": \"number\", \"Ubicacion\": \"number\", \"Revisor\": \"number\", \"FechaVistoBueno\": \"string\", \"Elementos\": [ { \"Id\": \"number\", \"Soporte\": \"number\", \"TipoBaja\": \"number\", \"Observaciones\": \"string\" } ] }', 'Formato para realizar la solicitud de bajas de varios elementos', 'SOL_BAJA', 13.0, now(), now(), true);")
}

// Reverse the migrations
func (m *AgregarSolicitudBajas_20200302_135939) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'SOL_BAJA';")

}
