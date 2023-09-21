//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package eth

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AccessList) DeepCopyInto(out *AccessList) {
	{
		in := &in
		*out = make(AccessList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessList.
func (in AccessList) DeepCopy() AccessList {
	if in == nil {
		return nil
	}
	out := new(AccessList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessListEntry) DeepCopyInto(out *AccessListEntry) {
	*out = *in
	if in.StorageKeys != nil {
		in, out := &in.StorageKeys, &out.StorageKeys
		*out = make([]Data32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessListEntry.
func (in *AccessListEntry) DeepCopy() *AccessListEntry {
	if in == nil {
		return nil
	}
	out := new(AccessListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlobsBundleV1) DeepCopyInto(out *BlobsBundleV1) {
	*out = *in
	if in.Blobs != nil {
		in, out := &in.Blobs, &out.Blobs
		*out = make([]Data, len(*in))
		copy(*out, *in)
	}
	if in.Commitments != nil {
		in, out := &in.Commitments, &out.Commitments
		*out = make([]Data, len(*in))
		copy(*out, *in)
	}
	if in.Proofs != nil {
		in, out := &in.Proofs, &out.Proofs
		*out = make([]Data, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlobsBundleV1.
func (in *BlobsBundleV1) DeepCopy() *BlobsBundleV1 {
	if in == nil {
		return nil
	}
	out := new(BlobsBundleV1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Block) DeepCopyInto(out *Block) {
	*out = *in
	if in.Number != nil {
		in, out := &in.Number, &out.Number
		*out = (*in).DeepCopy()
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(Data32)
		**out = **in
	}
	in.Difficulty.DeepCopyInto(&out.Difficulty)
	in.TotalDifficulty.DeepCopyInto(&out.TotalDifficulty)
	in.Size.DeepCopyInto(&out.Size)
	in.GasLimit.DeepCopyInto(&out.GasLimit)
	in.GasUsed.DeepCopyInto(&out.GasUsed)
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	if in.Transactions != nil {
		in, out := &in.Transactions, &out.Transactions
		*out = make([]TxOrHash, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Uncles != nil {
		in, out := &in.Uncles, &out.Uncles
		*out = make([]Data32, len(*in))
		copy(*out, *in)
	}
	if in.BaseFeePerGas != nil {
		in, out := &in.BaseFeePerGas, &out.BaseFeePerGas
		*out = (*in).DeepCopy()
	}
	if in.WithdrawalsRoot != nil {
		in, out := &in.WithdrawalsRoot, &out.WithdrawalsRoot
		*out = new(Data32)
		**out = **in
	}
	if in.Withdrawals != nil {
		in, out := &in.Withdrawals, &out.Withdrawals
		*out = make([]Withdrawal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ParentBeaconBlockRoot != nil {
		in, out := &in.ParentBeaconBlockRoot, &out.ParentBeaconBlockRoot
		*out = new(Data32)
		**out = **in
	}
	if in.ExcessBlobGas != nil {
		in, out := &in.ExcessBlobGas, &out.ExcessBlobGas
		*out = (*in).DeepCopy()
	}
	if in.BlobGasUsed != nil {
		in, out := &in.BlobGasUsed, &out.BlobGasUsed
		*out = (*in).DeepCopy()
	}
	if in.Nonce != nil {
		in, out := &in.Nonce, &out.Nonce
		*out = new(Data8)
		**out = **in
	}
	if in.MixHash != nil {
		in, out := &in.MixHash, &out.MixHash
		*out = new(Data)
		**out = **in
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = new(string)
		**out = **in
	}
	if in.Signature != nil {
		in, out := &in.Signature, &out.Signature
		*out = new(string)
		**out = **in
	}
	if in.SealFields != nil {
		in, out := &in.SealFields, &out.SealFields
		*out = new([]Data)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Data, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Block.
func (in *Block) DeepCopy() *Block {
	if in == nil {
		return nil
	}
	out := new(Block)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockNumberOrTag) DeepCopyInto(out *BlockNumberOrTag) {
	*out = *in
	in.number.DeepCopyInto(&out.number)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockNumberOrTag.
func (in *BlockNumberOrTag) DeepCopy() *BlockNumberOrTag {
	if in == nil {
		return nil
	}
	out := new(BlockNumberOrTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockSpecifier) DeepCopyInto(out *BlockSpecifier) {
	*out = *in
	if in.Number != nil {
		in, out := &in.Number, &out.Number
		*out = (*in).DeepCopy()
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(Tag)
		**out = **in
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(Data32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockSpecifier.
func (in *BlockSpecifier) DeepCopy() *BlockSpecifier {
	if in == nil {
		return nil
	}
	out := new(BlockSpecifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Condition) DeepCopyInto(out *Condition) {
	{
		in := &in
		*out = make(Condition, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in Condition) DeepCopy() Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Hashes) DeepCopyInto(out *Hashes) {
	{
		in := &in
		*out = make(Hashes, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hashes.
func (in Hashes) DeepCopy() Hashes {
	if in == nil {
		return nil
	}
	out := new(Hashes)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Log) DeepCopyInto(out *Log) {
	*out = *in
	if in.LogIndex != nil {
		in, out := &in.LogIndex, &out.LogIndex
		*out = (*in).DeepCopy()
	}
	if in.TxIndex != nil {
		in, out := &in.TxIndex, &out.TxIndex
		*out = (*in).DeepCopy()
	}
	if in.TxHash != nil {
		in, out := &in.TxHash, &out.TxHash
		*out = new(Data32)
		**out = **in
	}
	if in.BlockHash != nil {
		in, out := &in.BlockHash, &out.BlockHash
		*out = new(Data32)
		**out = **in
	}
	if in.BlockNumber != nil {
		in, out := &in.BlockNumber, &out.BlockNumber
		*out = (*in).DeepCopy()
	}
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make([]Data32, len(*in))
		copy(*out, *in)
	}
	if in.TxLogIndex != nil {
		in, out := &in.TxLogIndex, &out.TxLogIndex
		*out = (*in).DeepCopy()
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Log.
func (in *Log) DeepCopy() *Log {
	if in == nil {
		return nil
	}
	out := new(Log)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogFilter) DeepCopyInto(out *LogFilter) {
	*out = *in
	if in.FromBlock != nil {
		in, out := &in.FromBlock, &out.FromBlock
		*out = new(BlockNumberOrTag)
		(*in).DeepCopyInto(*out)
	}
	if in.ToBlock != nil {
		in, out := &in.ToBlock, &out.ToBlock
		*out = new(BlockNumberOrTag)
		(*in).DeepCopyInto(*out)
	}
	if in.BlockHash != nil {
		in, out := &in.BlockHash, &out.BlockHash
		*out = new(Data32)
		**out = **in
	}
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make([]Address, len(*in))
		copy(*out, *in)
	}
	if in.Topics != nil {
		in, out := &in.Topics, &out.Topics
		*out = make([][]Data32, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]Data32, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogFilter.
func (in *LogFilter) DeepCopy() *LogFilter {
	if in == nil {
		return nil
	}
	out := new(LogFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewHeadsNotificationParams) DeepCopyInto(out *NewHeadsNotificationParams) {
	*out = *in
	in.Result.DeepCopyInto(&out.Result)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewHeadsNotificationParams.
func (in *NewHeadsNotificationParams) DeepCopy() *NewHeadsNotificationParams {
	if in == nil {
		return nil
	}
	out := new(NewHeadsNotificationParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewHeadsResult) DeepCopyInto(out *NewHeadsResult) {
	*out = *in
	in.Number.DeepCopyInto(&out.Number)
	in.Difficulty.DeepCopyInto(&out.Difficulty)
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = (*in).DeepCopy()
	}
	in.GasLimit.DeepCopyInto(&out.GasLimit)
	in.GasUsed.DeepCopyInto(&out.GasUsed)
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	if in.BaseFeePerGas != nil {
		in, out := &in.BaseFeePerGas, &out.BaseFeePerGas
		*out = (*in).DeepCopy()
	}
	if in.WithdrawalsRoot != nil {
		in, out := &in.WithdrawalsRoot, &out.WithdrawalsRoot
		*out = new(Data32)
		**out = **in
	}
	if in.Nonce != nil {
		in, out := &in.Nonce, &out.Nonce
		*out = new(Data8)
		**out = **in
	}
	if in.MixHash != nil {
		in, out := &in.MixHash, &out.MixHash
		*out = new(Data)
		**out = **in
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = new(string)
		**out = **in
	}
	if in.Signature != nil {
		in, out := &in.Signature, &out.Signature
		*out = new(string)
		**out = **in
	}
	if in.SealFields != nil {
		in, out := &in.SealFields, &out.SealFields
		*out = new([]Data)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Data, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewHeadsResult.
func (in *NewHeadsResult) DeepCopy() *NewHeadsResult {
	if in == nil {
		return nil
	}
	out := new(NewHeadsResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewPendingTxBodyNotificationParams) DeepCopyInto(out *NewPendingTxBodyNotificationParams) {
	*out = *in
	in.Result.DeepCopyInto(&out.Result)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewPendingTxBodyNotificationParams.
func (in *NewPendingTxBodyNotificationParams) DeepCopy() *NewPendingTxBodyNotificationParams {
	if in == nil {
		return nil
	}
	out := new(NewPendingTxBodyNotificationParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NewPendingTxNotificationParams) DeepCopyInto(out *NewPendingTxNotificationParams) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NewPendingTxNotificationParams.
func (in *NewPendingTxNotificationParams) DeepCopy() *NewPendingTxNotificationParams {
	if in == nil {
		return nil
	}
	out := new(NewPendingTxNotificationParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Quantity.
func (in *Quantity) DeepCopy() *Quantity {
	if in == nil {
		return nil
	}
	out := new(Quantity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Signature) DeepCopyInto(out *Signature) {
	*out = *in
	in.r.DeepCopyInto(&out.r)
	in.s.DeepCopyInto(&out.s)
	in.v.DeepCopyInto(&out.v)
	in.chainId.DeepCopyInto(&out.chainId)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Signature.
func (in *Signature) DeepCopy() *Signature {
	if in == nil {
		return nil
	}
	out := new(Signature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transaction) DeepCopyInto(out *Transaction) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = (*in).DeepCopy()
	}
	if in.BlockHash != nil {
		in, out := &in.BlockHash, &out.BlockHash
		*out = new(Data32)
		**out = **in
	}
	if in.BlockNumber != nil {
		in, out := &in.BlockNumber, &out.BlockNumber
		*out = (*in).DeepCopy()
	}
	in.Gas.DeepCopyInto(&out.Gas)
	in.Nonce.DeepCopyInto(&out.Nonce)
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(Address)
		**out = **in
	}
	if in.Index != nil {
		in, out := &in.Index, &out.Index
		*out = (*in).DeepCopy()
	}
	in.Value.DeepCopyInto(&out.Value)
	in.V.DeepCopyInto(&out.V)
	in.R.DeepCopyInto(&out.R)
	in.S.DeepCopyInto(&out.S)
	if in.YParity != nil {
		in, out := &in.YParity, &out.YParity
		*out = (*in).DeepCopy()
	}
	if in.GasPrice != nil {
		in, out := &in.GasPrice, &out.GasPrice
		*out = (*in).DeepCopy()
	}
	if in.MaxFeePerGas != nil {
		in, out := &in.MaxFeePerGas, &out.MaxFeePerGas
		*out = (*in).DeepCopy()
	}
	if in.MaxPriorityFeePerGas != nil {
		in, out := &in.MaxPriorityFeePerGas, &out.MaxPriorityFeePerGas
		*out = (*in).DeepCopy()
	}
	if in.StandardV != nil {
		in, out := &in.StandardV, &out.StandardV
		*out = (*in).DeepCopy()
	}
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = new(Data)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = new(Data)
		**out = **in
	}
	if in.ChainId != nil {
		in, out := &in.ChainId, &out.ChainId
		*out = (*in).DeepCopy()
	}
	if in.Creates != nil {
		in, out := &in.Creates, &out.Creates
		*out = new(Address)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(Condition)
		if **in != nil {
			in, out := *in, *out
			*out = make([]byte, len(*in))
			copy(*out, *in)
		}
	}
	if in.AccessList != nil {
		in, out := &in.AccessList, &out.AccessList
		*out = new(AccessList)
		if **in != nil {
			in, out := *in, *out
			*out = make([]AccessListEntry, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.MaxFeePerBlobGas != nil {
		in, out := &in.MaxFeePerBlobGas, &out.MaxFeePerBlobGas
		*out = (*in).DeepCopy()
	}
	if in.BlobVersionedHashes != nil {
		in, out := &in.BlobVersionedHashes, &out.BlobVersionedHashes
		*out = make(Hashes, len(*in))
		copy(*out, *in)
	}
	if in.BlobBundle != nil {
		in, out := &in.BlobBundle, &out.BlobBundle
		*out = new(BlobsBundleV1)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transaction.
func (in *Transaction) DeepCopy() *Transaction {
	if in == nil {
		return nil
	}
	out := new(Transaction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransactionReceipt) DeepCopyInto(out *TransactionReceipt) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = (*in).DeepCopy()
	}
	in.TransactionIndex.DeepCopyInto(&out.TransactionIndex)
	in.BlockNumber.DeepCopyInto(&out.BlockNumber)
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = new(Address)
		**out = **in
	}
	in.CumulativeGasUsed.DeepCopyInto(&out.CumulativeGasUsed)
	in.GasUsed.DeepCopyInto(&out.GasUsed)
	if in.ContractAddress != nil {
		in, out := &in.ContractAddress, &out.ContractAddress
		*out = new(Address)
		**out = **in
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = make([]Log, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Root != nil {
		in, out := &in.Root, &out.Root
		*out = new(Data32)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = (*in).DeepCopy()
	}
	if in.EffectiveGasPrice != nil {
		in, out := &in.EffectiveGasPrice, &out.EffectiveGasPrice
		*out = (*in).DeepCopy()
	}
	if in.BlobGasPrice != nil {
		in, out := &in.BlobGasPrice, &out.BlobGasPrice
		*out = (*in).DeepCopy()
	}
	if in.BlobGasUsed != nil {
		in, out := &in.BlobGasUsed, &out.BlobGasUsed
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransactionReceipt.
func (in *TransactionReceipt) DeepCopy() *TransactionReceipt {
	if in == nil {
		return nil
	}
	out := new(TransactionReceipt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TxOrHash) DeepCopyInto(out *TxOrHash) {
	*out = *in
	in.Transaction.DeepCopyInto(&out.Transaction)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TxOrHash.
func (in *TxOrHash) DeepCopy() *TxOrHash {
	if in == nil {
		return nil
	}
	out := new(TxOrHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Uncle) DeepCopyInto(out *Uncle) {
	*out = *in
	if in.Number != nil {
		in, out := &in.Number, &out.Number
		*out = (*in).DeepCopy()
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = new(Data32)
		**out = **in
	}
	in.Difficulty.DeepCopyInto(&out.Difficulty)
	if in.TotalDifficulty != nil {
		in, out := &in.TotalDifficulty, &out.TotalDifficulty
		*out = (*in).DeepCopy()
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = (*in).DeepCopy()
	}
	in.GasLimit.DeepCopyInto(&out.GasLimit)
	in.GasUsed.DeepCopyInto(&out.GasUsed)
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	if in.Transactions != nil {
		in, out := &in.Transactions, &out.Transactions
		*out = make([]TxOrHash, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Uncles != nil {
		in, out := &in.Uncles, &out.Uncles
		*out = make([]Data32, len(*in))
		copy(*out, *in)
	}
	if in.Nonce != nil {
		in, out := &in.Nonce, &out.Nonce
		*out = new(Data8)
		**out = **in
	}
	if in.MixHash != nil {
		in, out := &in.MixHash, &out.MixHash
		*out = new(Data)
		**out = **in
	}
	if in.Step != nil {
		in, out := &in.Step, &out.Step
		*out = new(string)
		**out = **in
	}
	if in.Signature != nil {
		in, out := &in.Signature, &out.Signature
		*out = new(string)
		**out = **in
	}
	if in.SealFields != nil {
		in, out := &in.SealFields, &out.SealFields
		*out = new([]Data)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Data, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Uncle.
func (in *Uncle) DeepCopy() *Uncle {
	if in == nil {
		return nil
	}
	out := new(Uncle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Withdrawal) DeepCopyInto(out *Withdrawal) {
	*out = *in
	in.Index.DeepCopyInto(&out.Index)
	in.ValidatorIndex.DeepCopyInto(&out.ValidatorIndex)
	in.Amount.DeepCopyInto(&out.Amount)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Withdrawal.
func (in *Withdrawal) DeepCopy() *Withdrawal {
	if in == nil {
		return nil
	}
	out := new(Withdrawal)
	in.DeepCopyInto(out)
	return out
}
