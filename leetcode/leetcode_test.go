package leetcode

import (
	"testing"

	"github.com/yiranzai/go-utils/leetcode"
	"gotest.tools/v3/assert"
)

func Test_isAnagram(t *testing.T) {
	assert.Equal(t, isAnagram("first", "fisrt"), true)

	assert.Equal(t, isAnagram("firrst", "fisrtr"), true)

	assert.Equal(t, isAnagram("afirst", "fisrt"), false)

	assert.Equal(t, isAnagram("afirst", "fisrtb"), false)
}

func Test_minimumTotal(t *testing.T) {
	assert.Equal(t, minimumTotal([][]int{{-10}}), -10)
	assert.Equal(t, minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}), 11)
	assert.Equal(t, minimumTotal(bigTirangleArray), -8717)

}

//func Benchmark_minimumTotal(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		minimumTotal(bigTirangleArray)
//	}
//}

func Test_longestConsecutive(t *testing.T) {
	assert.Equal(t, longestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)
	assert.Equal(t, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), 9)
}

func Test_climbStairs(t *testing.T) {
	assert.Equal(t, climbStairs(2), 2)
	assert.Equal(t, climbStairs(3), 3)
}

func Test_canJump(t *testing.T) {
	assert.Equal(t, canJump([]int{0}), true)
	assert.Equal(t, canJump([]int{3, 2, 1, 0, 4}), false)
	assert.Equal(t, canJump([]int{2, 3, 1, 1, 4}), true)
}

func Test_jump(t *testing.T) {
	assert.Equal(t, jump([]int{0}), 0)
	assert.Equal(t, jump([]int{2, 3, 1, 1, 4}), 2)
	assert.Equal(t, jump([]int{2, 3, 2, 0, 1, 1, 1, 9, 0, 0, 0, 0, 0, 0, 1}), 6)
}

func Test_lengthOfLIS(t *testing.T) {
	assert.Equal(t, lengthOfLIS([]int{-2, -1}), 2)
	assert.Equal(t, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18, 4, 19}), 5)
}

