UPDATE movimientos_arka.formato_tipo_movimiento SET
    nombre = 'Solicitud de Bajas',
    formato = '{ "Elementos": [], "Ubicacion": "number", "Consecutivo": "string", "FuncionarioOrigen": "number" , "FuncionarioDestino": "number" }',
    descripcion = 'Solicitud de baja por hurto',
    codigo_abreviacion = 'BJ_HT',
    fecha_modificacion = now()
WHERE nombre = 'Baja por Hurto';

DELETE FROM movimientos_arka.formato_tipo_movimiento
    WHERE nombre = 'Baja por Daño';

DELETE FROM movimientos_arka.estado_movimiento
    WHERE nombre IN ('Baja En Trámite','Baja Rechazada','Baja En Comité','Baja Aprobada');

