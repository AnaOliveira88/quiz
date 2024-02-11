package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type questaostruc struct {
	PerguntaNum   int
	Pergunta      string
	Resposta1     string
	Resposta2     string
	Resposta3     string
	RespostaCerta int
}

func certas(arrayQuestoesRW []int) int {
	var j int
	for j = 0; j < len(arrayQuestoesRW); j++ {
		if arrayQuestoesRW[j] == 0 {
			return 0
		}
	}
	return 1
}
func mostraErradas(arrayQuestoesRW []int, q []questaostruc) {
	var j int
	for j = 0; j < len(arrayQuestoesRW); j++ {
		if arrayQuestoesRW[j] == 0 {
			var resp = q[j].RespostaCerta
			if resp == 1 {
				fmt.Println(q[j].PerguntaNum, "-", q[j].Pergunta, ", Resposta certa", q[j].Resposta1)
			} else {
				if resp == 2 {
					fmt.Println(q[j].PerguntaNum, "-", q[j].Pergunta, ", Resposta certa", q[j].Resposta2)
				} else {
					fmt.Println(q[j].PerguntaNum, "-", q[j].Pergunta, ", Resposta certa", q[j].Resposta3)
				}
			}

		}
	}
}

func questaofunc(nlinhas int, q []questaostruc, arrayQuestoesRW []int, sn int) int {
	var resposta, simNao, tudoCerto int
	var i int
	simNao = 2
	for i = 1; i < nlinhas; i++ {
		if sn == 1 {
			if arrayQuestoesRW[i] == 0 {
				fmt.Println("Pergunta ", q[i].PerguntaNum, "\n", q[i].Pergunta)
				fmt.Println("1", q[i].Resposta1)
				fmt.Println("2", q[i].Resposta2)
				fmt.Println("3", q[i].Resposta3)
				fmt.Scanf("%d", &resposta)
				if int(resposta) == int(q[i].RespostaCerta) {
					fmt.Println("CERTOOO")
					//acertou = 1
					arrayQuestoesRW[q[i].PerguntaNum] = 1
					fmt.Println(arrayQuestoesRW)
				} else {
					fmt.Println("Errado!!!")
				}
			}
		} else {
			fmt.Println("Pergunta ", q[i].PerguntaNum, "\n", q[i].Pergunta)
			fmt.Println("1", q[i].Resposta1)
			fmt.Println("2", q[i].Resposta2)
			fmt.Println("3", q[i].Resposta3)
			fmt.Scanf("%d", &resposta)
			if int(resposta) == int(q[i].RespostaCerta) {
				fmt.Println("CERTOOO")
				//acertou = 1
				arrayQuestoesRW[q[i].PerguntaNum] = 1
				fmt.Println(arrayQuestoesRW)
			} else {
				fmt.Println("Errado!!!")
			}
		}
	}
	tudoCerto = certas(arrayQuestoesRW)
	if tudoCerto == 1 {
		fmt.Println("Questoes respondidas corretamente. obrigada")
		return 1
	} else {
		fmt.Println("Pretendes repetir as questões erradas? Sim: 1, Nao: 0")
		for simNao != 0 || simNao != 1 {
			fmt.Scanf("%d", &simNao)
			if simNao == 1 {
				//repete so para incorrectas
				questaofunc(nlinhas, q, arrayQuestoesRW, simNao)
				return 1
			} else {
				if simNao == 0 {
					fmt.Println("obrigada. Estas foram as tuas questoes erradas:")
					mostraErradas(arrayQuestoesRW, q)
					return 0

				} else {
					fmt.Println("numero nao corresponde a nada")
					return 0
				}
			}
		}
	}
	return 1
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Fprintln(w, "<h1>hello world</h1>")
	w.Write(f)
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile("about.html")
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Fprintln(w, "<h1>hello world</h1>")
	w.Write(f)
}
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	http.ListenAndServe(":3002", nil)
	/*
		   	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		   		resp := []byte(`{"status": "ok"}`)
		   		rw.Header().Set("Content-Type", "application/json")
		   		rw.Header().Set("Content-Length", fmt.Sprint(len(resp)))
		   		rw.Write(resp)
		   	})

			log.Println("Server is available at http://localhost:8002")
			log.Fatal(http.ListenAndServe(":8002", handler))
	*/
	var nlinhas = 0
	var retorno int = 0
	//var arrayQuestoes []int // slice de inteiros
	//	var numero int

	//arrayQuestoes[0] "Pergunta" = "Acertou na Resposta ?"

	file, err := os.Open("rh.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	questoes := make([]questaostruc, len(records))
	nlinhas = len(records)
	// le as perguntas todas do xls e coloca no array questoes
	for k, record := range records {
		//fmt.Println(k)
		respostaCerta, _ := strconv.Atoi(record[5])
		perguntaNumero, _ := strconv.Atoi(record[0])
		questao := questaostruc{
			PerguntaNum:   perguntaNumero,
			Pergunta:      record[1],
			Resposta1:     record[2],
			Resposta2:     record[3],
			Resposta3:     record[4],
			RespostaCerta: respostaCerta,
		}
		//fmt.Println(questao)
		questoes[k] = questao // todas as perguntas e respostas e resposta certa estao aqui guardadas na struct
	}
	// Inicia um array com o numero de perguntas[i], tudo = 0 como se fosse ERRADO
	//flagg := 0
	fmt.Println("numero de linhas:", nlinhas)
	arrayQuestoesRW := make([]int, nlinhas, nlinhas)

	arrayQuestoesRW[0] = 1 // a primeira não é questao, é enunciado
	retorno = questaofunc(nlinhas, questoes, arrayQuestoesRW, 2)
	fmt.Println(retorno)

	/*
		for nlinhas > flag {
			// fazemos o seed do gerador de números aletórios
			rand.Seed(time.Now().UnixNano())
			// vamos gerar um número inteiro aleatório
			// na faixa de 1 a 10 (incluindo os dois números)
			numero = rand.Intn(nlinhas)
			fmt.Println(numero)
			arrayQuestoes[numero] = 0
			//fmt.Println(numero)
			//questaofunc(questoes[numero], arrayQuestoes)
			flag++
		}*/
	//arQuesRW = iniciaArray(nlinhas, questoes)
	//fmt.Println(arrayQuestoes)

}
