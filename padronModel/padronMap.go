package padronModel

var TaxPayerTypeMap = make(map[string]int)
var BrandSubjectMap = make(map[string]int)
var AlicuotTypeMap = make(map[string]int)
var PadronTypeMap = make(map[string]int)
var TaxMap = make(map[string]int)

func CompleteMaps() {
	if len(TaxPayerTypeMap) < 1 {
		completeTaxPayerMap()
	}
	if len(BrandSubjectMap) < 1 {
		completeBrandSubjectMap()
	}
	if len(AlicuotTypeMap) < 1 {
		completeAlicuotTypeMap()
	}
	if len(PadronTypeMap) < 1 {
		completePadronTypeMap()
	}
	if len(TaxMap) < 1 {
		completeTaxMap()
	}
}

func completeTaxPayerMap() {
	TaxPayerTypeMap["C"] = 1
	TaxPayerTypeMap["D"] = 2
}

func completeBrandSubjectMap() {
	BrandSubjectMap["S"] = 1
	BrandSubjectMap["N"] = 2
	BrandSubjectMap["B"] = 3
}

func completeAlicuotTypeMap() {
	AlicuotTypeMap["S"] = 1
	AlicuotTypeMap["N"] = 2
}

func completePadronTypeMap() {
	PadronTypeMap["R"] = 1
	PadronTypeMap["P"] = 1
}

func completeTaxMap() {
	TaxMap["R"] = 14
	TaxMap["P"] = 15
}
