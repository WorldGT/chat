package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type Userclaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var buf = make([]byte, 1024)
var uc = UserLogin{
	Email:    "stu01@qq.com",
	Password: "12346",
}

func login(conn net.Conn) {
	buf1 := []byte("/login")
	conn.Write(buf1)
	UCjsons, errs := json.Marshal(uc)
	if errs != nil {
		fmt.Println("json marshal error", errs)
	}
	fmt.Println("json data", string(UCjsons))

	conn.Write(UCjsons)

	// n, err := conn.Read(buf)
	// if err == io.EOF {
	// 	return
	// } else if err != nil {
	// 	fmt.Println("err1 :", err)
	// 	return
	// }
	// fmt.Println(string(buf[:n]))
}

func inputLogin() {
	fmt.Println("输入账号")
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	uc.Email = string(data)
	fmt.Println("输入密码")
	reader = bufio.NewReader(os.Stdin)
	data, _, _ = reader.ReadLine()
	uc.Password = string(data)

}

func sendMsd(conn net.Conn) {
	var input string

	for {

		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)
		if strings.ToUpper(input) == "Q" {
			return
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("err2 :", err)
			return
		}
	}
}

func main() {
	inputLogin()
	//	与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8085")
	if err != nil {
		fmt.Println("err0 :", err)
		return
	}
	login(conn)
	go sendMsd(conn)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println(err)
			error.Error(err)
			return
		} else if err != nil {
			fmt.Println("err1 :", err)
			return
		}
		fmt.Println(string(buf[:n]))

	}
}
