UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Traslado En Trámite',
    descripcion = 'Traslado solicitado',
    fecha_modificacion = now()
WHERE nombre = 'Traslado Por Confirmar';

DELETE FROM movimientos_arka.estado_movimiento 
WHERE nombre = 'Traslado Anulado';

ALTER TABLE movimientos_arka.estado_movimiento
    DROP CONSTRAINT IF EXISTS uq_nombre_estado_movimiento;
