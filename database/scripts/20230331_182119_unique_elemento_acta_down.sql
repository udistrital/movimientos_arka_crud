ALTER TABLE movimientos_arka.elementos_movimiento
DROP CONSTRAINT IF EXISTS uq_elemento_acta_id,
DROP CONSTRAINT IF EXISTS ck_valor_total_valor_residual_elementos_movimiento,
DROP CONSTRAINT IF EXISTS ck_valor_residual_elementos_movimiento;

UPDATE movimientos_arka.elementos_movimiento
SET elemento_acta_id = 0
WHERE elemento_acta_id = NULL;

ALTER TABLE movimientos_arka.elementos_movimiento
ALTER COLUMN elemento_acta_id SET NOT NULL;

ALTER TABLE movimientos_arka.novedad_elemento
DROP CONSTRAINT IF EXISTS ck_valor_libros_valor_residual_novedad_elemento,
DROP CONSTRAINT IF EXISTS ck_valor_residual_novedad_elemento;
