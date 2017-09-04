package main

import (
  "helloworld"
  "routes"
  "sayhi"
)

func main()  {
  sayhi.Sayhi()
  helloworld.Print_Hello_World()
  routes.Serve("9080")
}
