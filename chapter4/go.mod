module main

go 1.16

replace chapter4 => ../chapter4

require (
	chapter4 v0.0.0-00010101000000-000000000000
	github.com/headfirstgo/greeting v0.0.0-20190504033635-66e7507184ee // indirect
	github.com/headfirstgo/keyboard v0.0.0-20170926053303-9930bcf72703
)
