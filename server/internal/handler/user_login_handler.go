package handler

import (
	"chat_socket/server/internal/logic"
	"chat_socket/server/internal/types"
	"encoding/json"
	"fmt"
)

func UserLoginHandler(msg string) (resp *types.LoginReply, err error) {
	//登录检查
	msg1 := string(msg[6:])
	fmt.Println(msg1)
	buf2 := []byte(msg1)
	var res2 types.LoginRequest
	json.Unmarshal(buf2, &res2)

	b := types.LoginRequest{
		Email:    res2.Email,
		Password: res2.Password,
	}
	resp, err = logic.UserLogic(&b)
	if err != nil {
		fmt.Println("logic.UserLogic(&b)", err)
	}
	return resp, err
}
