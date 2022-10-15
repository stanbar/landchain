package types

type NotarialAct struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

type Party struct {
	PESEL     string `json:"pesel"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PubKey    string `json:"pubKey"`
}

type Request struct {
	CreationTime string  `json:"creationTime"`
	ID           string  `json:"id"`
	Parties      []Party `json:"parties"`
	Signature    string  `json:"signature"`
}