func Test_wordBreak(t *testing.T) {

	assert.Equal(
		t,
		wordBreak(
			"fohhemkkaecojceoaejkkoedkofhmohkcjmkggcmnami",
			[]string{
				"kfomka",
				"hecagbngambii",
				"anobmnikj",
				"c",
				"nnkmfelneemfgcl",
				"ah",
				"bgomgohl",
				"lcbjbg",
				"ebjfoiddndih",
				"hjknoamjbfhckb",
				"eioldlijmmla",
				"nbekmcnakif",
				"fgahmihodolmhbi",
				"gnjfe",
				"hk",
				"b",
				"jbfgm",
				"ecojceoaejkkoed",
				"cemodhmbcmgl",
				"j",
				"gdcnjj",
				"kolaijoicbc",
				"liibjjcini",
				"lmbenj",
				"eklingemgdjncaa",
				"m",
				"hkh",
				"fblb",
				"fk",
				"nnfkfanaga",
				"eldjml",
				"iejn",
				"gbmjfdooeeko",
				"jafogijka",
				"ngnfggojmhclkjd",
				"bfagnfclg",
				"imkeobcdidiifbm",
				"ogeo",
				"gicjog",
				"cjnibenelm",
				"ogoloc",
				"edciifkaff",
				"kbeeg",
				"nebn",
				"jdd",
				"aeojhclmdn",
				"dilbhl",
				"dkk",
				"bgmck",
				"ohgkefkadonafg",
				"labem",
				"fheoglj",
				"gkcanacfjfhogjc",
				"eglkcddd",
				"lelelihakeh",
				"hhjijfiodfi",
				"enehbibnhfjd",
				"gkm",
				"ggj",
				"ag",
				"hhhjogk",
				"lllicdhihn",
				"goakjjnk",
				"lhbn",
				"fhheedadamlnedh",
				"bin",
				"cl",
				"ggjljjjf",
				"fdcdaobhlhgj",
				"nijlf",
				"i",
				"gaemagobjfc",
				"dg",
				"g",
				"jhlelodgeekj",
				"hcimohlni",
				"fdoiohikhacgb",
				"k",
				"doiaigclm",
				"bdfaoncbhfkdbjd",
				"f",
				"jaikbciac",
				"cjgadmfoodmba",
				"molokllh",
				"gfkngeebnggo",
				"lahd",
				"n",
				"ehfngoc",
				"lejfcee",
				"kofhmoh",
				"cgda",
				"de",
				"kljnicikjeh",
				"edomdbibhif",
				"jehdkgmmofihdi",
				"hifcjkloebel",
				"gcghgbemjege",
				"kobhhefbbb",
				"aaikgaolhllhlm",
				"akg",
				"kmmikgkhnn",
				"dnamfhaf",
				"mjhj",
				"ifadcgmgjaa",
				"acnjehgkflgkd",
				"bjj",
				"maihjn",
				"ojakklhl",
				"ign",
				"jhd",
				"kndkhbebgh",
				"amljjfeahcdlfdg",
				"fnboolobch",
				"gcclgcoaojc",
				"kfokbbkllmcd",
				"fec",
				"dljma",
				"noa",
				"cfjie",
				"fohhemkka",
				"bfaldajf",
				"nbk",
				"kmbnjoalnhki",
				"ccieabbnlhbjmj",
				"nmacelialookal",
				"hdlefnbmgklo",
				"bfbblofk",
				"doohocnadd",
				"klmed",
				"e",
				"hkkcmbljlojkghm",
				"jjiadlgf",
				"ogadjhambjikce",
				"bglghjndlk",
				"gackokkbhj",
				"oofohdogb",
				"leiolllnjj",
				"edekdnibja",
				"gjhglilocif",
				"ccfnfjalchc",
				"gl",
				"ihee",
				"cfgccdmecem",
				"mdmcdgjelhgk",
				"laboglchdhbk",
				"ajmiim",
				"cebhalkngloae",
				"hgohednmkahdi",
				"ddiecjnkmgbbei",
				"ajaengmcdlbk",
				"kgg",
				"ndchkjdn",
				"heklaamafiomea",
				"ehg",
				"imelcifnhkae",
				"hcgadilb",
				"elndjcodnhcc",
				"nkjd",
				"gjnfkogkjeobo",
				"eolega",
				"lm",
				"jddfkfbbbhia",
				"cddmfeckheeo",
				"bfnmaalmjdb",
				"fbcg",
				"ko",
				"mojfj",
				"kk",
				"bbljjnnikdhg",
				"l",
				"calbc",
				"mkekn",
				"ejlhdk",
				"hkebdiebecf",
				"emhelbbda",
				"mlba",
				"ckjmih",
				"odfacclfl",
				"lgfjjbgookmnoe",
				"begnkogf",
				"gakojeblk",
				"bfflcmdko",
				"cfdclljcg",
				"ho",
				"fo",
				"acmi",
				"oemknmffgcio",
				"mlkhk",
				"kfhkndmdojhidg",
				"ckfcibmnikn",
				"dgoecamdliaeeoa",
				"ocealkbbec",
				"kbmmihb",
				"ncikad",
				"hi",
				"nccjbnldneijc",
				"hgiccigeehmdl",
				"dlfmjhmioa",
				"kmff",
				"gfhkd",
				"okiamg",
				"ekdbamm",
				"fc",
				"neg",
				"cfmo",
				"ccgahikbbl",
				"khhoc",
				"elbg",
				"cbghbacjbfm",
				"jkagbmfgemjfg",
				"ijceidhhajmja",
				"imibemhdg",
				"ja",
				"idkfd",
				"ndogdkjjkf",
				"fhic",
				"ooajkki",
				"fdnjhh",
				"ba",
				"jdlnidngkfffbmi",
				"jddjfnnjoidcnm",
				"kghljjikbacd",
				"idllbbn",
				"d",
				"mgkajbnjedeiee",
				"fbllleanknmoomb",
				"lom",
				"kofjmmjm",
				"mcdlbglonin",
				"gcnboanh",
				"fggii",
				"fdkbmic",
				"bbiln",
				"cdjcjhonjgiagkb",
				"kooenbeoongcle",
				"cecnlfbaanckdkj",
				"fejlmog",
				"fanekdneoaammb",
				"maojbcegdamn",
				"bcmanmjdeabdo",
				"amloj",
				"adgoej",
				"jh",
				"fhf",
				"cogdljlgek",
				"o",
				"joeiajlioggj",
				"oncal",
				"lbgg",
				"elainnbffk",
				"hbdi",
				"femcanllndoh",
				"ke",
				"hmib",
				"nagfahhljh",
				"ibifdlfeechcbal",
				"knec",
				"oegfcghlgalcnno",
				"abiefmjldmln",
				"mlfglgni",
				"jkofhjeb",
				"ifjbneblfldjel",
				"nahhcimkjhjgb",
				"cdgkbn",
				"nnklfbeecgedie",
				"gmllmjbodhgllc",
				"hogollongjo",
				"fmoinacebll",
				"fkngbganmh",
				"jgdblmhlmfij",
				"fkkdjknahamcfb",
				"aieakdokibj",
				"hddlcdiailhd",
				"iajhmg",
				"jenocgo",
				"embdib",
				"dghbmljjogka",
				"bahcggjgmlf",
				"fb",
				"jldkcfom",
				"mfi",
				"kdkke",
				"odhbl",
				"jin",
				"kcjmkggcmnami",
				"kofig",
				"bid",
				"ohnohi",
				"fcbojdgoaoa",
				"dj",
				"ifkbmbod",
				"dhdedohlghk",
				"nmkeakohicfdjf",
				"ahbifnnoaldgbj",
				"egldeibiinoac",
				"iehfhjjjmil",
				"bmeimi",
				"ombngooicknel",
				"lfdkngobmik",
				"ifjcjkfnmgjcnmi",
				"fmf",
				"aoeaa",
				"an",
				"ffgddcjblehhggo",
				"hijfdcchdilcl",
				"hacbaamkhblnkk",
				"najefebghcbkjfl",
				"hcnnlogjfmmjcma",
				"njgcogemlnohl",
				"ihejh",
				"ej",
				"ofn",
				"ggcklj",
				"omah",
				"hg",
				"obk",
				"giig",
				"cklna",
				"lihaiollfnem",
				"ionlnlhjckf",
				"cfdlijnmgjoebl",
				"dloehimen",
				"acggkacahfhkdne",
				"iecd",
				"gn",
				"odgbnalk",
				"ahfhcd",
				"dghlag",
				"bchfe",
				"dldblmnbifnmlo",
				"cffhbijal",
				"dbddifnojfibha",
				"mhh",
				"cjjol",
				"fed",
				"bhcnf",
				"ciiibbedklnnk",
				"ikniooicmm",
				"ejf",
				"ammeennkcdgbjco",
				"jmhmd",
				"cek",
				"bjbhcmda",
				"kfjmhbf",
				"chjmmnea",
				"ifccifn",
				"naedmco",
				"iohchafbega",
				"kjejfhbco",
				"anlhhhhg",
			},
		),
		true,
	)
	assert.Equal(t, wordBreak("aaaaaaa", []string{"aaaa", "aaa"}), true)
	assert.Equal(t, wordBreak("leetcode", []string{"leet", "code"}), true)
	assert.Equal(t, wordBreak("applepenapple", []string{"apple", "pen"}), true)
	assert.Equal(t, wordBreak("catsandog", []string{"cats", "cat", "dog", "sand", "and"}), false)

	assert.Equal(
		t,
		wordBreak(
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		),
		false,
	)
}

