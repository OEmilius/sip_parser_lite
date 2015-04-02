package sip_parser_lite

import (
	"fmt"
	"honnef.co/go/pcap"
	"os"
	"testing"
)

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

var responce string = "SIP/2.0 200 OK\r\n" +
	"Via: SIP/2.0/UDP bobspc.biloxi.com:5060;branch=z9hG4bKnashds7;received=192.0.2.4\r\n" +
	"To: Bob <sip:bob@biloxi.com>;tag=2493k59kd\r\n" +
	"From: Bob <sip:bob@biloxi.com>;tag=456248\r\n" +
	"Call-ID: 843817637684230@998sdasdh09\r\n" +
	"CSeq: 1826 REGISTER\r\n" +
	"Contact: <sip:bob@192.0.2.4>\r\n" +
	"Expires: 7200\r\n" +
	"Content-Length: 0\r\n\r\n"

var code_404 string = "SIP/2.0 404 Not Found\r\n" +
	"v: SIP/2.0/UDP 217.67.176.213:5060;branch=z9hG4bK8n8lnkevioynhbehuljhkxjci;received=217.67.176.213\r\n" +
	"f: <sip:217.67.176.213>;tag=6e8e9b4a\r\n" +
	"t: <sip:217.197.125.46;user=phone>;tag=as6891f9bf\r\n" +
	"i: SBCf7d8ea2f9ca93609520172b0b1863fbb@10.200.66.5\r\n" +
	"CSeq: 1 OPTIONS\r\n" +
	"Server: Asterisk PBX\r\n" +
	"Allow: INVITE, ACK, CANCEL, OPTIONS, BYE, REFER, SUBSCRIBE, NOTIFY, INFO, PUBLISH\r\n" +
	"k: replaces, timer\r\n" +
	"Accept: application/sdp\r\n" +
	"l: 0\r\n" +
	"\r\n\r\n"

func Test_request(t *testing.T) {
	//if sip_message, ok := Get_sip_msg(request); ok {
	if sip_message, ok := Get_sip_msg(code_404); ok {
		//fmt.Println("request line is ", sip_message.Headers[0])
		fmt.Println("type is ", sip_message.Sip_type)
		fmt.Println("method or code is", sip_message.Method_or_Code)
		fmt.Println("call_id is ", sip_message.Call_id)
		fmt.Println("from", sip_message.Headers["From"])
		fmt.Println("to", sip_message.Headers["To"])
		fmt.Println("CSeq", sip_message.Headers["CSeq"])
		fmt.Println(sip_message.Get_CSeq())

	} else {
		fmt.Println("this is not sip")
	}
}

func Test_response(t *testing.T) {
	if sip_message, ok := Get_sip_msg(responce); ok {
		fmt.Println("type is ", sip_message.Sip_type)
		fmt.Println("status_line:", sip_message.First_line)
		fmt.Println("method or code is", sip_message.Method_or_Code)
		fmt.Println("call_id is ", sip_message.Call_id)
		fmt.Println("from", sip_message.Headers["From"]) //но помни что возвращает тут слайс
		fmt.Println(sip_message.Get_from_host())
		fmt.Println("to", sip_message.Headers["To"])     //но помни что возвращает тут слайс
		fmt.Println("CSeq", sip_message.Headers["CSeq"]) //а тут уже строку
		fmt.Println(sip_message.Get_CSeq())
	}
}

func Test_readingbfile(t *testing.T) {
	if f, err := os.Open("big_sip.pcap"); err == nil {
		r := pcap.NewReader(f)
		r.ParseHeader()
		for i := 1; ; i++ {
			p, err := r.ReadPacket()
			if err != nil {
				fmt.Println(err)
				break // пакет не читается
			}
			may_be_sip := string(p.Data.Payload())
			real_sip := may_be_sip[42:] //отрезаем до эмпирически 42
			if s_msg, ok := Get_sip_msg(real_sip); ok {
				//t.Log("this packet ok", i)
				//fmt.Println("call_id :", s_msg.Call_id)
				fmt.Println("fh:", s_msg.Get_from_host(), "n:", i)
			} else {
				//t.Error("this is not sip packet N: ", i)
				//fmt.Println("this is not sip packet N: ", i)
			}
		}
		f.Close()
	} else {
		fmt.Println("cant open file")
		//t.Error("can't open file", err)
	}
}
