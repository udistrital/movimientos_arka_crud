UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Salida Aceptada',
    descripcion = 'Formato para marcar una salida como aceptada/aprobada',
    fecha_modificacion = now()
WHERE id = 3;

DELETE FROM movimientos_arka.estado_movimiento WHERE id = 9;
