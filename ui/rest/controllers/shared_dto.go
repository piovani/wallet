package controllers

type ReponseMsg struct {
	Msg string `json:"message"`
}

func NewReponseMsg(msg string) *ReponseMsg {
	return &ReponseMsg{Msg: msg}
}

type ReponseErr struct {
	Err string `json:"error message"`
}

func NewReponseErr(err error) *ReponseErr {
	return &ReponseErr{Err: err.Error()}
}
