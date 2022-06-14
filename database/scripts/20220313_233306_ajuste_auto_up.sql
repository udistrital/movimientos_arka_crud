INSERT INTO movimientos_arka.formato_tipo_movimiento (
		nombre,
		formato,
		descripcion,
		codigo_abreviacion,
		activo,
		fecha_modificacion,
		fecha_creacion)
VALUES
		('Ajuste Automático',
		'{ "Elementos": "[]int", "TrContable": "int", "Consecutivo": "string" }',
		'Ajuste de inventario y contable calculado por el sistema',
		'AAT',
		true,
		now(),
		now());

ALTER TABLE movimientos_arka.formato_tipo_movimiento
	ADD CONSTRAINT uq_codigo_abreviacion UNIQUE (codigo_abreviacion);
