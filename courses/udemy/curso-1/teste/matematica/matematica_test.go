package matematica

import "testing"

const defaulError = "Valor esperado %v, mas o retorno foi %v"

func TestMedia(t *testing.T) {
	valEsperado := 7.275
	val := Media(7.2, 9.9, 6.1, 5.9)

	if val != valEsperado {
		t.Errorf(defaulError, valEsperado, val)
	}

}
