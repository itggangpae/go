package main

import "fmt"

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

func (suite *ExampleTestSuite) SetupTest() {
	fmt.Printf("테스트 시작 전 수행")
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
