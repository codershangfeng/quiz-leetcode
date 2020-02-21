/* 792
# 792 Number of Matching Subsequences

Given string S and a dictionary of words words, find the number of words[i] that is a subsequence of S.

Example :
Input:
S = "abcde"
words = ["a", "bb", "acd", "ace"]
Output: 3
Explanation: There are three words in words that are a subsequence of S: "a", "acd", "ace".
Note:

All words in words and S will only consists of lowercase letters.
The length of S will be in the range of [1, 50000].
The length of words will be in the range of [1, 5000].
The length of words[i] will be in the range of [1, 50].

*/

package numberofmatchingsubsequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numMatchingSubseq(S string, words []string) int {
	endIdxOfS := len(S) - 1

	// Number of matched subsequences
	num := 0

NextWord:
	for _, word := range words {
		lps := 0 // Pillar for S
		rps := endIdxOfS
		matched := false // Flag represents subsequence matched state

		// Special Case: the length of word is 1 AND it is equal to the character in the start index or end index of S
		if len(word) == 1 {
			for i := 0; i < len(S); i++ {
				if word[lps] == S[i] {
					num++
					break
				}
			}
			continue
		}

		lpw := 0
		rpw := len(word) - 1
		for i, l := 0, len(word)/2; i <= l; i++ {
			for ; lps < rps; lps++ {
				if S[lps] == word[lpw] {
					break
				}
			}
			for ; rps > lps; rps-- {
				if S[rps] == word[rpw] {
					break
				}
			}

			offsetS := rps - lps
			offsetW := rpw - lpw

			if offsetS < offsetW {
				continue NextWord
			}

			if offsetW == 1 || offsetW == 0 {
				matched = true // The pillars meet together means each character of word found its occurences on S
				break
			}

			lpw++
			rpw--
		}

		if matched {
			num++
		}
	}
	return num
}

