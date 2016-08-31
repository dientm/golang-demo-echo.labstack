package beans

type Quote struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func NewQuote(name string, passwd string) (bs Quote) {
	bs.Username = name
	bs.Password = passwd
	return
}

func (b Quote) SetUsername(name string) {
	b.Username = name
}

func (b Quote) SetPassword(pw string) {
	b.Password = pw
}