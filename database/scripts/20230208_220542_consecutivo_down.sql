ALTER TABLE movimientos_arka.movimiento
    DROP COLUMN IF EXISTS consecutivo,
    DROP COLUMN IF EXISTS consecutivo_id,
    DROP COLUMN IF EXISTS fecha_corte,
    DROP CONSTRAINT IF EXISTS uq_consecutivo,
    DROP CONSTRAINT IF EXISTS uq_consecutivo_id;