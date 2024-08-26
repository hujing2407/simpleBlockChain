package TX

type TxInput struct {
	TxHash    []byte // tx hash
	Vout      int64  // the index of Storing TxOutput into Vout
	ScriptSig string // user_name
}
