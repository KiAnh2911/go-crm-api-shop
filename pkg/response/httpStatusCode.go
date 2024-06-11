package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvaild = 20003 // Email Invalid
)

// message
var Msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvaild: "Email is invalid",
}
