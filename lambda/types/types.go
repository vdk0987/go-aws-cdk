package types
import (
	"golang.org/x/crypto/bcrypt"
)
type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	PasswordHash string `json:"password"`
}

func newUser (registerUser RegisterUser) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), 10)

	if err!=nil{
		return User{}, err
	}

	return User{
		Username: registerUser.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

func validatePassword (hashedPassword, plainPassword string) bool {
	res := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return res == nil
}
