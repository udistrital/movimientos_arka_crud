-- Lo siguiente debería ser una forma segura de resetear el serial
-- (Referencia: https://stackoverflow.com/a/244265/3180052)
BEGIN;
LOCK TABLE movimientos_arka.formato_tipo_movimiento IN EXCLUSIVE MODE;
SELECT setval(
    pg_get_serial_sequence('movimientos_arka.formato_tipo_movimiento','id'),
    COALESCE((SELECT MAX(id)+1 FROM movimientos_arka.formato_tipo_movimiento), 14),
    false);
COMMIT;
-- Lo anterior equivale a (también funciona pero no es tan seguro):
-- ALTER SEQUENCE movimientos_arka.formato_tipo_movimiento_id_seq RESTART WITH 14
-- Otras formas de alterar secuencias:
-- https://stackoverflow.com/questions/8745051/postgres-manually-alter-sequence

INSERT INTO movimientos_arka.formato_tipo_movimiento (nombre,
    formato,
    descripcion,
    codigo_abreviacion, numero_orden, fecha_creacion, fecha_modificacion, activo)
VALUES
('Caja Menor',
'{ "acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "vigencia_ordenador": "string", "ordenador_gasto_id": "number", "solicitante_id": "number" }',
'Formato para guardar una Entrada por Caja Menor',
'ECM', 14.0, now(), now(), true),
-- (HU-8) Se toma como referencia la plantilla de la entrada tipo "Adquisicion", y se extiende, para el siguiente:
('Compra en el Extranjero',
'{"acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "contrato_id": "number", "vigencia_contrato": "number","importacion": "boolean", "num_reg_importacion": "string", "TRM": "number"}',
'Formato para guardar una Entrada por Compra en el Extranjero',
'ECE', 15.0, now(), now(), true),
-- (HU-9) Se toma como referencia la plantilla de la entrada tipo "Adquisicion", y se recorta, para el siguiente:
('Partes por Aprovechamientos',
'{"acta_recibido_id": "number", "consecutivo": "string", "documento_contable_id": "number", "vigencia_contrato": "number"}',
'Formato para guardar una Entrada por Nuevo Aprovechamiento',
'EPPA', 16.0, now(), now(), true);
