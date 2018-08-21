package compiler

var syscallss = map[string]map[string]string{
	"storage": {
		"GetContext": "System.Storage.GetContext",
		"Put":        "System.Storage.Put",
		"Get":        "System.Storage.Get",
		"Delete":     "System.Storage.Delete",
		"Find":       "System.Storage.Find",
	},
	"runtime": {
		"GetTrigger":   "System.Runtime.GetTrigger",
		"CheckWitness": "System.Runtime.CheckWitness",
		"Notify":       "System.Runtime.Notify",
		"Log":          "System.Runtime.Log",
		"GetTime":      "System.Runtime.GetTime",
		"Serialize":    "System.Runtime.Serialize",
		"Deserialize":  "System.Runtime.Deserialize",
	},
	"blockchain": {
		"GetHeight":      "System.Blockchain.GetHeight",
		"GetHeader":      "System.Blockchain.GetHeader",
		"GetBlock":       "System.Blockchain.GetBlock",
		"GetTransaction": "System.Blockchain.GetTransaction",
		"GetContract":    "System.Blockchain.GetContract",
		"GetAccount":     "Neo.Blockchain.GetAccount",
		"GetValidators":  "Neo.Blockchain.GetValidators",
		"GetAsset":       "Neo.Blockchain.GetAsset",
	},
	"header": {
		"GetIndex":         "System.Header.GetIndex",
		"GetHash":          "System.Header.GetHash",
		"GetPrevHash":      "System.Header.GetPrevHash",
		"GetTimestamp":     "System.Header.GetTimestamp",
		"GetVersion":       "Neo.Header.GetVersion",
		"GetMerkleRoot":    "Neo.Header.GetMerkleRoot",
		"GetConsensusData": "Neo.Header.GetConsensusData",
		"GetNextConsensus": "Neo.Header.GetNextConsensus",
	},
	"block": {
		"GetTransactionCount": "System.Block.GetTransactionCount",
		"GetTransactions":     "System.Block.GetTransactions",
		"GetTransaction":      "System.Block.GetTransaction",
	},
	"transaction": {
		"GetType":         "Neo.Transaction.GetType",
		"GetAttributes":   "Neo.Transaction.GetAttributes",
		"GetInputs":       "Neo.Transaction.GetInputs",
		"GetOutputs":      "Neo.Transaction.GetOutputs",
		"GetReferences":   "Neo.Transaction.GetReferences",
		"GetUnspentCoins": "Neo.Transaction.GetUnspentCoins",
		"GetScript":       "Neo.Transaction.GetScript",
	},
	"asset": {
		"GetAssetID":   "Neo.Asset.GetAssetID",
		"GetAssetType": "Neo.Asset.GetAssetType",
		"GetAmount":    "Neo.Asset.GetAmount",
	},
	"contract": {
		"GetScript": "Neo.Contract.GetScript",
		"IsPayable": "Neo.Contract.IsPayable",
	},
	"input": {
		"GetHash":  "Neo.Input.GetHash",
		"GetIndex": "Neo.Input.GetIndex",
	},
	"output": {
		"GetAssetID", "Neo.Output.GetAssetID",
		"GetValue", "Neo.Output.GetValue",
		"GetScriptHash", "Neo.Output.GetScriptHash",
	},
}
