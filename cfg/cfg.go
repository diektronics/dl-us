package cfg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	cfgpb "github.com/diektronics/dl-us/protos/cfg"
	"github.com/golang/protobuf/proto"
)

// stat is used to easily stub the function in unittests.
var stat = os.Stat

func GetConfig(cfgFile string) (*cfgpb.Config, error) {
	content, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("GetConfig Readfile: %v", err)
	}
	c := &cfgpb.Config{}
	if err := proto.UnmarshalText(string(content), c); err != nil {
		return nil, fmt.Errorf("GetConfig Unmarshal: %v", err)
	}

	if err := validate(c); err != nil {
		return nil, fmt.Errorf("GetConfig: Invalid configuration file: %v", err)
	}

	return c, nil
}

func validate(c *cfgpb.Config) error {
	allErrors := []string{}

	cv := reflect.ValueOf(*c)
	ct := reflect.TypeOf(*c)
	for i := 0; i < cv.NumField(); i++ {
		if cv.Field(i).Elem().Kind() == reflect.Struct {
			s := cv.Field(i).Elem()
			st := cv.Field(i).Elem().Type()
			for j := 0; j < s.NumField(); j++ {
				if s.Field(j).Kind() == reflect.String {
					if len(s.Field(j).String()) == 0 {
						allErrors = append(allErrors, fmt.Sprintf("%s.%s cannot be empty", ct.Field(i).Name, st.Field(j).Name))
					} else if strings.HasSuffix(st.Field(j).Name, "Path") {
						fmt.Println(st.Field(j).Name)
						if fi, err := stat(s.Field(j).String()); os.IsNotExist(err) {
							allErrors = append(allErrors, fmt.Sprintf("%q does not exist", s.Field(j).String()))
						} else if fi.Mode().Perm()&0111 == 0 {
							allErrors = append(allErrors, fmt.Sprintf("%q is not executable", s.Field(j).String()))
						}
					}
				}
			}
		}
	}

	if len(allErrors) != 0 {
		return errors.New(strings.Join(allErrors, ", "))
	}

	return nil
}
