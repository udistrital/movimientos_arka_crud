INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Bienes inmuebles registrados',
    'Adición de inmuebles al inventario',
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
    ('Registro bien inmueble',
    '{ "ConsevcutivoId": "int" }',
    'Adición de bienes inmuebles al inventario',
    'INM_REG',
    true,
    now(),
    now());
