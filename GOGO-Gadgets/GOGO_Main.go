/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter. (This is a GO Thing)

	Apr 22, 2022	- Some of the BEST CODE I EVER WROTE... check out GET_FUNC_PARAM_DYNAMIC!!!
	Aug 27, 2021	- Ripped out a bunch of stuff to make this smaller. They are 
	Jun 05, 2014    - Initial Rollout

*/

package main


import (
	// = = = = = Native Libraries
	// = = = = = PERSONAL Libraries
	// = = = = = 3rd Party Libraries

	// = = = = = Native Libraries
		"flag"		
		"math"
		"math/rand"
		"os"
		"os/exec"
		"runtime"
		"strconv"
		"strings"
		"time"
		"bufio"
		"unicode/utf8"

	// = = = = = PERSONAL Libraries


	// = = = = = 3rd Party Libraries

		"github.com/atotto/clipboard"
		"github.com/briandowns/spinner"
		"github.com/dustin/go-humanize"
		"github.com/fatih/color"
)

/*
	- - - -
	- - - -
	- - - - START OF GLOBALS WE NEED - - - - - -
	- - - -
	- - - -
*/

var SERIAL_NUM = "" // This is unique execution sid generated everytime a program starts. useful for troubleshooting in jenkins
var SHOW_SERIAL = false 		// If set, we Generate and show a serial number

var OSTYPE=""
var CURRENT_OS = ""	
var GOOS_VALUE = "" 		// Holds the current OS as reported by runtime.GOOS

var DEBUG_MODE = false		// Useful universal flag for enabling DEBUG_MODE code blocks
var ERROR_EXIT_CODE = -9999
var ERROR_CODE = ERROR_EXIT_CODE

var PROG_START_TIME string
var PROG_START_TIMEOBJ time.Time

var GLOBAL_CURR_DATE = ""		// Current Actual Date in the Timezone we specified
var GLOBAL_CURR_TIME = ""		// alias to CURR_DATE

var GLOBAL_DATE_OBJ time.Time		// Actual Global Date OBJ
var GLOBAL_TIME_OBJ time.Time		// alias



// -=-=-= COMMON COLOR GLOBAL references =-=-=-=-

var R = color.New(color.FgRed, color.Bold)
var G = color.New(color.FgGreen, color.Bold)
var Y = color.New(color.FgYellow, color.Bold)
var B = color.New(color.FgBlue, color.Bold)
var M = color.New(color.FgMagenta, color.Bold)
var C = color.New(color.FgCyan, color.Bold)
var W = color.New(color.FgWhite, color.Bold)

var R2 = color.New(color.FgRed)
var G2 = color.New(color.FgGreen)
var Y2 = color.New(color.FgYellow)
var B2 = color.New(color.FgBlue)
var M2 = color.New(color.FgMagenta)
var C2 = color.New(color.FgCyan)
var W2 = color.New(color.FgWhite)

var R3 = color.New(color.FgRed, color.Underline)
var G3 = color.New(color.FgGreen, color.Underline)
var Y3 = color.New(color.FgYellow, color.Underline)
var B3 = color.New(color.FgBlue, color.Underline)
var M3 = color.New(color.FgMagenta, color.Underline)
var C3 = color.New(color.FgCyan, color.Underline)
var W3 = color.New(color.FgWhite, color.Underline)


// Time Zone related stuff
var DEFAULT_ZONE_LOCATION_OBJ, _ = time.LoadLocation("Local")

var ZONE_LOCAL = ""
var ZONE_HOUR_OFFSET = "" 
var ZONE_UPPER = ""
var ZONE_FULL = ""

var USE_PST = false
var USE_CST = false
var USE_EST = false
var USE_MST = false
var USE_UTC = false

// These are used by SET_TIMEZONE_DEFAULTS.....and....GET_CURRENT_TIME 
var EST_OBJ, _ = time.LoadLocation("EST")
var CST_OBJ, _ = time.LoadLocation("America/Chicago")			// aka CST	}
var MST_OBJ, _ = time.LoadLocation("MST") 		// MDT / Mountain Standard
var PST_OBJ, _ = time.LoadLocation("America/Los_Angeles")		// aka PST
var UTC_OBJ, _ = time.LoadLocation("UTC") 


var DEFAULT_FLOAT = -999999999.999999999
var DEFAULT_INT = -999999999 

/*
	= = = = = = = = = = = = = = = = = = = = =
	End of GLOBALS definitions
	= = = = = = = = = = = = = = = = = = = = =
*/



