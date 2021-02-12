DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EIA'
AND EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EIA');
DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EID'
AND EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EID');
DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EBEMP'
AND EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EBEMP');
