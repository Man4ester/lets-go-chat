package repositories

var usersDataRep usersDataRepository

func AddUserRepository(rep *usersDataRepository){
	usersDataRep = *rep
}

func GetUSerRepository() usersDataRepository {
	return usersDataRep
}
