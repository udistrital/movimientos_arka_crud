INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre,
    formato,
    descripcion,
    codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)
VALUES
('Intangibles adquiridos',
'{ "acta_recibido_id": "number", "consecutivo": "string",  "documento_contable_id": "number", "contrato_id": "number", "vigencia_contrato": "number", "importacion": "boolean", "amortizacion": "number", "vida_util": "number"}',
'Formato para guardar una Entrada por Intangibles adquiridos',
'EIA', 18.0, now(), now(), true),
('Intangibles desarrollados',
'{"acta_recibido_id": "number", "consecutivo":"string",  "documento_contable_id":"number","contrato_id":"number", "vigencia_contrato":"number","importacion":"boolean","amortizacion":"number","vida_util":"number"}',
'Formato para guardar una Entrada de Intangibles Desarrollados',
'EID', 19.0, now(), now(), true),
('Provisional',
'{"acta_recibido_id": "number", "consecutivo":"string",  "documento_contable_id":"number","contrato_id":"number", "vigencia_contrato":"number","importacion":"boolean","amortizacion":"number","vida_util":"number"}',
'Formato para guardar una Entrada de Bienes Entregados de Manera Provisional',
'EBEMP', 20.0, now(), now(), true);
