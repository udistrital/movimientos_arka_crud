ALTER TABLE movimientos_arka.estado_movimiento
    ALTER COLUMN nombre TYPE VARCHAR(40);

INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Ajuste En Trámite',
    'Comprobante de ajuste en trámite',
    true,
    now(),
    now()),
    ('Ajuste Rechazado',
    'Comprobante de ajuste rechazado por funcionario de contabilidad o almacén',
    true,
    now(),
    now()),
    ('Ajuste Aprobado por Contabilidad',
    'Comprobante de ajuste aprobado por contabilidad, en espera de aprobación por almacén',
    true,
    now(),
    now()),
    ('Ajuste Aprobado por Almacén',
    'Comprobante de ajuste aprobado por almacén, en espera de aprobación por contabilidad',
    true,
    now(),
    now()),
    ('Ajuste Aprobado',
    'Comprobante de ajuste aprobado por almacén y contabilidad',
    true,
    now(),
    now());

INSERT INTO movimientos_arka.formato_tipo_movimiento (
    nombre,
    formato,
    descripcion,
    codigo_abreviacion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Ajuste Contable',
    '{ "TrContable": "[]map[string]interface{}", "TrContableId": "int", "RazonRechazo": "string", "AprobacionAlmacen": "string", "AprobacionContabilidad": "string" }',
    'Comprobante de ajuste',
    'AJ_CBE',
    true,
    now(),
    now());
