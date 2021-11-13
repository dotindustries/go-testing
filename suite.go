package testing

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type BDTestSuite struct {
	suite.Suite
}

func (s *BDTestSuite) AfterTest(suiteName, testName string) {
}

func (s *BDTestSuite) SetupTest() {
}

func (s *BDTestSuite) Scenario(name string, mockSetupFn func(mockCtrl *gomock.Controller), steps ...func(s **BDTestSuite)) {
	s.Run(name, func() {
		var mockCtrl *gomock.Controller
		if mockSetupFn != nil {
			mockCtrl = gomock.NewController(s.T())
			mockSetupFn(mockCtrl)
		}

		for _, step := range steps {
			s.NotPanics(func() {
				step(&s)
			})
		}

		if mockCtrl != nil {
			mockCtrl.Finish()
		}
	})
}

func (s *BDTestSuite) Given(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}

func (s *BDTestSuite) When(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}

func (s *BDTestSuite) Then(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}

func (s *BDTestSuite) And(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}
