package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Iniciando a instância http com a rota e uma função executada toda vez que é acessada
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Recebendo um array de nomes
		// ex: localhost:3000/?name=Jack&name=Shop
		// ficaria [Jack, Shop]
		names := r.URL.Query()["name"]
		var name string
		if len(names) > 1 {
			name = names[0]
		}

		// Transformo o nome em um objeto que vai ser { string: string}
		m := map[string]string{"name": name}

		// atribuo meu response write ao json encoder, dizendo que vou retornar um JSON
		enc := json.NewEncoder(w)

		// dou um pipe do json pro encoder escrito acima, retornando o json ao usuario
		enc.Encode(m)
	})

	// Iniciando o servidor na porta 3000
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
