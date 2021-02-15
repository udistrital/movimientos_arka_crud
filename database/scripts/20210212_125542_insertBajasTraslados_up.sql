INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre,
    formato,
    descripcion,
    codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)
VALUES
('Solicitud de Traslado',
'{ "Elementos": [ { "Funcionario": "number", "Ubicacion": "number", "ElementoActa": "number", "Cantidad": "number" } ] }',
'Formato para guardar una solicitud de Traslado',
'SOL_TRD', 21.0, now(), now(), true);