func Test_wordBreak2(t *testing.T) {

	assert.Equal(
		t,
		wordBreak2(
			"fohhemkkaecojceoaejkkoedkofhmohkcjmkggcmnami",
			[]string{
				"kfomka",
				"hecagbngambii",
				"anobmnikj",
				"c",
				"nnkmfelneemfgcl",
				"ah",
				"bgomgohl",
				"lcbjbg",
				"ebjfoiddndih",
				"hjknoamjbfhckb",
				"eioldlijmmla",
				"nbekmcnakif",
				"fgahmihodolmhbi",
				"gnjfe",
				"hk",
				"b",
				"jbfgm",
				"ecojceoaejkkoed",
				"cemodhmbcmgl",
				"j",
				"gdcnjj",
				"kolaijoicbc",
				"liibjjcini",
				"lmbenj",
				"eklingemgdjncaa",
				"m",
				"hkh",
				"fblb",
				"fk",
				"nnfkfanaga",
				"eldjml",
				"iejn",
				"gbmjfdooeeko",
				"jafogijka",
				"ngnfggojmhclkjd",
				"bfagnfclg",
				"imkeobcdidiifbm",
				"ogeo",
				"gicjog",
				"cjnibenelm",
				"ogoloc",
				"edciifkaff",
				"kbeeg",
				"nebn",
				"jdd",
				"aeojhclmdn",
				"dilbhl",
				"dkk",
				"bgmck",
				"ohgkefkadonafg",
				"labem",
				"fheoglj",
				"gkcanacfjfhogjc",
				"eglkcddd",
				"lelelihakeh",
				"hhjijfiodfi",
				"enehbibnhfjd",
				"gkm",
				"ggj",
				"ag",
				"hhhjogk",
				"lllicdhihn",
				"goakjjnk",
				"lhbn",
				"fhheedadamlnedh",
				"bin",
				"cl",
				"ggjljjjf",
				"fdcdaobhlhgj",
				"nijlf",
				"i",
				"gaemagobjfc",
				"dg",
				"g",
				"jhlelodgeekj",
				"hcimohlni",
				"fdoiohikhacgb",
				"k",
				"doiaigclm",
				"bdfaoncbhfkdbjd",
				"f",
				"jaikbciac",
				"cjgadmfoodmba",
				"molokllh",
				"gfkngeebnggo",
				"lahd",
				"n",
				"ehfngoc",
				"lejfcee",
				"kofhmoh",
				"cgda",
				"de",
				"kljnicikjeh",
				"edomdbibhif",
				"jehdkgmmofihdi",
				"hifcjkloebel",
				"gcghgbemjege",
				"kobhhefbbb",
				"aaikgaolhllhlm",
				"akg",
				"kmmikgkhnn",
				"dnamfhaf",
				"mjhj",
				"ifadcgmgjaa",
				"acnjehgkflgkd",
				"bjj",
				"maihjn",
				"ojakklhl",
				"ign",
				"jhd",
				"kndkhbebgh",
				"amljjfeahcdlfdg",
				"fnboolobch",
				"gcclgcoaojc",
				"kfokbbkllmcd",
				"fec",
				"dljma",
				"noa",
				"cfjie",
				"fohhemkka",
				"bfaldajf",
				"nbk",
				"kmbnjoalnhki",
				"ccieabbnlhbjmj",
				"nmacelialookal",
				"hdlefnbmgklo",
				"bfbblofk",
				"doohocnadd",
				"klmed",
				"e",
				"hkkcmbljlojkghm",
				"jjiadlgf",
				"ogadjhambjikce",
				"bglghjndlk",
				"gackokkbhj",
				"oofohdogb",
				"leiolllnjj",
				"edekdnibja",
				"gjhglilocif",
				"ccfnfjalchc",
				"gl",
				"ihee",
				"cfgccdmecem",
				"mdmcdgjelhgk",
				"laboglchdhbk",
				"ajmiim",
				"cebhalkngloae",
				"hgohednmkahdi",
				"ddiecjnkmgbbei",
				"ajaengmcdlbk",
				"kgg",
				"ndchkjdn",
				"heklaamafiomea",
				"ehg",
				"imelcifnhkae",
				"hcgadilb",
				"elndjcodnhcc",
				"nkjd",
				"gjnfkogkjeobo",
				"eolega",
				"lm",
				"jddfkfbbbhia",
				"cddmfeckheeo",
				"bfnmaalmjdb",
				"fbcg",
				"ko",
				"mojfj",
				"kk",
				"bbljjnnikdhg",
				"l",
				"calbc",
				"mkekn",
				"ejlhdk",
				"hkebdiebecf",
				"emhelbbda",
				"mlba",
				"ckjmih",
				"odfacclfl",
				"lgfjjbgookmnoe",
				"begnkogf",
				"gakojeblk",
				"bfflcmdko",
				"cfdclljcg",
				"ho",
				"fo",
				"acmi",
				"oemknmffgcio",
				"mlkhk",
				"kfhkndmdojhidg",
				"ckfcibmnikn",
				"dgoecamdliaeeoa",
				"ocealkbbec",
				"kbmmihb",
				"ncikad",
				"hi",
				"nccjbnldneijc",
				"hgiccigeehmdl",
				"dlfmjhmioa",
				"kmff",
				"gfhkd",
				"okiamg",
				"ekdbamm",
				"fc",
				"neg",
				"cfmo",
				"ccgahikbbl",
				"khhoc",
				"elbg",
				"cbghbacjbfm",
				"jkagbmfgemjfg",
				"ijceidhhajmja",
				"imibemhdg",
				"ja",
				"idkfd",
				"ndogdkjjkf",
				"fhic",
				"ooajkki",
				"fdnjhh",
				"ba",
				"jdlnidngkfffbmi",
				"jddjfnnjoidcnm",
				"kghljjikbacd",
				"idllbbn",
				"d",
				"mgkajbnjedeiee",
				"fbllleanknmoomb",
				"lom",
				"kofjmmjm",
				"mcdlbglonin",
				"gcnboanh",
				"fggii",
				"fdkbmic",
				"bbiln",
				"cdjcjhonjgiagkb",
				"kooenbeoongcle",
				"cecnlfbaanckdkj",
				"fejlmog",
				"fanekdneoaammb",
				"maojbcegdamn",
				"bcmanmjdeabdo",
				"amloj",
				"adgoej",
				"jh",
				"fhf",
				"cogdljlgek",
				"o",
				"joeiajlioggj",
				"oncal",
				"lbgg",
				"elainnbffk",
				"hbdi",
				"femcanllndoh",
				"ke",
				"hmib",
				"nagfahhljh",
				"ibifdlfeechcbal",
				"knec",
				"oegfcghlgalcnno",
				"abiefmjldmln",
				"mlfglgni",
				"jkofhjeb",
				"ifjbneblfldjel",
				"nahhcimkjhjgb",
				"cdgkbn",
				"nnklfbeecgedie",
				"gmllmjbodhgllc",
				"hogollongjo",
				"fmoinacebll",
				"fkngbganmh",
				"jgdblmhlmfij",
				"fkkdjknahamcfb",
				"aieakdokibj",
				"hddlcdiailhd",
				"iajhmg",
				"jenocgo",
				"embdib",
				"dghbmljjogka",
				"bahcggjgmlf",
				"fb",
				"jldkcfom",
				"mfi",
				"kdkke",
				"odhbl",
				"jin",
				"kcjmkggcmnami",
				"kofig",
				"bid",
				"ohnohi",
				"fcbojdgoaoa",
				"dj",
				"ifkbmbod",
				"dhdedohlghk",
				"nmkeakohicfdjf",
				"ahbifnnoaldgbj",
				"egldeibiinoac",
				"iehfhjjjmil",
				"bmeimi",
				"ombngooicknel",
				"lfdkngobmik",
				"ifjcjkfnmgjcnmi",
				"fmf",
				"aoeaa",
				"an",
				"ffgddcjblehhggo",
				"hijfdcchdilcl",
				"hacbaamkhblnkk",
				"najefebghcbkjfl",
				"hcnnlogjfmmjcma",
				"njgcogemlnohl",
				"ihejh",
				"ej",
				"ofn",
				"ggcklj",
				"omah",
				"hg",
				"obk",
				"giig",
				"cklna",
				"lihaiollfnem",
				"ionlnlhjckf",
				"cfdlijnmgjoebl",
				"dloehimen",
				"acggkacahfhkdne",
				"iecd",
				"gn",
				"odgbnalk",
				"ahfhcd",
				"dghlag",
				"bchfe",
				"dldblmnbifnmlo",
				"cffhbijal",
				"dbddifnojfibha",
				"mhh",
				"cjjol",
				"fed",
				"bhcnf",
				"ciiibbedklnnk",
				"ikniooicmm",
				"ejf",
				"ammeennkcdgbjco",
				"jmhmd",
				"cek",
				"bjbhcmda",
				"kfjmhbf",
				"chjmmnea",
				"ifccifn",
				"naedmco",
				"iohchafbega",
				"kjejfhbco",
				"anlhhhhg",
			},
		),
		true,
	)
	assert.Equal(t, wordBreak2("aaaaaaa", []string{"aaaa", "aaa"}), true)
	assert.Equal(t, wordBreak2("leetcode", []string{"leet", "code"}), true)
	assert.Equal(t, wordBreak2("applepenapple", []string{"apple", "pen"}), true)
	assert.Equal(t, wordBreak2("catsandog", []string{"cats", "cat", "dog", "sand", "and"}), false)

	assert.Equal(
		t,
		wordBreak2(
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		),
		false,
	)
}

