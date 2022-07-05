UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Depr Generada',
    descripcion = 'Depreciación generada a una determinada fecha de corte',
    fecha_modificacion = now()
WHERE nombre = 'Cierre En Curso';

UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Depr Rechazada',
    descripcion = 'Depreciación rechazada por funcionario de contabilidad',
    fecha_modificacion = now()
WHERE nombre = 'Cierre Rechazado';

UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Depr Aprobada',
    descripcion = 'Depreciación aprobada por funcionario de contabilidad',
    fecha_modificacion = now()
WHERE nombre = 'Cierre Aprobado';

DELETE FROM movimientos_arka.formato_tipo_movimiento
    WHERE codigo_abreviacion = 'CRR';
