package lamcproj

import (
	"math"
)

const nx int = 149 //X축 격자점 수
const ny int = 253 //y축 격자점 수

type lamc_parameter struct {
	re    float64 //사용할 지구반경 km
	grid  float64 //격자간격 km
	slat1 float64 //표준위도
	slat2 float64 //표준위도
	olon  float64 //기준점의 경도
	olat  float64 //기준점의 위도
	xo    float64 //기준점의 X좌표 [격자거리]
	yo    float64 //기준점의 Y좌료 [격자거리]
	first int     //시작여부  0=  시작
}

func Map_conv(lon, lat float64) (int, int) {
	lmap := lamc_parameter{}
	lmap.re = 6371.00877
	lmap.grid = 5.0
	lmap.slat1 = 30.0
	lmap.slat2 = 60.0
	lmap.olon = 126.0
	lmap.olat = 38.0
	lmap.xo = 210 / lmap.grid
	lmap.yo = 675 / lmap.grid

	lon1 := lon
	lat1 := lat

	return lamcproj(lon1, lat1, &lmap)

}

//Lambert Conformal Conic Projection
func lamcproj(lon, lat float64, lmap *lamc_parameter) (int, int) {
	PI := math.Asin(1.0) * 2.0
	DEGRAD := PI / 180.0

	re := lmap.re / lmap.grid
	slat1 := lmap.slat1 * DEGRAD
	slat2 := lmap.slat2 * DEGRAD
	olon := lmap.olon * DEGRAD
	olat := lmap.olat * DEGRAD

	sn := math.Tan(PI*0.25+slat2*0.5) / math.Tan(PI*0.25+slat1*0.5)
	sn = math.Log(math.Cos(slat1)/math.Cos(slat2)) / math.Log(sn)
	sf := math.Tan(PI*0.25 + slat1*0.5)
	sf = math.Pow(sf, sn) * math.Cos(slat1) / sn
	ro := math.Tan(PI*0.25 + olat*0.5)
	ro = re * sf / math.Pow(ro, sn)

	ra := math.Tan(PI*0.25 + lat*DEGRAD*0.5)
	ra = re * sf / math.Pow(ra, sn)
	theta := lon*DEGRAD - olon
	if theta > PI {
		theta -= 2.0 * PI
	}

	if theta < -PI {
		theta += 2.0 * PI
	}

	theta *= sn

	x := (ra * math.Sin(theta)) + lmap.xo
	y := (ro - ra*math.Cos(theta)) + lmap.yo

	return int(x + 1.5), int(y + 1.5)
}
