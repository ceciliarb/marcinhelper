package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "time"
  "github.com/enescakir/emoji"
)

func loopAlteracao(textos, opcoes []string, sep string) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("%v \tDigite o índice a ser alterado (ou 'kocham Ceci!' para sair):\n", emoji.Pencil)
  ind, _ :=  reader.ReadString('\n')
  ind =  strings.TrimSpace(ind)
  if strings.EqualFold(ind, "kocham ceci!") {
    fmt.Printf("%v \tKocham Marcin!!!\n", emoji.RedHeart)
    time.Sleep(2 * time.Second)
    os.Exit(0)

  } else {
    dict := make(map[string]string)
    for i, item := range opcoes {
      dict[item] = textos[i]
    }

    textoToSubs, ok := dict[ind]
    if !ok {
      fmt.Printf("%v \tO índice fornecido (%s) não existe!\n", emoji.PileOfPoo, ind)
    } else {
      fmt.Printf("%v \t(índice = %s / valor = %s)\n", emoji.MagnifyingGlassTiltedRight, ind, dict[ind])
      fmt.Printf("\t %v \t Digite o valor a ser alterado:\n", emoji.Pencil)
      val, _    := reader.ReadString('\n')
      textosStr := strings.Join(textos, sep)
      textosStr  = strings.Replace(textosStr, textoToSubs, strings.TrimSpace(string(val)), 1)
      textos     = strings.Split(string(textosStr), sep)
      fmt.Printf("%v \t", emoji.MagicWand)
      fmt.Println("(" + strings.Join(opcoes, sep) + ")(" + textosStr + ")\n")
    }
    loopAlteracao(textos, opcoes, sep)
  }
}

func loopInicial() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Printf("%v \t Insira a linha a ser alterada: \n", emoji.Pencil)
  linha, _ := reader.ReadString('\n')

  fmt.Printf("%v \t Insira o separador (ou enter para ','): \n", emoji.Pencil)
  sep, _ := reader.ReadString('\n')
  sep = strings.TrimSpace(sep)
  if sep == "" {
   sep = ","
  }

  // removendo parenteses externos
  input := strings.TrimSpace(string(linha))[1:len(linha)-2]
  // quebrando 2 partes
  input = string(strings.Replace(input, ")(", "|", -1))
  inputArr := strings.Split(input, "|")
  if len(inputArr) != 2 {
      fmt.Printf("%v \tFormato invalido!\n", emoji.PileOfPoo)
      loopInicial()
  } else {
    // a primeira parte representa as opcoes    
    opcoes := strings.Split(string(inputArr[0]), sep)    
    // a segunda parte representa os valores    
    textos := strings.Split(string(inputArr[1]), sep)    
 
    // print bonitinho    
    fmt.Printf("%v \t", emoji.MagnifyingGlassTiltedRight)    
    fmt.Println("(" + strings.Join(opcoes, sep) + ")")    
    fmt.Println("(")                      
    for _, item := range textos {         
      fmt.Println(item)                   
    }                                     
    fmt.Println(")")                      
                                        
    // loop de inputs de alteracao        
    loopAlteracao(textos, opcoes, sep)  
  }
}

func main() {
  loopInicial()
}

