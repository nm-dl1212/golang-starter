package usecase

import (
	"os"
	"rest-api/model"
	"rest-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {

	// パスワードからハッシュ値を算出する
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	// ハッシュ化したパスワードで新しくユーザーを作り直す
	newUser := model.User{Email: user.Email, Password: string(hash)}
	// uu -> urと飛んで，CreateUser()関数を呼び出す
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err // 失敗した場合は空のユーザーを返す
	}

	// レスポンスを作成する
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {

	// DBに保存されているユーザーと，引数のuserを照会する。
	storedUser := model.User{}
	// メールの照会
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// パスワードとハッシュ値の照会
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// jwtトークン(json web token)の生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
