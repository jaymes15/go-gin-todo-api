package chat

import "sort"

func getUserList() []string {
	var userList []string
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}

	}
	sort.Strings(userList)
	return userList
}
