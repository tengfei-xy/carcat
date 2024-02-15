package main

//
// aliyun
//

type aliyun_info struct {
	Config_str string `json:"config_str"`
	Plates     []aliyun_info_plates
	Request_id string `json:"request_id"`
	Success    bool   `json:"success"`
}
type aliyun_info_plates struct {
	Cls_name string  `json:"cls_name"`
	Cls_prob float64 `json:"cls_prob"`
	Detail   string  `json:"detail"`
	Prob     float64 `json:"prob"`
	Roi      aliyun_info_plates_ros
	Txt      string `json:"txt"`
}
type aliyun_info_plates_ros struct {
	H int `json:"h"`
	W int `json:"w"`
	X int `json:"x"`
	Y int `json:"y"`
}

//
// carcat
//
type carid struct {
	ID       string `json:"id"`
	CarID    string `json:"carid"`
	Name     string `json:"name"`
	Telphone string `json:"telphone"`
}
type basereq struct {
	Action string `json:"action"`
}
type baseres struct {
	Status  int    `json:"status"`
	Explain string `json:"explain"`
}
type LoginUserReq struct {
	basereq
	LoginID string `json:"loginid"`
}
type LoginUserRes struct {
	baseres
}
type UploadPicReq struct {
	basereq
	LoginID string `json:"loginid"`
	Data    string `json:"data"`
}
type UploadPicRes struct {
	baseres
	carid
	Update bool `json:"update"`
}
type UploadCarIDReq struct {
	baseres
	LoginID string `json:"loginid"`
	carid
}
type UploadCarIDRes struct {
	baseres
	carid
	Update bool `json:"update"`
}
type UpdateCarIDReq struct {
	baseres
	LoginID string `json:"loginid"`
	carid
}
type UpdateCarIDRes struct {
	baseres
}