func Test_longestCommonSubsequence(t *testing.T) {
	assert.Equal(t, longestCommonSubsequence("mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"), 6)
	assert.Equal(
		t,
		longestCommonSubsequence(
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		),
		0,
	)
	assert.Equal(t, longestCommonSubsequence("abcdefg", "bcdefgabc"), 6)
	assert.Equal(t, longestCommonSubsequence("abcdefg", "bcdefga"), 6)
	assert.Equal(t, longestCommonSubsequence("abcde", "ace"), 3)
	assert.Equal(t, longestCommonSubsequence("abc", "abc"), 3)
	assert.Equal(t, longestCommonSubsequence("abc", "def"), 0)
}

func Test_minDistance(t *testing.T) {
	assert.Equal(t, minDistance("distance", "springbok"), 9)
	assert.Equal(t, minDistance("horse", "ros"), 3)
	assert.Equal(t, minDistance("", ""), 0)
	assert.Equal(t, minDistance("", "b"), 1)
	assert.Equal(t, minDistance("a", "b"), 1)
	assert.Equal(t, minDistance("a", ""), 1)
	assert.Equal(t, minDistance("intention", "execution"), 5)
	assert.Equal(t, minDistance("abcdefg", "bcdefgabc"), 4)
	assert.Equal(t, minDistance("mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"), 15)
}

func Test_coinChange(t *testing.T) {
	assert.Equal(t, coinChange([]int{2}, 3), -1)
	assert.Equal(t, coinChange([]int{1}, 2), 2)

	assert.Equal(t, coinChange([]int{1, 2, 5}, 11), 3)

	assert.Equal(t, coinChange([]int{1}, 0), 0)
	assert.Equal(t, coinChange([]int{1}, 1), 1)
	assert.Equal(t, coinChange([]int{1, 2, 5, 6, 7, 8, 56, 21, 45, 13, 19}, 899), 18)
}

func Test_minPathSum(t *testing.T) {
	assert.Equal(t, minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}), 7)
	assert.Equal(t, minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}), 12)
}

func Test_minPathSum2(t *testing.T) {
	assert.Equal(t, minPathSum2([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}), 7)
	assert.Equal(t, minPathSum2([][]int{{1, 2, 3}, {4, 5, 6}}), 12)
}

func Test_uniquePaths(t *testing.T) {
	assert.Equal(t, uniquePaths(3, 3), 6)
	assert.Equal(t, uniquePaths(1, 1), 1)
	assert.Equal(t, uniquePaths(1, 7), 1)
	assert.Equal(t, uniquePaths(7, 3), 28)
	assert.Equal(t, uniquePaths(3, 7), 28)
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	assert.Equal(t, uniquePathsWithObstacles([][]int{{0, 0, 0}, {1, 0, 1}, {0, 0, 0}}), 1)
	assert.Equal(t, uniquePathsWithObstacles([][]int{{1, 0}}), 0)
	assert.Equal(t, uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}), 1)
	assert.Equal(t, uniquePathsWithObstacles([][]int{{1}}), 0)
	assert.Equal(t, uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}), 2)
	assert.Equal(t, uniquePathsWithObstacles([][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}}), 0)
	assert.Equal(t, uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 0}}), 1)
}

func Test_deleteDuplicates2(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 1})
	head = deleteDuplicates2(head)
	leetcode.LoopEqualList(t, head, []int{})

	head = leetcode.GenerateList([]int{1, 1, 1, 1, 2, 2, 3})
	head = deleteDuplicates2(head)
	leetcode.LoopEqualList(t, head, []int{3})

	var tn *leetcode.ListNode
	head = deleteDuplicates2(tn)
	leetcode.LoopEqualList(t, head, []int{})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5, 6, 7})
	head = deleteDuplicates2(head)
	leetcode.LoopEqualList(t, head, []int{1, 2, 3, 4, 5, 6, 7})

	head = leetcode.GenerateList([]int{1, 1})
	head = deleteDuplicates2(head)
	leetcode.LoopEqualList(t, head, []int{})

	head = leetcode.GenerateList([]int{1, 2, 2, 3, 4, 4})
	head = deleteDuplicates2(head)
	leetcode.LoopEqualList(t, head, []int{1, 3})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 4, 5, 6, 6, 7, 7})
	head = deleteDuplicates2(head)
	leetcode.LoopEqualList(t, head, []int{1, 2, 3, 5})
}

func Test_deleteDuplicates(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 1})
	head = deleteDuplicates(head)
	leetcode.LoopEqualList(t, head, []int{1})

	head = leetcode.GenerateList([]int{1, 1, 1, 1, 2, 2, 3})
	head = deleteDuplicates(head)
	leetcode.LoopEqualList(t, head, []int{1, 2, 3})

	var tn *leetcode.ListNode
	head = deleteDuplicates(tn)
	leetcode.LoopEqualList(t, head, []int{})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5, 6, 7})
	head = deleteDuplicates(head)
	leetcode.LoopEqualList(t, head, []int{1, 2, 3, 4, 5, 6, 7})

	head = leetcode.GenerateList([]int{1, 1})
	head = deleteDuplicates(head)
	leetcode.LoopEqualList(t, head, []int{})

	head = leetcode.GenerateList([]int{1, 2, 2, 3, 4, 4})
	head = deleteDuplicates(head)
	leetcode.LoopEqualList(t, head, []int{1, 2, 3, 4})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 4, 5, 6, 6, 7, 7})
	head = deleteDuplicates(head)
	leetcode.LoopEqualList(t, head, []int{1, 2, 3, 4, 5, 6, 7})
}

