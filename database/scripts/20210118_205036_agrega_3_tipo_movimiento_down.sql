IF EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECM')
    DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECM';
IF EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECE')
    DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ECE';
IF EXISTS (SELECT 1 FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ENA')
    DELETE FROM movimientos_arka.formato_tipo_movimiento WHERE codigo_abreviacion = 'ENA';