// Case11?
// func TestCase11(t *testing.T) {
// 	got := numMatchingSubseq("ricogwqznwxxcpueelcobbbkuvxxrvgyehsudccpsnuxpcqobtvwkuvsubiidjtccoqvuahijyefbpqhbejuisksutsowhufsygtwteiqyligsnbqglqblhpdzzeurtdohdcbjvzgjwylmmoiundjscnlhbrhookmioxqighkxfugpeekgtdofwzemelpyjsdeeppapjoliqlhbrbghqjezzaxuwyrbczodtrhsvnaxhcjiyiphbglyolnswlvtlbmkrsurrcsgdzutwgjofowhryrubnxkahocqjzwwagqidjhwbunvlchojtbvnzdzqpvrazfcxtvhkruvuturdicnucvndigovkzrqiyastqpmfmuouycodvsyjajekhvyjyrydhxkdhffyytldcdlxqbaszbuxsacqwqnhrewhagldzhryzdmmrwnxhaqfezeeabuacyswollycgiowuuudrgzmwnxaezuqlsfvchjfloczlwbefksxsbanrektvibbwxnokzkhndmdhweyeycamjeplecewpnpbshhidnzwopdjuwbecarkgapyjfgmanuavzrxricbgagblomyseyvoeurekqjyljosvbneofjzxtaizjypbcxnbfeibrfjwyjqrisuybfxpvqywqjdlyznmojdhbeomyjqptltpugzceyzenflfnhrptuugyfsghluythksqhmxlmggtcbdddeoincygycdpehteiugqbptyqbvokpwovbnplshnzafunqglnpjvwddvdlmjjyzmwwxzjckmaptilrbfpjxiarmwalhbdjiwbaknvcqovwcqiekzfskpbhgxpyomekqvzpqyirelpadooxjhsyxjkfqavbaoqqvvknqryhotjritrkvdveyapjfsfzenfpuazdrfdofhudqbfnzxnvpluwicurrtshyvevkriudayyysepzqfgqwhgobwyhxltligahroyshfndydvffd",
// 		[]string{"iowuuudrgzmw","azfcxtvhkruvuturdicnucvndigovkzrq","ylmmo","maptilrbfpjxiarmwalhbd","oqvuahijyefbpqhbejuisksutsowhufsygtwteiqyligsnbqgl","ytldcdlxqbaszbuxsacqwqnhrewhagldzhr","zeeab","cqie","pvrazfcxtvhkruvuturdicnucvndigovkzrqiya","zxnvpluwicurrtshyvevkriudayyysepzq","wyhxltligahroyshfn","nhrewhagldzhryzdmmrwn","yqbvokpwovbnplshnzafunqglnpjvwddvdlmjjyzmw","nhrptuugyfsghluythksqhmxlmggtcbdd","yligsnbqglqblhpdzzeurtdohdcbjvzgjwylmmoiundjsc","zdrfdofhudqbfnzxnvpluwicurrtshyvevkriudayyysepzq","ncygycdpehteiugqbptyqbvokpwovbnplshnzafun","gdzutwgjofowhryrubnxkahocqjzww","eppapjoliqlhbrbgh","qwhgobwyhxltligahroys","dzutwgjofowhryrubnxkah","rydhxkdhffyytldcdlxqbaszbuxs","tyqbvokpwovbnplshnzafunqglnpjvwddvdlmjjyzmwwxzjc","khvyjyrydhxkdhffyytldcdlxqbasz","jajekhvyjyrydhxkdhffyytldcdlxqbaszbuxsacqwqn","ppapjoliqlhbrbghq","zmwwxzjckmaptilrbfpjxiarm","nxkahocqjzwwagqidjhwbunvlchoj","ybfxpvqywqjdlyznmojdhbeomyjqptltp","udrgzmwnxae","nqglnpjvwddvdlmjjyzmww","swlvtlbmkrsurrcsgdzutwgjofowhryrubn","hudqbfnzxnvpluwicurr","xaezuqlsfvchjf","tvibbwxnokzkhndmdhweyeycamjeplec","olnswlvtlbmkrsurrcsgdzu","qiyastqpmfmuouycodvsyjajekhvyjyrydhxkdhffyyt","eiqyligsnbqglqblhpdzzeurtdohdcbjvzgjwyl","cgiowuuudrgzmwnxaezuqlsfvchjflocz","rxric","cygycdpehteiugqbptyqbvokpwovbnplshnzaf","g","surrcsgd","yzenflfnhrptuugyfsghluythksqh","gdzutwgjofowhryrubnxkahocqjzwwagqid","ddeoincygycdpeh","yznmojdhbeomyjqptltpugzceyzenflfnhrptuug","ejuisks","teiqyligsnbqglqblhpdzzeurtdohdcbjvzgjwylmmoi","mrwnxhaqfezeeabuacyswollycgio","qfskkpfakjretogrokmxemjjbvgmmqrfdxlkfvycwalbdeumav","wjgjhlrpvhqozvvkifhftnfqcfjmmzhtxsoqbeduqmnpvimagq","ibxhtobuolmllbasaxlanjgalgmbjuxmqpadllryaobcucdeqc","ydlddogzvzttizzzjohfsenatvbpngarutztgdqczkzoenbxzv","rmsakibpprdrttycxglfgtjlifznnnlkgjqseguijfctrcahbb","pqquuarnoybphojyoyizhuyjfgwdlzcmkdbdqzatgmabhnpuyh","akposmzwykwrenlcrqwrrvsfqxzohrramdajwzlseguupjfzvd","vyldyqpvmnoemzeyxslcoysqfpvvotenkmehqvopynllvwhxzr","ysyskgrbolixwmffygycvgewxqnxvjsfefpmxrtsqsvpowoctw","oqjgumitldivceezxgoiwjgozfqcnkergctffspdxdbnmvjago","bpfgqhlkvevfazcmpdqakonkudniuobhqzypqlyocjdngltywn","ttucplgotbiceepzfxdebvluioeeitzmesmoxliuwqsftfmvlg","xhkklcwblyjmdyhfscmeffmmerxdioseybombzxjatkkltrvzq","qkvvbrgbzzfhzizulssaxupyqwniqradvkjivedckjrinrlxgi","itjudnlqncbspswkbcwldkwujlshwsgziontsobirsvskmjbrq","nmfgxfeqgqefxqivxtdrxeelsucufkhivijmzgioxioosmdpwx","ihygxkykuczvyokuveuchermxceexajilpkcxjjnwmdbwnxccl","etvcfbmadfxlprevjjnojxwonnnwjnamgrfwohgyhievupsdqd","ngskodiaxeswtqvjaqyulpedaqcchcuktfjlzyvddfeblnczmh","vnmntdvhaxqltluzwwwwrbpqwahebgtmhivtkadczpzabgcjzx","yjqqdvoxxxjbrccoaqqspqlsnxcnderaewsaqpkigtiqoqopth","wdytqvztzbdzffllbxexxughdvetajclynypnzaokqizfxqrjl","yvvwkphuzosvvntckxkmvuflrubigexkivyzzaimkxvqitpixo","lkdgtxmbgsenzmrlccmsunaezbausnsszryztfhjtezssttmsr","idyybesughzyzfdiibylnkkdeatqjjqqjbertrcactapbcarzb","ujiajnirancrfdvrfardygbcnzkqsvujkhcegdfibtcuxzbpds","jjtkmalhmrknaasskjnixzwjgvusbozslrribgazdhaylaxobj","nizuzttgartfxiwcsqchizlxvvnebqdtkmghtcyzjmgyzszwgi","egtvislckyltpfogtvfbtxbsssuwvjcduxjnjuvnqyiykvmrxl","ozvzwalcvaobxbicbwjrububyxlmfcokdxcrkvuehbnokkzala","azhukctuheiwghkalboxfnuofwopsrutamthzyzlzkrlsefwcz","yhvjjzsxlescylsnvmcxzcrrzgfhbsdsvdfcykwifzjcjjbmmu","tspdebnuhrgnmhhuplbzvpkkhfpeilbwkkbgfjiuwrdmkftphk","jvnbeqzaxecwxspuxhrngmvnkvulmgobvsnqyxdplrnnwfhfqq","bcbkgwpfmmqwmzjgmflichzhrjdjxbcescfijfztpxpxvbzjch","bdrkibtxygyicjcfnzigghdekmgoybvfwshxqnjlctcdkiunob","koctqrqvfftflwsvssnokdotgtxalgegscyeotcrvyywmzescq","boigqjvosgxpsnklxdjaxtrhqlyvanuvnpldmoknmzugnubfoa","jjtxbxyazxldpnbxzgslgguvgyevyliywihuqottxuyowrwfar","zqsacrwcysmkfbpzxoaszgqqsvqglnblmxhxtjqmnectaxntvb","izcakfitdhgujdborjuhtwubqcoppsgkqtqoqyswjfldsbfcct","rroiqffqzenlerchkvmjsbmoybisjafcdzgeppyhojoggdlpzq","xwjqfobmmqomhczwufwlesolvmbtvpdxejzslxrvnijhvevxmc","ccrubahioyaxuwzloyhqyluwoknxnydbedenrccljoydfxwaxy","jjoeiuncnvixvhhynaxbkmlurwxcpukredieqlilgkupminjaj","pdbsbjnrqzrbmewmdkqqhcpzielskcazuliiatmvhcaksrusae","nizbnxpqbzsihakkadsbtgxovyuebgtzvrvbowxllkzevktkuu","hklskdbopqjwdrefpgoxaoxzevpdaiubejuaxxbrhzbamdznrr","uccnuegvmkqtagudujuildlwefbyoywypakjrhiibrxdmsspjl","awinuyoppufjxgqvcddleqdhbkmolxqyvsqprnwcoehpturicf"})

