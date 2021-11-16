package repositories

var usersDataRep usersDataRepository

func RegisterUserRepository(rep *usersDataRepository){
	usersDataRep = *rep
}

func GetUserRepository() UserRepository {
	return usersDataRep
}
