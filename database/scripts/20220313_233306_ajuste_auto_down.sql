DELETE FROM movimientos_arka.formato_tipo_movimiento
    WHERE codigo_abreviacion = 'AAT';

ALTER TABLE movimientos_arka.formato_tipo_movimiento
	DROP CONSTRAINT uq_codigo_abreviacion;
