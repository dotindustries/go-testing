package testing

import (
	"runtime/debug"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type TemporalWorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	Env *testsuite.TestWorkflowEnvironment
}

func (s *TemporalWorkflowTestSuite) AfterTest(suiteName, testName string) {
	// after Test methods on Suite -- could include multiple Scenarios
}

func (s *TemporalWorkflowTestSuite) SetupTest() {
	// setup per Test method on Suite -- could include multiple Scenarios
}

func (s *TemporalWorkflowTestSuite) Scenario(name string, mockSetupFn func(mockCtrl *gomock.Controller), steps ...func(suite **TemporalWorkflowTestSuite)) {
	s.Run(name, func() {
		var mockCtrl *gomock.Controller
		s.Env = s.NewTestWorkflowEnvironment()

		if mockSetupFn != nil {
			mockCtrl = gomock.NewController(s.T())
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
		s.Env.AssertExpectations(s.T())
		if mockCtrl != nil {
			mockCtrl.Finish()
		}
	})
}

func (s *TemporalWorkflowTestSuite) Given(fn func(suite **TemporalWorkflowTestSuite)) func(suite **TemporalWorkflowTestSuite) {
	return fn
}

func (s *TemporalWorkflowTestSuite) When(fn func(suite **TemporalWorkflowTestSuite)) func(suite **TemporalWorkflowTestSuite) {
	return fn
}

func (s *TemporalWorkflowTestSuite) Then(fn func(suite **TemporalWorkflowTestSuite)) func(suite **TemporalWorkflowTestSuite) {
	return fn
}

func (s *TemporalWorkflowTestSuite) And(fn func(suite **TemporalWorkflowTestSuite)) func(suite **TemporalWorkflowTestSuite) {
	return fn
}
