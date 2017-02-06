package main
import (
  "errors"
  "fmt"
  "strconv"
  "encoding/json"

  "github.com/hyperledger/fabric/core/chaincode/shim"
)

type BookChaincode struct {
}

var bookIndexStr = "_bookindex"

type Book struct {
  bookName string `json:"bookName"`
  userName string `json:"userName"`
}

func main() {
  err := shim.Start(new(BookChaincode))
  if err != nil {
    fmt.Printf("Error starting in Book Chaincode: %s", err)
  }
}

//==============================================================================
func (t *BookChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
  var Aval int
  var err error

  Aval, err = strconv.Atoi(args[0])
  //if err != nil {
    //return nil, errors.New("This is here Boy. Expecting Integer value for the asset holding")
  //}

  err = stub.PutState("abc", []byte(strconv.Itoa(Aval)))
  if err != nil {
    return nil, err
  }

  var empty []string
  jsonAsBytes, _ := json.Marshal(empty)
  err = stub.PutState(bookIndexStr, jsonAsBytes)
  if err != nil {
    return nil, err
  }

  return nil, nil
}

func (t *BookChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("Run is running " + function)
  return t.Invoke(stub, function, args)
}

func (t *BookChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("Invoke is running : " + function)

  if function == "init" {
    return t.Init(stub, "init", args)
  } else if function == "init_book" {
    return t.init_book(stub, args)
  } else if function == "write" {
    return t.Write(stub, args)
  }
  fmt.Println("Invoke didn't find any func: " + function)

  return nil, errors.New("Recieved unknown function invocation")
}

func (t *BookChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  fmt.Println("Query is running : " + function)

    if function == "read" {
    return t.read(stub, args)
  }
  fmt.Println("Didn't find any func: " + function)

  return nil, errors.New("Received unknown fucntion query")
}

func (t *BookChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var name, jsonResp string
  var err error

  if len(args) != 1 {
    return nil, errors.New("Incorrect number of arguments.")
  }

  name = args[0]
  valAsbytes, err := stub.GetState(name)
  if err != nil {
    jsonResp = "{'ERROR': Failed to get the state for " + name + "}"
    return nil, errors.New(jsonResp)
  }

  return valAsbytes, nil
}

func (t *BookChaincode) Write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var bookName, userName string
  var err error
  fmt.Println("Running Write()")

  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments for Write Function Call.")
  }

  bookName = args[0]
  userName = args[1]
  err = stub.PutState(bookName, []byte(userName))
  if err != nil {
    return nil, err
  }

  return nil, nil
}

func (t *BookChaincode) init_book(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var err error

  //    0         1
  //bookName  userName
  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments in Init_book")
  }

  fmt.Println("- Starting init book")

  bookName := args[0]
  userName := args[1]

  bookAsBytes, err := stub.GetState(bookName)
  if err != nil {
    return nil, errors.New("Failed to get the Book Details")
  }

  r := Book{}
  json.Unmarshal(bookAsBytes, &r)
  if r.bookName == bookName {
    fmt.Println("Book already exists: " + bookName)
    fmt.Println(r)

    return nil, errors.New("This book already exists.")
  }

  str := `{"bookName": "` + bookName + `", "userName": "`+ userName + `"}`
  err = stub.PutState(bookName, []byte(str))
  if err != nil {
    return nil, err
  }

  booksAsBytes, err := stub.GetState(bookIndexStr)
  if err != nil {
    return nil, errors.New("Failed to get the book index")
  }

  var bookIndex []string
  json.Unmarshal(booksAsBytes, &bookIndex)

  bookIndex = append(bookIndex, bookName)
  fmt.Println("Book Index:", bookIndex)
  jsonAsBytes, _ := json.Marshal(bookIndex)
  err = stub.PutState(bookIndexStr, jsonAsBytes)

  fmt.Println("- end init book")
  return nil, nil
}
