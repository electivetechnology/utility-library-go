package sanitise

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Phone(input string, defaultCountry string) string {
	log.Printf("Input: %v , DefaultCountry: %v", input, defaultCountry)

	output := removeAfterAlphaSlash(input)

	output = trimSpace(output)

	hasPrefix := hasPrefix(output)

	output = removeCharacters(output)

	outputCountryCode, outputAllowZero := hasCountryCode(output)
	log.Printf("Output country code: %v %v ", outputCountryCode, outputAllowZero)

	if hasPrefix && outputCountryCode != 0 {
		output = trimZeroAfterCode(output, outputCountryCode, outputAllowZero)
		log.Printf("Return hasExt && hasCountryCode: %v", output)
		return output
	}

	outputAndDefault := defaultCountry + output

	defaultCountryCode, defaultAllowZero := hasCountryCode(outputAndDefault)
	log.Printf("Output with default country code: %v %v", defaultCountryCode, defaultAllowZero)

	if defaultCountry != "" && defaultCountryCode != 0 {
		output = trimZeroAfterCode(outputAndDefault, defaultCountryCode, defaultAllowZero)
		log.Printf("Return defaultCountry && hasCountryCode: %v", output)
		return output
	}

	if outputCountryCode != 0 {
		output = trimZeroAfterCode(output, outputCountryCode, outputAllowZero)
		log.Printf("Return hasCountryCode: %v", output)
		return output
	}

	output = trimZeroLeft(output)
	log.Printf("Return: %v", output)

	return output
}

func trimStartAlpha(r rune) bool {
	return !unicode.IsNumber(r)
}

func trimZeroAfterCode(input string, code int, allowZero bool) string {
	if allowZero {
		return input
	}
	output := trimZeroLeft(input)
	codeString := strconv.FormatInt(int64(code), 10)
	withoutCode := strings.Replace(output, codeString, "", 1)
	log.Printf("Output after withoutCode: %v", withoutCode)
	output = codeString + trimZeroLeft(withoutCode)

	return output
}

func trimSpace(input string) string {
	output := strings.TrimSpace(input)
	log.Printf("Output after trimSpace: %v", output)

	return output
}

func trimZeroLeft(input string) string {
	output := strings.TrimLeftFunc(input, TrimByZero)
	log.Printf("Output after trimZeroLeft: %v", output)

	return output
}

func removeAfterAlphaSlash(input string) string {
	reg, err := regexp.Compile("[A-Za-z//]")

	if err != nil {
		log.Fatalf("Error removeAfterAlphaSlash: ", err)
		return input
	}

	for i, r := range input {

		find := reg.MatchString(string(r))

		if find {
			log.Printf("Output after removeAfterAlphaSlash: %v", input[:i])
			return input[:i]
		}
	}
	log.Printf("Output after removeAfterAlphaSlash: %v", input)
	return input
}

func removeCharacters(input string) string {
	reg, err := regexp.Compile("[^0-9]+")

	if err != nil {
		log.Fatalf("Error removeCharacters: ", err)
		return input
	}

	output := reg.ReplaceAllString(input, "")
	log.Printf("Output after ReplaceAllString: %v", output)

	return output
}

func hasPrefix(input string) bool {
	// Check for plus or 00
	reg, err := regexp.Compile("^(\\+|00)")

	if err != nil {
		log.Fatalf("Error hasPrefix: ", err)
	}

	find := reg.MatchString(input)

	if find {
		log.Printf("Output has prefix")
		return true
	}
	log.Printf("Output has no prefix")
	return false
}

func hasCountryCode(input string) (int, bool) {

	codeList := CodeList()

	for code, regex := range codeList {
		reg, err := regexp.Compile(regex)

		if err != nil {
			log.Fatalf("Error hasCountryCode: ", err)
		}

		find := reg.MatchString(input)

		if find {
			return code, !strings.Contains(regex, "[0]?")
		}
	}

	return 0, false
}

