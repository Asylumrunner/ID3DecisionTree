package main

import (
  "strings"
  "os"
  "bufio"
)

type Instance struct {
  attributes []string
  class string
}

type InstanceSet struct {
  Instances []*Instance
}

func CheckError(err error) {
  if err != nil{
    panic(err)
  }
}

func MakeInstance(attlist string) (*Instance, error) {
  i := new(Instance)

  a := strings.Split(attlist, " ")
  i.class = a[len(a)-1]

  a = a[:len(a)-1]
  i.attributes = a

  return i, nil
}

func TreeBuild() {

}

func main() {
  // First, open the training file and set it up for execution
  trainingFileName := os.Args[1]
  trainingFile, err := os.Open(trainingFileName)
  CheckError(err)
  defer trainingFile.Close()

  // We use bufio.scanner to read the training file line by line
  scanner := bufio.NewScanner(trainingFile)

  var set *InstanceSet = new(InstanceSet)

  // The first line is variable names. Discard it, we don't care.
  scanner.Scan()

  // Read every line of the training file after the first, and convert to Instances
  for scanner.Scan() {
    inst, readErr := MakeInstance(scanner.Text())
    CheckError(readErr)

    set.Instances = append(set.Instances, inst)
  }



}
