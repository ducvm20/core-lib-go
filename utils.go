package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

const (
	YMDHMS = "Y-m-d H:i:s"
	YMD    = "Y-m-d"
	YDM    = "Y-d-m"
	SG     = carbon.Saigon
)

func StructToMap(myStruct interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	jsonEnc, err := json.Marshal(myStruct)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(jsonEnc, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func DefaultEnv(variable_name, defaultValue string) (string, bool) {
	value, ok := os.LookupEnv(variable_name)
	if ok {
		return value, ok
	}
	fmt.Println("Không có biến môi trường:", variable_name)
	return defaultValue, ok
}
func Hash256(strData string) string {
	h := sha256.New()
	h.Write([]byte(strData))
	return hex.EncodeToString(h.Sum(nil))
}
func sampleBucket(cdf []float64) int {
	r := rand.Float64()
	bucket := 0
	for r > cdf[bucket] {
		bucket++
	}
	return bucket
}
func DistributePros(pdf []float64) int {
	cdf := make([]float64, len(pdf))
	cdf[0] = pdf[0]
	for i := 1; i < len(pdf); i++ {
		cdf[i] = cdf[i-1] + pdf[i]
	}
	return sampleBucket(cdf)
}
func ExecTime(start time.Time, funcName string, log *zap.Logger) {
	logStr := fmt.Sprintf("ExecTime of %s take dt=%d ms", funcName, Now().Sub(start).Milliseconds())
	if log != nil {
		log.Info(logStr)
	} else {
		fmt.Println(logStr)
	}
}
func Now() time.Time {
	return carbon.Now(SG).Carbon2Time()
}
func DateEqual(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
func Time2Str(t time.Time, format string) string {
	return carbon.Time2Carbon(t).Format(format, SG)
}
func Str2Time(timeAsString, format string) time.Time {
	return carbon.ParseByFormat(timeAsString, format, SG).Carbon2Time()
}
func DateBetween(start time.Time, end time.Time) bool {
	return Now().After(start) && Now().Before(end)
}
func EncodeBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}
func DecodeBase64(textEncoded string) (string, error) {
	enc, err := base64.StdEncoding.DecodeString(textEncoded)

	return string(enc), err

}
func Checkregex(input string, inputType string) bool {
	var sampleRegexp *regexp.Regexp
	switch inputType {
	case "phone":
		sampleRegexp = regexp.MustCompile(`^0[0-9]{9}$`)
	case "review":
		sampleRegexp = regexp.MustCompile(`^[0-9]{1}$`)
	case "timezone":
		sampleRegexp = regexp.MustCompile(`^\d*$`)
	case "pin":
		sampleRegexp = regexp.MustCompile(`^\d{6}$`)
	case "otp":
		sampleRegexp = regexp.MustCompile(`^\d{6}$`)
	case "email":
		sampleRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	case "permission":
		sampleRegexp = regexp.MustCompile(`^[0-9](:[0-9])*$`)
	case "address":
		sampleRegexp = regexp.MustCompile(`^[\p{L} 0-9\/,]*$`)
	case "text":
		sampleRegexp = regexp.MustCompile(`^[\p{L} ]*$`)
	case "text_number":
		sampleRegexp = regexp.MustCompile(`^[\p{L} 0-9]*$`)

	case "float":
		sampleRegexp = regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]*`)

	case "feedback":
		sampleRegexp = regexp.MustCompile(`^[\p{L} 0-9!@#$%&()+\-=\\{};':"\\,.<>\/?]*$`)
	case "appoinmentdate":
		sampleRegexp = regexp.MustCompile(`^[0-9]{4}:[0-9]{2}:[0-9]{2}$`)
	case "number":
		sampleRegexp = regexp.MustCompile(`^\d*$`)
	case "uuid":
		sampleRegexp = regexp.MustCompile(`^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$`)
	}

	return sampleRegexp.MatchString(input)
}
func StringRandom(lenStr int) string {
	var pool = "ASDFGHJKLQWERTYUIOPZXCVBNMasdfghjklzxcvbnmqwertyuiop1234567890"
	bytes := make([]byte, lenStr)
	for i := 0; i < lenStr; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)

}
func StringRandom2(lenStr int) string {
	var pool = "ASDFGHJKLQWERTYUIOPZXCVBNM1234567890"
	bytes := make([]byte, lenStr)
	for i := 0; i < lenStr; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)

}
func FloatToTime(input float64) time.Time {
	integ, decim := math.Modf(input)
	return time.Unix(int64(integ), int64(decim*(1e9)))
}
func StringToFloat64(input string) (float64, error) {
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
func StringToInt(input string) (int, error) {
	result, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return result, nil
}

var Weekdays = map[string]string{
	"Sunday":    "Chủ nhật",
	"Monday":    "Thứ hai",
	"Tuesday":   "Thứ ba",
	"Wednesday": "Thứ tư",
	"Thursday":  "Thứ năm",
	"Friday":    "Thứ sáu",
	"Saturday":  "Thứ bảy",
}

func DayInWeek(day time.Weekday) string {
	// fmt.Printf("day.String(): %v\n", day.String())
	return Weekdays[day.String()]
}
func TimeToTimeStr(start, end time.Time) string {
	format1 := "02 tháng 01, 2006 03:04 PM"
	format2 := "02 tháng 01, 2006 (03:04 PM - "
	format3 := "03:04 PM)"
	if start.Equal(end) {
		return DayInWeek(start.Weekday()) + ", " + start.Format(format1)
	} else {
		if start.Day() == end.Day() && start.Month() == end.Month() && start.Year() == end.Year() {
			return DayInWeek(start.Weekday()) + ", " + start.Format(format2) + end.Format(format3)
		} else {
			return DayInWeek(start.Weekday()) + ", " + start.Format(format1) + " - " + DayInWeek(end.Weekday()) + ", " + end.Format(format1)
		}
	}
}
func RemoveDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ConvertListToStringVN(input []string) string {
	lenStr := len(input)
	if lenStr == 1 {
		return input[0]
	}
	result := ""
	for index, value := range input {
		if index == 0 {
			result = value
		} else if index == lenStr-1 {
			result = result + " và " + value
		} else {
			result = result + ", " + value
		}
	}
	return result
}

func NewBrokers(useProduction bool) []string {
	if useProduction {
		return []string{"isc-kafka01:9092", "isc-kafka02:9092", "isc-kafka03:9092"}
	}
	return []string{"isc-kafka01:9092", "isc-kafka02:9092", "isc-kafka03:9092"}
}
func NewKafkaTopicName(useProduction bool) string {
	if useProduction {
		return "hifpt-hi-ecom-logs"
	}
	return "stag-hifpt-hi-ecom-logs"
}
func NewKafkaTopicNameAll(useProduction bool) string {
	if useProduction {
		return "hifpt-all-in-one-kafka"
	}
	return "stag-hifpt-all-in-one-kafka"
}
