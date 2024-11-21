package matrixtools

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {
	//if i < 0 || i >= len(m) {
	//return []float64{}
	//}
	return m[i]
}

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	col := make([]float64, len(m))
	for i := range m {
		if j < 0 || j >= len(m) {
			return m[i]
		}
		col[i] = m[i][j]
	}
	return col
}

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {
	if i < 0 || i >= len(m) || j < 0 || j >= len(m) {
		return
	}
	for k := range m[i] {
		m[i][k] += m[j][k]
	}
}

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {
	for k := range m[i] {
		m[i][k] = m[i][k] * factor
	}
}

// Transpose erwartet eine Matrix und liefert ihre Transponierte.
// D.h. alle Zeilen der ersten Matrix werden zu Spalten der Transponierten und umgekehrt.
func Transposed(m [][]float64) [][]float64 {
	transposed := make([][]float64, len(m[0]))

	return transposed
}
