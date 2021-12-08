package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body error %v \n", err)
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed %v \n", err)
		return
	}

	fmt.Fprintf(w, "success")
}

func SignUpWithContext(c *Context) {
	req := &signUpReq{}
	// 传入指针修改值
	err := c.ReadJson(req)
	if err != nil {
		_ = c.BadJson(&commonResponse{
			BizCode: 4,
			Msg:     fmt.Sprintf("invalid requiest: %v", err),
		})
		return
	}
	c.OkJson(&commonResponse{
		BizCode: 0,
		Msg: fmt.Sprintf("success, code is %d", 12),
	})
}

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}

func (c *Context) ReadJson(data interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func (c *Context) WriteJson(data interface{}, status int) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(bs)
	if err != nil {
		return err
	}
	c.W.WriteHeader(status)
	return nil
}

func (c *Context) OkJson(data interface{}) error {
	return c.WriteJson(data, http.StatusOK)
}

func (c *Context) BadJson(data interface{}) error {
	return c.WriteJson(data, http.StatusBadRequest)
}

type commonResponse struct {
	BizCode int8   `json:"biz_code"`
	Msg     string `json:"msg"`
}
