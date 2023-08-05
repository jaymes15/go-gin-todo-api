package users

type UserRes struct {
	ID       uint
	UserName string
}

type UsersRes struct {
	Data []UserRes
}

func ToUser(user UserModel) UserRes {
	return UserRes{
		ID:       user.ID,
		UserName: user.UserName,
	}
}
