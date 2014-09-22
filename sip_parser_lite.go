package sip_parser_lite

import (
	//"fmt"
	"strings"
)

/*
var request string = "INVITE sip:marconi@radio.org SIP/2.0\r\n" +
	"Max-Forwards: 70\r\n" +
	"To: G. Marconi <sip:Marconi@radio.org>\r\n" +
	"From: Nikola Tesla <sip:n.tesla@high-voltage.org>;tag=76341\r\n" +
	"Call-ID: 123456789@lab.high-voltage.org\r\n" +
	"Contact: <sip:n.tesla@lab.high-voltage.org>\r\n" +
	"CSeq: 1 INVITE\r\n" +
	"Subject: About That Power Outage...\r\n" +
	"Content-Type: application/sdp\r\n" +
	"Content-Length: 158\r\n" +
	"\r\n" +
	"v=0\r\n" +
	"o=Tesla 2890844526 2890844526 IN IP4 lab.high-voltage.org\r\n" +
	"s=Phone Call\r\n" +
	"c=IN IP4 100.101.102.103\r\n" +
	"t=0 0\r\n" +
	"m=audio 49170 RTP/AVP 0\r\n" +
	"a=rtpmap:0 PCMU/8000\r\n"
	//request = ""
var responce string = "SIP/2.0 200 OK\r\n" +
	"Via: SIP/2.0/UDP bobspc.biloxi.com:5060;branch=z9hG4bKnashds7;received=192.0.2.4\r\n" +
	"To: Bob <sip:bob@biloxi.com>;tag=2493k59kd\r\n" +
	"From: Bob <sip:bob@biloxi.com>;tag=456248\r\n" +
	"Call-ID: 843817637684230@998sdasdh09\r\n" +
	"CSeq: 1826 REGISTER\r\n" +
	"Contact: <sip:bob@192.0.2.4>\r\n" +
	"Expires: 7200\r\n" +
	"Content-Length: 0\r\n\r\n"
*/

type Sip_msg struct {
	Sip_type       string //response or request
	Method_or_Code string //200, 404, INVITE, PRACK
	Call_id        string
	Headers        Head_map
}

//Authorization: Digest username="352-S_C_Sputnik",realm="asterisk",nonce="6a7468ee",uri="sip:10.220.0.6",response="1864a03e13ada1fa380206216a175d32",algorithm=MD5

type Head_map map[string][]string

//var head_map = make(map[string][]string)

func (msg *Sip_msg) Get_CSeq() string {
	return msg.Headers["CSeq:"][1]
}

/*
func main() {
	fmt.Println("start")
	//get_headers_map(request)
	//get_headers_map(responce)

	if sip_message, ok := Get_sip_msg(request); ok {
		//fmt.Println("request line is ", sip_message.Headers[0])
		fmt.Println("type is ", sip_message.Sip_type)
		fmt.Println("method or code is", sip_message.Method_or_Code)
		fmt.Println("call_id is ", sip_message.Call_id)
		fmt.Println("from", sip_message.Headers["From:"])    //но помни что возвращает тут слайс
		fmt.Println("to", sip_message.Headers["To:"])        //но помни что возвращает тут слайс
		fmt.Println("CSeq", sip_message.Headers["CSeq:"][1]) //а тут уже строку
		fmt.Println(sip_message.Get_CSeq())

	}
	if sip_message, ok := Get_sip_msg(responce); ok {
		fmt.Println("type is ", sip_message.Sip_type)
		fmt.Println("method or code is", sip_message.Method_or_Code)
		fmt.Println("call_id is ", sip_message.Call_id)
		fmt.Println("CSeq", sip_message.Headers["CSeq:"][1]) //а тут уже строку
		fmt.Println(sip_message.Get_CSeq())
	}
}
*/
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
	if len(first_line_slice) == 3 && first_line_slice[2] == "SIP/2.0" {
		//fmt.Println("this is request")
		new_sip_msg.Sip_type = "REQUEST"
		new_sip_msg.Method_or_Code = first_line_slice[0]
	} else if first_line_slice[0] == "SIP/2.0" {
		//fmt.Println("this is response")
		new_sip_msg.Sip_type = "RESPONSE"
		new_sip_msg.Method_or_Code = first_line_slice[1]
	} else {
		return new_sip_msg, false
		//fmt.Println("this is not sip")

	}

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

}
