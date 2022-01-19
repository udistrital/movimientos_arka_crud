UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'EA',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_ADQ';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'EEP',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_EP';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ED',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_DN';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ESI',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_SI';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ET',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_TR';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'EPR',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_RP';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ECM',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_CM';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ECE',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_CE';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'EPPA',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_PPA';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'EIA',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_IA';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'EID',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_ID';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'EBEMP',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ENT_BEP';

DELETE FROM movimientos_arka.formato_tipo_movimiento
    WHERE codigo_abreviacion = 'ENT_AM';
