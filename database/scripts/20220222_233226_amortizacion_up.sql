INSERT INTO movimientos_arka.formato_tipo_movimiento (
		nombre,
		formato,
		descripcion,
		codigo_abreviacion,
		activo,
		fecha_modificacion,
		fecha_creacion)
VALUES
		('Amortizacion',
		'{ "Totales": "map[int]float64", "TrContable": "int", "FechaCorte": "string", "RazonRechazo": "string" }',
		'Liquidación de amortización a una fecha de corte determinada',
		'AMT',
		true,
		now(),
		now());
