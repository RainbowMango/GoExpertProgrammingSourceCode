package reporter

// WeatherReporter 定义天气预报接口。
// 实现者可以从任意渠道获取天气
type WeatherReporter interface {
	// Report 报告某个城市的当天天气
	Report(city string) string
}

func SuitTravel(city string, reporter WeatherReporter) bool {
	weather := reporter.Report(city)
	if weather == "SUNNY" { // 天气晴朗，适宜出行
		return true
	}

	if weather == "TYPHOON" { // 台风不宜出行
		return false
	}

	return false // 天气不明，不宜出行
}