func Test_maxDepth(t *testing.T) {
	var root *leetcode.TreeNode
	assert.Equal(t, maxDepth(root), 0)

	root = leetcode.GenerateTree([]interface{}{0, 3, 9, 20, nil, nil, 15, 7})
	assert.Equal(t, maxDepth(root), 3)
}

func Test_isBalanced(t *testing.T) {
	var root *leetcode.TreeNode

	assert.Equal(t, isBalanced(root), true)

	root = leetcode.GenerateTree([]interface{}{0, 1, 2, 2, 3, 3, nil, nil, 4, 4})
	assert.Equal(t, isBalanced(root), false)

	root = leetcode.GenerateTree([]interface{}{0, 3, 9, 20, nil, nil, 15, 7})
	assert.Equal(t, isBalanced(root), true)
}

func Test_maxPathSum(t *testing.T) {
	var root *leetcode.TreeNode

	root = leetcode.GenerateTree([]interface{}{0, 1})
	assert.Equal(t, maxPathSum(root), 1)

	root = leetcode.GenerateTree([]interface{}{0, -1, 0, 3})
	assert.Equal(t, maxPathSum(root), 3)

	root = leetcode.GenerateTree([]interface{}{0, 1, 0, 3})
	assert.Equal(t, maxPathSum(root), 4)

	root = leetcode.GenerateTree([]interface{}{0, -10, 9, 20, nil, nil, 15, 7})
	assert.Equal(t, maxPathSum(root), 42)
}

func Test_rotateRight(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	leetcode.LoopEqualList(t, rotateRight(head, 7), []int{4, 5, 1, 2, 3})

	head = leetcode.GenerateList([]int{0, 1, 2})
	leetcode.LoopEqualList(t, rotateRight(head, 7), []int{2, 0, 1})
}

func Test_lowestCommonAncestor(t *testing.T) {
	var head *leetcode.TreeNode

	head = leetcode.GenerateTree([]interface{}{0, 3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	assert.Equal(t, lowestCommonAncestor(head, &leetcode.TreeNode{Val: 5}, &leetcode.TreeNode{Val: 1}).Val, 3)
	assert.Equal(t, lowestCommonAncestor(head, &leetcode.TreeNode{Val: 5}, &leetcode.TreeNode{Val: 4}).Val, 5)

	head = leetcode.GenerateTree([]interface{}{0, 1, 2})
	assert.Equal(t, lowestCommonAncestor(head, &leetcode.TreeNode{Val: 1}, &leetcode.TreeNode{Val: 2}).Val, 1)
}

func Test_levelOrder(t *testing.T) {
	var head *leetcode.TreeNode

	head = leetcode.GenerateTree([]interface{}{0, 3, 9, 20, nil, nil, 15, 7})
	assert.DeepEqual(t, levelOrder(head), [][]int{{3},
		{9, 20},
		{15, 7},
	})
}

func Test_levelOrderBottom(t *testing.T) {
	var head *leetcode.TreeNode

	head = leetcode.GenerateTree([]interface{}{0, 3, 9, 20, nil, nil, 15, 7})
	assert.DeepEqual(t, levelOrderBottom(head), [][]int{
		{15, 7},
		{9, 20},
		{3},
	})
}

func Test_zigzagLevelOrder(t *testing.T) {
	var head *leetcode.TreeNode

	head = leetcode.GenerateTree([]interface{}{0, 3, 9, 20, nil, nil, 15, 7})
	assert.DeepEqual(t, zigzagLevelOrder(head), [][]int{{3},
		{20, 9},
		{15, 7},
	})
}

func Test_isValidBST(t *testing.T) {
	var head *leetcode.TreeNode

	head = leetcode.GenerateTree([]interface{}{0, 5, 14, nil, 1})
	assert.Equal(t, isValidBST(head), false)

	head = leetcode.GenerateTree(
		[]interface{}{0, 120, nil, 140, nil, nil, 130, 160, nil, nil, nil, nil, 119, 135, 150, 200},
	)
	assert.Equal(t, isValidBST(head), false)

	head = leetcode.GenerateTree([]interface{}{0, 120, 70, 140, 50, 100, 130, 160, 20, 55, 75, 110, 119, 135, 150, 200})
	assert.Equal(t, isValidBST(head), false)

	head = leetcode.GenerateTree([]interface{}{0, 32, 26, 47, 19, nil, nil, 56, nil, 27})
	assert.Equal(t, isValidBST(head), false)

	head = leetcode.GenerateTree([]interface{}{0, 2, 1, 3})
	assert.Equal(t, isValidBST(head), true)

	head = leetcode.GenerateTree([]interface{}{0, 5, 1, 4, nil, nil, 3, 6})
	assert.Equal(t, isValidBST(head), false)

	head = leetcode.GenerateTree([]interface{}{0, 5, 1, 6, nil, nil, 3, 7})
	assert.Equal(t, isValidBST(head), false)

	head = leetcode.GenerateTree([]interface{}{0, 5, 4, 6, 1, 3, 3, 7})
	assert.Equal(t, isValidBST(head), false)
}

func Test_Constructor(t *testing.T) {
	var head *leetcode.TreeNode

	head = leetcode.GenerateTree([]interface{}{0, 7, 3, 15, 1, 4, 9, 20})
	obj := BSTConstructor(head)
	assert.Equal(t, obj.Next(), 1)
	assert.Equal(t, obj.Next(), 3)
	assert.Equal(t, obj.Next(), 4)
	assert.Equal(t, obj.Next(), 7)
	assert.Equal(t, obj.HasNext(), true)
	assert.Equal(t, obj.Next(), 9)
	assert.Equal(t, obj.HasNext(), true)
	assert.Equal(t, obj.Next(), 15)
	assert.Equal(t, obj.HasNext(), true)
	assert.Equal(t, obj.Next(), 20)
	assert.Equal(t, obj.HasNext(), false)
}

func Test_insertIntoBST(t *testing.T) {
	var head *leetcode.TreeNode

	head = leetcode.GenerateTree([]interface{}{0, 4, 2, 7, 1, 3})
	head = insertIntoBST(head, 5)
	assert.Equal(t, isValidBST(head), true)

	head = leetcode.GenerateTree([]interface{}{0, 40, 20, 60, 10, 30, 50, 70})
	head = insertIntoBST(head, 25)
	assert.Equal(t, isValidBST(head), true)

	head = leetcode.GenerateTree([]interface{}{0, 4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil})
	head = insertIntoBST(head, 5)
	assert.Equal(t, isValidBST(head), true)
}

func Test_reverseList(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	leetcode.LoopEqualList(t, reverseList(head), []int{5, 4, 3, 2, 1})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})
	leetcode.LoopEqualList(t, reverseList(head), []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})
}

