package main

import (
	"database/sql"
	pnt "print"
)

func loginCept(id string) bool {
	switch id {
	case "17815918180":
		return true
	case "18248530181":
		return true
	default:
		return false
	}
}
func (lur *LoginUserReq) msgMain() []byte {
	var lus LoginUserRes
	if loginCept(lur.LoginID) {
		lus.Status = 0
	} else {
		lus.Status = 1
		lus.Explain = "登录失败"
	}
	return reParseJSON(lus)

}
func (upr *UploadPicReq) msgMain() []byte {
	var ups UploadPicRes
	if !loginCept(upr.LoginID) {
		ups.Status = 1
		return reParseJSON(ups)
	}
	carid, err := aliyunMain(upr.Data)
	if err == false {
		pnt.ErrorString("aliyun查询失败 " + upr.Data)

		ups.Status = 2
		ups.Explain = "识别失败，重新拍照上传"
		return reParseJSON(ups)
	}

	// 进行查询 carid
	qrerr := DB.QueryRow("SELECT id,carid,name,telphone FROM carid WHERE carid=?", carid).Scan(&ups.ID, &ups.CarID, &ups.Name, &ups.Telphone)

	if qrerr != nil && qrerr != sql.ErrNoRows {
		pnt.ErrorWString("mysql查询失败", qrerr)
		ups.Status = 4
		ups.Explain = "查询失败"
		return reParseJSON(ups)
	}
	ups.CarID = carid
	// 如果无记录
	if qrerr == sql.ErrNoRows {
		pnt.Infof("上传车牌:%s,无记录", ups.CarID)
		ups.Status = 3
		ups.Update = false
		return reParseJSON(ups)
	}

	// 如果有记录
	pnt.Infof("查找车牌成功,车牌:%s,车主:%s,手机号:%s", ups.CarID, ups.Name, ups.Telphone)
	ups.Status = 0
	ups.Update = true
	return reParseJSON(ups)
}
func (ulciq *UploadCarIDReq) msgMain() []byte {
	var ulcis UploadCarIDRes
	// 进行查询 carid
	qrerr := DB.QueryRow("SELECT id,carid,name,telphone FROM carid WHERE carid=?", ulciq.CarID).Scan(&ulcis.ID, &ulcis.CarID, &ulcis.Name, &ulcis.Telphone)

	if qrerr != nil && qrerr != sql.ErrNoRows {
		pnt.ErrorWString("mysql查询失败", qrerr)
		ulcis.Status = 2
		ulcis.Explain = "查询失败"
		return reParseJSON(ulcis)
	}

	// 如果无记录
	if qrerr == sql.ErrNoRows {
		if _, err := DB.Exec("INSERT INTO carid (carid,name,telphone) VALUES (?,?,?)", ulciq.CarID, ulciq.Name, ulciq.Telphone); err != nil {
			pnt.ErrorWString("插入carid失败", err)
			ulcis.Status = 3
			ulcis.Update = false
			ulcis.Explain = "提交失败"
			return reParseJSON(ulcis)
		}
		// 如果有记录
		pnt.Infof("插入成功,车牌:%s,车主:%s,手机号:%s", ulciq.CarID, ulciq.Name, ulciq.Telphone)
		ulcis.Status = 0
		ulcis.Explain = "提交成功"
		ulcis.Update = true
		return reParseJSON(ulcis)

	}

	return nil

}
func (udciq *UpdateCarIDReq) msgMain() []byte {

	var udcis UpdateCarIDRes
	if _, err := DB.Exec("UPDATE carid SET carid=?,name=?,telphone=? WHERE id=?", udciq.CarID, udciq.Name, udciq.Telphone, udciq.ID); err != nil {
		pnt.ErrorWString("更新carid失败", err)
		udcis.Status = 1
		udcis.Explain = "更新失败"
		return reParseJSON(udcis)
	}
	pnt.Infof("更新成功,车牌:%s,车主:%s,手机号:%s", udciq.CarID, udciq.Name, udciq.Telphone)

	udcis.Status = 0
	udcis.Explain = "更新完成"
	return reParseJSON(udcis)
}
