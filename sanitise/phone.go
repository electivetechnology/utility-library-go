package sanitise

import (
	"regexp"
	"strings"
)

func Phone(input string, defaultCountry string) string {
	log.Printf("Input: %v , DefaultCountry: %v", input, defaultCountry)

	output := removeAfterAlphaSlash(input)

	output = trimSpace(output)

	hasPrefix := hasPrefix(output)

	output = removeCharacters(output)

	output = trimZeroLeft(output)

	withDefault := defaultCountry + output

	if hasPrefix && hasCountryCode(output) {
		log.Printf("return hasExt && hasCountryCode: %v", output)
		return output
	}

	if defaultCountry != "" && hasCountryCode(withDefault) {
		log.Printf("return defaultCountry && hasCountryCode: %v", withDefault)
		return withDefault
	}

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
		log.Fatalf("removeAfterAlphaSlash error", err)
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
		log.Fatalf("removeCharacters error", err)
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
		log.Fatalf("hasPrefix error", err)
	}

	find := reg.MatchString(input)

	if find {
		log.Printf("Output has prefix")
		return true
	}
	log.Printf("Output has no prefix")
	return false
}

func hasCountryCode(input string) bool {

	codeList := CodeList()

	for _, regex := range codeList {
		reg, err := regexp.Compile(regex)

		if err != nil {
			log.Fatalf("hasCountryCode error", err)
		}

		find := reg.MatchString(input)

		if find {
			return true
		}
	}

	return false
}

