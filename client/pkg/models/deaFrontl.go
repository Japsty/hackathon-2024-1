package models

type Time struct {
	Hour   int `json:"Hour"`
	Minute int `json:"Minute"`
}

type Deal struct {
	TimeClosed Time `json:"TimeClosed"`
	OpenVal    int  `json:"OpenVal"`
	HighVal    int  `json:"HighVal"`
	LowVal     int  `json:"LowVal"`
	ClosedVal  int  `json:"ClosedVal"`
	VolumeVal  int  `json:"VolumeVal"`
}
