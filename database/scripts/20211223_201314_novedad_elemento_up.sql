-- -- Novedad Elemento

-- CREATE TABLE IF NOT EXISTS movimientos_arka.novedad_elemento (
--     id SERIAL NOT NULL,
--     vida_util NUMERIC(10,5) NOT NULL,
--     valor_libros NUMERIC(20,7) NOT NULL,
--     valor_residual NUMERIC(20,7) NOT NULL,
--     metadata JSONB,
--     movimiento_id INTEGER NOT NULL,
--     elemento_movimiento_id INTEGER NOT NULL,
--     activo BOOLEAN NOT NULL,
--     fecha_creacion TIMESTAMP NOT NULL,
--     fecha_modificacion TIMESTAMP NOT NULL,
--     CONSTRAINT pk_novedad_elemento PRIMARY KEY (id)
-- );

-- ALTER TABLE movimientos_arka.novedad_elemento
--     ADD CONSTRAINT fk_novedad_elemento_elementos_movimiento FOREIGN KEY (elemento_movimiento_id)
--     REFERENCES movimientos_arka.elementos_movimiento (id) MATCH FULL
--     ON DELETE RESTRICT ON UPDATE CASCADE;

-- ALTER TABLE movimientos_arka.novedad_elemento
--     ADD CONSTRAINT fk_novedad_elemento_movimiento FOREIGN KEY (movimiento_id)
--     REFERENCES movimientos_arka.movimiento (id) MATCH FULL
--     ON DELETE RESTRICT ON UPDATE CASCADE;

-- Elementos Movimiento

ALTER TABLE movimientos_arka.elementos_movimiento
    ADD COLUMN IF NOT EXISTS vida_util NUMERIC(10,5) DEFAULT 5,
    ADD COLUMN IF NOT EXISTS valor_residual NUMERIC(20,7) DEFAULT 0;

ALTER TABLE movimientos_arka.elementos_movimiento
    ALTER COLUMN vida_util DROP DEFAULT;

ALTER TABLE movimientos_arka.elementos_movimiento
    ALTER COLUMN valor_residual DROP DEFAULT;
