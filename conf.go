package lib

import (
	"github.com/qamarian-dtp/err"
	"github.com/qamarian-lib/viper"
	"github.com/spf13/afero"
	v "gopkg.in/asaskevich/govalidator.v9"
	"strconv"
)

func Conf () (sds, http map[string]string, funcErr error) {
	conf, errX := viper.NewFileViper (conf_ConfFileName, "yaml")
	if errX != nil {
		funcErr = err.New ("Unable to load conf.", nil, nil, errX)
		return
	}
	sds = map[string]string {}; http = map[string]string {}

// part a0
	tempSdsAddr := conf.GetString ("sds.net_addr")
	if ! v.IsHost (tempSdsAddr) {
		funcErr = err.New ("Conf data 'sds.net_addr': Invalid data.", nil, nil)
		return
	}
	sds ["net_addr"] = tempSdsAddr

// part a1
	tempSdsPort := conf.GetString ("sds.net_port")
	if ! v.IsPort (tempSdsPort) {
		funcErr = err.New ("Conf data 'sds.net_port': Invalid data.", nil, nil)
		return
	}
	if tempSdsPort == "0" {
		funcErr = err.New ("Conf data 'sds.net_port': Invalid data.", nil, nil)
		return
	}
	sds ["net_port"] = tempSdsPort

// part a2
	tempSdsUser := conf.GetString ("sds.user_name")
	if tempSdsUser == "" {
		funcErr = err.New ("Conf data 'sds.user_name': Invalid data.", nil, nil)
		return
	}
	sds ["user_name"] = tempSdsUser

// part a3
	tempSdsPass := conf.GetString ("sds.user_pass")
	if tempSdsPass == "" {
		funcErr = err.New ("Conf data 'sds.user_pass': Invalid data.", nil, nil)
		return
	}
	sds ["user_pass"] = tempSdsPass

// part a4
	tempSdsPubK := conf.GetString ("sds.pub_key")
	
	okA, errA := afero.Exists (afero.NewOsFs (), tempSdsPubK)
	if errA != nil {
		funcErr = err.New ("Conf data 'sds.pub_key': Unable to confirm existence of file.", nil, nil, errA)
		return
	}

	if okA == false {
		funcErr = err.New ("Conf data 'sds.pub_key': File does not exist.", nil, nil)
		return
	}

	sds ["pub_key"] = tempSdsPubK

// part a5
	tempSdsUpdatePass := conf.GetString ("sds.update_pass")
	if tempSdsUpdatePass == "" {
		funcErr = err.New ("Conf data 'sds.update_pass': Invalid data.", nil, nil)
		return
	}
	sds ["update_pass"] = tempSdsUpdatePass

// part b0
	tempHttpAddr := conf.GetString ("http.net_addr")
	if ! v.IsIPv4 (tempHttpAddr) {
		funcErr = err.New ("Conf data 'http.net_addr': Invalid data.", nil, nil)
		return
	}
	http ["net_addr"] = tempHttpAddr

// part b1
	tempHttpPort := conf.GetString ("http.net_port")
	if ! v.IsPort (tempHttpPort) {
		funcErr = err.New ("Conf data 'http.net_port': Invalid data.", nil, nil)
		return
	}
	if tempHttpPort == "0" {
		funcErr = err.New ("Conf data 'http.net_port': Invalid data.", nil, nil)
		return
	}
	http ["net_port"] = tempHttpPort

// part b2
	tempHttpTlsKey := conf.GetString ("http.tls_key")
	
	okB, errB := afero.Exists (afero.NewOsFs (), tempHttpTlsKey)
	if errB != nil {
		funcErr = err.New ("Conf data 'http.tls_key': Unable to confirm existence of file.", nil, nil, errB)
		return
	}

	if okB == false {
		funcErr = err.New ("Conf data 'http.tls_key': File does not exist.", nil, nil)
		return
	}

	http ["tls_key"] = tempHttpTlsKey

// part b3
	tempHttpTlsCrt := conf.GetString ("http.tls_crt")
	
	okC, errC := afero.Exists (afero.NewOsFs (), tempHttpTlsCrt)
	if errC != nil {
		funcErr = err.New ("Conf data 'http.tls_crt': Unable to confirm existence of file.", nil, nil, errC)
		return
	}

	if okC == false {
		funcErr = err.New ("Conf data 'http.tls_crt': File does not exist.", nil, nil)
		return
	}

	http ["tls_crt"] = tempHttpTlsCrt

// part b4
	tempHttpReadTmt, _ := strconv.Atoi (conf.GetString ("http.read_timeout"))
	if tempHttpReadTmt < 1 || tempHttpReadTmt > 960 {
		funcErr = err.New ("Conf data 'http.read_timeout': Invalid data.", nil, nil)
		return
	}
	http ["read_timeout"] = strconv.Itoa (tempHttpReadTmt)

// part b5
	tempHttpReadHeaderTmt, _ := strconv.Atoi (conf.GetString ("http.read_header_timeout"))
	if tempHttpReadHeaderTmt < 1 || tempHttpReadHeaderTmt > 960 {
		funcErr = err.New ("Conf data 'http.read_header_timeout': Invalid data.", nil, nil)
		return
	}
	http ["read_header_timeout"] = strconv.Itoa (tempHttpReadHeaderTmt)

// part b6
	tempHttpWrteTmt, _ := strconv.Atoi (conf.GetString ("http.wrte_timeout"))
	if tempHttpWrteTmt < 1 || tempHttpWrteTmt > 960 {
		funcErr = err.New ("Conf data 'http.wrte_timeout': Invalid data.", nil, nil)
		return
	}
	http ["wrte_timeout"] = strconv.Itoa (tempHttpWrteTmt)

// part b7
	tempHttpIdleTmt, _ := strconv.Atoi (conf.GetString ("http.idle_timeout"))
	if tempHttpIdleTmt < 1 || tempHttpIdleTmt > 960 {
		funcErr = err.New ("Conf data 'http.idle_timeout': Invalid data.", nil, nil)
		return
	}
	http ["idle_timeout"] = strconv.Itoa (tempHttpIdleTmt)

// end of parts
	return
}
var (
	conf_ConfFileName = "httpConf.yml"
)
