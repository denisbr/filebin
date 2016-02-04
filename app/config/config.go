package config

type Configuration struct {
	Host                string
	Port                int
	Readtimeout         int
	Writetimeout        int
	Maxheaderbytes      int
	Filedir             string
	Baseurl             string
	Tempdir             string
	Expiration          int64
	TriggerNewTag       string
	TriggerUploadFile   string
	TriggerDeleteTag    string
	TriggerDeleteFile   string
	TriggerExpireTag    string
	DefaultTagLength    int
	Workers             int
	Version             bool
	CacheInvalidation   bool
}

var Global Configuration

func init() {
	Global = Configuration{
		Host:           "127.0.0.1",
		Port:           31337,
		Readtimeout:    3600,
		Writetimeout:   3600,
		Maxheaderbytes: 1 << 20,
		// 7776000 = 3 months
		Expiration:          7776000,
		Baseurl:             "http://localhost:31337",
		Filedir:             "/srv/filebin/files",
		Tempdir:             "/tmp",
		DefaultTagLength:    16,
		Workers:             1,
		CacheInvalidation:   false,
	}
}
