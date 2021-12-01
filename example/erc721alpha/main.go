package main

import (
	"encoding/json"
	"github.com/xuperchain/contract-sdk-go/code"
	"github.com/xuperchain/contract-sdk-go/driver"
)

// 存进数据库中的相关key的设计
// tokeninfo ：<tokeninfo,<string, value>>
// idtoowner : <idtoowner,<id, owner>>
// idtouri  :  <idtouri, <id, uri>>
// idtoapproval : <idtoapproval, <id, approve>>
// ownertotokencount : <ownertotokencount, <owner, count>>
// <ownerToOperators, <owner,<operator, bool>>>
const (
	NFTINFO = "nftinfo"
	IDTOOWNER = "idtoowner"
	IDTOTOKENURI = "idtotokenuri"
	IDTOAPPROVAL = "idtoapproval"
	OWNERTOTOKENCOUNT = "ownertotokencount"
	OWNERTOOPETATORS = "ownertooperators"
)

type erc721 struct {
	// nft的相关信息
	NFTInfo map[string]string
	// idtoowner : <id, owner>
	IdToOwner map[int64]string
	// idtouri  : <id, uri>
	IdToTokenUri map[int64]string
	// idtoapproval  :   <id, approve>
	IdToApproval map[int64]string
	// ownertotokencount ： <owner, count>>
	OwnerToTokenCount map[string]int64
	// ownerToOperators ：<owner,<operator, bool>>
	OwnerToOperators map[string]map[string]bool
	// 合约运行上下文
	ctx code.Context
}

func newErc721() *erc721{
	token := &erc721{
		NFTInfo: map[string]string{},
		IdToOwner: map[int64]string{},
		IdToTokenUri: map[int64]string{},
		IdToApproval: map[int64]string{},
		OwnerToTokenCount: map[string]int64{},
		OwnerToOperators: map[string]map[string]bool{},
	}
	return token
}

