package utils

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
)

func PDFT() {
	c := gofpdf.InitType{
		Size: gofpdf.SizeType{
			Wd: 160,
			Ht: 80,
		},
	}

	// 获取二维码
	wechatUtils := new(WechatUtils)
	wechatUtils.SmallQrcodeData.Path = "card/money/detail?card_no=" + "v.CardNo"
	wechatUtils.SmallQrcodeData.Width = 430
	wechatUtils.Init("wx70618bbce8b722ab", "b4b911ea1ed1757a9b674c7091390f9c")
	wechatUtils.GetAccessToken()
	b, err := wechatUtils.GetSmallQrcode()
	if err != nil {
		log.Println(err.Error())
	}
	ioutil.WriteFile("logo.jpg", b, 0755)

	pdf := gofpdf.NewCustom(&c)
	pdf.AddUTF8Font("dayinzi", "", "ttf/pdf.ttf")
	//pdf.SetCellMargin(0)
	//pdf.SetLeftMargin(0)
	//pdf.SetTopMargin(0)
	//pdf.SetTopMargin(0)
	pdf.AddPage()
	pdf.Image("logo.jpg", 10, 10, 60, 60, false, "", 0, "")
	//pdf.Text(60, 10, "所属项目：哈哈哈哈哈哈哈哈哈哈")
	//pdf.SetY(pdf.SetY(0)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
	pdf.SetY(10) //水平居中的算法0)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
	pdf.SetX(70) //水平居中的算法
	pdf.SetFont("dayinzi", "", 20)
	pdf.MultiCell(90, 9,
		fmt.Sprintf("项目名称：%s \n\r产品名称：%s \n\r序号：%d/%d",
			"测试项目测试项目测试项目测试项目",
			"产品名称产品名称产品名称产品名称产品名称",
			1, 2), "", "Left", false)
	pdf.AddPage()
	pdf.Image("logo.jpg", 10, 10, 60, 60, false, "", 0, "")
	//pdf.Text(60, 10, "所属项目：哈哈哈哈哈哈哈哈哈哈")
	//pdf.SetY(pdf.SetY(0)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
	pdf.SetY(10) //水平居中的算法0)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
	pdf.SetX(70) //水平居中的算法
	pdf.SetFont("dayinzi", "", 20)
	pdf.MultiCell(90, 9,
		fmt.Sprintf("项目名称：%s \n\r产品名称：%s \n\r序号：%d/%d",
			"测试项目测试项目测试项目测试项目",
			"产品名称产品名称产品名称产品名称产品名称",
			1, 2), "", "Left", false)
	pdf.AddPage()
	pdf.Image("logo.jpg", 10, 10, 60, 60, false, "", 0, "")
	//pdf.Text(60, 10, "所属项目：哈哈哈哈哈哈哈哈哈哈")
	//pdf.SetY(pdf.SetY(0)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
	pdf.SetY(10) //水平居中的算法0)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
	pdf.SetX(70) //水平居中的算法
	pdf.SetFont("dayinzi", "", 20)
	pdf.MultiCell(90, 9,
		fmt.Sprintf("项目名称：%s \n\r产品名称：%s \n\r序号：%d/%d",
			"测试项目测试项目测试项目测试项目",
			"产品名称产品名称产品名称产品名称产品名称",
			1, 2), "", "Left", false)

	fileStr := "xx.pdf"
	err = pdf.OutputFileAndClose(fileStr)
	if err != nil {

	}
}
