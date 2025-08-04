package sefaz

import "encoding/xml"

type EnvelopeNfeAuth struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	Text    string   `xml:",chardata"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Nfe     string   `xml:"xmlns:nfe,attr"`
	Header  struct {
		Text        string `xml:",chardata"`
		NfeCabecMsg struct {
			Text        string `xml:",chardata"`
			CUF         string `xml:"cUF"`
			VersaoDados string `xml:"versaoDados"`
		} `xml:"nfe:nfeCabecMsg"`
	} `xml:"soap12:Header"`
	Body struct {
		Text        string `xml:",chardata"`
		NfeDadosMsg string `xml:"nfe:nfeDadosMsg"`
	} `xml:"soap12:Body"`
}

type EnviNFe struct {
	XMLName xml.Name `xml:"enviNFe"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Versao  string   `xml:"versao,attr"`
	IdLote  string   `xml:"idLote"`
	IndSinc string   `xml:"indSinc"`
	Nfe     Nfe      `xml:"NFe"`
}

type Nfe struct {
	Text   string `xml:",chardata"`
	InfNFe InfNFe `xml:"infNFe"`
}

type InfNFe struct {
	Text    string  `xml:",chardata"`
	Versao  string  `xml:"versao,attr"`
	ID      string  `xml:"Id,attr"`
	Ide     Ide     `xml:"ide"`
	Emit    Emit    `xml:"emit"`
	Dest    Dest    `xml:"dest"`
	Det     []Det   `xml:"det"`
	Total   Total   `xml:"total"`
	Transp  Transp  `xml:"transp"`
	Pag     Pag     `xml:"pag"`
	InfAdic InfAdic `xml:"infAdic"`
}

type Ide struct {
	Text     string `xml:",chardata"`
	CUF      string `xml:"cUF"`
	CNF      string `xml:"cNF"`
	NatOp    string `xml:"natOp"`
	Mod      string `xml:"mod"`
	Serie    string `xml:"serie"`
	NNF      string `xml:"nNF"`
	DhEmi    string `xml:"dhEmi"`
	TpNF     string `xml:"tpNF"`
	IdDest   string `xml:"idDest"`
	CMunFG   string `xml:"cMunFG"`
	TpImp    string `xml:"tpImp"`
	TpEmis   string `xml:"tpEmis"`
	CDV      string `xml:"cDV"`
	TpAmb    string `xml:"tpAmb"`
	FinNFe   string `xml:"finNFe"`
	IndFinal string `xml:"indFinal"`
	IndPres  string `xml:"indPres"`
	ProcEmi  string `xml:"procEmi"`
	VerProc  string `xml:"verProc"`
}

type Emit struct {
	Text      string  `xml:",chardata"`
	CNPJ      string  `xml:"CNPJ"`
	XNome     string  `xml:"xNome"`
	EnderEmit Address `xml:"enderEmit"`
	IE        string  `xml:"IE"`
	CRT       string  `xml:"CRT"`
}

type Dest struct {
	Text      string  `xml:",chardata"`
	CPF       string  `xml:"CPF"`
	XNome     string  `xml:"xNome"`
	EnderDest Address `xml:"enderDest"`
	IndIEDest string  `xml:"indIEDest"`
}

type Address struct {
	Text    string `xml:",chardata"`
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais"`
	XPais   string `xml:"xPais"`
	Fone    string `xml:"fone"`
}

type Det struct {
	Text    string  `xml:",chardata"`
	NItem   string  `xml:"nItem,attr"`
	Prod    Prod    `xml:"prod"`
	Imposto Imposto `xml:"imposto"`
}

type Prod struct {
	Text    string `xml:",chardata"`
	CProd   string `xml:"cProd"`
	XProd   string `xml:"xProd"`
	NCM     string `xml:"NCM"`
	CFOP    string `xml:"CFOP"`
	UCom    string `xml:"uCom"`
	QCom    string `xml:"qCom"`
	VUnCom  string `xml:"vUnCom"`
	VProd   string `xml:"vProd"`
	UTrib   string `xml:"uTrib"`
	QTrib   string `xml:"qTrib"`
	VUnTrib string `xml:"vUnTrib"`
	IndTot  string `xml:"indTot"`
}

type Imposto struct {
	Text string `xml:",chardata"`
	ICMS ICMS   `xml:"ICMS"`
}

type ICMS struct {
	Text   string `xml:",chardata"`
	ICMS00 ICMS00 `xml:"ICMS00"`
}

type ICMS00 struct {
	Text  string `xml:",chardata"`
	Orig  string `xml:"orig"`
	CST   string `xml:"CST"`
	ModBC string `xml:"modBC"`
	VBC   string `xml:"vBC"`
	PICMS string `xml:"pICMS"`
	VICMS string `xml:"vICMS"`
}

type Total struct {
	Text    string  `xml:",chardata"`
	ICMSTot ICMSTot `xml:"ICMSTot"`
}

type ICMSTot struct {
	Text   string `xml:",chardata"`
	VBC    string `xml:"vBC"`
	VICMS  string `xml:"vICMS"`
	VProd  string `xml:"vProd"`
	VFrete string `xml:"vFrete"`
	VNF    string `xml:"vNF"`
}

type Transp struct {
	Text       string     `xml:",chardata"`
	ModFrete   string     `xml:"modFrete"`
	Transporta Transporta `xml:"transporta"`
	VeicTransp VeicTransp `xml:"veicTransp"`
	Vol        Vol        `xml:"vol"`
}

type Transporta struct {
	Text   string `xml:",chardata"`
	CNPJ   string `xml:"CNPJ"`
	XNome  string `xml:"xNome"`
	IE     string `xml:"IE"`
	XEnder string `xml:"xEnder"`
	XMun   string `xml:"xMun"`
	UF     string `xml:"UF"`
}

type VeicTransp struct {
	Text  string `xml:",chardata"`
	Placa string `xml:"placa"`
	UF    string `xml:"UF"`
	RNTC  string `xml:"RNTC"`
}

type Vol struct {
	Text  string `xml:",chardata"`
	QVol  string `xml:"qVol"`
	Esp   string `xml:"esp"`
	Marca string `xml:"marca"`
	NVol  string `xml:"nVol"`
	PesoL string `xml:"pesoL"`
	PesoB string `xml:"pesoB"`
}

type Pag struct {
	Text   string `xml:",chardata"`
	DetPag DetPag `xml:"detPag"`
}

type DetPag struct {
	Text   string `xml:",chardata"`
	IndPag string `xml:"indPag"`
	TPag   string `xml:"tPag"`
	VPag   string `xml:"vPag"`
}

type InfAdic struct {
	Text   string `xml:",chardata"`
	InfCpl string `xml:"infCpl"`
}
