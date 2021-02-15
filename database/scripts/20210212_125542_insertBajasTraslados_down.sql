DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'SOL_TRD'
AND EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'SOL_TRD');
