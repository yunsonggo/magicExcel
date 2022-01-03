package param

type RegisterParam struct {
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CheckPass string `json:"check_pass" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
}

type LoginParam struct {
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
}

type ResetParam struct {
	OldPass   string `json:"old_pass"`
	Pass      string `json:"pass"`
	CheckPass string `json:"check_pass"`
}
