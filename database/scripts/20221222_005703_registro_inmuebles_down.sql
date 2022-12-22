DELETE FROM movimientos_arka.estado_movimiento 
WHERE nombre IN ('Bienes inmuebles registrados');

DELETE FROM movimientos_arka.formato_tipo_movimiento 
WHERE codigo_abreviacion IN ('INM_REG');
