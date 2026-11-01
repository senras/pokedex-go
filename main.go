package main

func main(){
	cfg := &config{}
	initConfig(cfg)
	startRepl(cfg)
}