/*
	SORTING Strings and Lists.. Cheat Sheet!

	var NAMES []string

	NAMES = append(NAMES, "Jenny")
	NAMES = append(NAMES, "Alec")
	NAMES = append(NAMES, "Kendal")
	NAMES = append(NAMES, "Carter")

	//4. To Sort the names in NAMES alphabetically, 
	sort.Strings(NAMES)


	// For sorting Structs Alphabetically 
	type STOCK_OBJ struct {

		SYMBOL		string
		Date		string
		Price		float64
		TIME_OBJ	time.Time
	}

	var STOCKS []STOCK_OBJ

	var S STOCK_OBJ

	S.SYMBOL = "NFLX"
	S.Date	= "12/05/2005"
	S.Price = 353.24
	STOCKS = append(STOCKS, S)

	S.SYMBOL = "AAPL"
	S.Date	= "10/10/2010"
	S.Price = 210.24
	STOCKS = append(STOCKS, S)

	S.SYMBOL = "ZNFL"
	S.Date	= "11/11/2002"
	S.Price = 53.24
	STOCKS = append(STOCKS, S)

	S.SYMBOL = "BAC"
	S.Date	= "01/05/2003"
	S.Price 98.24
	STOCKS = append(STOCKS, S)

	// At this point the list is unordered (items are entered in the order shown above)
	// To Sort ALPHABETICALLY do this: (if you want to sort in REVERSE alpha order, use > (greater than)) 
	
	//  (Note, the same applies to INTS and Floats)

    sort.Slice(STOCKS, func(i, j int) bool {
        return STOCKS[i].SYMBOL < STOCKS[j].SYMBOL
    })


	// Finally to sort the above slice /struct by way of the time.Time TIME_OBJ:
	
	sort.Slice(STOCKS, func(i, j int) bool { return (STOCKS)[i].TIME_OBJ.Before((STOCKS)[j].TIME_OBJ)})	

	(Will sort the slice by order of time/date the record was entered in say mongo)

	
	 
*/




type PARAM_OBJ struct {
	PARAM_VALUE	interface{}
	PARAM_INDEX	int
	PARAM_TYPE  string  // string, int float, struct  etc
}

/* 
Usage: GET_FUNC_PARAM(paramINDEX, "", &VAR_TO_ASSIGN_on_callback, params... )     (note ellipses AFTER params)

	NOTE:  First Parameter must be either the SPECIFIC INDEX of the parameter you want... or a KEYWORD that matches WHATEVER parameter that was passed
	RETURNED: FOUND (true|false), vALUE of the object (struct or string or int etc), and ALL the PARAM_OBJs incase you want to iterate them: EXAMPLE:


func TEST_FUNCTION(params ...interface{}) {

	valid, VALUE, _ := GET_FUNC_PARAM_DYNAMIC("cowboy", false, params...)

	if valid {
		C.Println("FOUND PARAM: ")
		Y.Println(VALUE)
	}
}	

Function Invoke Exampple:
    TEST_FUNCTION(4, 4.5, "terry", false, "cowboy", false, true, 3, client_key_HEADER_B, allHEADERS, mystrings, myints)

	(client_key_HEADER_B is a struct, allHEADERS is an array of struct, mystrings is []string): 
	
*/
func GET_FUNC_PARAM_DYNAMIC(FIND_USING interface{}, RESULT_to_ASSSIGN_TO *interface{}, BE_VERBOSE bool, params ...interface{}) (bool, []PARAM_OBJ ) {

	var ALL_PARAMS []PARAM_OBJ

	var INDEX_to_FIND = -9
	var SPECIFIC_STRING = ""

	var USE_INDEX = false
	var USE_KEYWORD = false

	// FIND_USING is either the INDEX of the param value we want to return... or the param that exactly matches the string we pass
	INT_val, IS_int := FIND_USING.(int)
	STRING_val, IS_string := FIND_USING.(string)			

	if IS_int {
		INDEX_to_FIND = INT_val
		USE_INDEX = true
		
	} else if IS_string {
		SPECIFIC_STRING = STRING_val
		USE_KEYWORD = true
	}
		
	for param_IND, arg := range params {
		
		var TMPOBJ PARAM_OBJ
		TMPOBJ.PARAM_INDEX = param_IND

		PARAM_string, FOUND_string := arg.(string)
		PARAM_int, FOUND_int := arg.(int)
		PARAM_float, FOUND_float := arg.(float64)
		PARAM_bool, FOUND_bool := arg.(bool)
		PARAM_struct, FOUND_struct := arg.(interface{})
		
		if FOUND_string {
			TMPOBJ.PARAM_TYPE = "string"
			TMPOBJ.PARAM_VALUE = PARAM_string
			ALL_PARAMS = append(ALL_PARAMS, TMPOBJ)
			
			continue
		}
		if FOUND_int {
			TMPOBJ.PARAM_TYPE = "int"
			TMPOBJ.PARAM_VALUE = PARAM_int
			ALL_PARAMS = append(ALL_PARAMS, TMPOBJ)

			continue
		}
		if FOUND_float {
			TMPOBJ.PARAM_TYPE = "float"
			TMPOBJ.PARAM_VALUE = PARAM_float
			ALL_PARAMS = append(ALL_PARAMS, TMPOBJ)

			continue
		}
		if FOUND_bool {
			TMPOBJ.PARAM_TYPE = "bool"
			TMPOBJ.PARAM_VALUE = PARAM_bool
			ALL_PARAMS = append(ALL_PARAMS, TMPOBJ)

			continue
		}

		if FOUND_struct {
			TMPOBJ.PARAM_TYPE = "struct_or_list"
			TMPOBJ.PARAM_VALUE = PARAM_struct
			ALL_PARAMS = append(ALL_PARAMS, TMPOBJ)
		}		
	}

	if BE_VERBOSE {
		C.Println(" Now in GET_FUNC_PARAM!!")		
		C.Println(" The Following Params were PASSED: ", len(ALL_PARAMS))			

		for _, x := range ALL_PARAMS {

			C.Println("TYPE: ", x.PARAM_TYPE)
			Y.Println("INDEX: ", x.PARAM_INDEX)
			G.Println(x.PARAM_VALUE)
			G.Println("")
		}
		PressAny()
	}



	//2. PHASE 2.. Look for a speciifc parameter to retun

	var RESULT PARAM_OBJ
	var FOUND = false
	for _, obj := range ALL_PARAMS {

		if USE_INDEX {
			if INDEX_to_FIND >= 0 {
				tmpindex := obj.PARAM_INDEX

				if tmpindex == INDEX_to_FIND {
					RESULT = obj
					FOUND = true
					break
				}
			}
			continue
		}

		if USE_KEYWORD {
			// Or if a specific STRING pattern was passed.. and the param value EXACTLY matches that pattern.. we return it (so we can do something about it)
			if SPECIFIC_STRING != "" {
				if obj.PARAM_TYPE == "string" {
					if obj.PARAM_VALUE == SPECIFIC_STRING {
						RESULT = obj
						FOUND = true
						break
					}
				}
			}
			continue
		}
	} //end of for


	// FINALLy.. return found, the param value, and optionally ALL_PARAMS objects
	if FOUND {
		if BE_VERBOSE {
			SHOW_BOX (" YAY! FOund the PARAM you were looking for!!")
			C.Print("Found at INDEX: ")
			G.Println(RESULT.PARAM_INDEX)
			W.Print(BOX_INDENT_SPACES, "TYPE: ")
			Y.Println( RESULT.PARAM_TYPE)
			W.Print(BOX_INDENT_SPACES, "VALUE: ")
			G.Println(RESULT.PARAM_VALUE)

			if RESULT.PARAM_TYPE == "struct_or_list" {
				M.Println("")
				M.Println(BOX_INDENT_SPACES, "Looks like this might be a STRUCT (golang Object)")
				Y.Println(BOX_INDENT_SPACES, "You should be able to ASSIGN this to whatever ")
				Y.Println(BOX_INDENT_SPACES, "custom Struct that is defined that MATCHES IT!!")
				PressAny()
			}
			
		}
	}	

	return FOUND, ALL_PARAMS

}


