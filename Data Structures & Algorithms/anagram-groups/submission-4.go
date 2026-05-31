import (
	"crypto/sha256"
	"encoding/hex"
)

func hashMap(m map[rune]int) (string, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(jsonData)
	return hex.EncodeToString(hash[:]), nil
}


func groupAnagrams(strs []string) [][]string {
	counts := map[string][]string{}

	for _, s := range strs {
		count := map[rune]int{}
		for _, c := range s {
			count[c]++
		}
		hash, _ := hashMap(count)
		_, ex := counts[hash]
		if ex {
			counts[hash] = append(counts[hash], s)
		} else {
			counts[hash] = []string{s}
		}
	}

	res := [][]string{}
	for _, v := range counts{
		res = append(res, v)
	}
	return res
}
