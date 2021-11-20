UPDATE movimientos_arka.formato_tipo_movimiento SET
    nombre = 'Solicitud de Traslado',
    descripcion = 'Formato para guardar una solicitud de Traslado',
    formato = '{ "Elementos": [ { "Funcionario": "number", "Ubicacion": "number", "ElementoActa": "number", "Cantidad": "number" } ] }',
    fecha_modificacion = now()
WHERE nombre = 'Traslado';

DELETE FROM movimientos_arka.estado_movimiento 
WHERE nombre IN ('Traslado Rechazado','Traslado Aceptado','Traslado En Trámite','Traslado Confirmado','Registro Kardex');

ALTER TABLE movimientos_arka.movimiento
    ALTER COLUMN detalle TYPE json;
