// Package protein implements translate RNA sequences into proteins.
package protein

// ErrStop needs for returning stop error.
var ErrStop error

// ErrInvalidBase needs for returning Invalid Base error.
var ErrInvalidBase error

// Codon2Protein maps codons to proteins.
var Codon2Protein = map[string]string{
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

// FromCodon converts codons to proteins.
func FromCodon(codon string) (string, error) {
	if protein, ok := Codon2Protein[codon]; ok && protein != "STOP" {
		return protein, nil
	}
	return "", ErrStop
}

// FromRNA translates proteins.
func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna)-2; i += 3 {
		s := rna[i : i+3]
		if protein, ok := Codon2Protein[s]; ok && protein != "STOP" {
			proteins = append(proteins, protein)
		} else {
			return proteins, ErrInvalidBase
		}
	}
	return proteins, nil
}
