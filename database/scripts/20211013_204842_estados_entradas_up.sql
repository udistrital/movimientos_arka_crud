-- PARTE 1 - Resetear secuencia dañada por registros con IDs quemados
-- Lo siguiente debería ser una forma segura de resetear el serial
-- y permitir que de ahora en más no se admita quemar los IDs
-- (o por lo menos para movimientos.tipo_movimiento)
-- (Referencia: https://stackoverflow.com/a/244265/3180052)
-- (Otra ref: https://hcmc.uvic.ca/blogs/index.php/how_to_fix_postgresql_error_duplicate_ke?blog=22)
-- Equivale a (también funciona pero no es tan seguro):
-- ALTER SEQUENCE movimientos.tipo_movimiento RESTART WITH 14
-- Otras formas de alterar secuencias:
-- https://stackoverflow.com/questions/8745051/postgres-manually-alter-sequence

BEGIN;
LOCK TABLE movimientos_arka.estado_movimiento IN EXCLUSIVE MODE;
SELECT setval(
    'movimientos_arka.estado_movimiento_id_seq',
    COALESCE((SELECT MAX(id)+1 FROM movimientos_arka.estado_movimiento), 15),
    false);
COMMIT;

UPDATE movimientos_arka.estado_movimiento SET
    nombre = 'Entrada En Trámite',
    descripcion = 'Entrada en estado de trámite',
    fecha_modificacion = now()
WHERE nombre = 'Entrada Aceptada';

INSERT INTO movimientos_arka.estado_movimiento (
    nombre,
    descripcion,
    activo,
    fecha_modificacion,
    fecha_creacion)
VALUES
    ('Salida Rechazada',
    'Salida en estado rechazada',
    true,
    now(),
    now()),
    ('Entrada Rechazada',
    'Entrada en estado rechazada',
    true,
    now(),
    now()),
    ('Entrada Aprobada',
    'Entrada en estado aprobada',
    true,
    now(),
    now());
