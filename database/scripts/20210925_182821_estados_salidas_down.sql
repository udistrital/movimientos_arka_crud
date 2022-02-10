UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Salida Aceptada',
    descripcion = 'Formato para marcar una salida como aceptada/aprobada',
    fecha_modificacion = now()
WHERE nombre = 'Salida En Trámite';

DELETE FROM movimientos_arka.estado_movimiento WHERE nombre = 'Salida Aprobada';
