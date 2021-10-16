UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Entrada En Trámite',
    descripcion = 'Entrada en estado de trámite',
    fecha_modificacion = now()
WHERE nombre = 'Entrada Aceptada';

INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Salida Rechazada',
    'Salida en estado rechazada',
    true,
    now(),
    now()),
    ('Entrada Rechazada',
    'Entrada en estado rechazada',
    true,
    now(),
    now()),
    ('Entrada Aprobada',
    'Entrada en estado aprobada',
    true,
    now(),
    now());
