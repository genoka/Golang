package main

import (
	"fmt"
	
	"golang.org/x/text/language"
	
	"golang.org/x/text/language/display"
	
	)
	
var userPrefs = []language.Tag{
	language.Make("gsw"),  //Swiss German
	language.Make("fr"),  //French
	}

var serverLangs = []language.Tag{
	language.AmericanEnglish,
	language.German,
}

var matcher =language.newMatcher (serverLangs)

func main()
{
	tag, index, confidence := matcher.Match(userPrefs...)
	
	fmt.Printf(" best Match : %s  (%s) index=%d  confidence=%v \n",
	
	display.English.Tags().Name(tag),
	display.Self.Name(tag),
	index, confidence )
	
	//best Match : German (DEUTSH) index=1  confidence=High
	
}

// var supported = []language.Tag{
//       language.English,            // en
//      language.French,             // fr
//      language.Dutch,              // nl
//      language.Make("nl-BE"),      // nl-BE
//      language.SimplifiedChinese,  // zh-Hans
//      language.TraditionalChinese, // zh-Hant
//      language.Russian,            // ru
//  }

//  en := display.English.Tags()
//    for _, t := range supported {
//        fmt.Printf("%-20s (%s)\n", en.Name(t), display.Self.Name(t))
//   }