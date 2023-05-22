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

func Initialize() {
	rand.Seed(time.Now().Unix())
}

func SetPositiveSeveritiesOnly() {
	severity = []string{"INFO", "DEBUG", "TRACE"}
}

func SetAllSeverities() {
	severity = []string{"ERROR", "WARN", "INFO", "DEBUG", "TRACE", ""}
}

func RandomSeverity() string {
	return severity[rand.Intn(len(severity))]
}

func RandomWord() string {
	return lipsumwords[rand.Intn(len(lipsumwords))]
}

func RandomPunctuation() string {
	return punctuation[rand.Intn(len(punctuation))]
}

func Words(count int) string {
	if count > 0 {
		return strings.TrimSpace(RandomWord() + " " + Words(count-1))
	} else {
		return ""
	}
}

func SentenceFragment() string {
	return Words(rand.Intn(10) + 3)
}

func Sentence() string {
	s := strings.Title(RandomWord()) + " "
	if rand.Intn(2) == 0 {
		for i := 0; i < rand.Intn(3); i++ {
			s += SentenceFragment() + ", "
		}
	}
	return SentenceFragment() + RandomPunctuation()
}

func Sentences(count int) string {
	if count > 0 {
		return Sentence() + " " + strings.TrimSpace(Sentences(count-1))
	} else {
		return ""
	}
}

func Paragraph() string {
	return Sentences(rand.Intn(10) + 2)
}
