package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnvVars_valid(t *testing.T) {
	// given
	enVars := []string{"APP_SERVER=test", "OS=windows"}

	// when
	result, err := GetEnvVars(enVars)

	// then
	assert.NoError(t, err)
	assert.Equal(t, result, map[string][]byte{"APP_SERVER": []byte("test")})
}

func TestGetEnvVars_invalid(t *testing.T) {
	// given
	enVars := []string{"APP_$%^4_=test"}

	// when
	result, err := GetEnvVars(enVars)

	// then
	assert.EqualError(t, err, "env: APP_$%^4_ does not match the required pattern")
	assert.Nil(t, result)
}

func TestReplaceEnvVars_valid(t *testing.T) {
	// given
	html := generateHtml(`<meta property="app:server-uri" content="__APP_SERVER_URI__">`)

	// when
	result, err := ReplaceEnvVars(html, map[string][]byte{
		"APP_SERVER_URI": []byte("http://example.com"),
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, string(result), string(generateHtml(`<meta property="app:server-uri" content="http://example.com">`)))
}

func TestReplaceEnvVars_valueNotInterpolated(t *testing.T) {
	// given
	html := generateHtml(`<meta property="app:server-uri" content="__APP_SERVER_URI__">`)

	// when
	result, err := ReplaceEnvVars(html, map[string][]byte{})

	// then
	assert.EqualError(t, err, "missing substituion for: __APP_SERVER_URI__")
	assert.Nil(t, result)
}

func TestReplaceEnvVars_missingInterpolationPoint(t *testing.T) {
	// given
	html := generateHtml(`<meta property="app:server-uri" content="__APP_SERVER_URI__">`)

	// when
	result, err := ReplaceEnvVars(html, map[string][]byte{
		"APP_SERVER_URI": []byte("http://example.com"),
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, string(result), string(generateHtml(`<meta property="app:server-uri" content="http://example.com">`)))
}

func TestReplaceEnvVars_duplicateInterpolationPoint(t *testing.T) {
	// given
	html := generateHtml(`<meta property="app:server-uri" content="__APP_SERVER_URI__">
<meta property="app:server-uri" content="__APP_SERVER_URI__">`)

	// when
	result, err := ReplaceEnvVars(html, map[string][]byte{
		"APP_SERVER_URI": []byte("http://example.com"),
	})

	// then
	assert.EqualError(t, err, "found multiple interpolation points for: APP_SERVER_URI")
	assert.Nil(t, result)
}

func TestReplaceEnvVars_twoInterpolationPoint(t *testing.T) {
	// given
	html := generateHtml(`<meta property="app:server-uri" content="__APP_SERVER_URI__">
<meta property="app:server-stage" content="__APP_SERVER_STAGE__">`)

	// when
	result, err := ReplaceEnvVars(html, map[string][]byte{
		"APP_SERVER_URI":   []byte("http://example.com"),
		"APP_SERVER_STAGE": []byte("dev"),
	})

	// then
	assert.NoError(t, err)
	assert.Equal(t, string(result), string(generateHtml(`<meta property="app:server-uri" content="http://example.com">
<meta property="app:server-stage" content="dev">`)))
}

func generateHtml(additional string) []byte {
	return []byte(fmt.Sprintf(`
<!doctype html>
<html>
 <head>
   <title>This is the title of the webpage!</title>
	%s
 </head>
 <body>
   <p>This is an example paragraph. Anything in the <strong>body</strong> tag will appear on the page, just like this <strong>p</strong> tag and its contents.</p>
 </body>
</html>`, additional))
}
