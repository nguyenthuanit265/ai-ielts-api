package models

type AuthResponse struct {
	//UserId       int      `json:"userId"`
	//FullName     string   `json:"fullName"`
	//Email        string   `json:"email"`
	//Roles        []string `json:"roles"`
	//Permissions  []string `json:"permissions"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type AuthClaim struct {
	Id          int      `json:"id,omitempty" mapstructure:"id,omitempty"`
	FullName    string   `json:"fullName,omitempty" mapstructure:"fullName,omitempty"`
	Email       string   `json:"email,omitempty" mapstructure:"email,omitempty"`
	Roles       []string `json:"roles,omitempty" mapstructure:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty" mapstructure:"permissions,omitempty"`
	Authorized  bool     `json:"authorized,omitempty" mapstructure:"authorized,omitempty"`
}
