UPDATE movimientos_arka.formato_tipo_movimiento
SET formato =
'{"acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "vigencia": "number", "supervisor": "string", "proveedor": "string"}'
WHERE codigo_abreviacion = 'EPPA';
