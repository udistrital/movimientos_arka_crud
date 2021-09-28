UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Salida En Trámite',
    descripcion = 'Salida en estado de trámite',
    fecha_modificacion = now()
WHERE nombre = 'Salida Aceptada';

INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES (
    'Salida Aprobada',
    'Salida en estado aprobada',
    true,
    now(),
    now());
