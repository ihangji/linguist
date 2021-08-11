/*
Copyright 2020-2021 Ihangji, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package linguist

import (
	"bytes"
	"log"
	"math"

	"github.com/ihangji/linguist/data"
	"github.com/ihangji/linguist/tokenizer"
	"github.com/jbrukh/bayesian"
)

var classifier *bayesian.Classifier
var classifierInitialized = false

// Gets the baysian.Classifier which has been trained on programming language
// samples from github.com/github/linguist after running the generator
//
// See also cmd/generate-classifier
func getClassifier() *bayesian.Classifier {
	// NOTE(tso): this could probably go into an init() function instead
	// but this lazy loading approach works, and it's conceivable that the
	// analyse() function might not invoked in an actual runtime anyway
	if !classifierInitialized {
		d, err := data.Asset("classifier")
		if err != nil {
			log.Panicln(err)
		}
		reader := bytes.NewReader(d)
		classifier, err = bayesian.NewClassifierFromReader(reader)
		if err != nil {
			log.Panicln(err)
		}
		classifierInitialized = true
	}
	return classifier
}

// Analyse returns the name of a programming language, or the empty string if one could
// not be determined.
//
// Uses Naive Bayesian Classification on the file contents provided.
//
// It is recommended to use LanguageByContents() instead of this function directly.
//
// Obtain hints from LanguageHints()
//
// NOTE(tso): May yield inaccurate results
func Analyse(contents []byte, hints []string) (language string) {
	document := tokenizer.Tokenize(contents)
	classifier := getClassifier()
	scores, idx, _ := classifier.LogScores(document)

	if len(hints) == 0 {
		return string(classifier.Classes[idx])
	}

	langs := map[string]struct{}{}
	for _, hint := range hints {
		langs[hint] = struct{}{}
	}

	bestScore := math.Inf(-1)
	bestAnswer := ""

	for id, score := range scores {
		answer := string(classifier.Classes[id])
		if _, ok := langs[answer]; ok {
			if score >= bestScore {
				bestScore = score
				bestAnswer = answer
			}
		}
	}
	return bestAnswer
}
