package users

type UserRespositoryInterface interface {
	Create(user UserModel) (UserModel, error)
	FindByID(id uint) UserModel
	FindByUserName(username string) UserModel
}
