package models

import (
	"strings"
	"testing"
)

func TestResolverAccionElementoSalidaSinRegistroPrevio(t *testing.T) {
	accion, err := resolverAccionElementoSalida(nil)
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if accion != insertarElementoSalida {
		t.Fatalf("esperaba accion insertar, obtuve %v", accion)
	}
}

func TestResolverAccionElementoSalidaReutilizaSalidaRechazada(t *testing.T) {
	elementoActaID := 41801
	accion, err := resolverAccionElementoSalida(&ElementosMovimiento{
		ElementoActaId: &elementoActaID,
		MovimientoId: &Movimiento{
			Id: 123,
			EstadoMovimientoId: &EstadoMovimiento{
				Nombre: "Salida Rechazada",
			},
		},
	})
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if accion != reutilizarElementoSalida {
		t.Fatalf("esperaba accion reutilizar, obtuve %v", accion)
	}
}

func TestEstadoSalidaPermiteReusoSalidaAnulada(t *testing.T) {
	if !estadoSalidaPermiteReuso("Salida anulada") {
		t.Fatal("esperaba que una salida anulada permitiera reutilizar el elemento")
	}
}

func TestResolverAccionElementoSalidaRechazaSalidaVigente(t *testing.T) {
	elementoActaID := 41801
	accion, err := resolverAccionElementoSalida(&ElementosMovimiento{
		ElementoActaId: &elementoActaID,
		MovimientoId: &Movimiento{
			Id: 8079,
			EstadoMovimientoId: &EstadoMovimiento{
				Nombre: "Salida Aprobada",
			},
		},
	})
	if err == nil {
		t.Fatal("esperaba error de negocio por salida vigente")
	}
	if accion != insertarElementoSalida {
		t.Fatalf("esperaba accion por defecto insertar, obtuve %v", accion)
	}
	if !strings.Contains(err.Error(), "41801") || !strings.Contains(err.Error(), "Salida Aprobada") {
		t.Fatalf("mensaje inesperado: %v", err)
	}
}
