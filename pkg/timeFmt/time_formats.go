package timeFmt

const (
	DefaultTimeFormat = "2006-01-02T15:04:05.999Z"
	DATEw             = "28-oct-90"             // dd-mmm-yy
	DATEwl            = "28-oct-1990"           // dd-mmm-yyyy
	ADATEw            = "10/28/90"              // mm/dd/yy
	ADATEwl           = "10/28/1990"            // mm/dd/yyyy
	EDATEw            = "28.10.90"              // dd.mm.yy
	EDATEwl           = "28.10.1990"            // dd.mm.yyyy
	JDATEw            = "90301"                 // yyddd
	JDATEwl           = "1990301"               // yyyyddd
	SDATEw            = "90/10/28"              // yy/mm/dd
	SDATEwl           = "1990/10/28"            // yyyy/mm/dd
	QYRw              = "4 Q 90"                // q Q yy
	QYRwl             = "4 Q 1990"              // q Q yyyy
	MOYRw             = "oct 90"                // mmm yy
	MOYRwl            = "oct 1990"              // mmm yyyy
	WKYRw             = "43 wk 90"              // ww WK yy
	WKYRwl            = "43 wk 1990"            // ww WK yyyy
	WKDAYw            = "su"                    // (name of the day)
	MONTHw            = "jan"                   // (name of the month)
	TIMEw             = "01:02"                 // hh:mm
	TIMEwl            = "11:50:30"              // hh:mm:ss
	TIMEwd            = "01:02:34.75"           // hh:mm:ss.s
	MTIMEw            = "02:34"                 // mm:ss
	MTIMEwd           = "02:34.75"              // mm:ss.s
	DTIMEw            = "20 08:03"              // ddd hh:mm
	DTIMEwl           = "113 04:20:18"          // ddd hh:mm:ss
	DATETIMEw         = "20-jun-1990 08:03"     // dd-mmm-yyyy hh:mm:ss
	DATETIMEwd        = "20-jun-1990 08:03:00"  // dd-mmm-yyyy hh:mm:ss.s
	YMDHMSw           = "1990-06-20 08:03"      // yyyy-mm-dd hh:mm
	YMDHMS            = "1987-02-04 21:13:49"   // yyyy-mm-dd hh:mm:ss
	YMDHMSwd          = "1990-06-20 08:03:00.0" // yyyy-mm-dd hh:mm:ss.s
)
