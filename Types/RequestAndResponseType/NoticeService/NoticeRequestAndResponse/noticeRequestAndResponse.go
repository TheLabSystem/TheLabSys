package NoticeRequestAndResponse

import (
	"TheLabSystem/Types/RequestAndResponseType/ErrNo"
	"TheLabSystem/Types/ServiceType/Notice"
)

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

type GetNoticeResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message string `json:"message"`
		Notice  []Notice.Notice
		Total   int
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

type UpdateNoticeRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNoticeResponse struct {
	Code ErrNo.ErrNo
	Data struct {
		Message string `json:"message"`
	}
}
