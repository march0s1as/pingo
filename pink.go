package main
import (
	"github.com/bogdanovich/dns_resolver"
	"github.com/fatih/color"
	"net/http"
	"os/exec"
	"strings"
	"regexp"
	"bufio"
	"log"
	"fmt"
	"os"
)

func clear() {
    out, err := exec.Command("clear").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    output := string(out[:])
    fmt.Println(output)
}

func ascii(){
	clear()
	color.Magenta(`   
	    ..uu.
           ?$""'?i           z'
           'M  .@"          x"
           'Z :#"  .   .    f 8M
           '&H?'  :$f U8   <  MP   x#'
           d#'    XM  $5.  $  M' xM"
         .!">     @  'f'$L:M  R.@!'
        +'  >     R  X  "NXF  R"*L
            k    'f  M   "$$ :E  5.
            %    '~  "    '  'K  'M
                             'E   'h
                              X     ' 
`)
}


func diretorio(){
	ascii()
	var numero int
	fmt.Println("[01] Díretórios")
	fmt.Println("[02] Subdomínios")
	fmt.Print(":/ ")
	fmt.Scan(&numero)
	if numero == 1 && numero == 01{	
		ascii()
		var site string
		fmt.Println("λ Qual site você deseja? (com http/https)")
		fmt.Print(":/ ")
		fmt.Scan(&site)
		ascii()
		fmt.Println("[" + site + "]")
		fmt.Println(" --- ")

		// -- wordlist --
		var wordlist string
		fmt.Println("λ Digite aqui o path da wordlist.")
		fmt.Print(":/ ")
		fmt.Scan(&wordlist)

		ascii()
		fmt.Println("[" + site + "]")
		fmt.Println(" --- ")
		// -- ler a wordlist --
		
		file, err := os.Open(wordlist)
	 
		if err != nil {
			fmt.Println("falha ao ler arquivo, favor verificar o path.")
		}
	 
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string
	 
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
	 
		file.Close()
	 
		for _, eachline := range txtlines {
			//fmt.Println(eachline)
			var sitezin = site + eachline
			req, err := http.Get(sitezin)
			if err != nil{
				fmt.Println("falha ao iniciar conexão com o site, favor verificar o path.")
			}
			//ua := req.Header.Get("Content-Length")
			//if req.Status != "404 Not Found" && req.Status != "404 File not found"{
			if req.Status == "200 OK"{
				fmt.Print(sitezin + " ")
				color.Green("[200]")
			} else if req.Status == "403 Forbidden"{
				fmt.Print(sitezin + " ")
				color.Yellow("[403]")
			} else if req.Status  == "301 Moved Permanently"{
				fmt.Print(sitezin + " ")
				color.Magenta("[301]")
			} else if req.Status == "500 Internal Server Error"{
				fmt.Print(sitezin + " ")
				color.Red("[500]")
			} else if req.Status != "404 Not Found"{
				fmt.Println("[" + req.Status + "] " + sitezin)
			}
		}
	} else if numero == 2 && numero == 02{

		var site string
		var wordlist string
		ascii()
		fmt.Println("λ Qual site você deseja? (com http/https).")
		fmt.Print(":/ ")
		fmt.Scan(&site)
		ascii()
		fmt.Println("λ Digite abaixo o path da wordlist.")
		fmt.Print(":/ ")
		fmt.Scan(&wordlist)
		ascii()

		file, err := os.Open(wordlist)
		if err != nil{
			fmt.Println("falha ao ler wordlist, favor verificar o path.")
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string
		for scanner.Scan(){
		txtlines = append(txtlines, scanner.Text())
		}
		file.Close()

		matched, err := regexp.MatchString(`http://`, site)
	   	if err != nil {
	        fmt.Println(err)
	        os.Exit(1)
	    }
	    if matched == true{
	    	// http
	    	site_sem_http := strings.TrimLeft(site, "http://")
	    	fmt.Println("site com tracinho -> ", site_sem_http )

	    	tracinho, err := regexp.MatchString(`/`, site_sem_http)
	    	if err != nil{
	    		fmt.Println(err)
	    		os.Exit(1)
	    	}
	    	if tracinho == true{
	    		site_sem_tracinho_http := strings.TrimSuffix(site_sem_http, "/")
	    		ascii()
	    		fmt.Println("realizando a bruteforce..")
	    		fmt.Println("---")
	    		resolver := dns_resolver.New([]string{"8.8.8.8", "8.8.4.4"})
	    		resolver.RetryTimes = 5
	    		for _, eachline := range txtlines {
				//fmt.Println(eachline)
				var site_subdomain = eachline + "." + site_sem_tracinho_http
				ip, err := resolver.LookupHost(site_subdomain)
			if err != nil{
				continue
			}
			_ = ip
			log.Println("http://" + site_subdomain)
		}	

	    	}

		} else{
			site_sem_https := strings.TrimLeft(site, "https://")
	    	fmt.Println("site com tracinho -> ", site_sem_https )

	    	tracinho, err := regexp.MatchString(`/`, site_sem_https)
	    	if err != nil{
	    		fmt.Println(err)
	    		os.Exit(1)
	    	}
	    	if tracinho == true{
	    		site_sem_tracinho_http := strings.TrimSuffix(site_sem_https, "/")
	    		ascii()
	    		fmt.Println("realizando a bruteforce..")
	    		fmt.Println("---")
	    		resolver := dns_resolver.New([]string{"8.8.8.8", "8.8.4.4"})
	    		resolver.RetryTimes = 5
	    		for _, eachline := range txtlines {
				//fmt.Println(eachline)
				var site_subdomain = eachline + "." + site_sem_tracinho_http
				ip, err := resolver.LookupHost(site_subdomain)
			if err != nil{
				continue
			}
			_ = ip
			log.Println("https://" + site_subdomain)
				}	
	    	}
		}
	}
}

func main(){
	diretorio()
}