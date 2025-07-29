package schemas

var NfeSchema = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "Generated schema for Root",
  "type": "object",
  "properties": {
    "ide": {
      "type": "object",
      "properties": {
        "cUF": {
          "type": "string"
        },
        "cNF": {
          "type": "string"
        },
        "natOp": {
          "type": "string"
        },
        "mod": {
          "type": "string"
        },
        "serie": {
          "type": "string"
        },
        "nNF": {
          "type": "string"
        },
        "dhEmi": {
          "type": "string"
        },
        "tpNF": {
          "type": "string"
        },
        "idDest": {
          "type": "string"
        },
        "cMunFG": {
          "type": "string"
        },
        "tpImp": {
          "type": "string"
        },
        "tpEmis": {
          "type": "string"
        },
        "cDV": {
          "type": "string"
        },
        "tpAmb": {
          "type": "string"
        },
        "finNFe": {
          "type": "string"
        },
        "indFinal": {
          "type": "string"
        },
        "indPres": {
          "type": "string"
        },
        "procEmi": {
          "type": "string"
        },
        "verProc": {
          "type": "string"
        }
      },
      "required": [
        "cUF",
        "cNF",
        "natOp",
        "mod",
        "serie",
        "nNF",
        "dhEmi",
        "tpNF",
        "idDest",
        "cMunFG",
        "tpImp",
        "tpEmis",
        "cDV",
        "tpAmb",
        "finNFe",
        "indFinal",
        "indPres",
        "procEmi",
        "verProc"
      ]
    },
    "emit": {
      "type": "object",
      "properties": {
        "CNPJ": {
          "type": "string"
        },
        "xNome": {
          "type": "string"
        },
        "enderEmit": {
          "type": "object",
          "properties": {
            "xLgr": {
              "type": "string"
            },
            "nro": {
              "type": "string"
            },
            "xBairro": {
              "type": "string"
            },
            "cMun": {
              "type": "string"
            },
            "xMun": {
              "type": "string"
            },
            "UF": {
              "type": "string"
            },
            "CEP": {
              "type": "string"
            },
            "cPais": {
              "type": "string"
            },
            "xPais": {
              "type": "string"
            },
            "fone": {
              "type": "string"
            }
          },
          "required": [
            "xLgr",
            "nro",
            "xBairro",
            "cMun",
            "xMun",
            "UF",
            "CEP",
            "cPais",
            "xPais",
            "fone"
          ]
        },
        "IE": {
          "type": "string"
        },
        "CRT": {
          "type": "string"
        }
      },
      "required": [
        "CNPJ",
        "xNome",
        "enderEmit",
        "IE",
        "CRT"
      ]
    },
    "dest": {
      "type": "object",
      "properties": {
        "CPF": {
          "type": "string"
        },
        "xNome": {
          "type": "string"
        },
        "enderDest": {
          "type": "object",
          "properties": {
            "xLgr": {
              "type": "string"
            },
            "nro": {
              "type": "string"
            },
            "xBairro": {
              "type": "string"
            },
            "cMun": {
              "type": "string"
            },
            "xMun": {
              "type": "string"
            },
            "UF": {
              "type": "string"
            },
            "CEP": {
              "type": "string"
            },
            "cPais": {
              "type": "string"
            },
            "xPais": {
              "type": "string"
            },
            "fone": {
              "type": "string"
            }
          },
          "required": [
            "xLgr",
            "nro",
            "xBairro",
            "cMun",
            "xMun",
            "UF",
            "CEP",
            "cPais",
            "xPais",
            "fone"
          ]
        },
        "indIEDest": {
          "type": "string"
        }
      },
      "required": [
        "CPF",
        "xNome",
        "enderDest",
        "indIEDest"
      ]
    },
    "det": {
      "type": "object",
      "properties": {
        "prod": {
          "type": "object",
          "properties": {
            "cProd": {
              "type": "string"
            },
            "cEAN": {
              "type": "string"
            },
            "xProd": {
              "type": "string"
            },
            "NCM": {
              "type": "string"
            },
            "CFOP": {
              "type": "string"
            },
            "uCom": {
              "type": "string"
            },
            "qCom": {
              "type": "string"
            },
            "vUnCom": {
              "type": "string"
            },
            "vProd": {
              "type": "string"
            },
            "cEANTrib": {
              "type": "string"
            },
            "uTrib": {
              "type": "string"
            },
            "qTrib": {
              "type": "string"
            },
            "vUnTrib": {
              "type": "string"
            },
            "indTot": {
              "type": "string"
            }
          },
          "required": [
            "cProd",
            "cEAN",
            "xProd",
            "NCM",
            "CFOP",
            "uCom",
            "qCom",
            "vUnCom",
            "vProd",
            "cEANTrib",
            "uTrib",
            "qTrib",
            "vUnTrib",
            "indTot"
          ]
        },
        "imposto": {
          "type": "object",
          "properties": {
            "ICMS": {
              "type": "object",
              "properties": {
                "ICMS00": {
                  "type": "object",
                  "properties": {
                    "orig": {
                      "type": "string"
                    },
                    "CST": {
                      "type": "string"
                    },
                    "modBC": {
                      "type": "string"
                    },
                    "vBC": {
                      "type": "string"
                    },
                    "pICMS": {
                      "type": "string"
                    },
                    "vICMS": {
                      "type": "string"
                    }
                  },
                  "required": [
                    "orig",
                    "CST",
                    "modBC",
                    "vBC",
                    "pICMS",
                    "vICMS"
                  ]
                }
              },
              "required": [
                "ICMS00"
              ]
            },
            "PIS": {
              "type": "object",
              "properties": {
                "PISAliq": {
                  "type": "object",
                  "properties": {
                    "CST": {
                      "type": "string"
                    },
                    "vBC": {
                      "type": "string"
                    },
                    "pPIS": {
                      "type": "string"
                    },
                    "vPIS": {
                      "type": "string"
                    }
                  },
                  "required": [
                    "CST",
                    "vBC",
                    "pPIS",
                    "vPIS"
                  ]
                }
              },
              "required": [
                "PISAliq"
              ]
            },
            "COFINS": {
              "type": "object",
              "properties": {
                "COFINSAliq": {
                  "type": "object",
                  "properties": {
                    "CST": {
                      "type": "string"
                    },
                    "vBC": {
                      "type": "string"
                    },
                    "pCOFINS": {
                      "type": "string"
                    },
                    "vCOFINS": {
                      "type": "string"
                    }
                  },
                  "required": [
                    "CST",
                    "vBC",
                    "pCOFINS",
                    "vCOFINS"
                  ]
                }
              },
              "required": [
                "COFINSAliq"
              ]
            }
          },
          "required": [
            "ICMS",
            "PIS",
            "COFINS"
          ]
        },
        "_nItem": {
          "type": "string"
        }
      },
      "required": [
        "prod",
        "imposto",
        "_nItem"
      ]
    },
    "total": {
      "type": "object",
      "properties": {
        "ICMSTot": {
          "type": "object",
          "properties": {
            "vBC": {
              "type": "string"
            },
            "vICMS": {
              "type": "string"
            },
            "vICMSDeson": {
              "type": "string"
            },
            "vFCP": {
              "type": "string"
            },
            "vBCST": {
              "type": "string"
            },
            "vST": {
              "type": "string"
            },
            "vFCPST": {
              "type": "string"
            },
            "vProd": {
              "type": "string"
            },
            "vFrete": {
              "type": "string"
            },
            "vSeg": {
              "type": "string"
            },
            "vDesc": {
              "type": "string"
            },
            "vII": {
              "type": "string"
            },
            "vIPI": {
              "type": "string"
            },
            "vPIS": {
              "type": "string"
            },
            "vCOFINS": {
              "type": "string"
            },
            "vOutro": {
              "type": "string"
            },
            "vNF": {
              "type": "string"
            }
          },
          "required": [
            "vBC",
            "vICMS",
            "vICMSDeson",
            "vFCP",
            "vBCST",
            "vST",
            "vFCPST",
            "vProd",
            "vFrete",
            "vSeg",
            "vDesc",
            "vII",
            "vIPI",
            "vPIS",
            "vCOFINS",
            "vOutro",
            "vNF"
          ]
        }
      },
      "required": [
        "ICMSTot"
      ]
    },
    "transp": {
      "type": "object",
      "properties": {
        "modFrete": {
          "type": "string"
        }
      },
      "required": [
        "modFrete"
      ]
    },
    "_versao": {
      "type": "string"
    },
    "_Id": {
      "type": "string"
    }
  },
  "required": [
    "ide",
    "emit",
    "dest",
    "det",
    "total",
    "transp",
    "_versao",
    "_Id"
  ]
}`
