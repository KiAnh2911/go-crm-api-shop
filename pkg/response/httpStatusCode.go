package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvaild = 20003 // Email is Invalid
	ErrorInvaildToken     = 30001 // Token is Invalid
)

// message
var Msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvaild: "Email is invalid",
	ErrorInvaildToken:     "Token is Invalid",
}
