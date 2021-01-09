package main

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//获取相差时间
func getHourDiffer(startTime, endTime string) int64 {
	a, _ := time.Parse("2006-01-02 15:04:05", startTime)
	b, _ := time.Parse("2006-01-02 15:04:05", endTime)
	d := b.Sub(a)
	logrus.Infof("距离域名过期还有 -> %d天",int64(d.Hours()/24))
	return int64(d.Hours() / 24)
}

func main() {
	domainName  := "https://www.baidu.com"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(domainName)
	defer resp.Body.Close()

	if err != nil {
		logrus.Error(domainName, " 请求失败")
		panic(err)
	}

	//fmt.Println(resp.TLS.PeerCertificates[0])
	certInfo := resp.TLS.PeerCertificates[0]
	logrus.Info("过期时间:", certInfo.NotAfter)
	logrus.Info("组织信息:", certInfo.Subject)

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	endTime := certInfo.NotAfter.Format("2006-01-02 15:04:05")
	getHourDiffer(nowTime,endTime)

}
