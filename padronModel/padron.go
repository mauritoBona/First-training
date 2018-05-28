package padronModel

import (
	"strconv"
	"strings"
)

type Padron struct {
	PadronTypeId    int
	TaxId           int
	PublicationDate string
	FromDate        string
	ToDate          string
	Cuit            string
	TaxPayerTypeId  int
	BrandSubjectId  int
	BrandAlicuotaId int
	Alicuota        float64
}

func (padron *Padron) BuildPadronFromCSV(padronCsv []string) {
	for val := range padronCsv {
		padron.setValueToPadron(padronCsv[val], val)
	}
}

func BuildInsertInitializeQuery() string {
	return "INSERT INTO padron(PADRON_TYPE_ID,TAX_ID,PUBLICATION_DATE,FROM_DATE,TO_DATE,CUIT,TAX_PAYER_TYPE_ID,BRAND_SUBJECT_ID,BRAND_ALICUOT_ID,ALICUOTA) VALUES "
}

func (padron *Padron) BuildInsertValueQuery() string {
	var insert strings.Builder
	insert.WriteString("(")
	insert.WriteString(strconv.Itoa(padron.PadronTypeId) + ",")
	insert.WriteString(strconv.Itoa(padron.TaxId) + ",STR_TO_DATE('" + padron.PublicationDate + "', '%d%m%Y')")
	insert.WriteString(",STR_TO_DATE('" + padron.FromDate + "', '%d%m%Y'),STR_TO_DATE('" + padron.ToDate + "', '%d%m%Y'),")
	insert.WriteString(padron.Cuit + "," + strconv.Itoa(padron.TaxPayerTypeId) + ",")
	insert.WriteString(strconv.Itoa(padron.BrandSubjectId) + "," + strconv.Itoa(padron.BrandAlicuotaId) + ",")
	insert.WriteString(strconv.FormatFloat(padron.Alicuota, 'f', 2, 64) + ")")
	return insert.String()
}

func (padron *Padron) setValueToPadron(value string, positionValue int) {
	switch positionValue {
	case 0:
		padron.PadronTypeId = PadronTypeMap[value]
		padron.TaxId = TaxMap[value]
	case 1:
		padron.PublicationDate = value
	case 2:
		padron.FromDate = value
	case 3:
		padron.ToDate = value
	case 4:
		padron.Cuit = value
	case 5:
		padron.TaxPayerTypeId = TaxPayerTypeMap[value]
	case 6:
		padron.BrandSubjectId = BrandSubjectMap[value]
	case 7:
		padron.BrandAlicuotaId = AlicuotTypeMap[value]
	case 8:
		alicuot, err := strconv.ParseFloat(value, 64)
		if err == nil {
			padron.Alicuota = alicuot
		}
	default:
		break
	}
}
