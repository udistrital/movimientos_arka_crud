-- Novedad Elemento

DROP TABLE IF EXISTS movimientos_arka.novedad_elemento;

-- Elementos Movimiento

ALTER TABLE movimientos_arka.elementos_movimiento
    DROP COLUMN IF EXISTS vida_util,
    DROP COLUMN IF EXISTS valor_residual;
