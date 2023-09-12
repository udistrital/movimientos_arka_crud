ALTER TABLE movimientos_arka.centro_costos
DROP COLUMN dependencia_id,
DROP COLUMN sede_id;

ALTER TABLE movimientos_arka.centro_costos
ADD COLUMN dependencia VARCHAR(150),
ADD COLUMN sede VARCHAR(50),
ALTER COLUMN nombre TYPE VARCHAR(100);

COMMENT ON COLUMN movimientos_arka.centro_costos.dependencia IS 'Dependencia asociada al centro de costos';
COMMENT ON COLUMN movimientos_arka.centro_costos.sede IS 'Sede asociada al centro de costos';