func CodeList() map[int]string {
	codes := make(map[int]string)

	codes[998] = "^998[1-9][0-9]{8}$"
	codes[996] = "^996[3-7][0-9]{8}$"
	codes[995] = "^995[1-9][0-9]{8}$"
	codes[994] = "^994[1-9][0-9]{8}$"
	codes[993] = "^993[1-9][0-9]{7}$"
	codes[992] = "^992[1-9][0-9]{8}$"
	codes[98] = "^98[1-9][0-9]{9}$"
	codes[977] = "^977[1-9][0-9]{7,9}$"
	codes[976] = "^976[1-9][0-9]{7}$"
	codes[975] = "^975[1-9][0-9]{6,7}$"
	codes[974] = "^974[1-9][0-9]{2,7}$"
	codes[973] = "^973[1-9][0-9]{7}$"
	codes[972] = "^972[1-9][0-9]{7,8}$"
	codes[971] = "^971[1-7,9][0-9]{7,8}$"
	codes[970] = "^970[1-9][0-9]{7,8}$"
	codes[968] = "^968[1-9][0-9]{7}$"
	codes[967] = "^967[1-9][0-9]{6,8}$"
	codes[966] = "^966[1-9][0-9]{8}$"
	codes[965] = "^965[1-9][0-9]{6,7}$"
	codes[964] = "^964[1-9][0-9]{8,9}$"
	codes[963] = "^963[1-9][0-9]{6,8}$"
	codes[962] = "^962[1-9][0-9]{7,8}$"
	codes[961] = "^961[1-9][0-9]{6,7}$"
	codes[960] = "^960[1-9][0-9]{6}$"
	codes[95] = "^95[1-9][0-9]{5,9}$"
	codes[94] = "^94[0-9]{9}$"
	codes[93] = "^93[1-9][0-9]{7,8}$"
	codes[92] = "^92[1-9][0-9]{9}$"
	codes[91] = "^91[1-9][0-9]{9}$"
	codes[90] = "^90[2-5][0-9]{9}$"
	codes[886] = "^886[1-9][0-9]{7,8}$"
	codes[880] = "^880[1-9][0-9]{9}$"
	codes[86] = "^86[1-9][0-9]{8,10}$"
	codes[856] = "^856[1-9][0-9]{7,9}$"
	codes[855] = "^855[1-9][0-9]{7,8}$"
	codes[853] = "^853[1-9][0-9]{7}$"
	codes[852] = "^852[1-9][0-9]{7}$"
	codes[84] = "^84[1-9][0-9]{7,10}$"
	codes[82] = "^82[1-9][0-9]{7,9}$"
	codes[81] = "^81[1-9][0-9]{8,9}$"
	codes[77] = "^77[0-9]{9}$"
	codes[66] = "^66[2-9][0-9]{7,8}$"
	codes[65] = "^65[1-9][0-9]{7}$"
	codes[599] = "^599[1-9][0-9]{6,7}$"
	codes[598] = "^598[1-9][0-9]{7}$"
	codes[597] = "^597[1-9][0-9]{5,6}$"
	codes[596] = "^596[1-9][0-9]{8}$"
	codes[595] = "^595[1-9][0-9]{4,9}$"
	codes[594] = "^594[1-9][0-9]{7,8}$"
	codes[593] = "^593[1-9][0-9]{7,8}$"
	codes[592] = "^592[1-9][0-9]{6}$"
	codes[591] = "^591[1-9][0-9]{6,7}$"
	codes[590] = "^590[1-9][0-9]{8}$"
	codes[58] = "^58[1-9][0-9]{9}$"
	codes[57] = "^57[1-9][0-9]{5,11}$"
	codes[56] = "^56[1-9][0-9]{7,9}$"
	codes[55] = "^55[1-9][0-9]{9,10}$"
	codes[54] = "^54[1-9][0-9]{6,11}$"
	codes[53] = "^53[1-9][0-9]{4,10}$"
	codes[52] = "^52[1-9][0-9]{5,9}$"
	codes[51] = "^51[1-9][0-9]{6,10}$"
	codes[509] = "^509[1-9][0-9]{7}$"
	codes[508] = "^508[1-9][0-9]{5}$"
	codes[507] = "^507[1-9][0-9]{6,7}$"
	codes[506] = "^506[1-9][0-9]{7}$"
	codes[505] = "^505[1-9][0-9]{7}$"
	codes[504] = "^504[1-9][0-9]{6,7}$"
	codes[503] = "^503[1-9][0-9]{7}$"
	codes[502] = "^502[1-9][0-9]{7}$"
	codes[501] = "^501[1-9][0-9]{6}$"
	codes[500] = "^500[1-9][0-9]{4}$"
	codes[49] = "^49[1-9][0-9]{2,11}$"
	codes[48] = "^48[1-9][0-9]{8}$"
	codes[47] = "^47[2-3,5-9][0-9]{7,11}$"
	codes[46] = "^46[1-9][0-9]{6,12}$"
	codes[45] = "^45[1-9][0-9]{7}$"
	codes[44] = "^44[1-9][0-9]{6,10}$"
	codes[43] = "^43[1-7][0-9]{3,10}$"
	codes[423] = "^423[1-7][0-9]{6}$"
	codes[421] = "^421[1-7][0-9]{8}$"
	codes[420] = "^420[1-7][0-9]{8}$"
	codes[41] = "^41[1-9][0-9]{8}$"
	codes[40] = "^40[1-9][0-9]{8}$"
	codes[39] = "^39[0-9]{6,11}$"
	codes[386] = "^386[1-9][0-9]{7}$"
	codes[385] = "^385[1-9][0-9]{5,9}$"
	codes[372] = "^372[3-9][0-9]{6,7}$"
	codes[371] = "^371[2-7][0-9]{7}$"
	codes[370] = "^370[1-9][0-9]{7}$"
	codes[36] = "^36[1-9][0-9]{6,9}$"
	codes[359] = "^359[1-9][0-9]{5,8}$"
	codes[358] = "^358[0-9]{6,9}$"
	codes[357] = "^357[2,9][0-9]{7}$"
	codes[356] = "^356[1-9][0-9]{7}$"
	codes[354] = "^354[3-8][0-9]{6,7}$"
	codes[353] = "^353[1,2,4-9][0-9]{7,8}$"
	codes[352] = "^352[1-9][0-9]{5,8}$"
	codes[351] = "^351[1-9][0-9]{8}$"
	codes[350] = "^350[0-9]{8}$"
	codes[34] = "^34[1-9][0-9]{2,8}$"
	codes[33] = "^33[1-9][0-9]{8}$"
	codes[32] = "^32[1-9][0-9]{7,11}$"
	codes[31] = "^31[1-9][0-9]{8}$"
	codes[30] = "^30[1-7][0-9]{8,9}$"
	codes[27] = "^27[1-9][0-9]{8}$"
	codes[264] = "^264[1-9][0-9]{7,8}$"
	codes[249] = "^249[1,9][0-9]{8}$"
	codes[234] = "^234[1-9][0-9]{7,9}$"
	codes[216] = "^216[1-9][0-9]{7}$"
	codes[20] = "^20[1-9][0-9]{9}$"
	codes[1] = "^1[2-9][0-9]{9}$"

	return codes
}
