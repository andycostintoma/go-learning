package main

import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

// tagger checks a single message and returns the appropriate tags.
func tagger(m sms) []string {
	content := strings.ToLower(m.content)
	tags := []string{}

	if strings.Contains(content, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(content, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

// tagMessages applies a tagging function to each message.
func tagMessages(messages []sms, f func(sms) []string) []sms {
	for i := range messages {
		messages[i].tags = f(messages[i])
		if messages[i].tags == nil {
			messages[i].tags = []string{}
		}
	}
	return messages
}
