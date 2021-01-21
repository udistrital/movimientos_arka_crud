DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECM'
AND EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECM');
DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECE'
AND EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECE');
DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EPPA'
AND EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'EPPA');
