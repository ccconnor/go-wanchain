

//-------INPUT PARAMS YOU SHOULD MODIFY TO YOURS--------------------

// This is update validator address.
var validatorAddr = ""

// Next period lock epochs.
var lockTime = 0

// baseAddr is the fund source account.(wallet)
var baseAddr = ""

// passwd is the fund source account password.
var passwd = ""

//-------INPUT PARAMS YOU SHOULD MODIFY TO YOURS--------------------

//------------------RUN CODE DO NOT MODIFY------------------
/////////////////////////////////register staker////////////////////////////////////////////////////////////////////////
personal.unlockAccount(baseAddr, passwd)
var cscDefinition = [{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"stakeAppend","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"},{"name":"lockEpochs","type":"uint256"}],"name":"stakeUpdate","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"secPk","type":"bytes"},{"name":"bn256Pk","type":"bytes"},{"name":"lockEpochs","type":"uint256"},{"name":"feeRate","type":"uint256"}],"name":"stakeIn","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"},{"name":"renewal","type":"bool"}],"name":"partnerIn","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"delegateAddress","type":"address"}],"name":"delegateIn","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"delegateAddress","type":"address"}],"name":"delegateOut","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}];
var contractDef = eth.contract(cscDefinition);
var cscContractAddr = "0x00000000000000000000000000000000000000DA";
var coinContract = contractDef.at(cscContractAddr);
var payload5 = coinContract.stakeUpdate.getData(validatorAddr, lockTime)
console.log("payload5: ", payload5)
var tx = eth.sendTransaction({from:baseAddr, to:cscContractAddr, value:'0x00', data:payload5, gas: 200000, gasprice:'0x' + (200000000000).toString(16)});
console.log("tx5= " + tx)
/////////////////////////////////unregister staker//////////////////////////////////////////////////////////////////////
