// Copyright © 2018 Kent Gibson <warthog618@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/warthog618/sms/encoding/tpdu"
	"github.com/warthog618/sms/ms/message"
	"github.com/warthog618/sms/ms/pdumode"
	"github.com/warthog618/sms/ms/sar"
)

func main() {
	var pm bool
	flag.BoolVar(&pm, "p", false, "PDU is prefixed with SCA (PDU mode)")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	udd, err := tpdu.NewUDDecoder()
	if err != nil {
		log.Fatal(err)
	}
	udd.AddAllCharsets()
	c := sar.NewCollector(time.Minute*5, func(arg1 error) {})
	x := message.NewReassembler(udd, c)
	defer x.Close()
	for _, a := range flag.Args() {
		b, err := hex.DecodeString(a)
		if err != nil {
			log.Fatal(err)
		}
		tb := b
		if pm {
			pd := pdumode.Decoder{}
			_, ntb, err := pd.Decode(b)
			if err != nil {
				log.Fatal(err)
			}
			tb = ntb
		}
		msg, err := x.Reassemble(tb)
		if err != nil {
			log.Printf("reassembly error: %v", err)
		}
		if msg != nil {
			fmt.Printf("%s: %s\n", msg.Number, msg.Msg)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "smsdeliver decodes and displays the message from one or more SMS Deliver TPDUs.\n"+
		"Usage: smsdeliver [-p] <pdu> [pdu...]\n")
	flag.PrintDefaults()
}
