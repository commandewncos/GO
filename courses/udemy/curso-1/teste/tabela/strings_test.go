package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - ídices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto, parte string
		esperado     int
	}{
		{texto: "Wilson Nascimento Costa",
			parte: "Wilson", esperado: 0},

		{texto: "Opa", parte: "opa", esperado: -1},
		{texto: "Epa", parte: "Epa", esperado: 0},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}

	}

}
