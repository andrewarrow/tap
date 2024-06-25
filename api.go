package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetData() {
	url := "https://api.elevenlabs.io/v1/text-to-speech/GBv7mTt0atIp3Br8iCZE/stream/with-timestamps"

	reader := strings.NewReader(example)
	req, _ := http.NewRequest("POST", url, reader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("xi-api-key", os.Getenv("E_API"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	asBytes, _ := io.ReadAll(resp.Body)
	asString := string(asBytes)
	fmt.Println(asString)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response from server:", resp.Status)
		return
	}

}

var example = `{
  "text": "For all the praises that are sung about how influential and accelerating computer science has been, my favorite perspective is that today, for any given field, there is a good chance that of all practices or techniques that have accelerated knowledge discovery within the field in the last few decades, computer science methods present a huge chunk of it. Another perspective I really love, is that computer science might be the set that is the highest common factor of pretty much all sciences that we have today. All of physics, biology, neuroscience saw a huge upthrust in progress fueled by the ability to run useful simulations and models.
Once you have a software-based abstraction over any study, it becomes incredibly easier to iterate and generate new hypotheses and understandings. So to say, computer science became a miraculous catalyst that fastens pretty much every single reaction.
In my opinion, The only analogues that produced a comparable effect are math and language.
Cross-pollination is one of the most beautiful meta-scientific phenomena that exists today, whenever a shiny new theory with a widely-effective and high explanatory theory comes around, it enables us to start taking specific issues, solve them with such a theory, and then extrapolate or generalise.

I believe we can do something similar if we get really serious and adept with Complex Systems.
Great progress in Complex Systems would imply a few things:
1. There is huge positive transfer.
Even if the two fields are separate in their own respect, modelling general behavior of how agents interact with each other, how incentives are aligned, and what really leads to "non-linearity" in the entire system can be fairly useful for all other fields too..

Once a nuanced understanding of some phenomena is developed using methods, it isn't difficult to quickly generalise and develop frameworks that would have broader applicability.

-Something like Figuring out how the human brain goes about self-regulation, bootstrapping would give us incredibly useful perspectives on developing generally intelligent machines.

-Figuring out how biological self-assembly or just self-assembly at any level would give us huge insights on how "emergence" really works, and why everything exists in a holonic sense i.e. a "whole" in its own is composed of many wholes themselves, and goes on to partake in a bigger "whole".

Subsequently, we may generalise these and go on to study interdependence in the layers of those hierarchies, figure out what incentives really lead to such structures, what dictates systemic behavior.

It's easy to imagine such an approach being useful in the study of macro-systems like big-scale economies, all while understanding what the actions of smaller entities imply at the broadest levels.
- We might go on to understand human behavior by figuring out how incentives stack together to influence strategies for agents in Systems small enough to be interpretable yet being composed of entities with large degrees of freedom and a general sense of non-linearity and stochasticness.
These are a few examples off the top of my head, another really interesting point is that we might be able to solve problems that at least today, due to a lack of theory or general consensus, we might deem impossibly hard to understand.

Consider some general problems we have today across a set of fields:
-How do we effectively model open free-markets?
-How does the body react to imbalances in some internally-relevant biomarkers?
-How do we effectively develop more effective climate models?
-How do we develop extremely effective precision medicine by understanding micro-effects on cellular mechanisms and pathways? ",
  "model_id": "eleven_multilingual_v2",
  "voice_settings": {
    "stability": 123,
    "similarity_boost": 123,
    "style": 123,
    "use_speaker_boost": true
  },
  "seed": 123238723
}`
