package main

func main()  {
	Server := NewServer()
	defer Server.Close()
	
	Server.Run()
}