package generator

import (
	"math/rand"
	"strings"
	"time"
)

var lipsumwords = []string{"a", "ac", "accumsan", "ad", "adipiscing", "aenean", "aliquam", "aliquet",
	"amet", "ante", "aptent", "arcu", "at", "auctor", "augue", "bibendum",
	"blandit", "class", "commodo", "condimentum", "congue", "consectetur",
	"consequat", "conubia", "convallis", "cras", "cubilia", "cum", "curabitur",
	"curae", "cursus", "dapibus", "diam", "dictum", "dictumst", "dignissim",
	"dis", "dolor", "donec", "dui", "duis", "egestas", "eget", "eleifend",
	"elementum", "elit", "enim", "erat", "eros", "est", "et", "etiam", "eu",
	"euismod", "facilisi", "facilisis", "fames", "faucibus", "felis",
	"fermentum", "feugiat", "fringilla", "fusce", "gravida", "habitant",
	"habitasse", "hac", "hendrerit", "himenaeos", "iaculis", "id", "imperdiet",
	"in", "inceptos", "integer", "interdum", "ipsum", "justo", "lacinia",
	"lacus", "laoreet", "lectus", "leo", "libero", "ligula", "litora",
	"lobortis", "lorem", "luctus", "maecenas", "magna", "magnis", "malesuada",
	"massa", "mattis", "mauris", "metus", "mi", "molestie", "mollis", "montes",
	"morbi", "mus", "nam", "nascetur", "natoque", "nec", "neque", "netus",
	"nibh", "nisi", "nisl", "non", "nostra", "nulla", "nullam", "nunc", "odio",
	"orci", "ornare", "parturient", "pellentesque", "penatibus", "per",
	"pharetra", "phasellus", "placerat", "platea", "porta", "porttitor",
	"posuere", "potenti", "praesent", "pretium", "primis", "proin", "pulvinar",
	"purus", "quam", "quis", "quisque", "rhoncus", "ridiculus", "risus",
	"rutrum", "sagittis", "sapien", "scelerisque", "sed", "sem", "semper",
	"senectus", "sit", "sociis", "sociosqu", "sodales", "sollicitudin",
	"suscipit", "suspendisse", "taciti", "tellus", "tempor", "tempus",
	"tincidunt", "torquent", "tortor", "tristique", "turpis", "ullamcorper",
	"ultrices", "ultricies", "urna", "ut", "varius", "vehicula", "vel", "velit",
	"venenatis", "vestibulum", "vitae", "vivamus", "viverra", "volutpat",
	"vulputate"}

var punctuation = []string{".", "?", "!"}

// var severity = []string{"ERROR", "WARN", "INFO", "DEBUG", "TRACE", ""}
var severity = []string{"INFO", "DEBUG", "TRACE"} // NOTE: This is the pretty view (everything alright! No problems!)

// Initialize must be called first before using this generator
func Initialize() {
	rand.Seed(time.Now().Unix())
}

// SetPositiveSeveritiesOnly forces the severity level of log lines to not have any errors or warnings
// This is the pretty view (everything alright! No problems!)
func SetPositiveSeveritiesOnly() {
	severity = []string{"INFO", "DEBUG", "TRACE"}
}

// SetAllSeverities opens up the log lines to have all possible severity types
func SetAllSeverities() {
	severity = []string{"ERROR", "WARN", "INFO", "DEBUG", "TRACE", ""}
}

// RandomSeverity returns a random severity level
// Examples could be ERROR or INFO
func RandomSeverity() string {
	return severity[rand.Intn(len(severity))]
}

// RandomWord returns a random lorem ipsum word
// Examples could be "varius" or "hendrerit"
func RandomWord() string {
	return lipsumwords[rand.Intn(len(lipsumwords))]
}

// RandomPunctuation returns a random punctuation
// Examples could be "." or "!"
func RandomPunctuation() string {
	return punctuation[rand.Intn(len(punctuation))]
}

// Words returns a string with a list of random lorem ipsum words that is count words long
// Example when count is 3: "maecenas ornare dapibus"
func Words(count int) string {
	if count > 0 {
		return strings.TrimSpace(RandomWord() + " " + Words(count-1))
	} else {
		return ""
	}
}

// SentenceFragment returns a string of words that could make up part of a sentence
// Most sentence fragments have at least 3 words
// But they could be larger, so this function could return up to 13 words in a sentence fragment
// Example: "non ridiculus integer iaculis iaculis"
func SentenceFragment() string {
	return Words(rand.Intn(10) + 3)
}

// Sentence constructs a full sentence, with fragments connected by commas, capitalization, and proper punctuation.
// Example: "Non ridiculus integer iaculis iaculis, consectetur dapibus consectetur viverra!"
func Sentence() string {
	s := strings.Title(RandomWord()) + " "
	if rand.Intn(2) == 0 {
		for i := 0; i < rand.Intn(3); i++ {
			s += SentenceFragment() + ", "
		}
	}
	return SentenceFragment() + RandomPunctuation()
}

// Sentences returns a string of count sentences
// Example with count set to 2: "Non ridiculus integer iaculis iaculis, consectetur dapibus consectetur viverra! Consequat at potenti risus."
func Sentences(count int) string {
	if count > 0 {
		return Sentence() + " " + strings.TrimSpace(Sentences(count-1))
	} else {
		return ""
	}
}

// Paragraph returns a series of sentences to reasonably create a paragraph
// Most paragraphs at least have two sentences
// A large paragraph could have around 12
func Paragraph() string {
	return Sentences(rand.Intn(10) + 2)
}
