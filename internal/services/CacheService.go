package services

var cacheRegistry = make(map[string] bool)


func AddUserToCache(user string){
	cacheRegistry[user]= true
}

func RemoveUserFromCache(user string){
	if _, ok := cacheRegistry[user]; ok {
		delete(cacheRegistry, user)
	}
}

func GetTotalActiveUsers() int {
	return 10
}