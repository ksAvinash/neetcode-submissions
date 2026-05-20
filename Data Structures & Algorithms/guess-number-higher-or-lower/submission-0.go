/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    i, j := 1, n
	for i < j {
		mid := int((i+j)/2)
		res := guess(mid)
		if res == 0 {
			return mid
		}
		if res == -1 {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return i
}
