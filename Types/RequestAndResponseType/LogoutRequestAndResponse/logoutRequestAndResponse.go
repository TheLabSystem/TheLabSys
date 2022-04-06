package LogoutRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type LogoutRequest struct {
}
type LogoutResponse struct {
	Code ErrNo.ErrNo `json:"Code" required:"True"`
}
