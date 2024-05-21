package main

func main() {
	sum := 100 + 010
	print(sum) // 108 と出力される

	sum = 100 + 0o10
	print(sum) // 108 と出力される

	bin := 0b10
	print(bin) // 2 と出力される

	hex := 0x10
	print(hex) // 16 と出力される

	ima := 3i
	print(ima) // 3i と出力される

    dec := 1_000_000_000
    print(dec)
}