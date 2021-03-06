package vars

const (
	GrantType     = "grant"
	VerifyType    = "verify"
	ChooseType    = "choose"
	ChallengeType = "challenge"
)

type Request struct {
	Type  string `json:"type"` // choose, verify
	Nonce int64  `json:"nonce"`
	Data  string `json:"data"`
}

type Response struct {
	Type          string `json:"type"` // challenge, grant
	Data          string `json:"data"`
	Complexity    int64  `json:"complexity"`
	WordsOfWisdom string `json:"words_of_wisdom"`
}

var Complexity int64 = 18

var Quotes = []string{
	"“The best way out is always through.”- Robert Frost",
	"Carpe Diem – Latin Proverb",
	"“Always Do What You Are Afraid To Do” – Ralph Waldo Emerson",
	"“Believe and act as if it were impossible to fail.” – Charles Kettering",
	"“Keep steadily before you the fact that all true success depends at last upon yourself.” – Theodore T. Hunger",
	"“The journey of a thousand miles begins with one step.” – Lao Tzu",
	"“Opportunity is always knocking. The problem is that most people have the self-doubt station in their heads turned up way too loud to hear it” – Brian Vaszily",
	"“You must be the change you wish to see in the world.” – Gandhi",
	"“Tough times never last, but tough people do.” – Dr. Robert Schuller",
	"“You must not only aim right, but draw the bow with all your might.” – Henry David Thoreau",
}
