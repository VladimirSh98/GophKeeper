package secret

type BankCard struct {
	Number      string `json:"number"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
	HolderName  string `json:"holder_name"`
	CVV         string `json:"cvv,omitempty"`
}

type LoginPass struct {
	Password string `json:"password"`
	Login    string `json:"login"`
}
