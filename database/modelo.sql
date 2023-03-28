-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.3
-- PostgreSQL version: 9.6
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: postgres | type: DATABASE --
-- DROP DATABASE IF EXISTS postgres;
CREATE DATABASE postgres
	ENCODING = 'UTF8'
	LC_COLLATE = 'en_US.UTF-8'
	LC_CTYPE = 'en_US.UTF-8'
	TABLESPACE = pg_default
	OWNER = postgres;
-- ddl-end --
COMMENT ON DATABASE postgres IS E'default administrative connection database';
-- ddl-end --


-- object: movimientos_arka | type: SCHEMA --
-- DROP SCHEMA IF EXISTS movimientos_arka CASCADE;
CREATE SCHEMA movimientos_arka;
-- ddl-end --
ALTER SCHEMA movimientos_arka OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,movimientos_arka;
-- ddl-end --

-- object: movimientos_arka.movimiento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS movimientos_arka.movimiento_id_seq CASCADE;
CREATE SEQUENCE movimientos_arka.movimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE movimientos_arka.movimiento_id_seq OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.movimiento CASCADE;
CREATE TABLE movimientos_arka.movimiento (
	id integer NOT NULL DEFAULT nextval('movimientos_arka.movimiento_id_seq'::regclass),
	observacion character varying(500),
	detalle jsonb NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	movimiento_padre_id integer,
	formato_tipo_movimiento_id integer NOT NULL,
	estado_movimiento_id integer NOT NULL,
	consecutivo character varying(20),
	consecutivo_id integer,
	fecha_corte date,
	CONSTRAINT pk_movimiento PRIMARY KEY (id),
	CONSTRAINT uq_consecutivo UNIQUE (consecutivo),
	CONSTRAINT uq_consecutivo_id UNIQUE (consecutivo_id)

);
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.movimiento.id IS E'Identificador de la tabla movimiento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_movimiento ON movimientos_arka.movimiento  IS E'Llave primaria de la tabla movimiento';
-- ddl-end --
ALTER TABLE movimientos_arka.movimiento OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.estado_movimiento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS movimientos_arka.estado_movimiento_id_seq CASCADE;
CREATE SEQUENCE movimientos_arka.estado_movimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE movimientos_arka.estado_movimiento_id_seq OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.estado_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.estado_movimiento CASCADE;
CREATE TABLE movimientos_arka.estado_movimiento (
	id integer NOT NULL DEFAULT nextval('movimientos_arka.estado_movimiento_id_seq'::regclass),
	nombre character varying(40) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	descripcion character varying(300) NOT NULL,
	CONSTRAINT pk_estado_movimiento PRIMARY KEY (id),
	CONSTRAINT uq_nombre_estado_movimiento UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.estado_movimiento IS E'Tabla para almacenar los posibles estados por los que puede pasar un movimiento de almacén';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.estado_movimiento.nombre IS E'Nombre del estado del movimiento';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.estado_movimiento.descripcion IS E'Descripcion del estado del movimiento';
-- ddl-end --
ALTER TABLE movimientos_arka.estado_movimiento OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.soporte_movimiento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS movimientos_arka.soporte_movimiento_id_seq CASCADE;
CREATE SEQUENCE movimientos_arka.soporte_movimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE movimientos_arka.soporte_movimiento_id_seq OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.soporte_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.soporte_movimiento CASCADE;
CREATE TABLE movimientos_arka.soporte_movimiento (
	id integer NOT NULL DEFAULT nextval('movimientos_arka.soporte_movimiento_id_seq'::regclass),
	documento_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	movimiento_id integer NOT NULL,
	CONSTRAINT pk_soporte_movimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.soporte_movimiento IS E'Tabla que almacena los documentos/soportes relacionados a un movimiento, ej: factura, concepto técnico de la red de datos, denuncio, entre otros.';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.soporte_movimiento.documento_id IS E'Referencia al API de documentos del Core';
-- ddl-end --
COMMENT ON CONSTRAINT pk_soporte_movimiento ON movimientos_arka.soporte_movimiento  IS E'Llave primaria de la tabla soporte_movimiento';
-- ddl-end --
ALTER TABLE movimientos_arka.soporte_movimiento OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.elementos_movimiento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS movimientos_arka.elementos_movimiento_id_seq CASCADE;
CREATE SEQUENCE movimientos_arka.elementos_movimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE movimientos_arka.elementos_movimiento_id_seq OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.elementos_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.elementos_movimiento CASCADE;
CREATE TABLE movimientos_arka.elementos_movimiento (
	id integer NOT NULL DEFAULT nextval('movimientos_arka.elementos_movimiento_id_seq'::regclass),
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
	elemento_catalogo_id integer,
	vida_util numeric(15,10),
	valor_residual numeric(20,7),
	CONSTRAINT pk_elementos_movimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.elementos_movimiento IS E'Tabla que almacena los elementos relacionados con un movimiento de almacén';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.elementos_movimiento.elemento_acta_id IS E'Referencia al elemento del acta de recibido';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.elementos_movimiento.fecha_creacion IS E'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.elementos_movimiento.fecha_modificacion IS E'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_elementos_movimiento ON movimientos_arka.elementos_movimiento  IS E'Llave primaria de la tabla elementos_movimiento';
-- ddl-end --
ALTER TABLE movimientos_arka.elementos_movimiento OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.formato_tipo_movimiento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS movimientos_arka.formato_tipo_movimiento_id_seq CASCADE;
CREATE SEQUENCE movimientos_arka.formato_tipo_movimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE movimientos_arka.formato_tipo_movimiento_id_seq OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.formato_tipo_movimiento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.formato_tipo_movimiento CASCADE;
CREATE TABLE movimientos_arka.formato_tipo_movimiento (
	id integer NOT NULL DEFAULT nextval('movimientos_arka.formato_tipo_movimiento_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	formato jsonb NOT NULL,
	descripcion character varying,
	codigo_abreviacion character varying(50),
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	CONSTRAINT pk_tipo_entrada PRIMARY KEY (id),
	CONSTRAINT uq_codigo_abreviacion UNIQUE (codigo_abreviacion)

);
-- ddl-end --
COMMENT ON TABLE movimientos_arka.formato_tipo_movimiento IS E'Tabla para parametrizar los campos de la entrada en el JSON.';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.formato_tipo_movimiento.formato IS E'Campos necesarios para cada tipo movimiento especificado en un JSON';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.formato_tipo_movimiento.fecha_creacion IS E'Campo para almacenar la fecha en que se realiza el registro. Permite llevar una trazabilidad.';
-- ddl-end --
COMMENT ON COLUMN movimientos_arka.formato_tipo_movimiento.fecha_modificacion IS E'Campo para almacenar las fechas en que se realizan cambios. Permite llevar una trazabilidad.';
-- ddl-end --
ALTER TABLE movimientos_arka.formato_tipo_movimiento OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.novedad_elemento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS movimientos_arka.novedad_elemento_id_seq CASCADE;
CREATE SEQUENCE movimientos_arka.novedad_elemento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE movimientos_arka.novedad_elemento_id_seq OWNER TO postgres;
-- ddl-end --

-- object: movimientos_arka.novedad_elemento | type: TABLE --
-- DROP TABLE IF EXISTS movimientos_arka.novedad_elemento CASCADE;
CREATE TABLE movimientos_arka.novedad_elemento (
	id integer NOT NULL DEFAULT nextval('movimientos_arka.novedad_elemento_id_seq'::regclass),
	vida_util numeric(15,10) NOT NULL,
	valor_libros numeric(20,7) NOT NULL,
	valor_residual numeric(20,7) NOT NULL,
	metadata jsonb,
	movimiento_id integer NOT NULL,
	elemento_movimiento_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_novedad_elemento PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE movimientos_arka.novedad_elemento OWNER TO postgres;
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

-- object: fk_movimiento_padre | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.movimiento DROP CONSTRAINT IF EXISTS fk_movimiento_padre CASCADE;
ALTER TABLE movimientos_arka.movimiento ADD CONSTRAINT fk_movimiento_padre FOREIGN KEY (movimiento_padre_id)
REFERENCES movimientos_arka.movimiento (id) MATCH FULL
ON DELETE CASCADE ON UPDATE CASCADE;
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

-- object: fk_novedad_elemento_elementos_movimiento | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.novedad_elemento DROP CONSTRAINT IF EXISTS fk_novedad_elemento_elementos_movimiento CASCADE;
ALTER TABLE movimientos_arka.novedad_elemento ADD CONSTRAINT fk_novedad_elemento_elementos_movimiento FOREIGN KEY (elemento_movimiento_id)
REFERENCES movimientos_arka.elementos_movimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_novedad_elemento_movimiento | type: CONSTRAINT --
-- ALTER TABLE movimientos_arka.novedad_elemento DROP CONSTRAINT IF EXISTS fk_novedad_elemento_movimiento CASCADE;
ALTER TABLE movimientos_arka.novedad_elemento ADD CONSTRAINT fk_novedad_elemento_movimiento FOREIGN KEY (movimiento_id)
REFERENCES movimientos_arka.movimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: "grant_CU_f2b0a53cd7" | type: PERMISSION --
GRANT CREATE,USAGE
   ON SCHEMA movimientos_arka
   TO postgres;
-- ddl-end --

-- object: "grant_rawdDxt_5601761752" | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE,REFERENCES,TRIGGER
   ON TABLE movimientos_arka.movimiento
   TO postgres;
-- ddl-end --

-- object: "grant_rawdDxt_1c03f39d85" | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE,REFERENCES,TRIGGER
   ON TABLE movimientos_arka.estado_movimiento
   TO postgres;
-- ddl-end --

-- object: "grant_rawdDxt_7fa52ac684" | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE,REFERENCES,TRIGGER
   ON TABLE movimientos_arka.soporte_movimiento
   TO postgres;
-- ddl-end --

-- object: "grant_rawdDxt_4412213fd8" | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE,REFERENCES,TRIGGER
   ON TABLE movimientos_arka.elementos_movimiento
   TO postgres;
-- ddl-end --

-- object: "grant_rawdDxt_6404033567" | type: PERMISSION --
GRANT SELECT,INSERT,UPDATE,DELETE,TRUNCATE,REFERENCES,TRIGGER
   ON TABLE movimientos_arka.formato_tipo_movimiento
   TO postgres;
-- ddl-end --


