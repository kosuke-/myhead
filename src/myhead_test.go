package main

import (
	"testing"

	log "bitbucket.org/groove-x/go-common/gxlog"
)

func TestHeadScanner(t *testing.T) {

	testpath := "../testdata/test1.txt"
	testnumber := 10

	retData, err := headScanner(testpath, testnumber)
	if err != nil {
		t.Fatalf("headScanner error: %+v", err)
	}
	require := "1\n22\n333\n4444\n55555\n666666\n7777777\n88888888\n999999999\n0000000000"
	if retData != require {
		log.Info(retData)
		t.Fatalf("headScanner retData error")

	}

	testpath = "../testdata/test2.txt"
	testnumber = 1

	retData, err = headScanner(testpath, testnumber)
	if err != nil {
		t.Fatalf("headScanner error: %+v", err)
	}
	require = "88888888"
	if retData != require {
		log.Info(retData)
		t.Fatalf("headScanner retdata error")

	}

	testpath = "../testdata/test3.txt"
	testnumber = 0

	retData, err = headScanner(testpath, testnumber)
	if err != nil {
		t.Fatalf("headScanner error: %+v", err)
	}
	require = ""
	if retData != require {
		log.Info(retData)
		t.Fatalf("headScanner retdata error")

	}

	testpath = "dammy.txt"
	testnumber = 10

	retData, err = headScanner(testpath, testnumber)
	if err == nil {
		t.Fatalf("headScanner is not error")
	}
	require = ""
	if retData != require {
		t.Fatalf("headScanner retdata error")

	}
}

func TestHeadPrint(t *testing.T) {

	data := make(map[string]string)
	data["head1"] = "1\n22\n333"
	data["head2"] = "333\n\n1"
	data["head3"] = "4444"

	headPrint(data)
}
