package main

import "gitlab.com/tsuchinaga/monorepo-test/products/libs/logger"

func main() {
	logger.Get("greet").Println("こんにちわーるど")
}