func Test_deepReverseList(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	leetcode.LoopEqualList(t, deepReverseList(head), []int{5, 4, 3, 2, 1})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})
	leetcode.LoopEqualList(t, deepReverseList(head), []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})

}

func Test_reverseBetween(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{3, 5})
	leetcode.LoopEqualList(t, reverseBetween(head, 1, 1), []int{3, 5})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})
	leetcode.LoopEqualList(t, reverseBetween(head, 1, 5), []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 1})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	leetcode.LoopEqualList(t, reverseBetween(head, 1, 5), []int{5, 4, 3, 2, 1})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	leetcode.LoopEqualList(t, reverseBetween(head, 2, 4), []int{1, 4, 3, 2, 5})
}

func Test_mergeTwoLists(t *testing.T) {
	leetcode.LoopEqualList(t,
		mergeTwoLists(
			leetcode.GenerateList([]int{1, 2, 4}),
			leetcode.GenerateList([]int{1, 3, 4})),
		[]int{1, 1, 2, 3, 4, 4})

	leetcode.LoopEqualList(t,
		mergeTwoLists(
			leetcode.GenerateList([]int{}),
			leetcode.GenerateList([]int{})),
		[]int{})

	leetcode.LoopEqualList(t,
		mergeTwoLists(
			leetcode.GenerateList([]int{1, 2, 4}),
			leetcode.GenerateList([]int{})),
		[]int{1, 2, 4})

	leetcode.LoopEqualList(t,
		mergeTwoLists(
			leetcode.GenerateList([]int{1, 2, 4}),
			leetcode.GenerateList([]int{0})),
		[]int{0, 1, 2, 4})

	leetcode.LoopEqualList(t,
		mergeTwoLists(
			leetcode.GenerateList([]int{1, 3, 5}),
			leetcode.GenerateList([]int{0, 2, 4})),
		[]int{0, 1, 2, 3, 4, 5})
}

func Test_partition(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{2, 1})
	leetcode.LoopEqualList(t, partition(head, 2), []int{1, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition(head, 3), []int{1, 2, 2, 4, 3, 5})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition(head, 2), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition(head, 9), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition(head, 0), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition(head, 1), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2, 6})
	leetcode.LoopEqualList(t, partition(head, 6), []int{1, 4, 3, 2, 5, 2, 6})
}

func Test_partition2(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{2, 1})
	leetcode.LoopEqualList(t, partition2(head, 2), []int{1, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition2(head, 3), []int{1, 2, 2, 4, 3, 5})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition2(head, 2), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition2(head, 9), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition2(head, 0), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
	leetcode.LoopEqualList(t, partition2(head, 1), []int{1, 4, 3, 2, 5, 2})

	head = leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2, 6})
	leetcode.LoopEqualList(t, partition2(head, 6), []int{1, 4, 3, 2, 5, 2, 6})
}

//func Benchmark_partition(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		head := leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
//		partition(head, 3)
//	}
//}

//func Benchmark_partition2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		head := leetcode.GenerateList([]int{1, 4, 3, 2, 5, 2})
//		partition2(head, 3)
//	}
//}

func Test_sortList(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{4, 2, 1, 3})
	leetcode.LoopEqualList(t, sortList(head), []int{1, 2, 3, 4})

	head = leetcode.GenerateList([]int{-1, 5, 3, 4, 0})
	leetcode.LoopEqualList(t, sortList(head), []int{-1, 0, 3, 4, 5})

	head = leetcode.GenerateList([]int{2, 1})
	leetcode.LoopEqualList(t, sortList(head), []int{1, 2})

	head = leetcode.GenerateList([]int{1, 2})
	leetcode.LoopEqualList(t, sortList(head), []int{1, 2})

	head = leetcode.GenerateList([]int{2})
	leetcode.LoopEqualList(t, sortList(head), []int{2})

	head = leetcode.GenerateList([]int{})
	leetcode.LoopEqualList(t, sortList(head), []int{})
}

//func Benchmark_sortList(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		sortList(leetcode.GenerateList(bigRepeatArray))
//	}
//}

func Test_reverseBits(t *testing.T) {
	assert.Equal(
		t,
		reverseBits(strToUint("00000010100101000001111010011100")),
		strToUint("00111001011110000010100101000000"),
	)
	assert.Equal(
		t,
		reverseBits(strToUint("11111111111111111111111111111101")),
		strToUint("10111111111111111111111111111111"),
	)

	assert.Equal(
		t,
		reverseBits2(strToUint("00000010100101000001111010011100")),
		strToUint("00111001011110000010100101000000"),
	)
	assert.Equal(
		t,
		reverseBits2(strToUint("11111111111111111111111111111101")),
		strToUint("10111111111111111111111111111111"),
	)
}

func strToUint(s string) uint32 {
	var sum uint32
	for i, r := range s {
		if r == '0' {
			continue
		}
		b := uint32(1) << (len(s) - 1 - i)
		sum += b
	}
	return sum
}

func Test_reorderList(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	reorderList(head)
	leetcode.LoopEqualList(t, head, []int{1, 5, 2, 4, 3})

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5, 6})
	reorderList(head)
	leetcode.LoopEqualList(t, head, []int{1, 6, 2, 5, 3, 4})

	head = leetcode.GenerateList([]int{1, 2, 3, 4})
	reorderList(head)
	leetcode.LoopEqualList(t, head, []int{1, 4, 2, 3})

	head = leetcode.GenerateList([]int{1})
	reorderList(head)
	leetcode.LoopEqualList(t, head, []int{1})

	head = leetcode.GenerateList([]int{1, 2})
	reorderList(head)
	leetcode.LoopEqualList(t, head, []int{1, 2})
}

func Test_hasCycle(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, hasCycle(head), false)

	node := head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head.Next

	assert.Equal(t, hasCycle(head), true)

	head = leetcode.GenerateList([]int{1, 2})
	assert.Equal(t, hasCycle(head), false)
	head.Next.Next = head
	assert.Equal(t, hasCycle(head), true)
}

