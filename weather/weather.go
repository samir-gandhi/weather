package weather

import "time"

type Values struct {
	Temperature float64 `json:"temperature"`
}

type Weather struct {
	Data struct {
		Timelines []struct {
			Timestep  string    `json:"timestep"`
			EndTime   time.Time `json:"endTime"`
			StartTime time.Time `json:"startTime"`
			Intervals []struct {
				StartTime time.Time `json:"startTime"`
				Values    Values    `json:"values"`
			} `json:"intervals"`
		} `json:"timelines"`
	} `json:"data"`
}

// type TemperatureValue struct {
// 	Temperature *float64
// }

// type Intervals struct {
// 	StartTime NonNullableTimeValue
// 	Values    *TemperatureValue
// }

// type Timelines struct {
// 	Timestep string
// 	EndTime NonNullableTimeValue
// 	StartTime NonNullableTimeValue
// 	Intervals []Intervals `json:"intervals"`
// }

// type Data struct {
// 	Timelines []Timelines `json:"timelines"`
// }

// type Weather struct {
// 	Data Data `json:"data"`
// }
