package NoticeRequestAndResponse

import "TheLabSystem/Types/RequestAndResponseType/ErrNo"

type AddNoticeRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AddNoticeResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message string `json:"message"`
	}
}

type GgetNoticeRequest struct {
}

type GetNoticeResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message string `json:"message"`
	}
}

type DeleteNoticeRequest struct {
	ID int `json:"id"`
}

type DeleteNoticeResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message string `json:"message"`
	}
}