/*
  This padds a string that is passed and makes it a total LENGH of TOTAL_FINAL_LEN
  Pass a character as a string param and it will pad with THAT char instead of space
*/
func PAD_STRING(input string, TOTAL_FINAL_LEN int, ALL_PARAMS ...string) string {

	// This is the default padd string
	var padString = " "

	//2. If you want to use another just pass it as a string
	for p, VAL := range ALL_PARAMS {

		if p == 0 {
			padString = VAL

			//2b. also add space to before and after input so the pad string char isnt running up against it
			input = " " + input + " "
		}

	}	

	// Courtesy of: https://gist.github.com/asessa/3aaec43d93044fc42b7c6d5f728cb039

	var padLength = TOTAL_FINAL_LEN
	var inputLength = len(input)
	var padStringLength = len(padString)

	length := (float64(padLength - inputLength)) / float64(2)
	repeat := math.Ceil(length / float64(padStringLength))
	output := strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]

	return output
}


// This has limited color support and if it receives a string with the color in |cyan| format..it prints in that color
func helper_SHOW_with_COLOR(input string, SHOW_OUTPUT bool) (string, string) {

	// var JUST_STRING = temps[2]
	var JUST_STRING = input
	var COLOR = ""

	if strings.Count(JUST_STRING, "|") == 2 {
		temps := strings.Split(input, "|")
		COLOR = temps[1]
		JUST_STRING = temps[2]

		if SHOW_OUTPUT {

			switch COLOR {
			case "cyan":
				C.Print(JUST_STRING)
				break

			case "green":
				G.Print(JUST_STRING)
				break

			case "yellow":
				Y.Print(JUST_STRING)
				break
			case "red":
				R.Print(JUST_STRING)
				break				
			default:
				W.Print(JUST_STRING)
				break
			}
		}

		COLOR = "|" + COLOR + "|"

	} else {
		if SHOW_OUTPUT {
			W.Print(JUST_STRING)
		}
	}

	return JUST_STRING, COLOR

} // end of


// Alias to SHOW_BOX
func SHOW_BOX_MESSAGE (ALL_PARAMS ...string) {
	SHOW_BOX(ALL_PARAMS...)
}

// Alias to SHOW_BOX
func SHOW_MESSAGE_BOX (ALL_PARAMS ...string) {
	SHOW_BOX(ALL_PARAMS...)
}