func CodeList() map[int]string {
	codes := make(map[int]string)

	codes[998] = "^998[0]?[1-9][0-9]{8}$"
	codes[996] = "^996[0]?[3-7][0-9]{8}$"
	codes[995] = "^995[0]?[1-9][0-9]{8}$"
	codes[994] = "^994[0]?[1-9][0-9]{8}$"
	codes[993] = "^993[0]?[1-9][0-9]{7}$"
	codes[992] = "^992[0]?[1-9][0-9]{8}$"
	codes[98] = "^98[0]?[1-9][0-9]{9}$"
	codes[977] = "^977[0]?[1-9][0-9]{7,9}$"
	codes[976] = "^976[0]?[1-9][0-9]{7}$"
	codes[975] = "^975[0]?[1-9][0-9]{6,7}$"
	codes[974] = "^974[0]?[1-9][0-9]{2,7}$"
	codes[973] = "^973[0]?[1-9][0-9]{7}$"
	codes[972] = "^972[0]?[1-9][0-9]{7,8}$"
	codes[971] = "^971[0]?[1-7,9][0-9]{7,8}$"
	codes[970] = "^970[0]?[1-9][0-9]{7,8}$"
	codes[968] = "^968[0]?[1-9][0-9]{7}$"
	codes[967] = "^967[0]?[1-9][0-9]{6,8}$"
	codes[966] = "^966[0]?[1-9][0-9]{8}$"
	codes[965] = "^965[0]?[1-9][0-9]{6,7}$"
	codes[964] = "^964[0]?[1-9][0-9]{8,9}$"
	codes[963] = "^963[0]?[1-9][0-9]{6,8}$"
	codes[962] = "^962[0]?[1-9][0-9]{7,8}$"
	codes[961] = "^961[0]?[1-9][0-9]{6,7}$"
	codes[960] = "^960[0]?[1-9][0-9]{6}$"
	codes[95] = "^95[0]?[1-9][0-9]{5,9}$"
	codes[94] = "^94[0-9]{9}$"
	codes[93] = "^93[0]?[1-9][0-9]{7,8}$"
	codes[92] = "^92[0]?[1-9][0-9]{9}$"
	codes[91] = "^91[0]?[1-9][0-9]{9}$"
	codes[90] = "^90[0]?[2-5][0-9]{9}$"
	codes[886] = "^886[0]?[1-9][0-9]{7,8}$"
	codes[880] = "^880[0]?[1-9][0-9]{9}$"
	codes[86] = "^86[0]?[1-9][0-9]{8,10}$"
	codes[856] = "^856[0]?[1-9][0-9]{7,9}$"
	codes[855] = "^855[0]?[1-9][0-9]{7,8}$"
	codes[853] = "^853[0]?[1-9][0-9]{7}$"
	codes[852] = "^852[0]?[1-9][0-9]{7}$"
	codes[84] = "^84[0]?[1-9][0-9]{7,10}$"
	codes[82] = "^82[0]?[1-9][0-9]{7,9}$"
	codes[81] = "^81[0]?[1-9][0-9]{8,9}$"
	codes[77] = "^77[0-9]{9}$"
	codes[66] = "^66[0]?[2-9][0-9]{7,8}$"
	codes[65] = "^65[0]?[1-9][0-9]{7}$"
	codes[63] = "^63[0]?[9][0-9]{8}$"
	codes[62] = "^62[0]?[1-9][0-9]{6,12}$"
	codes[60] = "^60[0]?[0-9]{9,10}$"
	codes[599] = "^599[0]?[1-9][0-9]{6,7}$"
	codes[598] = "^598[0]?[1-9][0-9]{7}$"
	codes[597] = "^597[0]?[1-9][0-9]{5,6}$"
	codes[596] = "^596[0]?[1-9][0-9]{8}$"
	codes[595] = "^595[0]?[1-9][0-9]{4,9}$"
	codes[594] = "^594[0]?[1-9][0-9]{7,8}$"
	codes[593] = "^593[0]?[1-9][0-9]{7,8}$"
	codes[592] = "^592[0]?[1-9][0-9]{6}$"
	codes[591] = "^591[0]?[1-9][0-9]{6,7}$"
	codes[590] = "^590[0]?[1-9][0-9]{8}$"
	codes[58] = "^58[0]?[1-9][0-9]{9}$"
	codes[57] = "^57[0]?[1-9][0-9]{5,11}$"
	codes[56] = "^56[0]?[1-9][0-9]{7,9}$"
	codes[55] = "^55[0]?[1-9][0-9]{9,10}$"
	codes[54] = "^54[0]?[1-9][0-9]{6,11}$"
	codes[53] = "^53[0]?[1-9][0-9]{4,10}$"
	codes[52] = "^52[0]?[1-9][0-9]{5,9}$"
	codes[51] = "^51[0]?[1-9][0-9]{6,10}$"
	codes[509] = "^509[0]?[1-9][0-9]{7}$"
	codes[508] = "^508[0]?[1-9][0-9]{5}$"
	codes[507] = "^507[0]?[1-9][0-9]{6,7}$"
	codes[506] = "^506[0]?[1-9][0-9]{7}$"
	codes[505] = "^505[0]?[1-9][0-9]{7}$"
	codes[504] = "^504[0]?[1-9][0-9]{6,7}$"
	codes[503] = "^503[0]?[1-9][0-9]{7}$"
	codes[502] = "^502[0]?[1-9][0-9]{7}$"
	codes[501] = "^501[0]?[1-9][0-9]{6}$"
	codes[500] = "^500[0]?[1-9][0-9]{4}$"
	codes[49] = "^49[0]?[1-9][0-9]{2,11}$"
	codes[48] = "^48[0]?[1-9][0-9]{8}$"
	codes[47] = "^47[0]?[2-3,5-9][0-9]{7,11}$"
	codes[46] = "^46[0]?[1-9][0-9]{6,12}$"
	codes[45] = "^45[0]?[1-9][0-9]{7}$"
	codes[44] = "^44[0]?[1-9][0-9]{6,10}$"
	codes[43] = "^43[0]?[1-7][0-9]{3,10}$"
	codes[423] = "^423[0]?[1-7][0-9]{6}$"
	codes[421] = "^421[0]?[1-7][0-9]{8}$"
	codes[420] = "^420[0]?[1-7][0-9]{8}$"
	codes[41] = "^41[0]?[1-9][0-9]{8}$"
	codes[40] = "^40[0]?[1-9][0-9]{8}$"
	codes[39] = "^39[0-9]{6,11}$"
	codes[386] = "^386[0]?[1-9][0-9]{7}$"
	codes[385] = "^385[0]?[1-9][0-9]{5,9}$"
	codes[372] = "^372[0]?[3-9][0-9]{6,7}$"
	codes[371] = "^371[0]?[2-7][0-9]{7}$"
	codes[370] = "^370[0]?[1-9][0-9]{7}$"
	codes[36] = "^36[0]?[1-9][0-9]{6,9}$"
	codes[359] = "^359[0]?[1-9][0-9]{5,8}$"
	codes[358] = "^358[0-9]{6,9}$"
	codes[357] = "^357[0]?[2,9][0-9]{7}$"
	codes[356] = "^356[0]?[1-9][0-9]{7}$"
	codes[354] = "^354[0]?[3-8][0-9]{6,7}$"
	codes[353] = "^353[0]?[1,2,4-9][0-9]{7,8}$"
	codes[352] = "^352[0]?[1-9][0-9]{5,8}$"
	codes[351] = "^351[0]?[1-9][0-9]{8}$"
	codes[350] = "^350[0-9]{8}$"
	codes[34] = "^34[0]?[1-9][0-9]{2,8}$"
	codes[33] = "^33[0]?[1-9][0-9]{8}$"
	codes[32] = "^32[0]?[1-9][0-9]{7,11}$"
	codes[31] = "^31[0]?[1-9][0-9]{8}$"
	codes[30] = "^30[0]?[1-7][0-9]{8,9}$"
	codes[27] = "^27[0]?[1-9][0-9]{8}$"
	codes[264] = "^264[0]?[1-9][0-9]{7,8}$"
	codes[249] = "^249[0]?[1,9][0-9]{8}$"
	codes[234] = "^234[0]?[1-9][0-9]{7,9}$"
	codes[216] = "^216[0]?[1-9][0-9]{7}$"
	codes[20] = "^20[0]?[1-9][0-9]{9}$"
	codes[1] = "^1[0]?[2-9][0-9]{9}$"

	return codes
}
