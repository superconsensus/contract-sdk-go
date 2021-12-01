# erc721合约

# 功能
* 通过initialize方法，设置第nfttoken信息
  * 注意initialize在部署合约的时候默认调用，二次调用无效
* 通过invoke方法，执行不同的交易功能
  * safeTransferFrom: userA将自己的某个收藏品token_id转给userB
  * transferFrom:  userA将自己的某个收藏品token_id转给userB    
  * approve: userB替userA将赋予权限的收藏品token_id卖给userC
  * setApprovalForAll:   userA将自己的所有收藏品token_id的售卖权限授予userB
  * mint： 铸币
  * burn: 销毁
* 通过query方法，执行不同的查询功能
  * balanceOf: userA的所有收藏品的数量
  * ownerOf: 查询token的所有者
  * getApproved: 获取被授权人
  * isApprovedForAll: 是否将所有资产授权给第三方
  * tokenURI： 查询token的uri
 
 
# 使用说明
## 1 initialize
initialize在部署合约的时候调用，相当给合约初始化。
``` json
{
    "module_name": "native",      // native或wasm
    "contract_name": "erc721",    // contract name
    "method_name": "initialize",  // initialize or query or invoke
    "args": {
        "from": "dudu",           // userName
        "supply": "1,2"          //  token_ids
    }
}

```
## 2 invoke
### 2.1 safeTransferFrom
```json
{
    "module_name": "native",      // native或wasm
    "contract_name": "erc721",    // contract name
    "method_name": "invoke",      // initialize or query or invoke
    "args": {
        "action": "safeTransferFrom",
        "_from": "userA",           // userName
        "_to": "userB",          //  token_ids
        "_tokenId": "1"
    }
}
```
**功能说明**：

safeTransferFrom 实现的功能是将userA的token转移给userB，

调用这个方法时调用者是token的所有者或者是被授权的第三方

**参数说明**：

_from : _tokenId的owner

_to   :  token接收方 

_tokenId : token的id 

### 2.2 transferFrom
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "invoke",
  "args": {
    "action": "transferFrom",
    "_from": "userA",
    "_to": "userB",
    "_tokenId": "go-erc721"
  }
}
```
**功能说明**：
transferFrom 实现的功能是将userA的token转移给userB，

调用这个方法时调用者是token的所有者或者是被授权的第三方

**参数说明**：
_from : _tokenId的owner

_to   :  token接收方 

_tokenId : token的id 

### 2.3 approve
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "invoke",
  "args": {
    "action": "approve",
    "_approved": "ureB",
    "_tokenId": "1"
  }
}
```
**功能说明**

approve 授权给其他人处理token

调用者要么是token的所有者，或者已经被授权处理所有资产的第三方

**参数述说明**

_approved ： 被授权处理的token的地址， 不能是token所有者（即不能给自己授权）

_tokenId :  token的id 
### 2.4 setApprovalForAll
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "invoke",
  "args": {
    "action": "setApprovalForAll",
    "_operator": "ureB",
    "_approved": "true"
  }
}
```
功能说明：

setApprovalForAll 将调用者的所有资产授权给第三方（是一个地址）处理。

参数说明：

_operator: 被授权处理的地址

_approved: 是否授权，true表示可以处理资产，false表示不可以处理资产

### 2.5 mint
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "invoke",
  "args": {
    "action": "mint",
    "_to": "nftalpha",
    "_tokenId": "go-erc721",
    "_uri": "http://boxi.com/666"
  }
}
```
功能说明：

mint 铸造一个token,没有设置任何权限前谁都可以铸造

参数说明：

_to : 接受地址

_tokenId : tokenid 全局唯一

_uri : token资源定位符

### 2.6 burn
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "invoke",
  "args": {
    "action": "burn",
    "_tokenId": "1"
  }
}
```
功能说明：

burn 销毁一个tokenId, 调用者必须是token的所有者

参数说明：

_tokenId : token的id

### 2.7 batchBatch
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "invoke",
  "args": {
    "action": "batchMint",
    "_to": "nftalpha",
    "_tokenIds": "1,2,3",
    "_uri": "http://boxi.com/666"
  }
}
```
功能说明：

批量铸造资产，暂时没有限制个数。

参数说明：

_to : 接受地址

_tokenIds : 需要传入的id, 使用逗号隔开拼接成字符串

_uri: 资源定位符


## 3 query
### 3.1 balanceOf
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "query",
  "args": {
    "action": "balanceOf",
    "_owenr": "ureB"
  }
}
```
功能说明：

balanceOf  查询一个地址拥有的nft数量
没有给合约方法设置前，任何人都可以调用

参数说明：

_owner: 想要查询地址

### 3.2 ownerOf 
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "query",
  "args": {
    "action": "ownerOf",
    "_tokenId": "1"
  }
}
```
功能说明：

ownerOf 查询一个nfttoken的所有者,没有给合约方法设置前，任何人都可以调用

参数说明：

_tokenId ： token的id


### 3.3 getApproved
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "query",
  "args": {
    "action": "getApproved",
    "_tokenId": "1"
  }
}
```
功能说明：

ownerOf 查询一个nfttoken的所有者,没有给合约方法设置前，任何人都可以调用

参数说明：

_tokenId ： token的id


### 3.4 isApprovedForAll
````json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "query",
  "args": {
    "action": "isApprovedForAll",
    "_owner": "nftalpha",
    "_operator": "go-erc721"
  }
}
````
功能说明：

isApprovedForAll 查询 _owner 是否授权给_operator

参数说明：

_owner: token的所有者

_operator： 被授权地址

### 3.5 tokenURI
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "query",
  "args": {
    "action": "tokenURI",
    "_tokenId": "go-erc721"
  }
}
```
功能说明：

tokenURI 查询token的uri

参数说明：

_tokenId: token的id

### 3.6 name
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "query",
  "args": {
    "action": "name"
  }
}
```
功能说明：

查询nft名字

### 3.7 symbol
```json
{
  "module_name": "native",
  "contract_name": "erc721",
  "method_name": "query",
  "args": {
    "action": "name"
  }
}
```
功能说明：

查询nft象征符号