/*
	This is a nice way of showing a message in a box
	Just pass each line you want in the box as a seperate parameter

╭――――――――――――――――――╮
│                  │
│                  │
│                  │
╰――――――――――――――――――╯
*/
var BOX_INDENT_SPACES = "         "
func SHOW_BOX(ALL_PARAMS ...string) {

	var lines []string

	var SPACE_PREFIX = "       "

	//1. if multiple lines are passed, lets iterate through them
	for _, VAL := range ALL_PARAMS {
		lines = append(lines, VAL)
	}

	//2. FIgure out which line is the LONGEST.. this is how we grow the box length
	largest_len := 0

	for _, l := range lines {

		var JUST_STRING, _ = helper_SHOW_with_COLOR(l, false)
		temp_len := len(JUST_STRING)

		if temp_len > largest_len {
			largest_len = temp_len
		}
	} //end of line len determine for

	largest_len += len(SPACE_PREFIX) + 4

	//3. Now drop top of box
	var BOX_TOP = "┌"
	var BOX_BOTTOM = "└"
	for x := 0; x < largest_len; x++ {
		BOX_TOP += "─"
		BOX_BOTTOM += "─"
	}
	//4. CLose the ends of the BOX
	BOX_TOP += "┐"
	BOX_BOTTOM += "┘"

	//5. MUST use the utf8 way to get the string length since it contains ASCII chars
	var BOX_LEN = utf8.RuneCountInString(BOX_TOP)
	BOX_LEN = BOX_LEN - 2 // We have to do BOXLEN-2 to account for the Right and Left angle Brakcets

	//6. Top of Box
	M.Println(SPACE_PREFIX + BOX_TOP)

	//7. Prints the Lines in between top and bottom
	for _, line := range lines {

		var temp_full_line, MCOLOR = helper_SHOW_with_COLOR(line, false)

		// Most likely the temp_full line is LESS than the BOX_LEN.. so lets padd it
		if len(temp_full_line) < BOX_LEN {
			temp_full_line = PAD_STRING(temp_full_line, BOX_LEN)
		}

		M.Print(SPACE_PREFIX + "│")

		helper_SHOW_with_COLOR(MCOLOR+temp_full_line, true)

		M.Println("│")
	}

	//8. Prints the BOTTOM of box.. we are DONE
	M.Println(SPACE_PREFIX + BOX_BOTTOM)

	//9. Add an indent so the next thing that is C.Println'd ... will be indented under the box
	C.Print(BOX_INDENT_SPACES)

} //end of SHOW_BOX_MESSAGE



func DO_EXIT() {

	SHOW_BOX("|red| DO_EXIT - DEBUG Forced Program Exit")
	os.Exit(ERROR_EXIT_CODE)
}



/*
   ADD_LEADING_ZERO: This takes in a number and returns a string with a leading 0
   If the number is already 10 or greater, it returns that same number as is
 
	SHOW_PRETTY_DATE is dependant on this
*/
func ADD_LEADING_ZERO( myNum int) string {

	RESULT := strconv.Itoa(myNum)

	if myNum <= 9 {
		RESULT = "0" + RESULT
	}

	return RESULT
}


// Easy way to get the UTC form of DATE_OBJ so there is no confusion.. Returns Time, String(pretty date) and Weekday all converted from the orig time
func GET_DB_DATE_UTC(input_DATE_OBJ time.Time) (time.Time, string, string, string) {

	result_DATE_OBJ := input_DATE_OBJ.In(UTC_OBJ)

	pretty, weekday := SHOW_PRETTY_DATE(result_DATE_OBJ)

	result_as_STRING := result_DATE_OBJ.String()

	return result_DATE_OBJ, result_as_STRING, pretty, weekday

}



/* SHOW_PRETTY_DATE Takes in a time.Time DATE_OBJ and returns a PRETTY formatted based on what you specify
   - Returns a STRING and a WEEKDAY
   - needs ADD_LEADING_ZERO
*/
func SHOW_PRETTY_DATE(input_DATE time.Time, EXTRA_ARGS...string) (string, string) {
	var output_FORMAT = "short"
	var SHOW_SECONDS = false

	//1. Parse out EXTRA_ARGS
	for _, VAL := range EXTRA_ARGS {

		//1c. If sec or seconds is passed, we also will show the seconds
		if VAL == "sec" || VAL == "seconds" {
			SHOW_SECONDS = true
			continue

		//1e. If short is passed, we show this format: Wednesday, 11/20/2001
		// If full is passed, we show this format: Wednesday, 11/20/2020 @ 13:56
		// if british or iso is passed, we show: 2015-05-30
		} else if VAL != "" {
			output_FORMAT = VAL
			continue
		}

	} // end of for

	//2. From this object, extract the M/D/Y HH:MM
	montemp := int(input_DATE.Month())
	daytemp := input_DATE.Day()

	hourtemp := input_DATE.Hour()
	mintemp := input_DATE.Minute()

	//3. Then, we add leading 0's as needed
	cMon := ADD_LEADING_ZERO(montemp)
	cDay := ADD_LEADING_ZERO(daytemp)
	cHour := ADD_LEADING_ZERO(hourtemp)
	cMin := ADD_LEADING_ZERO(mintemp)
	
	sectemp := input_DATE.Second()
	cSec := ADD_LEADING_ZERO(sectemp)


	//4. Thankfully we dont have to worry about this fuckery with the year!
	cYear := strconv.Itoa(input_DATE.Year())
	weekd := input_DATE.Weekday().String()

	/* 7. Here is the DEFAULT Pretty format that is returned

		09/26/1978 @ 13:58

			or (if SHOW_SECONDS is passed) 

		09/26/1978 @ 13:58:05
	*/
	result_TEXT := cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin
	if SHOW_SECONDS {
		result_TEXT += ":" + cSec
	}

	//8. SHORT Format is:  Wednesday, 11/20/2001
	if output_FORMAT == "short" {

		result_TEXT = weekd + ", " + cMon + "/" + cDay + "/" + cYear

	//9. FULL Format: //Wednesday, 11/20/2020 @ 13:56 EST (-5 Hours)
	} else if output_FORMAT == "full" {
		
		result_TEXT = weekd + ", " + cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin

		if SHOW_SECONDS {
			result_TEXT += ":" + cSec
		}

		result_TEXT += " " + ZONE_FULL
	
	//10. This is the british/iso format: 2020-09-26
	} else if output_FORMAT == "british" || output_FORMAT == "iso" {

		result_TEXT = cYear + "-" + cMon + "-" + cDay

	//11. This is JUSTDATE:  09/26/1988
	} else if output_FORMAT == "justdate" {

		result_TEXT = cMon + "/" + cDay + "/" + cYear
	
	
	//12. For use as a simple timestamp for a file suffix
	} else if output_FORMAT == "timestamp" {
	
		result_TEXT = weekd + "_" + cMon + "_" + cDay + "_" + cYear + "_" + cHour + "_" + cMin
	
	//13. SAME...but just the date
	} else if output_FORMAT == "datestamp" {
	
		result_TEXT = weekd + "_" + cMon + "_" + cDay + "_" + cYear

	}

	//12. As a bonus, we always return the weekday as a second variable

	return result_TEXT, weekd
} //end of func



