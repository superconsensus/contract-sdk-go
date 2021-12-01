package main

import (
	"github.com/xuperchain/contract-sdk-go/code"
	"strconv"
)

// Query类：仅仅是查询数据，读取合约数据库的内容，不涉及写数据库
func (e *erc721) Query(ctx code.Context) code.Response {
	e.setContext(ctx)
	action := string(ctx.Args()["action"])
	if action == "" {
		return code.Errors("Missing key: action")
	}
	switch action {
	case "name":
		return e.name(ctx)
	case"symbol":
		return e.symbol(ctx)
	case "balanceOf":
		// token余额
		return e.balanceOf(ctx)
	case "ownerOf":
		// 查询token的所有者
		return e.ownerOf(ctx)
	case "getApproved":
		// 授权者
		return e.getApproved(ctx)
	case "isApprovedForAll":
		// 是否将所有资产授权给第三方
		return e.isApprovedForAll(ctx)
	case "tokenURI":
		// 查询uri
		return e.tokenURI(ctx)
	default:
		return code.Errors("Invalid action " + action)
	}
}
// token名字
func (e *erc721) name(ctx code.Context) code.Response{
	e.fillNFTInfo()
	return code.OK([]byte(e.NFTInfo["name"]))
}

// token符号
func (e *erc721) symbol(ctx code.Context) code.Response{
	e.fillNFTInfo()
	return code.OK([]byte(e.NFTInfo["symbol"]))
}

// token余额
// args: _owenr 查询某个人的余额
func (e *erc721) balanceOf(ctx code.Context) code.Response{
	_owner := string(ctx.Args()["_owner"])
	if _owner == "" {
		return code.Errors("miss key : _owner")
	}
	e.fillOwnerToTokenCount()
	//require(_owner != address(0), ZERO_ADDRESS);
	return code.OK([]byte(strconv.FormatInt(e._getOwnerNFTCount(_owner), 10)))
}

// 查询token的所有者
// args: _tokenId
func (e *erc721) ownerOf(ctx code.Context) code.Response{
	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}
	e.fillIdToOwner()
	_owner, ok := e.IdToOwner[_tokenId]
	if !ok {
		return code.Errors("not valid _tokenid")
	}
	return code.OK([]byte(_owner))
}

// 查询授权者
// args: _tokenId
func (e *erc721) getApproved(ctx code.Context) code.Response{
	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}
	e.fillIdtoApproval()
	if !e.validNFTToken(_tokenId){
		return code.Errors("not validnfttoken ")
	}
	return code.OK([]byte(e.IdToApproval[_tokenId]))
}

// 查询授权信息
// args: _owner
//		_operator
func (e *erc721) isApprovedForAll(ctx code.Context) code.Response{
	//address _owner,
	//address _operator
	_owner := string(ctx.Args()["_owner"])
	if _owner == "" {
		return code.Errors("miss key : _owner")
	}
	_operator := string(ctx.Args()["_operator"])
	if _operator == "" {
		return code.Errors("miss key : _operator")
	}
	e.fillOwnerToOperators()

	operators, ok := e.OwnerToOperators[_owner]
	if !ok {
		return code.OK([]byte("false"))
	}else {
		return code.OK([]byte(strconv.FormatBool(operators[_operator])))
	}
}

// 查询uri
// args: _tokenId
func (e *erc721) tokenURI(ctx code.Context) code.Response{
	_tokenId, err := strconv.ParseInt(string(ctx.Args()["_tokenId"]), 10, 64)
	if err != nil {
		return code.Errors("_tokenId error")
	}
	e.fillIdTOTokenUri()
	return code.OK([]byte(e.IdToTokenUri[_tokenId]))
}

func (e *erc721) _getOwnerNFTCount(_owner string) int64{
	return e.OwnerToTokenCount[_owner]
}