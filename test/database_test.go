package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"light-cloud/src/core/model"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "define.DataSourceName")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*model.UserInfo, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dst.String())
}
