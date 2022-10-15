package types

type Owner struct {
	Share uint `json:"share"`
	Individual
}

type Mortgage struct {
	Creditor string `json:"creditor"`
	Amount   uint   `json:"amount"`
}

type Land struct {
	ID       string    `json:"id"`
	Address  string    `json:"address"`
	Owners   []Owner   `json:"owners"`
	Area     uint      `json:"area"`
	Mortgage *Mortgage `json:"mortgage"`
}

type NotarialAct struct {
	LandID       string     `json:"landID"`
	Notary       Individual `json:"notary"`
	CreationDate string     `json:"creationDate"`
	Description  string     `json:"description"`
	Owner        Individual `json:"owner"`
}

type State = string

const (
	Submitted          State = "SUBMITTED"
	Blocked            State = "BLOCKED"
	Rejected           State = "REJECTED"
	Accepted           State = "Accepted"
	WaitingForPayment  State = "WAITING_FOR_PAYMENT"
	WaitingForApproval State = "WAITING_FOR_APPROVAL"
)

type Request struct {
	LandID       string       `json:"landID"`
	State        State        `json:"state"`
	CreationTime string       `json:"creationTime"`
	Parties      []Individual `json:"parties"`
	Intermediary Individual   `json:"intermediary"`
}

type SignedRequest struct {
	Signatures []string `json:"signatures"`
	Request
}

type Individual struct {
	PESEL     string `json:"pesel"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PublicKey string `json:"publicKey"`
}