// 填充结构体 --- 读数据库
func (e *erc721) fill() {
	// 1 nftinfo
	nftinfo := map[string]string{}
	nftinfoBuf, err :=  e.ctx.GetObject([]byte(NFTINFO))
	if err != nil {
		// 一般第一次读取的时候都没有数据
		e.NFTInfo = nftinfo
	}else {
		_ = json.Unmarshal(nftinfoBuf, &nftinfo)
		e.NFTInfo = nftinfo
	}

	// 2 id to owner
	idtoowner := map[int64]string{}
	idtoownerBuf, err := e.ctx.GetObject([]byte(IDTOOWNER))
	if err != nil {
		e.IdToOwner = idtoowner
	}else {
		_ = json.Unmarshal(idtoownerBuf, &idtoowner)
		e.IdToOwner = idtoowner
	}

	// 3 id to tokenuri
	idtotokenuri := map[int64]string{}
	idtotokenuriBuf, err := e.ctx.GetObject([]byte(IDTOTOKENURI))
	if err != nil {
		e.IdToTokenUri = idtotokenuri
	}else {
		_ = json.Unmarshal(idtotokenuriBuf, &idtotokenuri)
		e.IdToTokenUri = idtotokenuri
	}

	// 4 id to approval
	idtoapproval := map[int64]string{}
	idtoapprovalBuf,err := e.ctx.GetObject([]byte(IDTOAPPROVAL))
	if err != nil {
		e.IdToApproval = idtoapproval
	}else {
		_ = json.Unmarshal(idtoapprovalBuf, &idtoapproval)
		e.IdToApproval = idtoapproval
	}

	// 5 owner to tokencount
	owenrtotokencount := map[string]int64{}
	owenrtotokencountBuf, err := e.ctx.GetObject([]byte(OWNERTOTOKENCOUNT))
	if err != nil {
		e.OwnerToTokenCount = owenrtotokencount
	}else {
		_ = json.Unmarshal(owenrtotokencountBuf, &owenrtotokencount)
		e.OwnerToTokenCount = owenrtotokencount
	}
	// 6 owner to operators
	ownertooperators := map[string]map[string]bool{}
	ownertooperatorsBuf, err := e.ctx.GetObject([]byte(OWNERTOOPETATORS))
	if err != nil {
		e.OwnerToOperators = ownertooperators
	}else{
		_ = json.Unmarshal(ownertooperatorsBuf, &ownertooperators)
		e.OwnerToOperators = ownertooperators
	}
}
// 拆分更细的颗粒度
// 1 nftinfo
func (e *erc721) fillNFTInfo() {
	// 1 nftinfo
	nftinfo := map[string]string{}
	nftinfoBuf, err :=  e.ctx.GetObject([]byte(NFTINFO))
	if err != nil {
		// 一般第一次读取的时候都没有数据
		e.NFTInfo = nftinfo
	}else {
		_ = json.Unmarshal(nftinfoBuf, &nftinfo)
		e.NFTInfo = nftinfo
	}
}
// 2 id to owner
func (e *erc721) fillIdToOwner(){
	// 2 id to owner
	idtoowner := map[int64]string{}
	idtoownerBuf, err := e.ctx.GetObject([]byte(IDTOOWNER))
	if err != nil {
		e.IdToOwner = idtoowner
	}else {
		_ = json.Unmarshal(idtoownerBuf, &idtoowner)
		e.IdToOwner = idtoowner
	}
}
// 3 id to tokenuri
func (e *erc721) fillIdTOTokenUri(){
	// 3 id to tokenuri
	idtotokenuri := map[int64]string{}
	idtotokenuriBuf, err := e.ctx.GetObject([]byte(IDTOTOKENURI))
	if err != nil {
		e.IdToTokenUri = idtotokenuri
	}else {
		_ = json.Unmarshal(idtotokenuriBuf, &idtotokenuri)
		e.IdToTokenUri = idtotokenuri
	}
}
// 4 id to approval
func (e *erc721) fillIdtoApproval(){
	idtoapproval := map[int64]string{}
	idtoapprovalBuf,err := e.ctx.GetObject([]byte(IDTOAPPROVAL))
	if err != nil {
		e.IdToApproval = idtoapproval
	}else {
		_ = json.Unmarshal(idtoapprovalBuf, &idtoapproval)
		e.IdToApproval = idtoapproval
	}
}
// 5 owner to tokencount
func(e *erc721) fillOwnerToTokenCount(){
	owenrtotokencount := map[string]int64{}
	owenrtotokencountBuf, err := e.ctx.GetObject([]byte(OWNERTOTOKENCOUNT))
	if err != nil {
		e.OwnerToTokenCount = owenrtotokencount
	}else {
		_ = json.Unmarshal(owenrtotokencountBuf, &owenrtotokencount)
		e.OwnerToTokenCount = owenrtotokencount
	}
}
// 6 owner to operators
func(e *erc721) fillOwnerToOperators(){
	ownertooperators := map[string]map[string]bool{}
	ownertooperatorsBuf, err := e.ctx.GetObject([]byte(OWNERTOOPETATORS))
	if err != nil {
		e.OwnerToOperators = ownertooperators
	}else{
		_ = json.Unmarshal(ownertooperatorsBuf, &ownertooperators)
		e.OwnerToOperators = ownertooperators
	}
}