func Test_detectCycle(t *testing.T) {
	var head *leetcode.ListNode
	var tail *leetcode.ListNode
	var node *leetcode.ListNode

	head = leetcode.GenerateList([]int{3, 2, 0, -4})
	assert.Equal(t, detectCycle(head), tail)

	node = head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head

	assert.Equal(t, detectCycle(head).Val, 3)

	head = leetcode.GenerateList([]int{3, 2, 0, -4})
	assert.Equal(t, detectCycle(head), tail)

	node = head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head.Next

	assert.Equal(t, detectCycle(head).Val, 2)

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, detectCycle(head), tail)

	node = head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head.Next

	assert.Equal(t, detectCycle(head).Val, 2)

	head = leetcode.GenerateList([]int{1, 2})
	assert.Equal(t, detectCycle(head), tail)
	head.Next.Next = head
	assert.Equal(t, detectCycle(head).Val, 1)
}

func Test_detectCycle2(t *testing.T) {
	var head *leetcode.ListNode
	var tail *leetcode.ListNode
	var node *leetcode.ListNode

	head = leetcode.GenerateList([]int{3, 2, 0, -4})
	assert.Equal(t, detectCycle2(head), tail)

	node = head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head

	assert.Equal(t, detectCycle2(head).Val, 3)

	head = leetcode.GenerateList([]int{3, 2, 0, -4})
	assert.Equal(t, detectCycle2(head), tail)

	node = head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head.Next

	assert.Equal(t, detectCycle2(head).Val, 2)

	head = leetcode.GenerateList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, detectCycle2(head), tail)

	node = head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head.Next

	assert.Equal(t, detectCycle2(head).Val, 2)

	head = leetcode.GenerateList([]int{1, 2})
	assert.Equal(t, detectCycle2(head), tail)
	head.Next.Next = head
	assert.Equal(t, detectCycle2(head).Val, 1)
}

func Test_searchMatrix(t *testing.T) {
	assert.Equal(t, searchMatrix([][]int{{1}}, 1), true)
	assert.Equal(t, searchMatrix([][]int{{1}}, 0), false)
	assert.Equal(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3), true)
	assert.Equal(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 60), true)
	assert.Equal(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 16), true)
	assert.Equal(t, searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 10), true)
}

func Test_isPalindrome(t *testing.T) {
	var head *leetcode.ListNode

	head = leetcode.GenerateList([]int{1, 2, 2, 1})
	assert.Equal(t, isPalindrome(head), true)

	head = leetcode.GenerateList([]int{1, 2, 5, 2, 1})
	assert.Equal(t, isPalindrome(head), true)

	head = leetcode.GenerateList([]int{1, 2, 1, 2})
	assert.Equal(t, isPalindrome(head), false)

	head = leetcode.GenerateList([]int{1, 2, 5, 1, 2})
	assert.Equal(t, isPalindrome(head), false)

	head = leetcode.GenerateList([]int{1, 2})
	assert.Equal(t, isPalindrome(head), false)

	head = leetcode.GenerateList([]int{1, 1})
	assert.Equal(t, isPalindrome(head), true)

	head = leetcode.GenerateList([]int{1})
	assert.Equal(t, isPalindrome(head), true)

	head = leetcode.GenerateList([]int{})
	assert.Equal(t, isPalindrome(head), true)

	head = leetcode.GenerateList([]int{1, 2, 2, 1})
	assert.Equal(t, isPalindrome2(head), true)

	head = leetcode.GenerateList([]int{1, 2, 5, 2, 1})
	assert.Equal(t, isPalindrome2(head), true)

	head = leetcode.GenerateList([]int{1, 2, 1, 2})
	assert.Equal(t, isPalindrome2(head), false)

	head = leetcode.GenerateList([]int{1, 2, 5, 1, 2})
	assert.Equal(t, isPalindrome2(head), false)

	head = leetcode.GenerateList([]int{1, 2})
	assert.Equal(t, isPalindrome2(head), false)

	head = leetcode.GenerateList([]int{1, 1})
	assert.Equal(t, isPalindrome2(head), true)

	head = leetcode.GenerateList([]int{1})
	assert.Equal(t, isPalindrome2(head), true)

	head = leetcode.GenerateList([]int{})
	assert.Equal(t, isPalindrome2(head), true)
}

func Test_copyRandomList(t *testing.T) {
	var head *RandomNode
	head = &RandomNode{Val: 7}
	node := head
	node.Next = &RandomNode{Val: 13, Random: head}
	node = node.Next
	node.Next = &RandomNode{Val: 11}
	a := node
	node = node.Next
	node.Next = &RandomNode{Val: 10}
	node = node.Next
	node.Next = &RandomNode{Val: 1, Random: head}
	b := node
	node = head
	node = node.Next.Next
	node.Random = b
	node = node.Next
	node.Random = a

	c := copyRandomList(head)

	for c != nil {
		assert.Equal(t, head == c, false)
		assert.Equal(t, head.Val, c.Val)
		c, head = c.Next, head.Next
	}
}

func Test_singleNumber(t *testing.T) {
	assert.Equal(t, singleNumber1([]int{1, 2, 2, 3, 3}), 1)
	assert.Equal(t, singleNumber1([]int{1, 1, 2, 2, 5, 3, 3}), 5)
	assert.Equal(t, singleNumber1([]int{1}), 1)
}

func Test_singleNumber2(t *testing.T) {
	assert.Equal(t, singleNumber([]int{2, 1, 2, 2, 3, 3, 3}), 1)
	assert.Equal(t, singleNumber([]int{2, 1, 1, 3, 2, 1, 2, 5, 3, 3}), 5)
	assert.Equal(t, singleNumber([]int{1}), 1)
}

func Test_singleNumber3(t *testing.T) {
	assert.DeepEqual(t, singleNumber3([]int{3, 2, 2, 4, 4, 5}), []int{3, 5})
	assert.DeepEqual(t, singleNumber3([]int{1, 1, 2, 2, 5, 3, 6, 3}), []int{5, 6})
	assert.DeepEqual(t, singleNumber3([]int{1, 6}), []int{1, 6})
}

func Test_hammingWeight(t *testing.T) {
	assert.Equal(t, hammingWeight(1), 1)
	assert.Equal(t, hammingWeight(2), 1)
	assert.Equal(t, hammingWeight(3), 2)
}

func Test_countBits(t *testing.T) {
	assert.DeepEqual(t, countBits(2), []int{0, 1, 1})
	assert.DeepEqual(t, countBits(5), []int{0, 1, 1, 2, 1, 2})
}