// This utilizes the HUMANIZE library and shows a HUMAN readable number of the passed variable
func ShowNum(innum int) string {

	result := humanize.Comma(int64(innum))

	// result = strconv.Itoa(result)
	return result
}


// Shows a pretty Number based on passed FLOAT
func ShowNum_FLOAT(innum float64) string {

        result := humanize.Comma(int64(innum))

        // result = strconv.Itoa(result)
        return result
}




// 64bit version of this.. not sure why im using this yet
func ShowNum64(innum int64) string {

	result := humanize.Comma(innum)

	// result = strconv.Itoa(result)
	return result
}

// When called, copies a specified string to the users CLIPBOARD
func CLIPBOARD_COPY(instring string) {
	clipboard.WriteAll(instring)
}

// This initializes some pretty necessary timezone defaults. 
func SET_TIMEZONE_DEFAULTS() {

	//1. By Deffault this will be Local
	ZONE_UPPER = "Local"		

		   if USE_EST {  DEFAULT_ZONE_LOCATION_OBJ = EST_OBJ
	} else if USE_CST {  DEFAULT_ZONE_LOCATION_OBJ = CST_OBJ
	} else if USE_MST {  DEFAULT_ZONE_LOCATION_OBJ = MST_OBJ
	} else if USE_PST {  DEFAULT_ZONE_LOCATION_OBJ = PST_OBJ
	// if needed, add more here.. before we hit UTC
	} else if USE_UTC {	 DEFAULT_ZONE_LOCATION_OBJ = UTC_OBJ   }
 
	//2. Also lets get the Timezone and OFFSET info	
	t := time.Now().In(DEFAULT_ZONE_LOCATION_OBJ)
	curr_zone, offset := t.Zone()

	//2b. see if it is negative, we need to convert this temporarily
	hprefix := "+"
	
	if offset < 0 {
		fixed_off := math.Abs(float64(offset))

		//2b. Convert fixed_off to an integar
		entry_FIXED_text := strconv.FormatFloat(fixed_off, 'f', 2, 64)
		entry_NUM_text := strings.Replace(entry_FIXED_text, ".", "", -1)
		offset, _ = strconv.Atoi(entry_NUM_text)

		hprefix = "-"
	}


	//3. NOw convert the offset seconds to hours
	off_hours := (offset / 60) / 60
	offstring := hprefix + strconv.Itoa(off_hours) + " hours"

	//4. Set ZONE_UPPER .. we use this in other places. 
	ZONE_UPPER = curr_zone
	ZONE_LOCAL = t.Location().String()
	
	ZONE_HOUR_OFFSET = offstring
	ZONE_FULL = " (" + curr_zone + " " + offstring + ")"
	
	//4. Get the current time to set the global time vaiables
	GLOBAL_CURR_DATE, GLOBAL_DATE_OBJ = GET_CURRENT_TIME("full")
	GLOBAL_CURR_TIME = GLOBAL_CURR_DATE
	GLOBAL_TIME_OBJ = GLOBAL_DATE_OBJ


} //end of func


// Shows the Local TIMEZONE info .. or if --zone is specified, uses that
func SHOW_ZONE_INFO() {

	W.Print("      [ Timezone: ")
	if ZONE_LOCAL == "Local" {
		C.Print(ZONE_LOCAL)
	}
	Y.Print(" " + ZONE_UPPER)
	G.Print(" (", ZONE_HOUR_OFFSET, ") ")
	W.Println(" ]")
} //end of func


