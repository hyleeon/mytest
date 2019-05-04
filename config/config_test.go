package config

import (
	"fmt"
	"testing"
)


func TestConfig(t *testing.T) {
	var conf = Config{}
	var value = conf.GetValue("abc");
	fmt.Println("conf get value:", value);
}