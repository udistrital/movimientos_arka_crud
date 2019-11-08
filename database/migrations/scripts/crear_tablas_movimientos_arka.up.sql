-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.2-beta
-- PostgreSQL version: 9.5
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: new_database | type: DATABASE --
-- -- DROP DATABASE IF EXISTS new_database;
-- CREATE DATABASE new_database;
-- -- ddl-end --
-- 

-- object: movimientos_arka | type: SCHEMA --
-- DROP SCHEMA IF EXISTS movimientos_arka CASCADE;
CREATE SCHEMA movimientos_arka;
-- ddl-end --

SET search_path TO pg_catalog,public,movimientos_arka;
-- ddl-end --

-- object: movimientos_arka.movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.movimiento CASCADE;
CREATE TABLE movimientos_arka.movimiento (
	id serial NOT NULL,
	observacion character varying(500),
	detalle json NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	movimiento_padre_id integer,
	formato_tipo_movimiento_id integer NOT NULL,
	estado_movimiento_id integer NOT NULL,
	CONSTRAINT pk_movimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.movimiento.id IS 'Identificador de la tabla movimiento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_movimiento ON movimientos_arka.movimiento  IS 'Llave primaria de la tabla movimiento';
-- ddl-end --

-- object: movimientos_arka.estado_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.estado_movimiento CASCADE;
CREATE TABLE movimientos_arka.estado_movimiento (
	id serial NOT NULL,
	nombre varchar(20) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	descripcion varchar(300) NOT NULL,
	CONSTRAINT pk_estado_movimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.estado_movimiento IS 'Tabla para almacenar los posibles estados por los que puede pasar un movimiento de almacén';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.estado_movimiento.nombre IS 'Nombre del estado del movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.estado_movimiento.descripcion IS 'Descripcion del estado del movimiento';
-- ddl-end --

-- object: movimientos_arka.soporte_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.soporte_movimiento CASCADE;
CREATE TABLE movimientos_arka.soporte_movimiento (
	id serial NOT NULL,
	documento_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	movimiento_id integer NOT NULL,
	CONSTRAINT pk_soporte_movimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.soporte_movimiento IS 'Tabla que almacena los documentos/soportes relacionados a un movimiento, ej: factura, concepto técnico de la red de datos, denuncio, entre otros.';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.soporte_movimiento.documento_id IS 'Referencia al API de documentos del Core';
-- ddl-end --
COMMENT ON CONSTRAINT pk_soporte_movimiento ON movimientos_arka.soporte_movimiento  IS 'Llave primaria de la tabla soporte_movimiento';
-- ddl-end --

-- object: movimientos_arka.elementos_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.elementos_movimiento CASCADE;
CREATE TABLE movimientos_arka.elementos_movimiento (
	id serial NOT NULL,
	elemento_acta_id integer NOT NULL,
	unidad numeric(6,2) NOT NULL,
	valor_unitario numeric(20,7) NOT NULL,
	valor_total numeric(20,7) NOT NULL,
	saldo_cantidad numeric(5,2) NOT NULL,
	saldo_valor numeric(20,7) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	movimiento_id integer,
	CONSTRAINT pk_elementos_movimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.elementos_movimiento IS 'Tabla que almacena los elementos relacionados con un movimiento de almacén';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.elementos_movimiento.elemento_acta_id IS 'Referencia al elemento del acta de recibido';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.elementos_movimiento.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.elementos_movimiento.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_elementos_movimiento ON movimientos_arka.elementos_movimiento  IS 'Llave primaria de la tabla elementos_movimiento';
-- ddl-end --

-- object: movimientos_arka.formato_tipo_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.formato_tipo_movimiento CASCADE;
CREATE TABLE movimientos_arka.formato_tipo_movimiento (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	formato json NOT NULL,
	descripcion varchar,
	codigo_abreviacion varchar(50),
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_tipo_entrada PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.formato_tipo_movimiento IS 'Tabla para parametrizar los campos de la entrada en el JSON.';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.formato_tipo_movimiento.formato IS 'Campos necesarios para cada tipo movimiento especificado en un JSON';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.formato_tipo_movimiento.fecha_creacion IS 'Campo para almacenar la fecha en que se realiza el registro. Permite llevar una trazabilidad.';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.formato_tipo_movimiento.fecha_modificacion IS 'Campo para almacenar las fechas en que se realizan cambios. Permite llevar una trazabilidad.';
-- ddl-end --

-- object: fk_movimiento_formato_tipo_movimiento | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.movimiento DROP CONSTRAINT IF EXISTS fk_movimiento_formato_tipo_movimiento CASCADE;
ALTER TABLE movimientos_arka.movimiento ADD CONSTRAINT fk_movimiento_formato_tipo_movimiento FOREIGN KEY (formato_tipo_movimiento_id)
REFERENCES movimientos_arka.formato_tipo_movimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_movimiento_estado_movimiento | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.movimiento DROP CONSTRAINT IF EXISTS fk_movimiento_estado_movimiento CASCADE;
ALTER TABLE movimientos_arka.movimiento ADD CONSTRAINT fk_movimiento_estado_movimiento FOREIGN KEY (estado_movimiento_id)
REFERENCES movimientos_arka.estado_movimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_soporte_movimiento_movimiento | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.soporte_movimiento DROP CONSTRAINT IF EXISTS fk_soporte_movimiento_movimiento CASCADE;
ALTER TABLE movimientos_arka.soporte_movimiento ADD CONSTRAINT fk_soporte_movimiento_movimiento FOREIGN KEY (movimiento_id)
REFERENCES movimientos_arka.movimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_elementos_movimiento_movimiento | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.elementos_movimiento DROP CONSTRAINT IF EXISTS fk_elementos_movimiento_movimiento CASCADE;
ALTER TABLE movimientos_arka.elementos_movimiento ADD CONSTRAINT fk_elementos_movimiento_movimiento FOREIGN KEY (movimiento_id)
REFERENCES movimientos_arka.movimiento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_movimiento_padre | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.movimiento DROP CONSTRAINT IF EXISTS fk_movimiento_padre CASCADE;
ALTER TABLE movimientos_arka.movimiento ADD CONSTRAINT fk_movimiento_padre FOREIGN KEY (movimiento_padre_id)
REFERENCES movimientos_arka.movimiento (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --


