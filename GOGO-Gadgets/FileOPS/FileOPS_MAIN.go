/*   GOGO File Ops GADGET - Useful file operation functions

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter. (This is a GO Thing)

	
	Aug 28, 2021    v1.23   - Initial Rollout

*/

package MODULE


import (

	// = = = = = Native Libraries

		"os"
		"net/http"
		"io"
		"io/ioutil"
		"sort"
		"bufio"
		"errors"

		// all for saving a struct to disk
		"encoding/json"
		"bytes"
		"sync"



	// = = = = = PERSONAL Libraries
		. "gopub.acedev.io/GOGO-Gadgets"
	

	// = = = = = 3rd Party Libraries

)




/*
 DownloadFile will download a url to a local file. It's efficient because it will
 write as it downloads and not load the whole file into memory.

  Courtesy of: https://golangcode.com/download-a-file-from-a-url/
*/
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}



/*
 Saves a Struct to DISK: 
 Inspired by:
   https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d
*/
var lock sync.Mutex
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
	return nil, err
	}
	return bytes.NewReader(b), nil
}



func SAVE_Struct_2_DISK(FULL_path_2_file string, v interface{}) error {
	Y.Print(" SAVING Struct to file: ")
	G.Println(FULL_path_2_file)

	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(FULL_path_2_file)
	if err != nil {
	  return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
	  return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}





func DOES_FILE_EXIST(FULL_path_2_file string) bool{

	//2. Determine if file exists:
	if _, err := os.Stat(FULL_path_2_file); err == nil {

		return true

	} else if errors.Is(err, os.ErrNotExist) {
		R.Print(" ERROR in DOES_FILE_EXIST: ")
		Y.Println(err)
		// path/to/whatever does *not* exist
		return false

	} else {
		R.Print(" ERROR in DOES_FILE_EXIST (weird ERROR):")
		Y.Println(err)
		return false
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}	

} //end of 


// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func LOAD_Struct_from_FILE(FULL_path_2_file string, v interface{}, bequiet bool) error {

	if bequiet == false {
		Y.Print(" Reading STRUCT from: ")
		W.Println(FULL_path_2_file)
	}

	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(FULL_path_2_file)
	if err != nil {
		R.Print(" ERROR! during OPEN")
		Y.Println(err)
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}






// Opens a file and returns a file object
func OPEN_FILE(path_to_file string) *os.File {

	file_obj, err := os.Open(path_to_file)
	if err != nil {
		R.Print(" ** ERROR ** Cannot open the file: ")
		W.Println(path_to_file)
		Y.Println(err.Error())
	}

	return file_obj

} //end of func

// func AppendStringToFile(path string, text string, OVERWRITE_FILE_FIRST bool) bool {
func WRITE_FILE(FULL_FILENAME string, ALL_PARAMS ...string) bool {

	var QUIET_flag = false

	var LINE_OUTPUT = ""
	var doOverwrite = false

	for x, VAL := range ALL_PARAMS {

		//1. First param specified, means we OVERWRITE the file (instead of the default which is append)
		if x == 0 {
			LINE_OUTPUT = VAL

		} else if x == 1 {
			if VAL == "overwrite" || VAL == "true" || VAL == "yes" {
				doOverwrite = true
				continue
			}			

		// If we want QUIET mode
		} else if x == 2 {
			QUIET_flag = true
		}
	} // end of for

	//2. If set, we erase the file before writing to it
	if doOverwrite {
		if FILE_EXISTS(FULL_FILENAME) {
			err := os.Remove(FULL_FILENAME)
			if err != nil {
				R.Println(" WRITE_FILE weird error checking for exists: ", err)
				return false
			}
		}
	}

	file_OBJ, err := os.OpenFile(FULL_FILENAME, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		R.Print("*** WRITE_FILE ERROR: ")
		Y.Println(err)
		return false
	}


	//3. Otherwise, write the line we have TO the file specified
	datawriter := bufio.NewWriter(file_OBJ)     
	_, _ = datawriter.WriteString(LINE_OUTPUT + "\n")
	datawriter.Flush()
	file_OBJ.Close()

	if QUIET_flag == false {
		SHOW_BOX("Writing TO file named: ", "|green|" + FULL_FILENAME)	
	}


	return true
}


// Gets a list of files in a directory
func Get_FILE_LIST(dirname_or_mask string) []string {

	var results []string

	//1. Read from the specified directory
	files, err := ioutil.ReadDir(dirname_or_mask)
	if err != nil {
		R.Print(" *** ERROR IN GET FILES LIST: ", err)
		return results
	}

	//2. Now lets add these file names to the directory list

	for _, f := range files {

		results = append(results, f.Name())

	} //end of for loop

	//5. Sort the list alphabetically
	sort.Strings(results)

	return results
} //end of get_File_LIST



func MAKE_DIR(input_DIR string) {
	SHOW_BOX(" CREATING Directory:", "|yellow|" + input_DIR)

	err := os.MkdirAll(input_DIR, os.ModePerm)
	if err != nil {
		R.Print(" Error Creating Directory: ")
		Y.Println(err)
	}
/*
	err := os.Mkdir(input_DIR, 0777)	
	// old style
	if err != nil {
		R.Println(" ERROR! Cant MAKE Directory: ", input_DIR, err)
	}
*/


}

func REMOVE_DIR(input_DIR string) {

	SHOW_BOX(" REMOVING Directory:", "|yellow|" + input_DIR)
	err := os.RemoveAll(input_DIR)	

	if err != nil {
		R.Println(" ERROR! Cant REMOVE Directory: ", input_DIR, err)
	}	
}


// This checks to see if a file or directory exists
func FILE_EXISTS(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
} // end of fileExist

func DIR_EXISTS(path string) bool {
	return FILE_EXISTS(path)
}


/* 
   Kept here as filler /template /example / Reference
.. anything you put in this function will run  when the module is imported

*/
func init() {




} // end of main
	

