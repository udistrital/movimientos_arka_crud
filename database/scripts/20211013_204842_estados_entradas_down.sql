UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Entrada Aceptada',
    descripcion = 'Formato para marcar una entrada como aceptada/aprobada',
    fecha_modificacion = now()
WHERE nombre = 'Entrada En Trámite';

DELETE FROM movimientos_arka.estado_movimiento
WHERE nombre IN ('Salida Rechazada', 'Entrada Rechazada', 'Entrada Aprobada');
