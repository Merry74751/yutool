package common

import "time"

var (
	Epoch         int64 = 1597075200000
	machineID     int64                
	sn            int64                
	lastTimeStamp int64                 
)

func init() {
	lastTimeStamp = time.Now().UnixNano()/1e6 - Epoch
}

func SetMachineID(mid int64) {
	machineID = mid << 12
}

func GetSnowflakeID() int64 {
	curTimeStamp := time.Now().UnixNano()/1e6 - Epoch
	if curTimeStamp == lastTimeStamp {
		sn++
		if sn > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano()/1e6 - Epoch
			lastTimeStamp = curTimeStamp
			sn = 0
		}

		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22

		id := rightBinValue | machineID | sn
		return id
	} else if curTimeStamp > lastTimeStamp {
		sn = 0
		lastTimeStamp = curTimeStamp
		rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
		rightBinValue <<= 22
		return rightBinValue | machineID | sn
	}
	return 0
}
