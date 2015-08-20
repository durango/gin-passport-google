package GinPassportGoogle

type Profile struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Hd            string `json:"hd"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

func (profile *Profile) FirstName() string {
	return profile.FamilyName
}

func (profile *Profile) LastName() string {
	return profile.GivenName
}
