package constant

type ResponseStatus int
type Headers int
type General int

// Constant Api Constant
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	Unauthorized
	InvalidRequest
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "ALREADY_CONFIRMED", "DATA_ALREADY_EXIST", "UNAUTHORIZED", "ACCESS_DENIED", "SHOUL_CHANGE_PASSWORD", "UNDEFINED_CONFIGURATION"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Invitation already confirmed!", "Data Already Exists", "Unauthorized", "Access Denied", "Please change password!", "Undefined configuration!"}[r-1]
}
