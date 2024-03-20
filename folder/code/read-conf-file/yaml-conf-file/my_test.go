package yconf

import (
	"testing"
)

func TestInitSetting(t *testing.T) {

	if err := InitSetting(); err != nil {
		t.Error(err)
	}

	t.Logf("%+v \n", Conf)
}
