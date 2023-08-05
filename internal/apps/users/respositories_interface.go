package users

type UserRespositoryInterface interface {
	Create(user UserModel) (UserModel, error)
	FindByID(id int) UserModel
	FindByUserName(username string) UserModel
}
