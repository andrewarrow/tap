package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetData() {
	url := "https://api.elevenlabs.io/v1/text-to-speech/GBv7mTt0atIp3Br8iCZE/stream/with-timestamps"

	readerPost := strings.NewReader(example)
	req, _ := http.NewRequest("POST", url, readerPost)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("xi-api-key", os.Getenv("E_API"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	//audio_base64

	reader := bufio.NewReader(resp.Body)

	var buffer bytes.Buffer

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading data:", err)
			return
		}

		fmt.Println(len(line))
		buffer.Write(line)
		var m map[string]any
		err = json.Unmarshal(buffer.Bytes(), &m)
		if err == nil {
			fmt.Println("m", len(m))
			buffer.Reset()
		}
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response from server:", resp.Status)
		return
	}

}

func processChunk(data []byte) {

}

var example = `{
  "text": "Oceans cover more than 70% of the Earth's surface and play a crucial role in regulating the planet's climate. These vast bodies of water are home to a diverse array of marine life, from tiny plankton to massive whales. Oceans are divided into five main regions: the Pacific, Atlantic, Indian, Southern, and Arctic. They are interconnected and influence weather patterns, ocean currents, and global temperatures. The health of the oceans is vital for the survival of many species, including humans. Overfishing, pollution, and climate change pose significant threats to marine ecosystems. Coral reefs, often called the \"rainforests of the sea,\" are particularly vulnerable. They provide habitat for countless species and protect coastlines from erosion. Efforts to conserve and protect the oceans are essential to maintaining biodiversity and ensuring sustainable resources for future generations. Oceanography, the study of the oceans, helps scientists understand the complex interactions within marine environments and develop strategies to address environmental challenges. The beauty and mystery of the oceans continue to inspire awe and wonder, reminding us of the importance of preserving these majestic natural resources.",
  "model_id": "eleven_multilingual_v2",
  "voice_settings": {
    "stability": 0.6,
    "similarity_boost": 0.6,
    "use_speaker_boost": true
  },
  "seed": 123238723
}`
