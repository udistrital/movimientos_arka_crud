INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre,
    formato,
    descripcion,
    codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)
VALUES
('Caja Menor',
'{ "acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "vigencia_ordenador": "string", "ordenador_gasto_id": "number", "solicitante_id": "number" }',
'Formato para guardar una Entrada por Caja Menor',
'ECM', 14.0, now(), now(), true),
('Compra en el Extranjero',
'{"acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "vigencia_ordenador": "string", "ordenador_gasto_id": "number", "solicitante_id": "number" }',
'Formato para guardar una Entrada por Compra en el Extranjero',
'ECE', 15.0, now(), now(), true),
('Nuevo Aprovechamiento',
'{"acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "vigencia_ordenador": "string", "ordenador_gasto_id": "number", "solicitante_id": "number" }',
'Formato para guardar una Entrada por Nuevo Aprovechamiento',
'ENA', 16.0, now(), now(), true);