/*

	Gets the current Time/Date ...defaults to LOCAL
	However if you specify EST/PST/MST ...you get THAT time
*/
func GET_CURRENT_TIME(EXTRA_ARGS ...string) (string, time.Time) {

	//1. Default ot the local machines time zone
	dateOBJ := time.Now()
	dateOBJ = dateOBJ.In(DEFAULT_ZONE_LOCATION_OBJ)
	var output_FORMAT = "full"		// Full is the default format we return the time string in

	//2. Now, see if flags were specified. Iterate through them
	for _, VAL := range EXTRA_ARGS {

		VAL = strings.ToLower(VAL)

		switch VAL {
			case "est":
				dateOBJ = dateOBJ.In(EST_OBJ)

			case "cst":
				dateOBJ = dateOBJ.In(CST_OBJ)

			case "mst":
				dateOBJ = dateOBJ.In(MST_OBJ)
			case "mdt":
				dateOBJ = dateOBJ.In(MST_OBJ)							

			case "pst":
				dateOBJ = dateOBJ.In(PST_OBJ)

			case "utc":
				dateOBJ = dateOBJ.In(UTC_OBJ)
	
		} //end of switch

		//3. If full, short, british or iso specified, set the output format
		if VAL == "short" || VAL == "full" || VAL == "british" || VAL == "iso" || VAL == "justdate" {
			output_FORMAT = VAL
		}

	} //end of for

	result, _ := SHOW_PRETTY_DATE(dateOBJ, output_FORMAT)

	return result, dateOBJ

} //end of func

// Shows the amount of time a program ran (and start and end time)
func SHOW_START_and_END_TIME() {

	endTime, endOBJ := GET_CURRENT_TIME()
	difftemp := endOBJ.Sub(PROG_START_TIMEOBJ)
	TIME_DIFF := difftemp.String()
	


	Y.Println("\n\n ****************************************************** ")


	W.Print("              Start Time:")
	B.Println(" " + PROG_START_TIME)
	Y.Print("                End Time:")
	M.Println(" " + endTime)
	C.Print("      Total PROGRAM DURATION: ")
	G.Println(" ", TIME_DIFF)
	C.Println("******************************************************")
}

var SPINNER_SPEED = 100
var SPINNER_CHAR = 4
var spinOBJ = spinner.New(spinner.CharSets[14], 100*time.Millisecond)

//	Creates a cool "im busy right now" status spinner so you know the program is running 
func START_Spinner() {

	sduration := time.Duration(SPINNER_SPEED)

	spinOBJ = spinner.New(spinner.CharSets[SPINNER_CHAR], sduration*time.Millisecond)
	spinOBJ.Start()
}

func STOP_Spinner() {

	spinOBJ.Stop()
}

// this is a simple sleep function
func Sleep(seconds int, ALL_PARAMS ...bool) {

	var showOutput = false

	for x, BOOL_VAL := range ALL_PARAMS {

		//1. First Param is allthat is used
		if x == 0 {
			showOutput = BOOL_VAL
			continue
		}
	} // end of for	

	if showOutput == true {
		secText := ""
		suffix := "seconds"
		sectemp := seconds

		if seconds >= 119 {
			sectemp = seconds / 60
			suffix = "minutes"
		}
		secText = strconv.Itoa(sectemp)
		C.Println("        ** Sleeping for: " + secText + " ", suffix, "...")
	}

	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)

} //end of sleep function


var SHOW_WHAT_WAS_TYPED = false
func GET_USER_INPUT() string {
	reader := bufio.NewReader(os.Stdin)
	userTEMP, _ := reader.ReadString('\n')
	userTEMP = strings.TrimSuffix(userTEMP, "\n")

	if SHOW_WHAT_WAS_TYPED {
		Y.Print("\n     You Typed: ")
		W.Print(userTEMP)
		Y.Println("**")
	}

	return userTEMP

} //end of

func GET_INPUT() string {
	return GET_USER_INPUT()
}

// This takes IN a string and returns a shuffle of the characters contained in it
func SHUFFLE_STRING(input_STRING string) string {

	//1. Get the length of the string
	slen := len(input_STRING)

	stringRUNE := []rune(input_STRING)

	shuffledString_RESULT := make([]rune, slen)

	for i := range shuffledString_RESULT {
		shuffledString_RESULT[i] = stringRUNE[rand.Intn(slen)]
	}
	return string(shuffledString_RESULT)
} // end of genSESSION

func VERIFICATION_PROMPT(warning_TEXT string, required_input string) {

	M.Println("\n      - - - - - - - - WARNING - - - - - - - - - - - - - -")
	
	for x := 0; x < 3; x++ {
		C.Println("")
		C.Println("      ", warning_TEXT)
		C.Print("       Type: ")
		G.Print(required_input)
		C.Println(" To Continue")
		Y.Print("       RESPONSE: ")
		userResponse := GET_USER_INPUT()

		if strings.Contains(userResponse, required_input) {
			return
		} else {
			R.Println("\n ! ! ! ! ! ! INVALID RESPONSE  ! ! ! ! ! !")
		    M.Println("\n     - - - - - - - - - - - - - - - - - - - - - - - - -")			
		}
	} //end of for
	

	//2. If we get this far without a valid response, we will exit the program without proceeding
	os.Exit(ERROR_EXIT_CODE)


} //end of prompt

func PROMPT(warning_TEXT string, required_input string) {
	VERIFICATION_PROMPT(warning_TEXT, required_input)
}



/*
 NeW VERSION that accepts Multi Param Types 
 Test with:  ERROR_FOUND_TEMP(true, "exitForMe", 5)
 */
func ERROR_FOUND_TEMP(ALL_PARAMS ...interface{}) {


    for _, param := range ALL_PARAMS {

		Y.Print("PARAM: ")
		W.Println(param)
    }
}



