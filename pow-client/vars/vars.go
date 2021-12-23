package vars

const (
	VerifyType    = "verify"
	ChooseType    = "choose"
	ChallengeType = "challenge"
)

type Request struct {
	Type         string `json:"type"` // challenge, grant
	Data         string `json:"data"`
	Complexity   int64  `json:"complexity"`
	WordOfWisdom string `json:"words_of_wisdom"`
}

type Response struct {
	Type  string `json:"type"` // choose, verify
	Nonce int64  `json:"nonce"`
	Data  string `json:"data"`
}
