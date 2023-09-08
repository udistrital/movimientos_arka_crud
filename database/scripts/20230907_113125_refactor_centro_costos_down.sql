ALTER TABLE movimientos_arka.centro_costos
DROP COLUMN dependencia,
DROP COLUMN sede;

ALTER TABLE movimientos_arka.centro_costos
ADD COLUMN dependencia_id INTEGER,
ADD COLUMN sede_id INTEGER,
ALTER COLUMN nombre TYPE TEXT;

COMMENT ON COLUMN movimientos_arka.centro_costos.dependencia_id IS 'Dependencia asociada al centro de costos';
COMMENT ON COLUMN movimientos_arka.centro_costos.sede_id IS 'Sede asociada al centro de costos';
