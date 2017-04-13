package cfg

import (
	"errors"
	"os"
	"testing"

	cfgpb "github.com/diektronics/dl-us/protos/cfg"
	"github.com/prashantv/gostub"
)

type fakeFileInfo struct {
	os.FileInfo

	mode os.FileMode
}

func (fi *fakeFileInfo) Mode() os.FileMode { return fi.mode }

func TestValidate(t *testing.T) {
	var (
		goodDb = &cfgpb.Config_Db{
			User:     "user",
			Server:   "server",
			Password: "password",
			Database: "database",
		}
		badDb = &cfgpb.Config_Db{
			User:     "",
			Server:   "server",
			Password: "password",
			Database: "database",
		}
		goodMail = &cfgpb.Config_Mail{
			Addr:      "addr",
			Port:      42,
			Recipient: "recipient@example.com",
			Sender:    "sender@example.com",
			Password:  "password",
		}
		badMail = &cfgpb.Config_Mail{
			Port:      42,
			Recipient: "recipient@example.com",
			Sender:    "sender@example.com",
			Password:  "password",
		}
		goodDownload = &cfgpb.Config_Download{
			Dir:           "/path/to/download",
			PlowdownPath:  "/path/to/plowdown",
			PlowprobePath: "/path/to/plowprobe",
			LinkRegexp:    "some regexp",
			Feed:          "http://example.com/feed",
		}
		goodBackend = &cfgpb.Config_Backend{Port: 43}
		goodWeb     = &cfgpb.Config_Web{Port: 44}
		tests       = []struct {
			desc  string
			input *cfgpb.Config
			mode  os.FileMode
			fiErr error
			want  error
		}{
			{
				desc: "all is good",
				input: &cfgpb.Config{
					Db:       goodDb,
					Mail:     goodMail,
					Download: goodDownload,
					Backend:  goodBackend,
					Web:      goodWeb,
				},
				mode: 0777,
				want: nil,
			},
			{
				desc: "plowdown, plowprobe do not exist",
				input: &cfgpb.Config{
					Db:       goodDb,
					Mail:     goodMail,
					Download: goodDownload,
					Backend:  goodBackend,
					Web:      goodWeb,
				},
				fiErr: os.ErrNotExist,
				want:  errors.New(`"/path/to/plowdown" does not exist, "/path/to/plowprobe" does not exist`),
			},
			{
				desc: "plowdown, plowprobe are not executable",
				input: &cfgpb.Config{
					Db:       goodDb,
					Mail:     goodMail,
					Download: goodDownload,
					Backend:  goodBackend,
					Web:      goodWeb,
				},
				mode: 0,
				want: errors.New(`"/path/to/plowdown" is not executable, "/path/to/plowprobe" is not executable`),
			},
			{
				desc: "Db is missing fields",
				input: &cfgpb.Config{
					Db:       badDb,
					Mail:     goodMail,
					Download: goodDownload,
					Backend:  goodBackend,
					Web:      goodWeb,
				},
				mode: 0777,
				want: errors.New("Db.User cannot be empty"),
			},
			{
				desc: "Db and Mail are missing fields",
				input: &cfgpb.Config{
					Db:       badDb,
					Mail:     badMail,
					Download: goodDownload,
					Backend:  goodBackend,
					Web:      goodWeb,
				},
				mode: 0777,
				want: errors.New("Db.User cannot be empty, Mail.Addr cannot be empty"),
			},
			{
				desc: "Db and Mail are missing fields and plow bins are not executable",
				input: &cfgpb.Config{
					Db:       badDb,
					Mail:     badMail,
					Download: goodDownload,
					Backend:  goodBackend,
					Web:      goodWeb,
				},
				mode: 0,
				want: errors.New(`Db.User cannot be empty, Mail.Addr cannot be empty, "/path/to/plowdown" is not executable, "/path/to/plowprobe" is not executable`),
			},
		}
		fi    = &fakeFileInfo{}
		fiErr error
		stubs = gostub.Stub(&stat, func(string) (os.FileInfo, error) { return fi, fiErr })
	)
	defer stubs.Reset()

	for _, ts := range tests {
		fi.mode = ts.mode
		fiErr = ts.fiErr
		if got, want := validate(ts.input), ts.want; !errMatch(got, want) {
			t.Errorf("%s: unexpected error, got: %s, want: %s", ts.desc, got, want)
		}
	}
}

func errMatch(got, want error) bool {
	if got == nil && want != nil {
		return false
	}
	if got != nil && want == nil {
		return false
	}
	if got != nil && want != nil && got.Error() != want.Error() {
		return false
	}

	return true
}
