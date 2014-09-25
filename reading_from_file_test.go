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
	"\r\n"

var my_200_OK string = "SIP/2.0 200 OK\r\n" +
	"Via: SIP/2.0/UDP 10.100.100.123:7648;branch=z9hG4bK-d8754z-782db835a0f52820-1---d8754z-;received=83.69.212.78;rport=16798\r\n" +
	"Call-ID: OTlhNGZlZDE5ZmYzOWFlZGQ3MGEyOWM4YTY0ZDVjMzc\r\n" +
	"From: \"254-mastertel\"<sip:254-mastertel@mte.master-office.net>;tag=44ac1757\r\n" +
	"To: \"84957870777\"<sip:84957870777@89.106.175.198>;tag=as6d510d85\r\n" +
	"CSeq: 2 INVITE\r\n" +
	"Allow: INVITE,ACK,CANCEL,OPTIONS,BYE,REFER,SUBSCRIBE,NOTIFY,INFO,PUBLISH,UPDATE\r\n" +
	"Server: Master-office\r\n" +
	"Supported: replaces,timer\r\n" +
	"Contact: <sip:84957870777@89.106.175.198:5060;user=phone>\r\n" +
	"Content-Length: 263\r\n" +
	"Content-Type: application/sdp\r\n" +
	"\r\n" +
	"v=0\r\n" +
	"o=root 795608413 795608413 IN IP4 89.106.175.198\r\n" +
	"s=Asterisk PBX 1.8.26.0\r\n" +
	"c=IN IP4 89.106.175.198\r\n" +
	"t=0 0\r\n" +
	"m=audio 46288 RTP/AVP 8 0 101\r\n" +
	"a=rtpmap:8 PCMA/8000\r\n" +
	"a=rtpmap:0 PCMU/8000\r\n" +
	"a=rtpmap:101 telephone-event/8000\r\n" +
	"a=fmtp:101 0-16\r\n" +
	"a=ptime:20\r\n" +
	"a=sendrecv\r\n" +
	"\r\n"

func Test_request(t *testing.T) {
	//if sip_message, ok := Get_sip_msg(request); ok {
	//if sip_message, ok := Get_sip_msg(code_404); ok {
	if sip_message, ok := Get_sip_msg(my_200_OK); ok {
		//fmt.Println("request line is ", sip_message.Headers[0])
		fmt.Println("type is ", sip_message.Sip_type)
		fmt.Println("method or code is", sip_message.Method_or_Code)
		fmt.Println("call_id is ", sip_message.Call_id)
		fmt.Println("from", sip_message.Headers["From:"])    //но помни что возвращает тут слайс
		fmt.Println("to", sip_message.Headers["To:"])        //но помни что возвращает тут слайс
		fmt.Println("CSeq", sip_message.Headers["CSeq:"][1]) //а тут уже строку
		fmt.Println(sip_message.Get_CSeq())
	}
}

func Test_response(t *testing.T) {
	if sip_message, ok := Get_sip_msg(responce); ok {
		fmt.Println("type is ", sip_message.Sip_type)
		fmt.Println("method or code is", sip_message.Method_or_Code)
		fmt.Println("call_id is ", sip_message.Call_id)
		fmt.Println("CSeq", sip_message.Headers["CSeq:"][1]) //а тут уже строку
		fmt.Println(sip_message.Get_CSeq())
	}
}

func Test_readingfile(t *testing.T) {
	if f, err := os.Open("with_reg.pcap"); err == nil {
		r := pcap.NewReader(f)
		r.ParseHeader()
		for i := 0; ; i++ {
			p, err := r.ReadPacket()
			if err != nil {
				fmt.Println(err)
				break // пакет не читается
			}
			may_be_sip := string(p.Data.Payload())
			real_sip := may_be_sip[42:] //отрезаем до эмпирически 42
			if s_msg, ok := Get_sip_msg(real_sip); ok {
				t.Log("this packet ok")
				//fmt.Println("call_id :", s_msg.Call_id)
				fmt.Println("method or code:", s_msg.Method_or_Code)
			} else {
				t.Error("this is not sip packet N: ", i)
			}
		}
		f.Close()
	} else {
		fmt.Println("cant open file")
		t.Error("can't open file", err)
	}
}

/*
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
				fmt.Println("call_id :", s_msg.Call_id)
			} else {
				//t.Error("this is not sip packet N: ", i)
				fmt.Println("this is not sip packet N: ", i)
			}
		}
		f.Close()
	} else {
		fmt.Println("cant open file")
		//t.Error("can't open file", err)
	}
}
*/
