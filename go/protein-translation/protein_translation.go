package protein

import (
	"errors"
)

var Condon = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	ErrStop        = errors.New("Stop it")
	ErrInvalidBase = errors.New("No such condon exist")
)

func FromCodon(input string) (string, error) {
	if found := Condon[input]; found == "" {
		return "", ErrInvalidBase
	} else if found == "STOP" {
		return "", ErrStop
	} else {
		return found, nil
	}
}

func FromRNA(input string) ([]string, error) {
	var result []string
	for i := 0; i < len(input); i += 3 {
		codon := input[i : i+3]
		if protein, err := FromCodon(codon); err != nil {
			if err == ErrStop {
				return result, nil
			}
			return result, ErrInvalidBase
		} else {
			result = append(result, protein)
		}
	}
	return result, nil
}