// 	assert.Equal(t, 51, got)
// }

func TestCase10(t *testing.T) {
	got := numMatchingSubseq("ziyyxhtfbmgvnmvsqusddpvmrfskzaexergyqpjjusjfbzakugrjkglikryksufbftcjsmlzjmuwqdvlndetuzsyacmzdyulnrmv",
		[]string{"n", "kugrjkglikryksufbftcjsmlzjmuwqdvlndetuzsyacmzdyul", "muwqdvlndetuzsyacmzdyuln", "vmrfskzaexergyqpjjusjfbzakugrjk", "skzaexergyqpjjusjfbzakugrjkglikryksufbftcjsmlz", "pywquphugcnfuvrduhrcoxxratptwmqmxwttykvkoskejgjhdy", "rutbpwjugodzrgwzkvjzaafusqnsoiyvgszstwtlabgoeygfrx", "ruhodvpdipmlacbuiuwmhpuvooraxrunwwlzctzsgacjupowdb", "effkefvkcztipbojcttajvlexwuazmzqtfwubehuawmivyhzpw", "geguatbhflhixqeikcnaxyugmepetztyblyrwfwnhzxofjonqs"})

	assert.Equal(t, 5, got)
}

func TestCase9(t *testing.T) {
	got := numMatchingSubseq("fcovcoijujlnkkwigaecrlmxlqsoimttsvzfcnhspmooedfewducsqpdpguglpvytkrjfwlshgycpghlelqvnvuicpnkeykzyruayeysqejahyzvarkdaspqdxitmwpfssjbadpwyxqhuoyrmcclmqfvllfepvqyokqurqojeysjgwjqbpejigwkzjhoedkqjotliijncgpturborvxogxqitvrwxhphqzbwrtdvncmxnxbbvfetbagddfwjmvhtxxqacxcwlzdowkqnxkpkwssflculgbucoayflldcudcgbzvbrkkypufrgkstlwaxhyslzpwcrnufxbmksmnmwupfcwmvlwtgspjdoggpsslnaluhkikpftpkzaxmzcqapnaoikpgbzzsxottvealvbzeayipucbwziikmmrzcxdqlsywgtgnnwtnmxvvhuafgqhkywuviwajmyqnkiatclcpxnqznezywqhixqufuqkxerxgzpgkogoytzokgbtcutfyigpxuxcbkdcsasnyftetveqikiqrtemndqrywmsrrwlpbfyigfjmeiiqqqekmzwrycnuprdvvybqxzvodigemwabvriqqmcrtapnhbxaaglvxbmcfccdcoqdaoawavyxzeqjvskeuclolvdzrilqssfqwvwkxiquxmzaqcpufyirhyqpgsatingzyzszhulweualwczgxodiwueidnwhkhsnwswataqkwewjatxfbctijicbxgzajgyfpzscjddixijoogjqjodvvxxdsybkvwcmlqexsgpmoxnwfihpxoruqbycshchbhjbknjcztryhgplzcbjdpilatprzqalnevrruoxbcvajdddoqpvmwycajxjwfrzqllqtuxqebkwyxvrkwtwwzuexraaxiatdnldxpticwrjxmxbizcrouvopejdcmcswwnrozuuzkgfflcatxsjpxjxfwjfklragzjtwzddxemxysuo",
		[]string{"xvvhuafgqhkywuviwajmyqnkiatclcpxnqz", "mzcqapnaoik", "qqmcrtapnhbxaaglvx", "fvllfepvqyokq", "xzvodig", "yeysqejahyzvarkdaspqdxitmwpfssjba", "clcpxnqznezywqhixqufuqkxerxgzpgkogoytzo", "ec", "vrwxhphqzbwrtdvncmxnxbb", "pilatp", "mlqexsgpmoxnwfihp", "kwewjatxfbctijicbxgzajgyfpz", "wcmlqexsgpmoxnwfihpxoruqbycshchbhjbknjcztryhgplz", "vodigemwabvriqqmcrtapnhbxaaglvxbmcfc", "c", "hgplzcbjdpilatp", "rzqalnevrruoxbcvajdddoqpv", "bagddfwjmvhtxx", "bkwyxvrkwtwwzuexraaxiatdnldxpticwrjx", "qbycshchbhjbknjcztry", "uzkgfflcatxsjpxjxfwjf", "mcswwnrozuuzkgfflcatxsjpxjxfwjfklragzjtwzddxem", "pxuxcbkdcsasnyftetveq", "gbtcutfyigpx", "xcwlzdowkqnxkpkwssflcul", "hbxaagl", "cclmqfvllfepvqyokqurqojeysjgwjqbpejigwkzj", "gkstlwaxhyslzpwcr", "bagddfwjmvhtxxqacxcwlzdowkqnxkpkwssflculgbu", "vytkrjfwlshgycpgh", "xuxcbkdcsasnyftetveqikiqrtemndqrywms", "vuicpnkeykzyruayeysqe", "rruoxbcvajdddoqpvmwycajxjwfrzqllqtuxqebkwyxvr", "pjdoggpsslnaluhkikpftp", "xitmwp", "mzwrycnuprdvvyb", "wewjatxfbctijicbxgzajgyfpzscjddixijoog", "dnwhkhsnwswataqkwewjatxfbc", "qxzvodigemwabvriqqmcrtapnhbxaagl", "zcqapnaoikpgbzzsxottvealvbzeayipucbwziikmmrzcx", "bycshchbhjbknjcztryhgplzcbjd", "aoik", "pkzaxmzcqapnaoikpgbzz", "lcpxnqznezywqhixqufuqkxerxgzpgkogoytzokgb", "kpkwssflculgbucoayflldcudcgbzvbrkkypufrgkstlwa", "vvhuaf", "qnxkpkwssflculgbucoay", "gyfpzscjddixijoog", "ypufrgkstlwaxhyslzpwcrnufx", "cgbzvbrkk", "nxzpmwpqsrdzyinqhhgrtenlmsrzaowfhaxlzkpgahcsjhylig", "quldmwqctytfdnjhqgfznfbneauhsgyecpkpaogczgdzebwrqx", "zklbybcnhnuklxegeaafedilyqdwjweynyesjpzmolmmkfjekk", "kvbiingwvrmxultsxqeigdfrrraopnjraqqzeyaqsoxgooihqi", "vbnqopmssqsmzxwiufoqyurhrzaarrgjxvotbnikenoehkrzqb", "zdsspapcabxuqeostucqgwdtxizqvzzzzcheyjtcigzqgzvfud", "ervognhjrvexgyrrerdvhomsejmncpbhsafqegxihtrzgikcgh", "nzvrkrwercdpbygmdwnrplfcldbsynbhwwjskuiukwginjzshq", "gpyssipujykheiowxjvdjtwwcwclryvwulofqevgjxsedmtobb", "pvdqfkmerqxxfxythlprfpvkakcpyzdzffknihrorpdhoheltq", "mmpmnotihocwgphilbjwzycnimnyctbevchkpiizmayctoczzo", "yfybopgonhglpmpflysombyhzhkrebsfvtqivhvvdxzimlfotx", "uieepqzpftgetqzkiyogdjchxxienucxffdnlkltumwjsxpykf", "krrkusbjisoiiztzfdkmpqwrxeziulwuycsohglqtqwogtsebh", "yfphycvkzpepsmoqotxtypclmwnawbfewlbjysrqmanmxcusah", "wvkarpvctrpmnlqvbnpfrcprydlwfbvywsxjxssinuunbspyai", "szhyuxyhtatzgpndjidprdyravclcmekniomstmafqasjykxpw", "ojtbmwmjmxjcovjlwydnardhuzrmmxdfsbbzzffhonhytgphnx", "wvgxfethxgmpngwavzkvzwauqoqexwvdkualcwesjwmmpninaq", "xylolqutvozyupmasxpiyiaehjaebddaroouwgupodapubnigv", "metltcwpqgrdxfubkbxdvxrxfqjefjxrosefjcdemyidjjgqvj", "nehivcgzaarnhcfsbpaztfepowqbwzspxvtqvfhlsbdpkgvikz", "xbbniitconjfhuwnzdlhtfkrvqypoydqzkdhtcyvcksetyujwc", "gomlhkxmehtnixxpwkkexrehjugingmofyokgtisguqxvzzxuk", "xkrdfwdakjaopfkefzbuxyhhhgwbgnaofcglefzreneqwsbhng", "uiytrpwamoflxxlslmqxhslzyxjfdcmbudhifpaekcqjbgxpck", "glhxdmdkgfnkqddpzzzvwtvskmodraztwepqbijcihgufnhlyv", "gbdddksdrihjcrlmrvlfhfwmcimakaiahccwgasbnvyipszmbl", "qiqwajunbwvkxlponldsgqgydcrnyivayctiofahqskvzwadij", "kofvmrzbcddeuviiunquxgxvuyafqdrwlfteqlicwyxsxuwxcy", "bebhvtwdvwldlybrrjmerwncbixxyjvzzhehenoajnbmirfepr", "bvrupgghcbybbwpsxsvpttlammdsvftlmacoqcqbatibbksipk", "qdrtudlyxzpysovmtmthjydignxoarxlkuhyiaelzsbvkzuttt", "uthnngxpvlnvsnwpyrkbgngqiiktkcidhgvelchaydltyavrnx", "oiiioklpuceznstfdsyjgzytpmptdgtcbxqrxdvqbejjollqpr", "xcbmnkqkoxbkybrzwksyalmjveyuagbddgwyyiljusytpxkiwh", "ipzfmgdlvgxgivoxbklzvgpbiwxjglzslyqorzvknqovpawfao", "mmuubdscyuinjnajxygmpjmoocwfpdrwuifavbujihrkqoyfyu", "lerobprjnzxlaotfkvsnjbiseptjhrpujgmtmdcwtxhbrjjbqs", "vvfjyxdveauyncdffmjlrlktgmdvafbjgpbkrmuugvomgtgaqz", "rumgbodihspvvnitwofdemujlbuqzyqsttagsdhqatfkorbwcb", "azluxlzxifauywmmvqbrfbhztybqqrzmdcqbgqsbgueobutyce", "qscqyzajodzsokgnipithbmbhpvtemekdajiynjuluqlueefpc", "taxyvtdnjsridiqgquvenxtldbsngobauegbepybnwkrixjevq", "sddxvltlyiwguiompafvzxracnwlwibkrfwpgdtefrgkwjczsv", "whatzhclkbobfqbiirngefwoeuwkqivwfjqakbyfwvhulwzhfs", "zhanwkbyfdlfvqfxyotynjrrixhhaiclkedzwxubpqeukmuurx", "boqwgjbzivdfzkpgdhawwilihkigagkogfcezmrhzdschzemve", "mvihlxoeowdvswgsgancprchxotwjfecvjtmmjpiwnzuowyanp", "zabqbmsgavihemexsnsotleomqftzzfmskonjwqfrhjcdlamsq"})

	assert.Equal(t, 51, got)
}

