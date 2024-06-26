package mocks

type CEP struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	DDD        string `json:"ddd"`
}

var Ceps = []CEP{
	{
		Cep:        "26556-030",
		Logradouro: "Rua Augusto Cardoso",
		Bairro:     "Coréia",
		Localidade: "Mesquita",
		UF:         "RJ",
		DDD:        "021",
	},
	{
		Cep:        "01001-000",
		Logradouro: "Praça da Sé",
		Bairro:     "Sé",
		Localidade: "São Paulo",
		UF:         "SP",
		DDD:        "11",
	},
}
