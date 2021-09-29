package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "github.com/kyokomi/emoji/v2"
)

func loopAlteracao(textos, opcoes []string, sep string) {
  reader := bufio.NewReader(os.Stdin)
  emoji.Println(":pencil2: Digite o índice a ser alterado (ou 'kocham Ceci!' para sair):")
  ind, _ :=  reader.ReadString('\n')
  ind =  strings.TrimSpace(ind)
  if strings.EqualFold(ind, "kocham ceci!") {
    emoji.Println(":heart: Kocham Marcin!!!")
    os.Exit(0)

  } else {
    dict := make(map[string]string)
    for i, item := range opcoes {
      dict[item] = textos[i]
    }

    textoToSubs, ok := dict[ind]
    if !ok {
      emoji.Printf(":poo: O índice fornecido (%s) não existe!\n", ind)
    } else {
      emoji.Printf(":magnifying_glass_tilted_right: (índice = %s / valor = %s)\n", ind, dict[ind])
      emoji.Println("  :pencil2: Digite o valor a ser alterado:")
      val, _    := reader.ReadString('\n')
      textosStr := strings.Join(textos, sep)
      textosStr  = strings.Replace(textosStr, textoToSubs, strings.TrimSpace(string(val)), 1)
      textos     = strings.Split(string(textosStr), sep)
      emoji.Println(":magic_wand: (" + strings.Join(opcoes, sep) + ")(" + textosStr + ")")
    }
    loopAlteracao(textos, opcoes, sep)
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  emoji.Println(":pencil2: Insira a linha a ser alterada: ")
  linha, _ := reader.ReadString('\n')

  emoji.Println(":pencil2: Insira o separador (ou enter para ','): ")
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
  // a primeira parte representa as opcoes
  opcoes := strings.Split(string(inputArr[0]), sep)
  // a segunda parte representa os valores
  textos := strings.Split(string(inputArr[1]), sep)

  // print bonitinho
  emoji.Println(":magnifying_glass_tilted_right:")
  fmt.Println("(" + strings.Join(opcoes, sep) + ")")
  fmt.Println("(")
  for _, item := range textos {
    fmt.Println(item)
  }
  fmt.Println(")")

  // loop de inputs de alteracao
  loopAlteracao(textos, opcoes, sep)


}
