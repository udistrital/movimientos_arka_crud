DELETE FROM movimientos_arka.estado_movimiento
    WHERE nombre IN ('Ajuste En Trámite', 'Ajuste Rechazado', 'Ajuste Aprobado por Contabilidad', 'Ajuste Aprobado por Almacén', 'Ajuste Aprobado');

ALTER TABLE movimientos_arka.estado_movimiento
    ALTER COLUMN nombre TYPE VARCHAR(20);

DELETE FROM movimientos_arka.formato_tipo_movimiento
    WHERE nombre = 'Ajuste Contable';
