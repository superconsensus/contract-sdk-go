package main

import "github.com/xuperchain/contract-sdk-go/code"


// 设置token信息
// 初始化-在部署的时候直接调用的方法
func (e *erc721) Initialize(ctx code.Context) code.Response {
	e.setContext(ctx)
	// 主要用于设置token的基本信息
	// name symbol
	_name := string(ctx.Args()["_name"])
	if _name == "" {
		return code.Errors("miss key : _name")
	}
	_symbol := string(ctx.Args()["_symbol"])
	if _symbol == "" {
		return code.Errors("miss key : _symbol")
	}
	e.fillNFTInfo()
	if _, ok := e.NFTInfo["name"]; !ok {
		e.NFTInfo["name"] = _name
	}
	if _, ok := e.NFTInfo["symbol"]; !ok {
		e.NFTInfo["symbol"] = _symbol
	}
	e.commitNFTInfo()
	return code.OK(nil)
}
