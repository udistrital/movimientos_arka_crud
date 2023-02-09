ALTER TABLE movimientos_arka.movimiento
    ADD COLUMN IF NOT EXISTS consecutivo VARCHAR(20),
    ADD COLUMN IF NOT EXISTS consecutivo_id INTEGER,
    ADD COLUMN IF NOT EXISTS fecha_corte DATE,
    ADD CONSTRAINT uq_consecutivo UNIQUE (consecutivo),
    ADD CONSTRAINT uq_consecutivo_id UNIQUE (consecutivo_id);

UPDATE movimientos_arka.movimiento
SET (
    consecutivo,
    consecutivo_id,
    fecha_corte
    ) = (
    COALESCE(detalle ->>'consecutivo', detalle ->>'Consecutivo'),
    CAST(detalle ->>'ConsecutivoId' AS INTEGER),
    CAST(detalle ->>'FechaCorte' AS DATE)
);