// Useful error handling object.. alternete way of doing if err != nil
func ERROR_FOUND(err error, ALL_PARAMS ...string) bool {

	var MESSAGE = ""
	exit_on_error := false

	for x, VAL := range ALL_PARAMS {

		//1. First Param is the extra message to append to ERROR
		if x == 0 {
			MESSAGE = VAL
			continue
		}

		//2. Second param if specified will make the program OS_EXIT
		if x == 1 && VAL == "yes" || VAL == "exit" {
			exit_on_error = true
			continue
		}
	} // end of for


	if err != nil {

		R.Println(" = = =")
		R.Print(" = = ERROR_FOUND Error: ")
		M.Println(MESSAGE)
		Y.Println(err)
		R.Println(" = = =")

		//3. Finally if this is specified, we exit the whole program

		if exit_on_error {
			os.Exit(-9)
		}

		return true
	}

	return false
}


var PAGE_COUNT = 0
var PAGE_MAX = 5

// This is a basic Paging routine that prompts you to PressAny key
// after x number of items have been shown
func Pager(tmax int) {
	PAGE_MAX = tmax
	PAGE_COUNT++

	if PAGE_COUNT == PAGE_MAX {
		C.Print("   - - PAGER - -")
		PressAny()
		PAGE_COUNT = 0
	}

} //end of Pager

// Simple PressAny Key function
func PressAny() {

	W.Println("")
	W.Println("         ...Press Enter to Continue...")
	W.Println("")

	//1. New way of doing PAK
	b := make([]byte, 10)
	if _, err := os.Stdin.Read(b); err != nil {
		R.Println("Fatal error in PressAny Key: ", err)
	}

} // end of func


// This gets the platform we are running on (mac, linux, windows)
func GET_CURRENT_OS_INFO() {

	if runtime.GOOS == "linux" {
		OSTYPE="Linux"

	//2. Otherwise see if this is MAC
	} else if runtime.GOOS == "darwin" {
		OSTYPE="MAC"

	//3. otherwise.. its windows.. it wins by default!!	
	} else if runtime.GOOS == "windows" {
		OSTYPE="Windows"

	//4. If we get this far, means we have some weird unrecognizable OS:
	} else {
		OSTYPE="- - UNKNOWN OS - -"
	}

	//5. Another courtesy Alias
	CURRENT_OS = OSTYPE
	GOOS_VALUE = runtime.GOOS

} //end of getOsType
func Show_TOTAL_PROG_RUNTIME() {

	endTime, endOBJ := GET_CURRENT_TIME()

	dtemp := endOBJ.Sub(PROG_START_TIMEOBJ)
	DIFF := dtemp.String()

	//12. DISPPLAY Status on this Threads Performance (and metrics)
	R.Println("")
	W.Print(" = = = = = = = = = = = = = ")
	G.Print(" Program run Complete! ")
	W.Println(" = = = = = = = = = = = = = ")

	
	C.Println("")
	W.Print("         STARTED:")
	B.Println(" " + PROG_START_TIME)
	Y.Print("        ENDED on:")
	M.Println(" " + endTime)
	C.Print("  Total DURATION: ")
	G.Println(DIFF)

	W.Println(" = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")

}

// Returns a randomly generated number within a given range (returns a STRING AND an int)
func GenRandomRange(min int, max int) (int, string) {

	resultNum := rand.Intn(max-min) + min
	resultText := strconv.Itoa(resultNum)


	// Always return a string with a 0 prefix
	if resultNum < 10 {
		resultText = "0" + resultText
	}


	return resultNum, resultText

} //end of genRandomRange


// Return a random delay when we hit too many CALLS

func GENERATE_RETRY_SLEEP_and_WAIT(retry_count int, MAX_SLEEP_VAL int, MAX_RETRY_ATTEMPTS int) {

	rnum, sec_text := GenRandomRange(5, MAX_SLEEP_VAL)
	retry_count = retry_count + 1
	parent_FUNC := GET_PARENT_FUNC(2)	// play with the integer you ass to get the proper calling function name

	retry_text := ShowNum(retry_count)
	max_text := ShowNum(MAX_RETRY_ATTEMPTS)

	SHOW_BOX_MESSAGE("|red|RETRY ERROR in: ", "|yellow|" + parent_FUNC, "Retry in " + sec_text + " seconds..", "|yellow|" + retry_text + " of " + max_text )

	Sleep(rnum, true)

}


// This gets the name of the PARENT calling function
// pass dnum of 1 to it for starters.. if you arent getting the right caller name.. increment by 1
func GET_PARENT_FUNC(dnum int) string {

	var result = ""
	
	pc, _, _, ok := runtime.Caller(dnum)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		result = details.Name()
	}	

	return result
} //end of




// This is a simplified way of executing external commands... just pass the command and its parameters, it returns the output.. (or the error)
func ComExec(command_INPUT string, VERBOSE_MODE bool) (string, []string) {

	parts := strings.Fields(command_INPUT)
	head := parts[0]
	parts = parts[1:len(parts)]

	if VERBOSE_MODE {
		Y.Print("EXECUTING: ")
		justString := strings.Join(parts, " ")
		W.Print(head + " ")
		Y.Println(justString)

	}

	cmd := exec.Command(head, parts...)
	etext := ""
	out, err := cmd.CombinedOutput()
	if err != nil {

		Y.Println("\n     WARNING: Execution Error!", err.Error())
		Y.Println("")
		etext = err.Error()
	}

	otemp := string(out)
	outText := strings.Split(otemp, "\n")

	outText = append(outText, etext)

	if VERBOSE_MODE {
		M.Println(otemp)
	}

	return otemp, outText

} // end of comExec

