package main

import "fmt"

func strMul(a string, b string) int {
	result := 0
	if len(a) > 0 && len(b) > 0 {
		// fmt.Println(a, " - ", b)
		countA := map[uint8]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
		for i := 0; i < len(a); i++ {
			countA[a[i]] = 1
		}
		for i := 0; i < len(b); i++ {
			if countA[b[i]] > 0 {
				return result
			}
		}
		result = len(a) * len(b)
	}
	return result
}

func maxProduct(words []string) int {
	result := 0
	if len(words) > 0 {
		for i := 0; i < len(words); i++ {
			for j := i + 1; j < len(words); j++ {
				tmp := 0
				if len(words[i]) > len(words[j]) {
					tmp = strMul(words[j], words[i])
				} else {
					tmp = strMul(words[i], words[j])
				}
				if result < tmp {
					result = tmp
				}
			}
		}

	}
	return result
}

func main() {
	words := []string{"db", "effbfccbecb", "eeafcededcdfff", "defdbbcdad", "faadbbdeacfdbc", "bceddfccbaefaabaeea", "baaeefacecafafcdafec", "aedaacbcefabcdfcfb", "dfebdeabadbdcb", "fbeebdbeebcacce", "ceffdeedabafdbbddcede", "ccceddfecca", "bdfabbd", "cecaaeaedfb", "daabceddfaeecedfcfcf", "affafffece", "ab", "bcfcdfbdfd", "beef", "baafefedc", "affb", "fcedeaecedcc", "bbdbbdeecbaa", "cedacad", "eeffcddaaacaf", "dcdaefa", "edbceeffbdacbe", "fde", "abaebbf", "aabaacf", "efcaaaaaaeafcdaaac", "accbfedfddb", "ecedbeededadaedbd", "afbeaceddfdfef", "adacfddddcfb", "dfdbfedacda", "afbcbdbdfcbaaee", "fbcbadf", "befdbbaddcdaab", "fdeedcfadcec", "dacdeecadcfeecad", "ddafabfcbb", "edddcdfddbccebe", "bdcac", "beb", "fccaaeaf", "bcbfbfdaffdc", "bbefe", "ebefcecdb", "cd", "aaebfeedaddccaecdedbd", "fd", "bb", "ebafccfaccc", "eafdbacefdbcdaeeaad", "fdedeacddcdfaf", "ceffeaefabefedffdb", "ffcfababedbbcc", "baeffdefedff", "eacfefbeebbbbbfc", "fdacee", "ffbfbeefedb", "fdcfedacebbcfbedac", "debdcdafaadcad", "bdfcfccbc", "bdaacccb", "eabab", "dddf", "aedcccbcbfdbcbff", "edfaccccdcfb", "ccadaeeacdcbbfcfdd", "aeaeffec", "db", "dbaceefbeabbfabb", "ececfffeb", "eaadecaaacdfabb", "ebdabcaedb", "acccecfabbbebeebcef", "fcfd", "aaadddfeee", "ddfbebfbeacdcfedfbbff", "bafcbeeedeadfeba", "ddadbdedfbddaafceeffb", "eccccc", "bcceeb", "abcbbcedfef", "edaeccbfed", "aaebccfcdf", "dbfdcbdbecbcecadefb", "eefeeceb", "bbd", "afececaedddfdbcdea", "dcbc", "eabcafeadadbabfeaecdd", "bddddeffbefafda", "ddbefac", "feeecbaeffadbfbbb", "bcbcbcccadfbdadadf", "fbceeeaebfbfad", "ecfbbcadeebd", "dbcadfedcafdb", "aba", "dbdafffcebddffabebbb", "dbfaff", "dbfffcdeb", "eddbbedcdffbb", "dceaacbfed", "bbfffbdcaecb", "abdbae", "cccdecadfcbafaeffe", "ebcacdbcfacffbdccccf", "febaafaed", "dcbeeafda", "ebadfdddbaeadfea", "faeddccfbbb", "adfe", "caecfdaecafabafacbeb", "ceadfdbaefceaeadb", "edcdabddfe", "abddcb", "fb", "afebffcefaaddadfececd", "cebbfdbdabfdcd", "ee", "eccfdeb", "caedeaeddececd", "afbaccabdbcdaafcaaa", "dafbbafeeb", "fabafadaad", "dbaebcec", "aacccddfea", "ecaaf", "afefacff", "dccdda", "ffcfbcffc", "bbbaa", "ccecdbbdafcfda", "dddfdcdabdfeabaaa", "acdbfecddebbdfffbb", "abcecdffadcbcbdbfef", "acabedcecddbd", "adbeedbbaacdb", "fecbeaacdcabefddfdaba", "aeaceabdcccd", "bebbcbbefadcdedbb", "fdfcfdbfff", "feeeeabdecb", "cacbfccddabccbbcc", "eedeaffcafdaccebd", "addcfeadbe", "feebbdfffbfbadefdeae", "afa", "cfeeffcd", "efebddbcf", "cbeffecdaccbecfa", "ffdec", "ececfcb", "ddebff", "dacaeaafdbaacbdbcec", "aabebffafbafffa", "acbcebefcafaaceff", "debaaeceaadccffeedc", "bdccfdaadfcafcecba", "faadedd", "becbeecfbacbadaaad", "cffebccdeedfdf", "fdedbfcdddeceb", "acceecbedaceadabbefe", "edaddcbacbcebefbefab", "fae", "baceecabbdbc", "bbdddefaaabddf", "eafedeafcd", "cecacfacbceaaba", "dbbdbead", "aceffbedb", "ddcbbadcbeefbfdfcabbc", "abaa", "ecbabfd", "adbabcbdbccdeaafb", "abafedefcbbcbfde", "ddcbcdbaeccfbc", "febddccdfdcfbcddc", "fceefeccafcfd", "fdcdc", "cabfacfccccdfbbfbbec", "ceeceebdffacaf", "fedeacebbeceecadd", "ee", "ebefaecdfdedbbaefffd", "baadadcdeffffdeafa", "bcedc", "daffeacf", "dccccffedfafecebc", "eefcebccebbb", "fefdfcddabfc", "accdfabbdafacfdfbaba", "bbf", "ddfebafbbbbaedfacf", "dcfeebcbaad", "cfaffccaeebfbffaaac", "eeaeddfecfafbecddbefc", "efbdddfdfaaebefaef", "ad", "becbbcebf", "eeefbbfccabcdd", "fcebfdeecbbccffbfafc", "caf", "bcbfdebc", "febab", "abeeefebac", "ecdbccacaeef", "ccaecbaadaffc", "aeccfecebdadbdfda", "abaebbddfeccecdfeabc", "beeaaefccdffafbf", "eebdb", "eff", "cac", "eda", "bdabbafdaa", "dccffceff", "aecfdfdacaabdceacf", "add", "eecbbfeaaaadbd", "aecfcbcddaba", "ebfefceddcaec", "cfdcdcaedffaadaab", "fedf", "dcdfdfcbfaadebeee", "aaeab", "fcfeecaceeecfb", "cadbedeccfefefaabddc", "bbceeebcaf", "beecbdda", "bcbabceefa", "abca", "cbafb", "cbabefddfadd", "dfffdaabdbfcffa", "cfbe", "efcebadeea", "cddad", "ceadfadfccf"}
	fmt.Println(maxProduct(words))
}
