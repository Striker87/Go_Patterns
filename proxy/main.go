package main

import (
	"fmt"
	"patterns/proxy/pkg"
)

/*
Заместитель - структурный паттерн проектирования, который позволяет контролировать процесс общения с реальным объектом. Заместитель перехватывает вызовы к оригинальному объекту, позволяя сделать что-то до или после передачи вызова.
*/

var (
	admin = "admin"
	user  = "user"

	users = map[string]bool{
		admin: true,
		user:  false,
	}
)

func main() {
	proxy := pkg.ProxyDataBase{
		Users: users,
		Db:    &pkg.DataBase{},
	}

	adminData, err := proxy.GetData(admin)
	fmt.Printf("From [%s] Data: [%v] Error: [%v]\n", admin, adminData, err)

	userData, err := proxy.GetData(user)
	fmt.Printf("From [%s] Data: [%v] Error: [%v]\n", user, userData, err)
}
