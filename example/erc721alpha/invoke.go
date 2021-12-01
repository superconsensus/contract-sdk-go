package main

import (
	"github.com/xuperchain/contract-sdk-go/code"
	"strconv"
	"strings"
)

// Invoke类：执行会产生新的数据，或者会发生数据的转移
func (e *erc721) Invoke(ctx code.Context) code.Response {
	e.setContext(ctx)
	action := string(ctx.Args()["action"])
	if action == "" {
		return code.Errors("Missing key: action")
	}
	switch action {
	case "safeTransferFrom":
		return e.safeTransferFrom(ctx)
	case "transferFrom":
		// 转帐
		return e.transferForm(ctx)
	case "approve":
		// 授权给某个地址
		return e.approve(ctx)
	case "setApprovalForAll":
		// 设置将所有资产授权给第三方
		return e.setApprovalForAll(ctx)
	case "mint":
		// 铸造
		return e.mint(ctx)
	case "burn":
		return e.burn(ctx)
	case "batchMint":
		// 批量铸造
		return e.batchMint(ctx)
	default:
		return code.Errors("Invalid action " + action)

	}
}
// 转账
func (e *erc721) safeTransferFrom(ctx code.Context) code.Response{
	_from := string(ctx.Args()["_from"])
	if _from == "" {
		return code.Errors("miss key : _from")
	}
	_to := string(ctx.Args()["_to"])
	if _to == "" {
		return code.Errors("miss key : _to")
	}
	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}

	e.fillIdToOwner()
	e.fillIdtoApproval()
	e.fillOwnerToOperators()
	e.fillOwnerToTokenCount()

	if !e._safeTransferFrom(_from, _to, _tokenId){
		code.Errors("_safeTransferFrom error")
	}

	e.commitIdToOwner()
	e.commitIdToApproval()
	e.commitOwenrToOperators()
	e.commitOwnerToTokenCount()

	return code.OK([]byte("safeTransferFrom success"))
}

// 转账
func (e *erc721) transferForm(ctx code.Context) code.Response{
	_from := string(ctx.Args()["_from"])
	if _from == "" {
		return code.Errors("miss key : _from")
	}
	_to := string(ctx.Args()["_to"])
	if _to == "" {
		return code.Errors("miss key : _to")
	}
	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}
	e.fillIdToOwner()
	e.fillIdtoApproval()
	e.fillOwnerToOperators()
	e.fillOwnerToTokenCount()

	if !e.canTransfer(_tokenId) {
		return code.Errors("can not transfer")
	}
	if !e.validNFTToken(_tokenId){
		return code.Errors("not validnfttoken ")
	}

	tokenOwner := e.IdToOwner[_tokenId]
	if tokenOwner != _from {
		return code.Errors("not owner")
	}

	e._transfer(_to, _tokenId)

	e.commitIdToOwner()
	e.commitIdToApproval()
	e.commitOwenrToOperators()
	e.commitOwnerToTokenCount()
	return code.OK(nil)
}

// 授权
func (e *erc721) approve(ctx code.Context) code.Response{

	_approved := string(ctx.Args()["_approved"])
	if _approved == "" {
		return code.Errors("miss key : _approved")
	}

	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}


	e.fillIdToOwner()
	e.fillIdtoApproval()
	e.fillOwnerToOperators()

	if !e.canOperate(_tokenId) {
		return code.Errors("msg.send can not operate")
	}

	if !e.validNFTToken(_tokenId){
		return code.Errors("not valid nfttoken ")
	}

	tokenOwner := e.IdToOwner[_tokenId]
	if tokenOwner == _approved {
		return code.Errors("_approved is not myself")
	}

	e.IdToApproval[_tokenId] = _approved;

	e.commitIdToOwner()
	e.commitIdToApproval()
	e.commitOwenrToOperators()

	return code.OK(nil)
}

// 设置第三方的权力，能否可以操作nfttoken
// Enables or disables approval for a third party ("operator")
func (e *erc721) setApprovalForAll(ctx code.Context) code.Response{
	//address _operator
	//bool _approved
	// 设置调用者的
	_operator := string(ctx.Args()["_operator"])
	if _operator == "" {
		return code.Errors("miss key : _operator")
	}

	//  Enables or disables
	approved := string(ctx.Args()["_approved"])
	if approved == "" {
		return code.Errors("miss key : _approved")
	}
	_approved, err := strconv.ParseBool(approved)
	if err != nil {
		return code.Errors("_approved error")
	}

	e.fillOwnerToOperators()
	// 是否回空指针,写回会发生，读不会

	//e.OwnerToOperators[e.ctx.Initiator()][_operator] = _approved
	operators,ok := e.OwnerToOperators[e.ctx.Initiator()];
	if  !ok {
		// 没有数据
		operators = map[string]bool{}
		operators[_operator] = _approved
	}else{
		operators[_operator] = _approved
	}
	e.OwnerToOperators[e.ctx.Initiator()] = operators

	e.commitOwenrToOperators()
	return code.OK(nil)
}