func TestCase8(t *testing.T) {
	got := numMatchingSubseq("iwdlcxpyagegrcnrcylxolxlnhhwnxyzltiscrjztiivnpnzlubzpueihinsqdfvypdteztiodbhaqhxskupwulvkzhczdyoouym",
		[]string{"xgfypskupoouym", "xryxskuzd"})

	assert.Equal(t, 2, got)
}

func TestCase7(t *testing.T) {
	got := numMatchingSubseq("iwdlcxpyagegrcnrcylxolxlnhhwnxyzltiscrjztiivnpnzlubzpueihinsqdfvypdteztiodbhaqhxskupwulvkzhczdyoouym",
		[]string{"hhwnxyzltiscrjztiivnpnzlubzpueihinsqdfvyp", "vnpnzlubzpueihinsqdfvypdteztiodbha", "rcnrcylxolxlnhhwnxyzltiscrjztiivnpnzlubzpueihi", "dfvypdteztiodbhaqhxskupwulvk", "zltiscrjztii", "wdmbatbcewwittubryrqwwrvfkrmniomofygybeqfzusrgeart", "myzfexqmzxnbmmnhmpbddqhrwrobqzjiwdzzpyzodejysuuquc", "wxvrcbihbasohfvuwuxleesqeujxvjfvgwnhltenbspdgzsdrs", "nztyysfhfbfcihyeaqdarqxfpjunevabzafvbmpbtenarvyizv", "nivufheyodfjuggrbndyojeahrzgptikjfqufhwyhzyyjteahx"})

	assert.Equal(t, 5, got)
}

func TestCase6(t *testing.T) {
	got := numMatchingSubseq("bdeepborrtqjmbuxfccdboyiedlqduslwabxqmvjkojrprwsytvueqsziiuqurigxsseuzvutxpgyreozjmfgliznagkfblqgyhxw",
		[]string{"exwqkcihfcifqhpprnsckbyxdpjbxahyrgdziduzbugfzlxmqy", "rtqjmburpwkfxw"})

	assert.Equal(t, 1, got)
}

func TestCase5(t *testing.T) {
	got := numMatchingSubseq("abcce", []string{"e"})

	assert.Equal(t, 1, got)
}

func TestCase4(t *testing.T) {
	got := numMatchingSubseq("abcce", []string{"ce"})

	assert.Equal(t, 1, got)
}

func TestCase3(t *testing.T) {
	got := numMatchingSubseq("abcde", []string{"z", "c"})

	assert.Equal(t, 1, got)
}

func TestCase2(t *testing.T) {
	got := numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"})

	assert.Equal(t, 2, got)
}

func TestCase1(t *testing.T) {
	got := numMatchingSubseq("abcde", []string{"a", "acd", "bb", "ace"})

	assert.Equal(t, 3, got)
}
