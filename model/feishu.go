package model

//接受飞书请求的结构体,会巨丑无比
type CommonEvent struct {
	Header struct {
		EventType string `json:"event_type"`
		Token     string `json:"token"`
	} `json:"header"`

	Event struct {
		Sender struct {
			Sender_id struct {
				UserID string `json:"user_id"`
			} `json:"sender_id"`
		} `json:"sender"`
	} `json:"event"`
}

type FirstPost struct {
	Clg   string `json:"challenge"`
	Token string `json:"token"`
	Type  string `json:"type"`
}

type EventContent struct {
	Event struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"event"`
}

type FSUserIDResp struct {
	Data struct {
		EmployeeID   string `json:"employee_id"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

type FSUserAuth struct {
	Data struct {
		User struct {
			Name    string `json:"name"`
			EmpType int    `json:"employee_type"`
		} `json:"user"`
	} `json:"data"`
}
