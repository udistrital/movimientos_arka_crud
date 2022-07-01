UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Cierre En Curso',
    descripcion = 'Cierre de sistema en espera de revisión',
    fecha_modificacion = now()
WHERE nombre = 'Depr Generada';

UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Cierre Rechazado',
    descripcion = 'Cierre de sistema rechazado',
    fecha_modificacion = now()
WHERE nombre = 'Depr Rechazada';

UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Cierre Aprobado',
    descripcion = 'Cierre de sistema aprobado',
    fecha_modificacion = now()
WHERE nombre = 'Depr Aprobada';

INSERT INTO movimientos_arka.formato_tipo_movimiento (
		nombre,
		formato,
		descripcion,
		codigo_abreviacion,
		activo,
		fecha_modificacion,
		fecha_creacion)
VALUES
		('Cierre',
		'{ "ConsecutivoId": "int", "Consecutivo": "string", "FechaCorte": "string", "RazonRechazo": "string" }',
		'Cierre del sistema a una fecha de corte determinada',
		'CRR',
		true,
		now(),
		now());
