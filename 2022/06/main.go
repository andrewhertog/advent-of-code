package main

import "fmt"

var input = "qnjjqgjqgglqqwrrvvcwwtjjzpjjfwwtwzwmzmrmttrvrpppbtbllhrrwtwzwjwzznfzfvvvngnrrhzrhhsmsrrsrrtqrrpzpnpfphfhllwlclfcfrrqffhvhsvswvswwngnzzcmmnrnzrnnwhwgwvggpcggnrntrrgmgdgmmgzmzssprrmddjndnsnrsrddnmdmqddrzznwznndfnndjjmgjmjmmnpmnmrmzmfzmmwhwzwrwqqqhshhncnvvhfvhfhzzvttrnrdnnwddptphpthhbgbgmbbvppnhphmmncnwcwqwppgddsqqjqppcrcncvnccjsjzztczttjzzwwhshnsnttrctchthnnswnnfzfhhsfhhdccgwcggbrblrlwlrlnntsscffpnnwffqgqdqhqhphhvtvmvhmmnssqbqfqfwwbfwbfwwzwgwgpgwwchcshslhssqsppqdqhhnssrzzvbbhjjqdjdgjgttvzzmvzvvnmvvfbbncbcqcpqpbbcjcncjnccjfjpjdjvdddbzzfrfwfpwpfwppvjvfjjhfhvvqttfhffscsvvbqbzqqjssbmsbshsrshsphhwhvwhwcwwnjjlzzfgfrgrwggfddpqpqqcmmtwtcwttgjjnfnnjvvmfmccvhchjcjmccmmwgghccnffjccqczqqbdbrdrsdsffmzzndnncbnbgbsshbhvbvsbvsbbvjbvbdvvqpqnpngpnggbbndbbdzdhzdhhjcjjmbbqgqbbfrrvzvcvffhggcbcjbbqtbbszbbvvcnvnccldddlrrtqrttpjpnjpnpznnqmnnrrmmsdmdbbffcwwllnffssspvvrjjwhwmmqhmmhjjtnjtjgttnrtnnsgsbssffstssjvjwwwwnjngjnnnlndndqdqqwdwcdcsccsgssgrrqwqqqhgqgfgffmnffzjjssqsfqqgddgcgtgrrvgrrdhdcdgdbdsswgwwpwggdfggnjnnvffzhhcqclqclcdlldbldlffdrrzttnhthdthhrthhgcgllzqlzzfhhmshssdvvtnnpngpprqrrrbnbssnhnnqdndsnshhzjjrnjnmjjwfffpbpnbbwddpwptwwpbpffvggmjmgghrrnqqhnhsnhsnnnlddmjddrdrhrchhgmgfmfdfcfpprsrzzzmhzhmhrmmrwmmpdmpmcclmlvvlffhlhhczccsvvptvpvwvdvqqdvqdddvnvsnswwspwpqqpddpvvsqsjjwnwrrjbjtttdhhhwbhbzzqrqfrrzlrzrqrbqrqzzfczfzrrqffnrnhrrcrtrvrffvflfsllcjcvvmqmlqqwjwgglhlvldltlqqdcddtssltsllsnndvdnddfqdffprrbssgccqmqrrnsrnnlzlhhqrhqhcqcjqcqpcqqmvqqgbgpbggnfggdpgpvpgvvsrrhqhshmsmwsmwssvnvtnttphpdpdspsvslshhlplmplmmljjvnnsrslrrvttqqprrptrptpnnmmflltntrrnjjcljcllwcllddbgbjbrrtwrtrftfctftsslsggggzqqhhrvvdqqwqbqtthdhrrpwpqpmmqcmqmllsmllhvfwbfvgzqgbhlfhqrqtzfpplgjtgngzrdfltnqlwsbhmwdfvrdjlgwftjvmdsgdgnswhrtmzgfqfsfnczjnmgqcfzvwlbgzsrpcbwwtrmbqtqhdmhmscqgjgpqdqcrvdvwplpdzsjqbvgpqcvfspqrcsjlhrqpjmdszjpqhmdwtddwqhbwsrlcjpzwsjjvbzcllqfwbhfvjqbqfbsrbgvgchdmgqjnvdrzlmrnlpzrljgjvtrdtqnzbnhpgpgjvwttmnfbpvrtmpstbmtwdwfzvznrwmspftgvrmdfwqltzzmlgrvvwjgdblnnbjzjfqpldsqbhrstnhrjqfcmzcrtqcpqmmfqzndgjwtcgnwrgdznzdgzbvmrlvrjvjgmfcrmrjbpjwqvhprbphnqsbpcpnflpgnpgggqpgrwghfpfvdljjqnvqgbvcpjbjlqghjppfhzfzczmcwnhrjzzrwlfqhvdwqrcbvprclnmqmqwdrhgtswwqhqtcjhndrmcrzdjnvwsstzplhcdzwzqbjjthsmjrmpfbqljlmnvmfddqflnhrfhchzgnlbcpvppnspsdtnqmfhrztznhltmqlrwzgczqmjggvnphwpvrwcnqhsdfglblvvdlqfnghtzgngtjnrzsjgwcglsrczlfqpmmzsrqwcclslzjpvpfgcwzwsfmwwswdsvnvtmwzmhwzvnnnldnjljbwlglpjvrggbmvhctsggtjgrmjrglnhtjzwhrvtmfmphgmpdmvlhzrhrtjcwswlnnjphsschrdvstmflzqsbqmqggjjdsdtjhcvjlzclbpvnbjgngnzvtdgszdtrwpnjmgwpcjpbsvrsmdjdjdbhqqvplbmfhmbgmtggvtltqjphlggscdbzncqhqqlmflmjbpnhbnqjcbfmggnftjjbzbrqwcgvfhsddrqgrvnztzzqzwjrjprmqctlttmgltdgrcvnbrbzrsbbqnsrslrswvmrqlctvsdfmcrdddcdwvfmwwcbqsdhmmgplzhnsfvhtctmvrzntvptpnbhhvjlphmtcctlhdvtvvqbmbczjqfvnqwzbgjzlgnqhvbjjnvrgfprvllltqdwvvmchrdqfczmmbmcjwtwgqbfjzgpcfzhhnwcghqltrlhlntjcrthpsbmcdzzqljdrdhfctwrzlvhdjvgmpscssmhdggrnhfbtfpqtsbqnqjfjwlbdfcpzfjtspnrjzshjcrrzwclddqqbplwghrcjhfqgblfrlrdhmdqmrzdnqthvvzlcjmqzqtgbcwcprpnjtbhsqmzlprhmssltfvvgqzzpnjtsmplhfpsnznjwrrgvmbbvjzzwctmwcwcwjwvmlpcmlmrfqbdjpclsqjnsnndzqfcghqsmhqcjwjjbbwmsttwnththhhgrlhrtrcjppvmjlqtqhpbhpsdhzcqpmqqbvwrlvjrnmlfvtntwghtphwzqmbmmpgvfrqsjwffdbcjjbwrrntrhnjwjfprzjcpnwvtwjcdhppspbdbwnzdbhvpqljbrfnpnbqdtqzdjlbvgbvvhcjrqwlqmhnhjfppllrsnsvnmttvrsdrqjfrdvhplwbhhnvglvgcvtzblfswzrrbwpjnslthjmgsdgwqrfjfplmbgsdcltzmdhrbpqggnjrcddltvdbrlhtzwfghbpnlmghwqdtpctmqbggfzflfzjgcbvfqtztsmgdvdlrrdrtmmjbtdtjvvjvlwwlwnzhlfqtrmdgvlflmbffnzlvtccnmcddrqfhgzlhngrfmsqdrnshshfhhtnwmnqmjtzggdfmpgcsrhjzhhvlppfnqmjsvtcwwjdvdrgrldbdnrqtvjjnddgmfdtqrncdlfznfpbqldztzzgpvwhvnqrrzdwsrdvcnslmwcdztvqzgbshpqrdbnmlfltpbsfrnbbcwvwvjnrdspwrwzpljqwwvqhbhhqqjnlrgwqvvsjrscbfqvhbwtjvwzmwzvftlwmtbmglvqvlnwrzvwclvzjcvgfstjpwdbtbhmnqzfgpqmthpchrrqfpqfflwhwrhmdzvgbzlmfgzvhrldprlgrvjcvjhslghqddfvqzgplmtfqjzhpctvtgrzjhwrjzbsqqrcztrstbvdjbtszrbvvzvwzfclsnrftcqtqdghdrnhglzptzwtqtgnfsqhmvnqqwtnwszgtvpnndbzcwgwhpcdlsdpptvlcqtmfhfjzbvfsvnjvqrrjqbggnjqfrrrvqmzscvlltfwzdlzqstlrrmqqmbmbtgbqjsvbncblddgzvlstmrpmrlfcqpvwthgbhlvfcfcwvslcjnztfwgdgwjfrwrzbbszmtvzchvstcfrgqllfrdccrjmtspshgqbzcldddqgjnpdsctmjphcwqvjtvmqlvvpzdgjdlpwdrjbshhrgjtglwnzlzsqngspbzgbwpgmfcv"

// var input = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

func main() {
	res, idx := findMarker(input)
	fmt.Println(res)
	fmt.Println(idx)
}

func findMarker(s string) (string, int) {
	offset := 14
	for i := 0; i < len(s); i++ {
		if i >= offset {
			if !isDuplicate(s[i-offset : i]) {
				return s[i-offset : i], i

			}
		}
	}
	return "", 0
}

func isDuplicate(s string) bool {
	count := 0
	for _, v := range s {
		for _, u := range s {
			if u == v {
				count += 1
			}
		}
	}
	if count > 14 {
		return true
	}
	return false
}
