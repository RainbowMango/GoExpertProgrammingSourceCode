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

	// 期望调用Report()，且参数为"beijing"时输出"SUNNY"
	fakeReporter.EXPECT().Report("beijing").Return("SUNNY")

	// 测试1：晴朗天气适合出行
	if SuitTravel("beijing", fakeReporter) != true {
		t.Errorf("expected true")
	}

	// 期望调用Report()，且参数为"shanghai"时输出"TYPHOON"
	fakeReporter.EXPECT().Report("shanghai").Return("TYPHOON")

	// 测试2：台风天气不适合出行
	if SuitTravel("shanghai", fakeReporter) != false {
		t.Errorf("expected false")
	}
}
