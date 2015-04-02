package sip_parser_lite

import (
	//	"fmt"
	"strings"
)

type Sip_msg struct {
	Sip_type       string //response or request
	First_line     string //request of status line
	Method_or_Code string //200, 404, INVITE, PRACK
	Call_id        string
	Headers        Head_map
}

//Authorization: Digest username="352-S_C_Sputnik",realm="asterisk",nonce="6a7468ee",uri="sip:10.220.0.6",response="1864a03e13ada1fa380206216a175d32",algorithm=MD5

//type Head_map map[string][]string
type Head_map map[string]string

//var head_map = make(map[string][]string)

func (msg *Sip_msg) Get_CSeq() string {
	space_back_pos := strings.LastIndex(msg.Headers["CSeq"], " ")
	return msg.Headers["CSeq"][space_back_pos:]
}

func (msg *Sip_msg) Get_from_host() string {
	if _, ok := msg.Headers["From"]; ok != true {
		return ""
	}

	fs := msg.Headers["From"]
	if start := strings.Index(fs, "@"); start > 0 {
		if end := strings.Index(fs[start:], ">"); end > 3 {
			return fs[start+1 : start+end]
		} else if end := strings.Index(fs[start:], ";"); end > 3 {
			return fs[start+1 : start+end]
		}
	}
	return ""
}

func Get_sip_msg(s string) (Sip_msg, bool) {
	end := strings.Index(s, "\r\n\r\n")
	h := Head_map{}
	new_sip_msg := Sip_msg{}
	if end < 4 {
		return new_sip_msg, false
	}
	headers_slice := strings.Split(s[:end], "\r\n")
	first_line_slice := strings.Fields(headers_slice[0])
	//fmt.Println("------first_line", first_line_slice)
	if len(first_line_slice) >= 3 && first_line_slice[2] == "SIP/2.0" {
		//fmt.Println("this is request")
		new_sip_msg.Sip_type = "REQUEST"
		new_sip_msg.Method_or_Code = first_line_slice[0]
		new_sip_msg.First_line = strings.Join(first_line_slice, " ")
	} else if first_line_slice[0] == "SIP/2.0" {
		//fmt.Println("this is response")
		new_sip_msg.Sip_type = "RESPONSE"
		new_sip_msg.Method_or_Code = first_line_slice[1]
		new_sip_msg.First_line = strings.Join(first_line_slice, " ")
	} else {
		return new_sip_msg, false
		//fmt.Println("this is not sip")
	}
	/*
		for i, _ := range headers_slice {
			header_string_slice := strings.Fields(headers_slice[i])
			h[header_string_slice[0]] = header_string_slice[1:]
		}
		if _, ok := h["Call-ID:"]; ok {
			new_sip_msg.Call_id = h["Call-ID:"][0]
		} else if _, ok := h["i:"]; ok {
			new_sip_msg.Call_id = h["i:"][0]
		} else {
			return new_sip_msg, false
		}
		//new_sip_msg.Call_id = h["Call-ID:"][0]
		new_sip_msg.Headers = h
		//fmt.Println(new_sip_msg)
		return new_sip_msg, true
	*/
	for i := 1; i < len(headers_slice); i++ {
		two_dots := strings.Index(headers_slice[i], ":")
		header := headers_slice[i][:two_dots]
		h[header] = headers_slice[i][two_dots:]

	}
	if _, ok := h["Call-ID"]; ok {
		new_sip_msg.Call_id = h["Call-ID"]
	} else if _, ok := h["i"]; ok {
		new_sip_msg.Call_id = h["i"]
	} else {
		return new_sip_msg, false
	}
	new_sip_msg.Headers = h
	return new_sip_msg, true
}
