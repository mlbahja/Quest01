package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func combn(n int) {
	// 0 >>>9
	if n < 0 || n > 10 {
		return
	}
	if n == 1 {
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
	if n == 2 {
		for i := '0'; i <= '8'; i++ {
			for j := '1'; j <= '9'; j++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				if i != '8' || j != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 3 {
		for a := '0'; a <= '7'; a++ {
			for b := '1'; b <= '8'; b++ {
				for c := '2'; c <= '9'; c++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if a != '7' || b != '8' || c != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 4 {
		for a := '0'; a <= '6'; a++ {
			for b := '1'; b <= '7'; b++ {
				for c := '2'; c <= '8'; c++ {
					for d := '3'; d <= '9'; d++ {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune(d)
						if a != '6' || b != '7' || c != '8' || d != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 5 {
		for a := '0'; a <= '5'; a++ {
			for b := '1'; b <= '6'; b++ {
				for c := '2'; c <= '7'; c++ {
					for d := '3'; d <= '8'; d++ {
						for e := '4'; e <= '9'; e++ {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
							z01.PrintRune(d)
							z01.PrintRune(e)
							if a != '5' || b != '6' || c != '7' || d != '8' || e != '9' {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 6 {
		for a := '0'; a <= '4'; a++ {
			for b := '1'; b <= '5'; b++ {
				for c := '2'; c <= '6'; c++ {
					for d := '3'; d <= '7'; d++ {
						for e := '4'; e <= '8'; e++ {
							for f := '5'; f <= '9'; f++ {
								z01.PrintRune(a)
								z01.PrintRune(b)
								z01.PrintRune(c)
								z01.PrintRune(d)
								z01.PrintRune(e)
								z01.PrintRune(f)
								if a != '4' || b != '5' || c != '6' || d != '7' || e != '8' || f != '9' {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 7 {
		for a := '0'; a <= '3'; a++ {
			for b := '1'; b <= '4'; b++ {
				for c := '2'; c <= '5'; c++ {
					for d := '3'; d <= '6'; d++ {
						for e := '4'; e <= '7'; e++ {
							for f := '5'; f <= '8'; f++ {
								for g := '6'; g <= '9'; g++ {
									z01.PrintRune(a)
									z01.PrintRune(b)
									z01.PrintRune(c)
									z01.PrintRune(d)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune(g)
									if a != '3' || b != '4' || c != '5' || d != '6' || e != '7' || f != '8' || g != '9' {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 8 {
		for a := '0'; a <= '2'; a++ {
			for b := '1'; b <= '3'; b++ {
				for c := '2'; c <= '4'; c++ {
					for d := '3'; d <= '5'; d++ {
						for e := '4'; e <= '6'; e++ {
							for f := '5'; f <= '7'; f++ {
								for g := '6'; g <= '8'; g++ {
									for h := '7'; h <= '9'; h++ {
										z01.PrintRune(a)
										z01.PrintRune(b)
										z01.PrintRune(c)
										z01.PrintRune(d)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(g)
										z01.PrintRune(h)
										if a != '2' || b != '3' || c != '4' || d != '5' || e != '6' || f != '7' || g != '8' || h != '9' {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
	if n == 9 {
		for a := '0'; a <= '1'; a++ {
			for b := '1'; b <= '2'; b++ {
				for c := '2'; c <= '3'; c++ {
					for d := '3'; d <= '4'; d++ {
						for e := '4'; e <= '5'; e++ {
							for f := '5'; f <= '6'; f++ {
								for g := '6'; g <= '7'; g++ {
									for h := '7'; h <= '8'; h++ {
										for k := '8'; k <= '9'; k++ {
											z01.PrintRune(a)
											z01.PrintRune(b)
											z01.PrintRune(c)
											z01.PrintRune(d)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(g)
											z01.PrintRune(h)
											z01.PrintRune(k)
											if a != '1' || b != '2' || c != '3' || d != '4' || e != '5' || f != '6' || g != '7' || h != '8' || k != '9' {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		z01.PrintRune('\n')
	}
}
func main() {
    combn(1)
	fmt.Println()
	combn(2)
	fmt.Println()
	combn(3)
	fmt.Println()
	combn(4)
	fmt.Println()
	combn(5)
	fmt.Println()
	combn(6)
	fmt.Println()
	combn(7)
	fmt.Println()
	combn(8)
	fmt.Println()
	combn(9)
	fmt.Println()
	combn(0)
	fmt.Println()
}
