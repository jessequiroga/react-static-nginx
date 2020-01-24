package pkg

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var prefix = "APP_"
var validEnvVarRe = regexp.MustCompile(`^\w[\w\d_]*\w$`)

// GetEnvVars Find all environment variables matching the prefix
func GetEnvVars(envVars []string) (map[string][]byte, error) {
	var envs = make(map[string][]byte)

	for _, e := range envVars {
		pair := strings.SplitN(e, "=", 2)

		var key = pair[0]
		var val = []byte(pair[1])

		if strings.HasPrefix(key, prefix) {
			envs[key] = val
		}

		if !validEnvVarRe.Match([]byte(key)) {
			return nil, fmt.Errorf("env: %s does not match the required pattern", key)
		}
	}

	return envs, nil
}

func ReplaceEnvVars(data []byte, vars map[string][]byte) ([]byte, error) {
	for key, value := range vars {
		var re = regexp.MustCompile(fmt.Sprintf("__%s__", key))

		matches := re.FindAllIndex(data, -1)

		if matches == nil {
			return nil, fmt.Errorf("could not find subsitution point for: %s, Please add %s to the html", key, substitutionPoint(key))
		}

		if len(matches) > 1 {
			return nil, fmt.Errorf("found multiple interpolation points for: %s", key)
		}

		data = re.ReplaceAll(data, value)
	}

	err := findLeftoverMatches(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func findLeftoverMatches(data []byte) error {
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		// ...
	}
	var f func(*html.Node) error

	f = func(n *html.Node) error {
		// Find Meta tags in the head
		if n.Type == html.ElementNode && n.Data == "meta" && n.Parent.Type == html.ElementNode && n.Parent.Data == "head" {
			//
			for _, a := range n.Attr {
				if a.Key == "content" && strings.HasPrefix(a.Val, "__") && strings.HasSuffix(a.Val, "__") {
					return fmt.Errorf("missing substituion for: %s", a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			err := f(c)

			if err != nil {
				return err
			}
		}

		return nil
	}

	return f(doc)
}

func substitutionPoint(env string) string {
	return fmt.Sprintf(`<meta property="app:server-uri" content="__%s__">`, env)
}
