package temporaltest

import (
	"runtime/debug"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type TemporalActivityTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	Env *testsuite.TestActivityEnvironment
}

func (s *TemporalActivityTestSuite) AfterTest(suiteName, testName string) {
	// after Test methods on Suite -- could include multiple Scenarios
}

func (s *TemporalActivityTestSuite) SetupTest() {
	// setup per Test method on Suite -- could include multiple Scenarios
}

func (s *TemporalActivityTestSuite) Scenario(name string, mockSetupFn func(mockCtrl *gomock.Controller), steps ...func(suite **TemporalActivityTestSuite)) {
	s.Run(name, func() {
		mockCtrl := gomock.NewController(s.T())
		s.Env = s.NewTestActivityEnvironment()

		if mockSetupFn != nil {
			mockSetupFn(mockCtrl)
		}

		for _, step := range steps {
			func() {
				defer func() {
					if r := recover(); r != nil {
						s.T().Error("step paniced", r, string(debug.Stack()))
					}
				}()
				step(&s)
			}()
		}
		mockCtrl.Finish()
	})
}

func (s *TemporalActivityTestSuite) Given(fn func(suite **TemporalActivityTestSuite)) func(suite **TemporalActivityTestSuite) {
	return fn
}

func (s *TemporalActivityTestSuite) When(fn func(suite **TemporalActivityTestSuite)) func(suite **TemporalActivityTestSuite) {
	return fn
}

func (s *TemporalActivityTestSuite) Then(fn func(suite **TemporalActivityTestSuite)) func(suite **TemporalActivityTestSuite) {
	return fn
}

func (s *TemporalActivityTestSuite) And(fn func(suite **TemporalActivityTestSuite)) func(suite **TemporalActivityTestSuite) {
	return fn
}
