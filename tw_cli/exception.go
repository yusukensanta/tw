package tw_cli

import (
)

const (
    UNKNOWN_PARAM      = 1
    INSUFFICIENT_PARAM = 2
)

type TwError struct {
    Msg string
    Code int
}

func GenerateError(errCode int){

}