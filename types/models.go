package types

type Individual struct {
	PESEL     string `json:"pesel"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PublicKey string `json:"publicKey"`
}

type Owner struct {
	Share uint `json:"share"`
	Individual
}

type Mortgage struct {
	Creditor string `json:"creditor"`
	Amount   uint   `json:"amount"`
}

type NotarialAct struct {
	LandID       string     `json:"landID"`
	Notary       Individual `json:"notary"`
	CreationDate string     `json:"creationDate"`
	Description  string     `json:"description"`
	Owner        Individual `json:"owner"`
}

type Land struct {
	ID          string      `json:"id"`
	Address     string      `json:"address"`
	Owner       Owner       `json:"owner"`
	Area        uint        `json:"area"`
	Mortgage    Mortgage    `json:"mortgage"`
	NotarialAct NotarialAct `json:"notarialAct"`
}

type State = string

const (
	Submitted          State = "SUBMITTED"
	Blocked            State = "BLOCKED"
	Rejected           State = "REJECTED"
	Accepted           State = "ACCEPTED"
	WaitingForPayment  State = "WAITING_FOR_PAYMENT"
	WaitingForApproval State = "WAITING_FOR_APPROVAL"
)

type Request struct {
	LandID              string       `json:"landID"`
	CreationTime        string       `json:"creationTime"`
	Parties             []Individual `json:"parties"`
	Intermediary        Individual   `json:"intermediary"`
	NewOwner            Owner        `json:"newOwner"`
	PreviousNotarialAct NotarialAct  `json:"previousNotarialAct"`
	NewNotarialAct      NotarialAct  `json:"newNotarialAct"`
}

type SignedRequest struct {
	ID         string   `json:"id"`
	State      State    `json:"state"`
	Signatures []string `json:"signatures"`
	Request
}
