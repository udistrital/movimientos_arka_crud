UPDATE movimientos_arka.formato_tipo_movimiento SET
    nombre = 'Traslado',
    descripcion = 'Traslado de bien',
    formato = '{ "Elementos": [], "Ubicacion": "number", "Consecutivo": "string", "FuncionarioOrigen": "number" , "FuncionarioDestino": "number" }',
    fecha_modificacion = now()
WHERE nombre = 'Solicitud de Traslado';

INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Traslado Rechazado',
    'Traslado rechazado por almacen',
    true,
    now(),
    now()),
    ('Traslado Aprobado',
    'Traslado aprobado por almacén',
    true,
    now(),
    now()),
    ('Traslado En Trámite',
    'Traslado solicitado',
    true,
    now(),
    now()),
    ('Traslado Confirmado',
    'Traslado confirmado por funcionario destino',
    true,
    now(),
    now()),
    ('Registro Kardex',
    'Movimiento de bodega de consumo',
    true,
    now(),
    now());

ALTER TABLE movimientos_arka.movimiento
    ALTER COLUMN detalle TYPE jsonb;
