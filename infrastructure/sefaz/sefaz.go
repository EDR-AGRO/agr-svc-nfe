package sefaz

import "fmt"

type StateUrls struct {
	NfeAuth string
}

var stateUrls = map[string]StateUrls{
	"AM":     {NfeAuth: "https://nfe.sefaz.am.gov.br/services2/services/NfeAutorizacao4"},
	"BA":     {NfeAuth: "https://nfe.sefaz.ba.gov.br/webservices/NFeAutorizacao4/NFeAutorizacao4.asmx"},
	"GO":     {NfeAuth: "https://nfe.sefaz.go.gov.br/nfe/services/NFeAutorizacao4"},
	"MG":     {NfeAuth: "https://nfe.fazenda.mg.gov.br/nfe2/services/NFeAutorizacao4"},
	"MS":     {NfeAuth: "https://nfe.sefaz.ms.gov.br/ws/NFeAutorizacao4"},
	"MT":     {NfeAuth: "https://nfe.sefaz.mt.gov.br/nfews/v2/services/NfeAutorizacao4"},
	"PE":     {NfeAuth: "https://nfe.sefaz.pe.gov.br/nfe-service/services/NFeAutorizacao4"},
	"PR":     {NfeAuth: "https://nfe.sefa.pr.gov.br/nfe/NFeAutorizacao4"},
	"RS":     {NfeAuth: "https://nfe.sefazrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx"},
	"SP":     {NfeAuth: "https://nfe.fazenda.sp.gov.br/ws/nfeautorizacao4.asmx"},
	"SVAN":   {NfeAuth: "https://www.sefazvirtual.fazenda.gov.br/NFeAutorizacao4/NFeAutorizacao4.asmx"},
	"SVRS":   {NfeAuth: "https://nfe.svrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx"},
	"SVC-AN": {NfeAuth: "https://www.sefazvirtual.fazenda.gov.br/NFeAutorizacao4/NFeAutorizacao4.asmx"},
	"SVC-RS": {NfeAuth: "https://nfe.svrs.rs.gov.br/ws/NfeAutorizacao/NFeAutorizacao4.asmx"},
}

var dependentVirtualService = map[string]string{
	"AC": "SVRS",
	"AL": "SVRS",
	"AP": "SVRS",
	"CE": "SVRS",
	"DF": "SVRS",
	"ES": "SVRS",
	"PA": "SVRS",
	"PB": "SVRS",
	"PI": "SVRS",
	"RJ": "SVRS",
	"RN": "SVRS",
	"RO": "SVRS",
	"RR": "SVRS",
	"SC": "SVRS",
	"SE": "SVRS",
	"TO": "SVRS",
	"MA": "SVAN",
}

var contingencyRelation = map[string]string{
	"AC": "SVC-AN",
	"AL": "SVC-AN",
	"AP": "SVC-AN",
	"CE": "SVC-AN",
	"DF": "SVC-AN",
	"ES": "SVC-AN",
	"MG": "SVC-AN",
	"PA": "SVC-AN",
	"PB": "SVC-AN",
	"PI": "SVC-AN",
	"RJ": "SVC-AN",
	"RN": "SVC-AN",
	"RO": "SVC-AN",
	"RR": "SVC-AN",
	"RS": "SVC-AN",
	"SC": "SVC-AN",
	"SE": "SVC-AN",
	"SP": "SVC-AN",
	"TO": "SVC-AN",
	"AM": "SVC-RS",
	"BA": "SVC-RS",
	"GO": "SVC-RS",
	"MA": "SVC-RS",
	"MS": "SVC-RS",
	"MT": "SVC-RS",
	"PE": "SVC-RS",
	"PR": "SVC-RS",
}

var stateNumber = map[string]string{
	"RO": "11",
	"AC": "12",
	"AM": "13",
	"RR": "14",
	"PA": "15",
	"MA": "21",
	"PI": "22",
	"CE": "23",
	"RN": "24",
	"PB": "25",
	"PE": "26",
	"AL": "27",
	"SE": "28",
	"BA": "29",
	"MG": "31",
	"ES": "32",
	"RJ": "33",
	"SP": "35",
	"PR": "41",
	"SC": "42",
	"RS": "43",
	"MS": "50",
	"MT": "51",
	"GO": "52",
	"DF": "53",
}

func buildSefazNfe(xml EnviNFe) *EnvelopeNfeAuth {

	envelope := new(EnvelopeNfeAuth)

	/*set envelope*/
	envelope.Soap12 = "http://www.w3.org/2003/05/soap-envelope"
	envelope.Nfe = "http://www.portalfiscal.inf.br/nfe/wsdl/NFeAutorizacao4"

	/*set header*/
	envelope.Header.NfeCabecMsg.CUF = xml.Nfe.InfNFe.Ide.CUF
	envelope.Header.NfeCabecMsg.VersaoDados = "4.00"

	/*set builder*/
	envelope.Body.NfeDadosMsg = fmt.Sprintf(`<![CDATA[%s]]>`, xml)

	return envelope
}