// 将数据写回数据库
func (e *erc721) commit() {
	// 1 nftinfo
	nftinfJSON, _ := json.Marshal(e.NFTInfo)
	_ = e.ctx.PutObject([]byte(NFTINFO), nftinfJSON)

	// 2 id to owner
	idtoownerJSON, _ := json.Marshal(e.IdToOwner)
	_ = e.ctx.PutObject([]byte(IDTOOWNER), idtoownerJSON)

	// 3 id to tokenuri
	idtotokenuriJSON, _ := json.Marshal(e.IdToTokenUri)
	_ = e.ctx.PutObject([]byte(IDTOTOKENURI), idtotokenuriJSON)

	// 4 id to approval
	idtoapprovalJSON, _ := json.Marshal(e.IdToApproval)
	_ = e.ctx.PutObject([]byte(IDTOAPPROVAL), idtoapprovalJSON)

	// 5 owner to tokencount
	ownertotokencountJSON, _ := json.Marshal(e.OwnerToTokenCount)
	_ = e.ctx.PutObject([]byte(OWNERTOTOKENCOUNT), ownertotokencountJSON)

	// 6 owner to operators
	ownertooperatorsJSON, _ := json.Marshal(e.OwnerToOperators)
	_ = e.ctx.PutObject([]byte(OWNERTOOPETATORS), ownertooperatorsJSON)
}
// 拆分更细的颗粒度
// 1 nftinfo
func (e *erc721) commitNFTInfo(){
	// 1 nftinfo
	nftinfJSON, _ := json.Marshal(e.NFTInfo)
	_ = e.ctx.PutObject([]byte(NFTINFO), nftinfJSON)
}
// 2 id to owner
func (e *erc721) commitIdToOwner(){
	idtoownerJSON, _ := json.Marshal(e.IdToOwner)
	_ = e.ctx.PutObject([]byte(IDTOOWNER), idtoownerJSON)
}
// 3 id to tokenuri
func (e *erc721) commitIdToTokenURI(){
	idtotokenuriJSON, _ := json.Marshal(e.IdToTokenUri)
	_ = e.ctx.PutObject([]byte(IDTOTOKENURI), idtotokenuriJSON)
}
// 4 id to approval
func (e *erc721) commitIdToApproval(){
	idtoapprovalJSON, _ := json.Marshal(e.IdToApproval)
	_ = e.ctx.PutObject([]byte(IDTOAPPROVAL), idtoapprovalJSON)
}
// 5 owner to tokencount
func (e *erc721) commitOwnerToTokenCount(){
	ownertotokencountJSON, _ := json.Marshal(e.OwnerToTokenCount)
	_ = e.ctx.PutObject([]byte(OWNERTOTOKENCOUNT), ownertotokencountJSON)
}
// 6 owner to operators
func (e *erc721) commitOwenrToOperators(){
	ownertooperatorsJSON, _ := json.Marshal(e.OwnerToOperators)
	_ = e.ctx.PutObject([]byte(OWNERTOOPETATORS), ownertooperatorsJSON)
}
// 设置上下文
func (e *erc721) setContext(ctx code.Context) {
	e.ctx = ctx
}



// 是否可以操作
func (e *erc721) canOperate(_tokenId int64) bool{
	// 拥有者
	tokenOwner, ok := e.IdToOwner[_tokenId]
	if !ok {
		// 没有这个id
		return false
	}
	// 1 调用者是所有者
	if tokenOwner == e.ctx.Initiator(){
		return true
	}


	// 2 第三方授权者
	operators, ok  := e.OwnerToOperators[tokenOwner]
	if !ok {
		return false
	}
	return operators[e.ctx.Initiator()]
}

// 是否可以转账
func (e *erc721) canTransfer(_tokenId int64) bool {
	tokenOwner, ok := e.IdToOwner[_tokenId]
	if !ok {
		// 没有这个 _tokenId
		return false
	}
	// 1 调用者是所有者
	if tokenOwner == e.ctx.Initiator(){
		return true
	}

	// 2 授权
	operator, _ := e.IdToApproval[_tokenId]
	if operator == e.ctx.Initiator(){
		return true
	}

	// 3 第三方授权
	operators, ok  := e.OwnerToOperators[tokenOwner]
	if !ok {
		return false
	}
	return operators[e.ctx.Initiator()]
}

// 是否是有效的token
func (e *erc721) validNFTToken(_tokenId int64) bool {
	tokenOwner, ok := e.IdToOwner[_tokenId]
	if !ok {
		// 没有这个 _tokenId
		return false
	}
	return tokenOwner != ""
}


func main() {
	driver.Serve(newErc721())
}

