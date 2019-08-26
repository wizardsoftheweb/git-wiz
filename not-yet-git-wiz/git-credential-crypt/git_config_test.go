package git_credentials_store

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestGitConfig(t *testing.T) { TestingT(t) }

type GitConfigSuite struct {
	BaseSuite
}

var _ = Suite(&GitConfigSuite{})

func (s *GitConfigSuite) SetUpTest(c *C) {
}

func (s *GitConfigSuite) TearDownTest(c *C) {
}

func (s *GitConfigSuite) TestCollectGitConfig(c *C) {
	result := CollectGitConfig()
	c.Assert(result, Not(IsNil))
}

func (s *GitConfigSuite) TestCheckPathValue(c *C) {
	matrix := []struct {
		input  string
		result bool
	}{
		{
			"credential.usehttppath true",
			true,
		},
		{
			"credential.usehttppath false",
			false,
		},
	}
	for _, entry := range matrix {
		theyDontMatter := CheckPathValue(entry.input)
		c.Assert(theyDontMatter, Equals, entry.result)
	}
}
