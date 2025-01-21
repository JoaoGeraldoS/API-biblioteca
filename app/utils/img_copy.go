package utils

import (
	"fmt"
	"io"
	"os"
)

func CopyImg(caminhoOrigem, caminhoDestino string) error {
	arquivoOrigem, err := os.Open(caminhoOrigem)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo de origem: %v", err)
	}
	defer arquivoOrigem.Close()

	arquivoDestino, err := os.Create(caminhoDestino)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de destino: %v", err)
	}
	defer arquivoDestino.Close()

	_, err = io.Copy(arquivoDestino, arquivoOrigem)
	if err != nil {
		return fmt.Errorf("erro ao copiar arquivo: %v", err)
	}

	return nil
}
