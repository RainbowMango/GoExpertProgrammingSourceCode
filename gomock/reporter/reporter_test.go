package reporter

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rainbowmango/goexpertprogrammingsourcecode/gomock/mockreporter"
)

func TestSuitTravel(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	fakeReporter := mockreporter.NewMockWeatherReporter(ctl)

	// 测试1：晴朗天气适合出行
	SuitTravel("beijing", fakeReporter)

	// 测试2：台风天气不适合出行
}
