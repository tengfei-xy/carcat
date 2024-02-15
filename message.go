package main

import (
	"env"
	pnt "print"
)

func msgMain(msg []byte) []byte {
	if len(msg) < 22 {
		return nil
	}
	msgtype := string(msg[11:22])

	switch msgtype {

	case env.ActLoginUser:
		pnt.Json(string(msg))
		var lur LoginUserReq
		parseJSON(&msg, &lur)
		return lur.msgMain()

	case env.ActUploadPic:
		var upr UploadPicReq
		parseJSON(&msg, &upr)
		return upr.msgMain()
	case env.ActUploadCarID:
		pnt.Json(string(msg))
		var ulciq UploadCarIDReq
		parseJSON(&msg, &ulciq)
		return ulciq.msgMain()

	case env.ActUpdateCarID:
		pnt.Json(string(msg))
		var udciq UpdateCarIDReq
		parseJSON(&msg, &udciq)
		return udciq.msgMain()

	// 默认
	default:
		return nil
	}
	return nil
}