func GET_VAR_TYPE(ALL_PARAMS ...interface{}) string {

	for _, param := range ALL_PARAMS {
		_, IS_INT := param.(int)
		_, IS_FLOAT := param.(float64) 
		_, IS_STRING := param.(string) 

		if IS_INT {
			return "int"
		}

		if IS_FLOAT {
			return "float"
		}

		if IS_STRING {
			return "string"
		}
	}

	return " ( Unknown Type ) "
}


func IS_STRING(param interface{}) bool {

	result := GET_VAR_TYPE(param)
	if result == "string" { return true }
	return false
}
func IS_INT(param interface{}) bool {

	result := GET_VAR_TYPE(param)
	if result == "int" { return true }
	return false
}
func IS_FLOAT(param interface{}) bool {

	result := GET_VAR_TYPE(param)
	if result == "float" { return true }
	return false
}

// These are the characters used to generate the serial
//var sessTEMPLATE = []rune("GRZBJHUFLEKVXMNTQPSOADWYC527183469")

// This generates a serial.. usually used discern between multiple execution runs like in jenkins
func GenSerial(serial_length int) {

	result := SHUFFLE_STRING("grzbjhuflcekivxmntqpsoadwy527183469")

	part_ONE := result[0:4]
	part_TWO := result[3:serial_length]

	SERIAL_NUM = part_ONE + "-" + part_TWO

} // end of GenSerial



var ENABLE_DEFAULT_PARAMS = false

func SETUP_DEFAULT_COMMAND_LINE_PARAMS() {

	// This is too useful NOT to have everywhere
	
	flag.BoolVar(&DEBUG_MODE,       "debug", DEBUG_MODE,         "  If specified we run in DEBUG MODE. Code that checks for this run when this is set")

	//1. First lets get some default command line params we always use
	if ENABLE_DEFAULT_PARAMS {
		//2. These are variables used for modifying the TIMEZONE (if ever even needed)
		flag.BoolVar(&USE_EST, "est", USE_EST,         "  Force Timezone to be EST")
		flag.BoolVar(&USE_CST, "cst", USE_CST,         "  Force Timezone to be CST")
		flag.BoolVar(&USE_CST, "mst", USE_MST,         "  Force TZ to be MST/Mountain")
		flag.BoolVar(&USE_CST, "mdt", USE_MST,         "  Force TZ to be MST/Mountain")
		flag.BoolVar(&USE_PST, "pst", USE_PST,         "  Force Timezone to be PST")
		flag.BoolVar(&USE_UTC, "utc", USE_UTC,         "  Force Timezone to be UTC")

		//3. In case we want to show serial numbers the program automatically generates program runs	
		flag.BoolVar(&SHOW_SERIAL, "serial", SHOW_SERIAL, "  If specified we Show a RUNTIME serial number")
	}
	
	//4. And finally, Very important.. This is the final flag.Parse that is run in the program
	flag.Parse()

} //end of setup default command line params


/* 
  MUST ALWAYS CALL THIS in the MAIN of every program.. 
  This is how command line params get initted
   Also make sure it is the LAST 'init' type function called (for example.. BEFORE AWS_INIT
*/
func MASTER_INIT(PROGNAME string, ALL_PARAMS ...float64) {

	var VERSION = DEFAULT_FLOAT

	//2. Now, see if version was passed
	for x, VAL := range ALL_PARAMS {

		if x == 0 {
			VERSION = VAL
		}

	} //end of for


	//1. Setup default COmmand line params
	SETUP_DEFAULT_COMMAND_LINE_PARAMS()

	//2. Load defaults we need for proper Timezone and DateMath operations
	SET_TIMEZONE_DEFAULTS()

	//3. And, Always init the random number seeder
	rand.Seed(time.Now().UTC().UnixNano())	

	//3b. And get current OS Data
	GET_CURRENT_OS_INFO()

	//4. Setup the prog start time globals
	PROG_START_TIME, PROG_START_TIMEOBJ = GET_CURRENT_TIME()
	
	SHOW_BOX_MESSAGE(PROGNAME, "|yellow|Running On: " + CURRENT_OS)

	if VERSION != DEFAULT_FLOAT {
		C.Print("                 Version: ")	
		Y.Println(VERSION)
		
	}
	SHOW_ZONE_INFO()

	//7. If specified we show the unique execution serial. This is useful when running from within Jenkins
	if SHOW_SERIAL {
		GenSerial(10)
		SHOW_BOX_MESSAGE("Generated EXEC Serial: ", "|cyan|" + SERIAL_NUM)
	}


	if DEBUG_MODE {
		SHOW_BOX_MESSAGE("|yellow|RUNNNING in DEBUG MODE")
	}
	

	W.Println("")

} //end of func


/* 
   Kept here as filler /template /example / Reference
.. anything you put in this function will run  when the module is imported

*/
func init() {
	
	// Add stuff here

} // end of main
