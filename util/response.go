package util

type SingleRes struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Data any `json:"data"`
}