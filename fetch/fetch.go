// gopl.io/ch1/fetch
// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// este programa recebe URLs como argumentos da linha de comando, faz uma solicitação HTTP GET para cada URL, lê e imprime o conteúdo da resposta. Se ocorrerem erros durante o processo, ele imprime mensagens de erro apropriadas e termina com um código de saída diferente de zero.
