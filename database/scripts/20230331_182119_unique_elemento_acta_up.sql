-- ALTER TABLE movimientos_arka.elementos_movimiento
-- ALTER COLUMN elemento_acta_id DROP NOT NULL;

-- UPDATE movimientos_arka.elementos_movimiento
-- SET elemento_acta_id = NULL
-- WHERE elemento_acta_id = 0;

-- ALTER TABLE movimientos_arka.elementos_movimiento
-- ADD CONSTRAINT uq_elemento_acta_id UNIQUE (elemento_acta_id),
-- ADD CONSTRAINT ck_valor_total_valor_residual_elementos_movimiento CHECK(valor_total >= valor_residual),
-- ADD CONSTRAINT ck_valor_residual_elementos_movimiento CHECK(valor_residual >= 0);

-- ALTER TABLE movimientos_arka.novedad_elemento
-- ADD CONSTRAINT ck_valor_libros_valor_residual_novedad_elemento CHECK(valor_libros >= valor_residual),
-- ADD CONSTRAINT ck_valor_residual_novedad_elemento CHECK(valor_residual >= 0);
