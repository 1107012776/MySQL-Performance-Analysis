package main

import (
	"MySQL-PA/mysql"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_RootConfigDir(t *testing.T) {
	p := mysql.Explain{}
	assert.Equal(t, p.Analyze(), true)
}
