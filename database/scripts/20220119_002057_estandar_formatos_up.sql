UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_ADQ',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'EA';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_EP',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'EEP';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_DN',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ED';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_SI',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ESI';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_TR',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ET';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_RP',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'EPR';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_CM',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ECM';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_CE',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'ECE';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_PPA',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'EPPA';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_IA',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'EIA';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_ID',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'EID';

UPDATE movimientos_arka.formato_tipo_movimiento SET
		codigo_abreviacion = 'ENT_BEP',
		fecha_modificacion = now()
WHERE codigo_abreviacion = 'EBEMP';

INSERT INTO movimientos_arka.formato_tipo_movimiento (
		nombre,
		formato,
		descripcion,
		codigo_abreviacion,
		activo,
		fecha_modificacion,
		fecha_creacion)
VALUES
		('Adiciones y Mejoras',
		'{"acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "vigencia_ordenador": "string"}',
		'Formato para guardar una entrada por adiciones y mejoras',
		'ENT_AM',
		true,
		now(),
		now());
