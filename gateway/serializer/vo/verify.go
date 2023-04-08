package vo

type ApplyEmailVerify struct {
	Email string `json:"email"`
	Ok    bool   `json:"ok"`
}
