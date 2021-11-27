UPDATE movimientos_arka.formato_tipo_movimiento SET
    nombre = 'Baja por Hurto',
    formato = '{ "Elementos": [], "Ubicacion": "number", "Consecutivo": "string", "FuncionarioOrigen": "number" , "FuncionarioDestino": "number" }',
    descripcion = 'Solicitud de baja por hurto',
    codigo_abreviacion = 'BJ_HT',
    fecha_modificacion = now()
WHERE nombre = 'Solicitud de Bajas';

INSERT INTO movimientos_arka.formato_tipo_movimiento (
    nombre,
    formato,
    descripcion,
    codigo_abreviacion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Baja por Daño',
    '{ "Elementos": [], "Consecutivo": "string", "FuncionarioSolicitante": "number" }',
    'Solicitud de baño porque no funciona',
    'BJ_DÑ',
    true,
    now(),
    now());

INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Baja En Trámite',
    'Solicitud de baja registrada',
    true,
    now(),
    now()),
    ('Baja Rechazada',
    'Solicitud de baja rechazada por oficina de almacén o comité de inventarios',
    true,
    now(),
    now()),
    ('Baja En Comité',
    'Baja aprobada por almacén y en espera de su revisión por el comité de inventario',
    true,
    now(),
    now()),
    ('Baja Aprobada',
    'Traslado confirmado por funcionario destino',
    true,
    now(),
    now());
