package dto

func NewInfNFe() *InfNFe {
	return &InfNFe{}
}

type InfNFe struct {
	Ide    Ide    `json:"ide"`
	Emit   Emit   `json:"emit"`
	Dest   Dest   `json:"dest"`
	Det    Det    `json:"det"`
	Total  Total  `json:"total"`
	Transp Transp `json:"transp"`
	Versao string `json:"_versao"`
	ID     string `json:"_Id"`
}

type Ide struct {
	CUF      string `json:"cUF"`
	CNF      string `json:"cNF"`
	NatOp    string `json:"natOp"`
	Mod      string `json:"mod"`
	Serie    string `json:"serie"`
	NNF      string `json:"nNF"`
	DhEmi    string `json:"dhEmi"`
	TpNF     string `json:"tpNF"`
	IDDest   string `json:"idDest"`
	CMunFG   string `json:"cMunFG"`
	TpImp    string `json:"tpImp"`
	TpEmis   string `json:"tpEmis"`
	CDV      string `json:"cDV"`
	TpAmb    string `json:"tpAmb"`
	FinNFe   string `json:"finNFe"`
	IndFinal string `json:"indFinal"`
	IndPres  string `json:"indPres"`
	ProcEmi  string `json:"procEmi"`
	VerProc  string `json:"verProc"`
}

type Emit struct {
	CNPJ      string    `json:"CNPJ"`
	XNome     string    `json:"xNome"`
	EnderEmit EnderEmit `json:"enderEmit"`
	IE        string    `json:"IE"`
	CRT       string    `json:"CRT"`
}

type EnderEmit struct {
	XLgr    string `json:"xLgr"`
	Nro     string `json:"nro"`
	XBairro string `json:"xBairro"`
	CMun    string `json:"cMun"`
	XMun    string `json:"xMun"`
	UF      string `json:"UF"`
	CEP     string `json:"CEP"`
	CPais   string `json:"cPais"`
	XPais   string `json:"xPais"`
	Fone    string `json:"fone"`
}

type Dest struct {
	CPF       string    `json:"CPF"`
	XNome     string    `json:"xNome"`
	EnderDest EnderDest `json:"enderDest"`
	IndIEDest string    `json:"indIEDest"`
}

type EnderDest struct {
	XLgr    string `json:"xLgr"`
	Nro     string `json:"nro"`
	XBairro string `json:"xBairro"`
	CMun    string `json:"cMun"`
	XMun    string `json:"xMun"`
	UF      string `json:"UF"`
	CEP     string `json:"CEP"`
	CPais   string `json:"cPais"`
	XPais   string `json:"xPais"`
	Fone    string `json:"fone"`
}

type Det struct {
	Prod    Prod    `json:"prod"`
	Imposto Imposto `json:"imposto"`
	NItem   string  `json:"_nItem"`
}

type Prod struct {
	CProd    string `json:"cProd"`
	CEAN     string `json:"cEAN"`
	XProd    string `json:"xProd"`
	NCM      string `json:"NCM"`
	CFOP     string `json:"CFOP"`
	UCom     string `json:"uCom"`
	QCom     string `json:"qCom"`
	VUnCom   string `json:"vUnCom"`
	VProd    string `json:"vProd"`
	CEANTrib string `json:"cEANTrib"`
	UTrib    string `json:"uTrib"`
	QTrib    string `json:"qTrib"`
	VUnTrib  string `json:"vUnTrib"`
	IndTot   string `json:"indTot"`
}

type Imposto struct {
	ICMS   ICMS   `json:"ICMS"`
	PIS    PIS    `json:"PIS"`
	COFINS COFINS `json:"COFINS"`
}

type ICMS struct {
	ICMS00 ICMS00 `json:"ICMS00"`
}

type ICMS00 struct {
	Orig  string `json:"orig"`
	CST   string `json:"CST"`
	ModBC string `json:"modBC"`
	VBC   string `json:"vBC"`
	PICMS string `json:"pICMS"`
	VICMS string `json:"vICMS"`
}

type PIS struct {
	PISAliq PISAliq `json:"PISAliq"`
}

type PISAliq struct {
	CST  string `json:"CST"`
	VBC  string `json:"vBC"`
	PPIS string `json:"pPIS"`
	VPIS string `json:"vPIS"`
}

type COFINS struct {
	COFINSAliq COFINSAliq `json:"COFINSAliq"`
}

type COFINSAliq struct {
	CST     string `json:"CST"`
	VBC     string `json:"vBC"`
	PCOFINS string `json:"pCOFINS"`
	VCOFINS string `json:"vCOFINS"`
}

type Total struct {
	ICMSTot ICMSTot `json:"ICMSTot"`
}

type ICMSTot struct {
	VBC        string `json:"vBC"`
	VICMS      string `json:"vICMS"`
	VICMSDeson string `json:"vICMSDeson"`
	VFCP       string `json:"vFCP"`
	VBCST      string `json:"vBCST"`
	VST        string `json:"vST"`
	VFCPST     string `json:"vFCPST"`
	VProd      string `json:"vProd"`
	VFrete     string `json:"vFrete"`
	VSeg       string `json:"vSeg"`
	VDesc      string `json:"vDesc"`
	VII        string `json:"vII"`
	VIPI       string `json:"vIPI"`
	VPIS       string `json:"vPIS"`
	VCOFINS    string `json:"vCOFINS"`
	VOutro     string `json:"vOutro"`
	VNF        string `json:"vNF"`
}

type Transp struct {
	ModFrete string `json:"modFrete"`
}
