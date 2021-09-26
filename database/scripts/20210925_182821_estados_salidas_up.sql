UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Salida En Trámite',
    descripcion = 'Salida en estado de trámite',
    fecha_modificacion = now()
WHERE id = 3;

INSERT INTO movimientos_arka.estado_movimiento (
    id,
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES (
    9,
    'Salida Aprobada',
    'Salida en estado aprobada',
    true,
    now(),
    now());
