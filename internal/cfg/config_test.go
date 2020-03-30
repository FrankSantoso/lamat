package cfg

import (
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

type testConf struct {
	name string
	path string
}

func TestParseConfig(t *testing.T) {
	t.Log("Given the need to parse config file")
	{
		testConfigsErr := []testConf{
			{
				name: "Nonexistent path",
				path: "./test-configs/testconf",
			},
		}
		testConfigSuccess := []testConf{
			{
				name: "Return default config on empty file",
				path: "./test-configs/testenv",
			},
		}
		t.Logf("\tTest for problematic config file")
		{
			for _, tc := range testConfigsErr {
				f := func(t *testing.T) {
					conf, err := ReadConfig(tc.path)
					if err == nil {
						t.Fatalf("\t%s\tShould return error on: %s, config: %v",
							failed, tc.name, conf)
					}
					t.Logf("\t%s\tShould return error on: %s => %v",
						success, tc.name, err)
				}
				t.Run(tc.name, f)
			}
		}
		t.Logf("\tTest for good configs")
		{
			for _, tc := range testConfigSuccess {
				f := func(t *testing.T) {
					conf, err := ReadConfig(tc.path)
					if err != nil {
						t.Fatalf("\t%s\tShould be able to return good config"+
							": %s on: %s", failed, tc.name, tc.path)
					}
					t.Logf("\t%s\tShould be able to return good config"+
						": %s on %v", success, tc.name, conf)
				}
				t.Run(tc.name, f)
			}
		}
	}
}
