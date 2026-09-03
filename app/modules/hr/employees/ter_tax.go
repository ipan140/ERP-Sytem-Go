package employees

import (
	"math"
	"strings"
)

// TERCategory returns "A", "B", or "C" based on Indonesian PTKP Status
// Sesuai PP 58 / 2023 & PMK 168 / 2023:
// Kategori A: TK/0, TK/1, K/0
// Kategori B: TK/2, TK/3, K/1, K/2
// Kategori C: K/3
func GetTERCategory(ptkp string) string {
	cleaned := strings.ToUpper(strings.TrimSpace(ptkp))
	switch cleaned {
	case "TK/0", "TK/1", "K/0":
		return "A"
	case "TK/2", "TK/3", "K/1", "K/2":
		return "B"
	case "K/3":
		return "C"
	default:
		// Default aman bagi karyawan yang belum isi PTKP adalah TK/0 (Kategori A)
		return "A"
	}
}

// GetTERRate mengembalikan tarif desimal (contoh: 0.015 untuk 1.5%) berdasarkan Bruto Bulanan dan Kategori TER
func GetTERRate(category string, bruto float64) float64 {
	if bruto <= 0 {
		return 0
	}

	switch category {
	case "A":
		switch {
		case bruto <= 5400000:
			return 0.0
		case bruto <= 5650000:
			return 0.0025
		case bruto <= 5950000:
			return 0.005
		case bruto <= 6300000:
			return 0.0075
		case bruto <= 6750000:
			return 0.01
		case bruto <= 7500000:
			return 0.0125
		case bruto <= 8550000:
			return 0.015
		case bruto <= 9650000:
			return 0.0175
		case bruto <= 10050000:
			return 0.02
		case bruto <= 10350000:
			return 0.0225
		case bruto <= 10700000:
			return 0.025
		case bruto <= 11050000:
			return 0.03
		case bruto <= 11600000:
			return 0.035
		case bruto <= 12500000:
			return 0.04
		case bruto <= 13750000:
			return 0.05
		case bruto <= 15100000:
			return 0.06
		case bruto <= 16950000:
			return 0.07
		case bruto <= 19750000:
			return 0.08
		case bruto <= 24150000:
			return 0.09
		case bruto <= 26450000:
			return 0.10
		case bruto <= 28000000:
			return 0.11
		case bruto <= 30050000:
			return 0.12
		case bruto <= 32400000:
			return 0.13
		case bruto <= 35400000:
			return 0.14
		case bruto <= 39100000:
			return 0.15
		case bruto <= 43850000:
			return 0.16
		case bruto <= 47800000:
			return 0.17
		case bruto <= 51400000:
			return 0.18
		case bruto <= 56300000:
			return 0.19
		case bruto <= 62200000:
			return 0.20
		case bruto <= 68600000:
			return 0.21
		case bruto <= 77500000:
			return 0.22
		case bruto <= 89000000:
			return 0.23
		case bruto <= 101900000:
			return 0.24
		case bruto <= 114000000:
			return 0.25
		case bruto <= 126600000:
			return 0.26
		case bruto <= 140500000:
			return 0.27
		case bruto <= 156000000:
			return 0.28
		case bruto <= 173200000:
			return 0.29
		case bruto <= 192300000:
			return 0.30
		case bruto <= 213500000:
			return 0.31
		case bruto <= 237000000:
			return 0.32
		case bruto <= 263000000:
			return 0.33
		default:
			return 0.34
		}

	case "B":
		switch {
		case bruto <= 6200000:
			return 0.0
		case bruto <= 6500000:
			return 0.0025
		case bruto <= 6850000:
			return 0.005
		case bruto <= 7300000:
			return 0.0075
		case bruto <= 9200000:
			return 0.01
		case bruto <= 10750000:
			return 0.015
		case bruto <= 12550000:
			return 0.02
		case bruto <= 13050000:
			return 0.03
		case bruto <= 14370000:
			return 0.04
		case bruto <= 15400000:
			return 0.05
		case bruto <= 16800000:
			return 0.06
		case bruto <= 18100000:
			return 0.07
		case bruto <= 19800000:
			return 0.08
		case bruto <= 24150000:
			return 0.09
		case bruto <= 26450000:
			return 0.10
		case bruto <= 28000000:
			return 0.11
		case bruto <= 30050000:
			return 0.12
		case bruto <= 32400000:
			return 0.13
		case bruto <= 35400000:
			return 0.14
		case bruto <= 39100000:
			return 0.15
		case bruto <= 43850000:
			return 0.16
		case bruto <= 47800000:
			return 0.17
		case bruto <= 51400000:
			return 0.18
		case bruto <= 56300000:
			return 0.19
		case bruto <= 62200000:
			return 0.20
		case bruto <= 68600000:
			return 0.21
		case bruto <= 77500000:
			return 0.22
		case bruto <= 89000000:
			return 0.23
		case bruto <= 101900000:
			return 0.24
		case bruto <= 114000000:
			return 0.25
		case bruto <= 126600000:
			return 0.26
		case bruto <= 140500000:
			return 0.27
		case bruto <= 156000000:
			return 0.28
		case bruto <= 173200000:
			return 0.29
		case bruto <= 192300000:
			return 0.30
		case bruto <= 213500000:
			return 0.31
		case bruto <= 237000000:
			return 0.32
		case bruto <= 263000000:
			return 0.33
		default:
			return 0.34
		}

	case "C":
		switch {
		case bruto <= 6600000:
			return 0.0
		case bruto <= 6950000:
			return 0.0025
		case bruto <= 7350000:
			return 0.005
		case bruto <= 7800000:
			return 0.0075
		case bruto <= 8850000:
			return 0.01
		case bruto <= 9800000:
			return 0.0125
		case bruto <= 10950000:
			return 0.015
		case bruto <= 11200000:
			return 0.0175
		case bruto <= 12050000:
			return 0.02
		case bruto <= 12950000:
			return 0.03
		case bruto <= 14150000:
			return 0.04
		case bruto <= 15150000:
			return 0.05
		case bruto <= 16300000:
			return 0.06
		case bruto <= 17450000:
			return 0.07
		case bruto <= 18700000:
			return 0.08
		case bruto <= 24150000:
			return 0.09
		case bruto <= 26450000:
			return 0.10
		case bruto <= 28000000:
			return 0.11
		case bruto <= 30050000:
			return 0.12
		case bruto <= 32400000:
			return 0.13
		case bruto <= 35400000:
			return 0.14
		case bruto <= 39100000:
			return 0.15
		case bruto <= 43850000:
			return 0.16
		case bruto <= 47800000:
			return 0.17
		case bruto <= 51400000:
			return 0.18
		case bruto <= 56300000:
			return 0.19
		case bruto <= 62200000:
			return 0.20
		case bruto <= 68600000:
			return 0.21
		case bruto <= 77500000:
			return 0.22
		case bruto <= 89000000:
			return 0.23
		case bruto <= 101900000:
			return 0.24
		case bruto <= 114000000:
			return 0.25
		case bruto <= 126600000:
			return 0.26
		case bruto <= 140500000:
			return 0.27
		case bruto <= 156000000:
			return 0.28
		case bruto <= 173200000:
			return 0.29
		case bruto <= 192300000:
			return 0.30
		case bruto <= 213500000:
			return 0.31
		case bruto <= 237000000:
			return 0.32
		case bruto <= 263000000:
			return 0.33
		default:
			return 0.34
		}
	}

	return 0.05 // default fallback
}

// CalculatePPh21TER menghitung nilai nominal potongan PPh 21 TER bulanan
func CalculatePPh21TER(bruto float64, ptkp string) (float64, string, float64) {
	category := GetTERCategory(ptkp)
	rate := GetTERRate(category, bruto)
	nominal := math.Round(bruto * rate)
	return nominal, category, rate
}
