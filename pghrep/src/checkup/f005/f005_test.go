package f005

import (
	"fmt"
	"testing"

	checkup ".."
)

func TestF005Success(t *testing.T) {
	fmt.Println(t.Name())
	var report F005Report
	var hostResult F005ReportHostResult
	hostResult.Data.IndexBloatTotal = F005IndexBloatTotal{
		Count:                104,
		ExtraSizeBytesSum:    25526329344,
		RealSizeBytesSum:     25526329344,
		BloatSizeBytesSum:    25526329344,
		LiveDataSizeBytesSum: 457681690624,
		BloatRatioPercentAvg: 5.367978121989509,
		BloatRatioAvg:        1.0389965180727816,
	}

	hostResult.Data.IndexBloat = map[string]F005IndexBloat{
		"index_1": F005IndexBloat{
			Num:               1,
			IsNa:              "",
			IndexName:         "index_1",
			TableName:         "table_1",
			ExtraSizeBytes:    0,
			ExtraRatioPercent: 0.0,
			Extra:             "",
			BloatSizeBytes:    0,
			BloatRatioPercent: 0.0,
			RealSizeBytes:     0,
			LiveDataSize:      "",
			LiveDataSizeBytes: 0,
			LastVaccuum:       "",
			Fillfactor:        100.0,
			OverridedSettings: false,
			BloatRatio:        1.0,
		},
		"index_2": F005IndexBloat{
			Num:               2,
			IsNa:              "",
			IndexName:         "index_2",
			TableName:         "table_2",
			ExtraSizeBytes:    0,
			ExtraRatioPercent: 0.0,
			Extra:             "",
			BloatSizeBytes:    0,
			BloatRatioPercent: 0.0,
			RealSizeBytes:     0,
			LiveDataSize:      "",
			LiveDataSizeBytes: 0,
			LastVaccuum:       "",
			Fillfactor:        100.0,
			OverridedSettings: false,
			BloatRatio:        1.0,
		},
	}
	hostResult.Data.DatabaseSizeBytes = 9828637208
	report.Results = F005ReportHostsResults{"test-host": hostResult}
	result := F005Process(report)
	if result.P1 || result.P2 || result.P3 {
		t.Fatal("TestF005Success failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}

func TestF005TotalExcess(t *testing.T) {
	fmt.Println(t.Name())
	var report F005Report
	var hostResult F005ReportHostResult
	hostResult.Data.IndexBloatTotal = F005IndexBloatTotal{
		Count:                104,
		ExtraSizeBytesSum:    25526329344,
		RealSizeBytesSum:     25526329344,
		BloatSizeBytesSum:    25526329344,
		LiveDataSizeBytesSum: 457681690624,
		BloatRatioPercentAvg: 25.367978121989509,
		BloatRatioAvg:        1.0389965180727816,
		TableSizeBytesSum:    25526329344 * 5,
	}
	hostResult.Data.IndexBloat = map[string]F005IndexBloat{}
	hostResult.Data.DatabaseSizeBytes = 25526329344 * 4
	report.Results = F005ReportHostsResults{"test-host": hostResult}
	result := F005Process(report)
	if !result.P1 {
		t.Fatal("TestF005TotalExcess failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}

func TestF005Warnig(t *testing.T) {
	fmt.Println(t.Name())
	var report F005Report
	var hostResult F005ReportHostResult
	hostResult.Data.IndexBloatTotal = F005IndexBloatTotal{
		Count:                104,
		ExtraSizeBytesSum:    25526329344,
		RealSizeBytesSum:     25526329344,
		BloatSizeBytesSum:    25526329344,
		LiveDataSizeBytesSum: 457681690624,
		BloatRatioPercentAvg: 5.367978121989509,
		BloatRatioAvg:        1.0389965180727816,
	}
	hostResult.Data.DatabaseSizeBytes = 25526329344 * 4
	hostResult.Data.IndexBloat = map[string]F005IndexBloat{
		"index_1": F005IndexBloat{
			Num:               1,
			IsNa:              "",
			IndexName:         "index_1",
			TableName:         "table_1",
			ExtraSizeBytes:    340475904,
			ExtraRatioPercent: 86.9861866889912,
			Extra:             "",
			BloatSizeBytes:    340475904,
			BloatRatioPercent: 86.9861866889912,
			RealSizeBytes:     391413760,
			LiveDataSize:      "",
			LiveDataSizeBytes: 50937856,
			LastVaccuum:       "",
			Fillfactor:        100.0,
			OverridedSettings: false,
			BloatRatio:        7.684142811193309,
		},
		"index_2": F005IndexBloat{
			Num:               2,
			IsNa:              "",
			IndexName:         "index_2",
			TableName:         "table_2",
			ExtraSizeBytes:    3915776,
			ExtraRatioPercent: 59.3788819875776,
			Extra:             "",
			BloatSizeBytes:    3915776,
			BloatRatioPercent: 59.3788819875776,
			RealSizeBytes:     6594560,
			LiveDataSize:      "",
			LiveDataSizeBytes: 2678784,
			LastVaccuum:       "",
			Fillfactor:        100.0,
			OverridedSettings: false,
			BloatRatio:        2.46177370030581,
		},
	}

	report.Results = F005ReportHostsResults{"test-host": hostResult}
	result := F005Process(report)
	if !result.P2 {
		t.Fatal("TestF005SWarnig failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}

func TestF005Critical(t *testing.T) {
	fmt.Println(t.Name())
	var report F005Report
	var hostResult F005ReportHostResult
	hostResult.Data.IndexBloatTotal = F005IndexBloatTotal{
		Count:                104,
		ExtraSizeBytesSum:    25526329344,
		RealSizeBytesSum:     25526329344,
		BloatSizeBytesSum:    25526329344,
		LiveDataSizeBytesSum: 457681690624,
		BloatRatioPercentAvg: 5.367978121989509,
		BloatRatioAvg:        1.0389965180727816,
	}
	hostResult.Data.DatabaseSizeBytes = 25526329344 * 4
	hostResult.Data.IndexBloat = map[string]F005IndexBloat{
		"index_1": F005IndexBloat{
			Num:               1,
			IsNa:              "",
			IndexName:         "index_1",
			TableName:         "table_1",
			ExtraSizeBytes:    340475904,
			ExtraRatioPercent: 96.9861866889912,
			Extra:             "",
			BloatSizeBytes:    340475904,
			BloatRatioPercent: 96.9861866889912,
			RealSizeBytes:     391413760,
			LiveDataSize:      "",
			LiveDataSizeBytes: 50937856,
			LastVaccuum:       "",
			Fillfactor:        100.0,
			OverridedSettings: false,
			BloatRatio:        7.684142811193309,
		},
		"index_2": F005IndexBloat{
			Num:               2,
			IsNa:              "",
			IndexName:         "index_2",
			TableName:         "table_2",
			ExtraSizeBytes:    3915776,
			ExtraRatioPercent: 59.3788819875776,
			Extra:             "",
			BloatSizeBytes:    3915776,
			BloatRatioPercent: 59.3788819875776,
			RealSizeBytes:     6594560,
			LiveDataSize:      "",
			LiveDataSizeBytes: 2678784,
			LastVaccuum:       "",
			Fillfactor:        100.0,
			OverridedSettings: false,
			BloatRatio:        2.46177370030581,
		},
	}

	report.Results = F005ReportHostsResults{"test-host": hostResult}
	result := F005Process(report)
	if !result.P1 {
		t.Fatal("TestF005SWarnig failed")
	}
	checkup.PrintConclusions(result)
	checkup.PrintReccomendations(result)
}