func Test_rangeBitwiseAnd(t *testing.T) {
	assert.Equal(t, rangeBitwiseAnd((1<<32)-2, (1<<32)-1), (1<<32)-2)

	assert.Equal(t, rangeBitwiseAnd(5, 7), 4)
	assert.Equal(t, rangeBitwiseAnd(5, 10), 0)
	assert.Equal(t, rangeBitwiseAnd((1<<30)-1, (1<<31)-1), 0)

	assert.Equal(t, rangeBitwiseAnd((1<<31)-2, (1<<31)-1), (1<<31)-2)
}

//func Benchmark_rangeBitwiseAnd2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		rangeBitwiseAnd2(5, 7)
//		rangeBitwiseAnd2(5, 10)
//		rangeBitwiseAnd2((1<<30)-1, (1<<31)-1)
//		rangeBitwiseAnd2((1<<31)-2, (1<<31)-1)
//	}
//}

//func Benchmark_rangeBitwiseAnd(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		rangeBitwiseAnd(5, 7)
//		rangeBitwiseAnd(5, 10)
//		rangeBitwiseAnd((1<<30)-1, (1<<31)-1)
//		rangeBitwiseAnd((1<<31)-2, (1<<31)-1)
//	}
//}

func Test_LRUConstructor(t *testing.T) {
	lru := LRUConstructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	assert.Equal(t, lru.Get(1), 1)
	lru.Put(3, 3)
	assert.Equal(t, lru.Get(2), -1)
	lru.Put(4, 4)
	assert.Equal(t, lru.Get(1), -1)
	assert.Equal(t, lru.Get(3), 3)
	assert.Equal(t, lru.Get(4), 4)

	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//	[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//	[null,null,null,0,null,-1,null,-1,3,4]
	lru = LRUConstructor(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	assert.Equal(t, lru.Get(1), 0)
	lru.Put(3, 3)
	assert.Equal(t, lru.Get(2), -1)
	lru.Put(4, 4)
	assert.Equal(t, lru.Get(1), -1)
	assert.Equal(t, lru.Get(3), 3)
	assert.Equal(t, lru.Get(4), 4)

}

func Test_LRUConstructor2(t *testing.T) {
	lru := ConstructorLRU(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	assert.Equal(t, lru.Get(1), 1)
	lru.Put(3, 3)
	assert.Equal(t, lru.Get(2), -1)
	lru.Put(4, 4)
	assert.Equal(t, lru.Get(1), -1)
	assert.Equal(t, lru.Get(3), 3)
	assert.Equal(t, lru.Get(4), 4)

	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//	[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//	[null,null,null,0,null,-1,null,-1,3,4]
	lru = ConstructorLRU(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	assert.Equal(t, lru.Get(1), 0)
	lru.Put(3, 3)
	assert.Equal(t, lru.Get(2), -1)
	lru.Put(4, 4)
	assert.Equal(t, lru.Get(1), -1)
	assert.Equal(t, lru.Get(3), 3)
	assert.Equal(t, lru.Get(4), 4)

}

func Test_LFUConstructor(t *testing.T) {
	var lfu LFUCache

	// ["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
	// [[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]

	lfu = ConstructorLFU(3)
	lfu.Put(2, 2)
	lfu.Put(1, 1)
	assert.Equal(t, lfu.Get(2), 2)
	assert.Equal(t, lfu.Get(1), 1)
	assert.Equal(t, lfu.Get(2), 2)
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(3), -1)
	assert.Equal(t, lfu.Get(2), 2)
	assert.Equal(t, lfu.Get(1), 1)
	assert.Equal(t, lfu.Get(4), 4)

	//  输入
	//  ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	//	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
	//	输出：
	//	[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	lfu = ConstructorLFU(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 1)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	assert.Equal(t, lfu.Get(3), 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), -1)
	assert.Equal(t, lfu.Get(3), 3)
	assert.Equal(t, lfu.Get(4), 4)

	//  ["LFUCache","put","put","get","put","get","put","get","get","get"]
	//	[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//	[null,null,null,0,null,-1,null,0,-1,4]
	lfu = ConstructorLFU(2)
	lfu.Put(1, 0)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 0)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), 0)
	assert.Equal(t, lfu.Get(3), -1)
	assert.Equal(t, lfu.Get(4), 4)
}

func Test_LFUConstructor2(t *testing.T) {
	var lfu LFUCache2

	// ["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
	// [[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]

	lfu = LFUConstructor(3)
	lfu.Put(2, 2)
	lfu.Put(1, 1)
	assert.Equal(t, lfu.Get(2), 2)
	assert.Equal(t, lfu.Get(1), 1)
	assert.Equal(t, lfu.Get(2), 2)
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(3), -1)
	assert.Equal(t, lfu.Get(2), 2)
	assert.Equal(t, lfu.Get(1), 1)
	assert.Equal(t, lfu.Get(4), 4)

	//  输入
	//  ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	//	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
	//	输出：
	//	[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	lfu = LFUConstructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 1)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	assert.Equal(t, lfu.Get(3), 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), -1)
	assert.Equal(t, lfu.Get(3), 3)
	assert.Equal(t, lfu.Get(4), 4)

	//  ["LFUCache","put","put","get","put","get","put","get","get","get"]
	//	[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//	[null,null,null,0,null,-1,null,0,-1,4]
	lfu = LFUConstructor(2)
	lfu.Put(1, 0)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 0)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), 0)
	assert.Equal(t, lfu.Get(3), -1)
	assert.Equal(t, lfu.Get(4), 4)
}

func Test_LFUConstructor3(t *testing.T) {
	var lfu LFUCache3

	// ["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
	// [[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]

	lfu = LFConstructor(3)
	lfu.Put(2, 2)
	lfu.Put(1, 1)
	assert.Equal(t, lfu.Get(2), 2)
	assert.Equal(t, lfu.Get(1), 1)
	assert.Equal(t, lfu.Get(2), 2)
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(3), -1)
	assert.Equal(t, lfu.Get(2), 2)
	assert.Equal(t, lfu.Get(1), 1)
	assert.Equal(t, lfu.Get(4), 4)

	//  输入
	//  ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	//	[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
	//	输出：
	//	[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

	lfu = LFConstructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 1)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	assert.Equal(t, lfu.Get(3), 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), -1)
	assert.Equal(t, lfu.Get(3), 3)
	assert.Equal(t, lfu.Get(4), 4)

	//  ["LFUCache","put","put","get","put","get","put","get","get","get"]
	//	[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//	[null,null,null,0,null,-1,null,0,-1,4]
	lfu = LFConstructor(2)
	lfu.Put(1, 0)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 0)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), 0)
	assert.Equal(t, lfu.Get(3), -1)
	assert.Equal(t, lfu.Get(4), 4)
}
