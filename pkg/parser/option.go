package parser

import (
	"encoding/json"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

type Option func(*Parser)

func isJSON(s string) bool {
	var n = len(s)
	if n > 5 && strings.ToLower(string(s[n-5:])) == ".json" {
		return true
	}
	return false
}

func isYAML(s string) bool {
	var n = len(s)
	if n > 4 && strings.ToLower(string(s[n-4:])) == ".yml" {
		return true
	}
	if n > 5 && strings.ToLower(string(s[n-5:])) == ".yaml" {
		return true
	}
	return false
}

func WithInput(input string) Option {
	input = strings.TrimSpace(input)
	return func(p *Parser) {
		p.Input = input
		if isJSON(input) {
			WithJSONDecoder()(p)
		}
		if isYAML(input) {
			WithYAMLDecoder()(p)
		}
	}
}

func WithOutput(output string) Option {
	output = strings.TrimSpace(output)
	return func(p *Parser) {
		p.Output = output
		if isJSON(output) {
			WithJSONEncoder()(p)
		}
		if isYAML(output) {
			WithYAMLEncoder()(p)
		}
	}
}

func WithJSONDecoder() Option {
	return func(p *Parser) {
		p.NewDecoder = func(r io.Reader) Decoder {
			return json.NewDecoder(r)
		}
	}
}

func WithJSONEncoder() Option {
	return func(p *Parser) {
		p.NewEncoder = func(w io.Writer) Encoder {
			return json.NewEncoder(w)
		}
	}
}

func WithYAMLDecoder() Option {
	return func(p *Parser) {
		p.NewDecoder = func(r io.Reader) Decoder {
			return yaml.NewDecoder(r)
		}
	}
}

func WithYAMLEncoder() Option {
	return func(p *Parser) {
		p.NewEncoder = func(w io.Writer) Encoder {
			return yaml.NewEncoder(w)
		}
	}
}