// mint
func (e *erc721) mint(ctx code.Context) code.Response{
	_to := string(ctx.Args()["_to"])
	if _to == "" {
		return code.Errors("miss key : _to")
	}

	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}

	_uri := string(ctx.Args()["_uri"])
	if _uri == "" {
		return code.Errors("miss key : _uri")
	}

	e.fillIdToOwner()
	e.fillIdTOTokenUri()
	e.fillOwnerToTokenCount()

	if _, ok := e.IdToOwner[_tokenId]; ok {
		return code.Errors("_tokenId exist")
	}
	e._addNFToken(_to, _tokenId)
	e._setTokenUri(_tokenId, _uri)

	e.commitIdToOwner()
	e.commitIdToTokenURI()
	e.commitOwnerToTokenCount()
	return code.OK(nil)
}

// batch
func (e *erc721) batchMint(ctx code.Context) code.Response{
	_to := string(ctx.Args()["_to"])
	if _to == "" {
		return code.Errors("miss key : _to")
	}

	tokenIdsStr := string(ctx.Args()["_tokenIds"])
	if tokenIdsStr == "" {
		return code.Errors("Missing key: supply")
	}

	_tokenIds := make([]int64,0)
	for _, s := range strings.Split(tokenIdsStr, ","){
		num, _ := strconv.ParseInt(s, 10, 64)
		_tokenIds = append(_tokenIds, num)
	}


	_uri := string(ctx.Args()["_uri"])
	if _uri == "" {
		return code.Errors("miss key : _uri")
	}

	e.fillIdToOwner()
	e.fillIdTOTokenUri()
	e.fillOwnerToTokenCount()

	for i :=0; i < len(_tokenIds); i++{
		if _, ok := e.IdToOwner[_tokenIds[i]]; ok {
			return code.Errors("_tokenId exist")
		}
		e._addNFToken(_to, _tokenIds[i])
		e._setTokenUri(_tokenIds[i], _uri)
	}

	e.commitIdToOwner()
	e.commitIdToTokenURI()
	e.commitOwnerToTokenCount()

	return code.OK(nil)
}

// burn
func (e *erc721) burn(ctx code.Context) code.Response{

	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}

	e.fill()
	tokenOwner, ok := e.IdToOwner[_tokenId]
	if !ok {
		code.Errors("tokenId error")
	}
	if tokenOwner != e.ctx.Initiator() {
		code.Errors("调用者必选是token所有者")
	}
	e._clearApproval(_tokenId);
	e._removeNFToken(tokenOwner, _tokenId);
	delete(e.IdToTokenUri, _tokenId)

	e.commit()

	return code.OK(nil)
}





func (e *erc721) _safeTransferFrom(_from, _to string, _tokenId int64) bool {
	if !e.canTransfer(_tokenId) {
		return false
	}
	if !e.validNFTToken(_tokenId){
		return false
	}

	tokenOwner, _:= e.IdToOwner[_tokenId]
	if tokenOwner != _from {
		return false
	}
	if _to == ""{
		return false
	}
	e._transfer(_to, _tokenId)
	return true
}
func (e *erc721) _transfer(_to string, _tokenId int64) {

	from := e.IdToOwner[_tokenId]
	e._clearApproval(_tokenId)
	e._removeNFToken(from, _tokenId)
	e._addNFToken(_to, _tokenId)
	//_clearApproval(_tokenId);
	//_removeNFToken(from, _tokenId);
	//_addNFToken(_to, _tokenId);
}
func (e *erc721) _clearApproval(_tokenId int64)  {
	delete(e.IdToApproval, _tokenId)
}
func (e *erc721) _removeNFToken(_from string, _tokenId int64){
	//require(idToOwner[_tokenId] == _from, NOT_OWNER);
	//ownerToNFTokenCount[_from] -= 1;
	//delete idToOwner[_tokenId];
	tokeOwnr, _ := e.IdToOwner[_tokenId]
	if tokeOwnr != _from {
		return
	}
	e.OwnerToTokenCount[_from] -= 1
	delete(e.IdToOwner, _tokenId)
}
func (e *erc721) _addNFToken(_to string, _tokenId int64){
	if _, ok := e.IdToOwner[_tokenId]; ok {
		// tokenid 存在
		return
	}
	//idToOwner[_tokenId] = _to;
	//ownerToNFTokenCount[_to] += 1;
	e.IdToOwner[_tokenId] = _to
	e.OwnerToTokenCount[_to] += 1
}
func (e *erc721) _setTokenUri(_tokenId int64, _uri string){
	e.IdToTokenUri[_tokenId] = _uri
}
