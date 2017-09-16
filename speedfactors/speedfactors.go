// ibmservicecreds converts JSON-formtted IBM creds into a Golang struct
package speedfactors

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type Speedfactors struct {
	Diarizer    map[string]float64 `json:"diarizer"`
	Transcriber map[string]float64 `json:"transcriber"`
}

func ReadCredsFromPath(path string) (Speedfactors, error) {
	file, err := os.Open(path)
	if err != nil {
		return Speedfactors{}, err
	}
	return ReadCredsFromReader(file)
}
func ReadCredsFromReader(r io.Reader) (creds Speedfactors, err error) {
	buf := []byte{}
	buf, err = ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(buf, &creds)
	if err != nil {
		panic(err)
	}
	return creds, err
}

// the following functions return time.Duration whose underlying type is int64 (nanosecs),
func (sf Speedfactors) EstimatedDiarizationTime(diarizer string, soundFileSizeinBytes int64) time.Duration {
	if val, ok := sf.Diarizer[diarizer]; ok {
		return time.Duration(int64(val*float64(soundFileSizeinBytes)*1e9) / 32000)
	} else {
		panic(fmt.Sprintf("I couldn't find Diarizer[\"%s\"]!", diarizer))
	}
}

func (sf Speedfactors) EstimatedTranscriptionTime(transcriber string, soundFileSizeinBytes int64) time.Duration {
	if val, ok := sf.Transcriber[transcriber]; ok {
		return time.Duration(int64(val*float64(soundFileSizeinBytes)*1e9) / 32000)
	} else {
		panic(fmt.Sprintf("I couldn't find Transcriber[\"%s\"]!", transcriber))
	}
}