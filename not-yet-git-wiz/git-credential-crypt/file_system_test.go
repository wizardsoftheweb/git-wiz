package git_credentials_store

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	. "gopkg.in/check.v1"
)

func TestFileSystem(t *testing.T) { TestingT(t) }

type FileSystemSuite struct {
	BaseSuite
	fileContents  []byte
	fileToWrite   string
	homeDirectory string
}

var currentUser, _ = user.Current()

var _ = Suite(&FileSystemSuite{
	fileContents:  []byte("test data"),
	fileToWrite:   "",
	homeDirectory: currentUser.HomeDir,
})

func (s *FileSystemSuite) SetUpTest(c *C) {
	s.fileToWrite = filepath.Join(s.workingDirectory, "test.file")
}

func (s *FileSystemSuite) TearDownTest(c *C) {
	_ = os.Remove(s.fileToWrite)
	pathTidier = tidyPath
}

func (s *FileSystemSuite) TestTidyPath(c *C) {
	var tidyPathData = [][]string{
		{"/", "/"},
		{"/some/dir", "/some", "dir"},
		{fmt.Sprintf("%s/%s", s.workingDirectory, "some/dir"), "some", "dir"},
		{s.homeDirectory, "~"},
		{filepath.Join(s.homeDirectory, "test"), "~/test"},
	}
	for _, value := range tidyPathData {
		result, err := tidyPath(value[1:]...)
		c.Assert(err, Not(ErrorMatches), "*")
		c.Assert(result, Equals, value[0])
	}
}

func (s *FileSystemSuite) TestEnsureDirectoryExistsWorksWithCwd(c *C) {
	err := EnsureDirectoryExists(s.workingDirectory)
	c.Assert(err, IsNil)
}

func (s *FileSystemSuite) TestEnsureDirectoryExistsCreatesDirectories(c *C) {
	additionalPathComponents := []string{"some", "dir"}
	fullPath := filepath.Join(
		append(
			[]string{s.workingDirectory},
			additionalPathComponents...,
		)...,
	)
	err := EnsureDirectoryExists(additionalPathComponents...)
	c.Assert(err, IsNil)
	_, err = os.Stat(fullPath)
	c.Assert(os.IsNotExist(err), Equals, false)
}

func (s *FileSystemSuite) TestLoadFilePathError(c *C) {
	pathTidier = s.brokenPathTidier
	_, err := LoadFile(s.workingDirectory)
	c.Assert(err, ErrorMatches, s.errorMessage)
}

func (s *FileSystemSuite) TestLoadFileThatDoesntExist(c *C) {
	_, err := LoadFile(filepath.Join(s.workingDirectory, "random", "file"))
	c.Assert(os.IsNotExist(err), Equals, true)
}

func (s *FileSystemSuite) TestLoadFileNonEmpty(c *C) {
	contents, err := LoadFile(s.currentFilename)
	c.Assert(err, IsNil)
	c.Assert(contents, Not(Equals), []byte{})
}

func (s *FileSystemSuite) TestWriteFilePathError(c *C) {
	pathTidier = s.brokenPathTidier
	err := WriteFile(s.fileContents, 0600, s.workingDirectory)
	c.Assert(err, ErrorMatches, s.errorMessage)
}

func (s *FileSystemSuite) TestWriteFileSuccess(c *C) {
	perm := os.FileMode(uint32(0640))
	err := WriteFile(s.fileContents, perm, s.fileToWrite)
	c.Assert(err, IsNil)
	contents, err := ioutil.ReadFile(s.fileToWrite)
	c.Assert(err, IsNil)
	c.Assert(contents, DeepEquals, s.fileContents)
	stats, _ := os.Stat(s.fileToWrite)
	c.Assert(stats.Mode().String(), Equals, perm.String())
}

func (s *FileSystemSuite) TestDoesPathExistNo(c *C) {
	newPath := s.currentFilename[1:]
	c.Assert(DoesPathExist(newPath), Equals, false)
}

func (s *FileSystemSuite) TestDoesPathExistYes(c *C) {
	c.Assert(DoesPathExist(s.currentFilename), Equals, true)
}

func (s *FileSystemSuite) TestDoesPathExistBadPath(c *C) {
	pathTidier = s.brokenPathTidier
	c.Assert(DoesPathExist(s.currentFilename), Equals, false)
}
