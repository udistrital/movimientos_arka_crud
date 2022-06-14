UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Traslado Por Confirmar',
    descripcion = 'Solicitud de traslado en espera de ser confirmada por el funcionario destino',
    fecha_modificacion = now()
WHERE nombre = 'Traslado En Trámite';

INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Traslado Anulado',
    'Traslado anulado por error o no aprobación de almacén',
    true,
    now(),
    now());

ALTER TABLE movimientos_arka.estado_movimiento
    ADD CONSTRAINT uq_nombre_estado_movimiento UNIQUE (nombre);
