// Package strand implements RNA complement (per RNA transcription).
package strand

var dna2rna = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA returns replacing each nucleotide with its complement.
func ToRNA(dna string) (rna string) {
	for i := 0; i < len(dna); i++ {
		rna += dna2rna[string(dna[i])]
	}
	return
}
