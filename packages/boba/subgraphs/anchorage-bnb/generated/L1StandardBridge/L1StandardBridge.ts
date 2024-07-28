// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.

import {
  ethereum,
  JSONValue,
  TypedMap,
  Entity,
  Bytes,
  Address,
  BigInt,
} from "@graphprotocol/graph-ts";

export class ERC20BridgeFinalized extends ethereum.Event {
  get params(): ERC20BridgeFinalized__Params {
    return new ERC20BridgeFinalized__Params(this);
  }
}

export class ERC20BridgeFinalized__Params {
  _event: ERC20BridgeFinalized;

  constructor(event: ERC20BridgeFinalized) {
    this._event = event;
  }

  get localToken(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get remoteToken(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get from(): Address {
    return this._event.parameters[2].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[3].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[4].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[5].value.toBytes();
  }
}

export class ERC20BridgeInitiated extends ethereum.Event {
  get params(): ERC20BridgeInitiated__Params {
    return new ERC20BridgeInitiated__Params(this);
  }
}

export class ERC20BridgeInitiated__Params {
  _event: ERC20BridgeInitiated;

  constructor(event: ERC20BridgeInitiated) {
    this._event = event;
  }

  get localToken(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get remoteToken(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get from(): Address {
    return this._event.parameters[2].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[3].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[4].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[5].value.toBytes();
  }
}

export class ERC20DepositInitiated extends ethereum.Event {
  get params(): ERC20DepositInitiated__Params {
    return new ERC20DepositInitiated__Params(this);
  }
}

export class ERC20DepositInitiated__Params {
  _event: ERC20DepositInitiated;

  constructor(event: ERC20DepositInitiated) {
    this._event = event;
  }

  get l1Token(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get l2Token(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get from(): Address {
    return this._event.parameters[2].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[3].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[4].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[5].value.toBytes();
  }
}

export class ERC20WithdrawalFinalized extends ethereum.Event {
  get params(): ERC20WithdrawalFinalized__Params {
    return new ERC20WithdrawalFinalized__Params(this);
  }
}

export class ERC20WithdrawalFinalized__Params {
  _event: ERC20WithdrawalFinalized;

  constructor(event: ERC20WithdrawalFinalized) {
    this._event = event;
  }

  get l1Token(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get l2Token(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get from(): Address {
    return this._event.parameters[2].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[3].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[4].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[5].value.toBytes();
  }
}

export class ETHBridgeFinalized extends ethereum.Event {
  get params(): ETHBridgeFinalized__Params {
    return new ETHBridgeFinalized__Params(this);
  }
}

export class ETHBridgeFinalized__Params {
  _event: ETHBridgeFinalized;

  constructor(event: ETHBridgeFinalized) {
    this._event = event;
  }

  get from(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[2].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[3].value.toBytes();
  }
}

export class ETHBridgeInitiated extends ethereum.Event {
  get params(): ETHBridgeInitiated__Params {
    return new ETHBridgeInitiated__Params(this);
  }
}

export class ETHBridgeInitiated__Params {
  _event: ETHBridgeInitiated;

  constructor(event: ETHBridgeInitiated) {
    this._event = event;
  }

  get from(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[2].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[3].value.toBytes();
  }
}

export class ETHDepositInitiated extends ethereum.Event {
  get params(): ETHDepositInitiated__Params {
    return new ETHDepositInitiated__Params(this);
  }
}

export class ETHDepositInitiated__Params {
  _event: ETHDepositInitiated;

  constructor(event: ETHDepositInitiated) {
    this._event = event;
  }

  get from(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[2].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[3].value.toBytes();
  }
}

export class ETHWithdrawalFinalized extends ethereum.Event {
  get params(): ETHWithdrawalFinalized__Params {
    return new ETHWithdrawalFinalized__Params(this);
  }
}

export class ETHWithdrawalFinalized__Params {
  _event: ETHWithdrawalFinalized;

  constructor(event: ETHWithdrawalFinalized) {
    this._event = event;
  }

  get from(): Address {
    return this._event.parameters[0].value.toAddress();
  }

  get to(): Address {
    return this._event.parameters[1].value.toAddress();
  }

  get amount(): BigInt {
    return this._event.parameters[2].value.toBigInt();
  }

  get extraData(): Bytes {
    return this._event.parameters[3].value.toBytes();
  }
}

export class Initialized extends ethereum.Event {
  get params(): Initialized__Params {
    return new Initialized__Params(this);
  }
}

export class Initialized__Params {
  _event: Initialized;

  constructor(event: Initialized) {
    this._event = event;
  }

  get version(): i32 {
    return this._event.parameters[0].value.toI32();
  }
}

export class L1StandardBridge extends ethereum.SmartContract {
  static bind(address: Address): L1StandardBridge {
    return new L1StandardBridge("L1StandardBridge", address);
  }

  MESSENGER(): Address {
    let result = super.call("MESSENGER", "MESSENGER():(address)", []);

    return result[0].toAddress();
  }

  try_MESSENGER(): ethereum.CallResult<Address> {
    let result = super.tryCall("MESSENGER", "MESSENGER():(address)", []);
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toAddress());
  }

  OTHER_BRIDGE(): Address {
    let result = super.call("OTHER_BRIDGE", "OTHER_BRIDGE():(address)", []);

    return result[0].toAddress();
  }

  try_OTHER_BRIDGE(): ethereum.CallResult<Address> {
    let result = super.tryCall("OTHER_BRIDGE", "OTHER_BRIDGE():(address)", []);
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toAddress());
  }

  deposits(param0: Address, param1: Address): BigInt {
    let result = super.call("deposits", "deposits(address,address):(uint256)", [
      ethereum.Value.fromAddress(param0),
      ethereum.Value.fromAddress(param1),
    ]);

    return result[0].toBigInt();
  }

  try_deposits(param0: Address, param1: Address): ethereum.CallResult<BigInt> {
    let result = super.tryCall(
      "deposits",
      "deposits(address,address):(uint256)",
      [ethereum.Value.fromAddress(param0), ethereum.Value.fromAddress(param1)],
    );
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toBigInt());
  }

  l2TokenBridge(): Address {
    let result = super.call("l2TokenBridge", "l2TokenBridge():(address)", []);

    return result[0].toAddress();
  }

  try_l2TokenBridge(): ethereum.CallResult<Address> {
    let result = super.tryCall(
      "l2TokenBridge",
      "l2TokenBridge():(address)",
      [],
    );
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toAddress());
  }

  messenger(): Address {
    let result = super.call("messenger", "messenger():(address)", []);

    return result[0].toAddress();
  }

  try_messenger(): ethereum.CallResult<Address> {
    let result = super.tryCall("messenger", "messenger():(address)", []);
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toAddress());
  }

  otherBridge(): Address {
    let result = super.call("otherBridge", "otherBridge():(address)", []);

    return result[0].toAddress();
  }

  try_otherBridge(): ethereum.CallResult<Address> {
    let result = super.tryCall("otherBridge", "otherBridge():(address)", []);
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toAddress());
  }

  paused(): boolean {
    let result = super.call("paused", "paused():(bool)", []);

    return result[0].toBoolean();
  }

  try_paused(): ethereum.CallResult<boolean> {
    let result = super.tryCall("paused", "paused():(bool)", []);
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toBoolean());
  }

  superchainConfig(): Address {
    let result = super.call(
      "superchainConfig",
      "superchainConfig():(address)",
      [],
    );

    return result[0].toAddress();
  }

  try_superchainConfig(): ethereum.CallResult<Address> {
    let result = super.tryCall(
      "superchainConfig",
      "superchainConfig():(address)",
      [],
    );
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toAddress());
  }

  version(): string {
    let result = super.call("version", "version():(string)", []);

    return result[0].toString();
  }

  try_version(): ethereum.CallResult<string> {
    let result = super.tryCall("version", "version():(string)", []);
    if (result.reverted) {
      return new ethereum.CallResult();
    }
    let value = result.value;
    return ethereum.CallResult.fromValue(value[0].toString());
  }
}

export class ConstructorCall extends ethereum.Call {
  get inputs(): ConstructorCall__Inputs {
    return new ConstructorCall__Inputs(this);
  }

  get outputs(): ConstructorCall__Outputs {
    return new ConstructorCall__Outputs(this);
  }
}

export class ConstructorCall__Inputs {
  _call: ConstructorCall;

  constructor(call: ConstructorCall) {
    this._call = call;
  }
}

export class ConstructorCall__Outputs {
  _call: ConstructorCall;

  constructor(call: ConstructorCall) {
    this._call = call;
  }
}

export class BridgeERC20Call extends ethereum.Call {
  get inputs(): BridgeERC20Call__Inputs {
    return new BridgeERC20Call__Inputs(this);
  }

  get outputs(): BridgeERC20Call__Outputs {
    return new BridgeERC20Call__Outputs(this);
  }
}

export class BridgeERC20Call__Inputs {
  _call: BridgeERC20Call;

  constructor(call: BridgeERC20Call) {
    this._call = call;
  }

  get _localToken(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _remoteToken(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[2].value.toBigInt();
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[3].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[4].value.toBytes();
  }
}

export class BridgeERC20Call__Outputs {
  _call: BridgeERC20Call;

  constructor(call: BridgeERC20Call) {
    this._call = call;
  }
}

export class BridgeERC20ToCall extends ethereum.Call {
  get inputs(): BridgeERC20ToCall__Inputs {
    return new BridgeERC20ToCall__Inputs(this);
  }

  get outputs(): BridgeERC20ToCall__Outputs {
    return new BridgeERC20ToCall__Outputs(this);
  }
}

export class BridgeERC20ToCall__Inputs {
  _call: BridgeERC20ToCall;

  constructor(call: BridgeERC20ToCall) {
    this._call = call;
  }

  get _localToken(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _remoteToken(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _to(): Address {
    return this._call.inputValues[2].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[3].value.toBigInt();
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[4].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[5].value.toBytes();
  }
}

export class BridgeERC20ToCall__Outputs {
  _call: BridgeERC20ToCall;

  constructor(call: BridgeERC20ToCall) {
    this._call = call;
  }
}

export class BridgeETHCall extends ethereum.Call {
  get inputs(): BridgeETHCall__Inputs {
    return new BridgeETHCall__Inputs(this);
  }

  get outputs(): BridgeETHCall__Outputs {
    return new BridgeETHCall__Outputs(this);
  }
}

export class BridgeETHCall__Inputs {
  _call: BridgeETHCall;

  constructor(call: BridgeETHCall) {
    this._call = call;
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[0].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[1].value.toBytes();
  }
}

export class BridgeETHCall__Outputs {
  _call: BridgeETHCall;

  constructor(call: BridgeETHCall) {
    this._call = call;
  }
}

export class BridgeETHToCall extends ethereum.Call {
  get inputs(): BridgeETHToCall__Inputs {
    return new BridgeETHToCall__Inputs(this);
  }

  get outputs(): BridgeETHToCall__Outputs {
    return new BridgeETHToCall__Outputs(this);
  }
}

export class BridgeETHToCall__Inputs {
  _call: BridgeETHToCall;

  constructor(call: BridgeETHToCall) {
    this._call = call;
  }

  get _to(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[1].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[2].value.toBytes();
  }
}

export class BridgeETHToCall__Outputs {
  _call: BridgeETHToCall;

  constructor(call: BridgeETHToCall) {
    this._call = call;
  }
}

export class DepositERC20Call extends ethereum.Call {
  get inputs(): DepositERC20Call__Inputs {
    return new DepositERC20Call__Inputs(this);
  }

  get outputs(): DepositERC20Call__Outputs {
    return new DepositERC20Call__Outputs(this);
  }
}

export class DepositERC20Call__Inputs {
  _call: DepositERC20Call;

  constructor(call: DepositERC20Call) {
    this._call = call;
  }

  get _l1Token(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _l2Token(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[2].value.toBigInt();
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[3].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[4].value.toBytes();
  }
}

export class DepositERC20Call__Outputs {
  _call: DepositERC20Call;

  constructor(call: DepositERC20Call) {
    this._call = call;
  }
}

export class DepositERC20ToCall extends ethereum.Call {
  get inputs(): DepositERC20ToCall__Inputs {
    return new DepositERC20ToCall__Inputs(this);
  }

  get outputs(): DepositERC20ToCall__Outputs {
    return new DepositERC20ToCall__Outputs(this);
  }
}

export class DepositERC20ToCall__Inputs {
  _call: DepositERC20ToCall;

  constructor(call: DepositERC20ToCall) {
    this._call = call;
  }

  get _l1Token(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _l2Token(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _to(): Address {
    return this._call.inputValues[2].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[3].value.toBigInt();
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[4].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[5].value.toBytes();
  }
}

export class DepositERC20ToCall__Outputs {
  _call: DepositERC20ToCall;

  constructor(call: DepositERC20ToCall) {
    this._call = call;
  }
}

export class DepositETHCall extends ethereum.Call {
  get inputs(): DepositETHCall__Inputs {
    return new DepositETHCall__Inputs(this);
  }

  get outputs(): DepositETHCall__Outputs {
    return new DepositETHCall__Outputs(this);
  }
}

export class DepositETHCall__Inputs {
  _call: DepositETHCall;

  constructor(call: DepositETHCall) {
    this._call = call;
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[0].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[1].value.toBytes();
  }
}

export class DepositETHCall__Outputs {
  _call: DepositETHCall;

  constructor(call: DepositETHCall) {
    this._call = call;
  }
}

export class DepositETHToCall extends ethereum.Call {
  get inputs(): DepositETHToCall__Inputs {
    return new DepositETHToCall__Inputs(this);
  }

  get outputs(): DepositETHToCall__Outputs {
    return new DepositETHToCall__Outputs(this);
  }
}

export class DepositETHToCall__Inputs {
  _call: DepositETHToCall;

  constructor(call: DepositETHToCall) {
    this._call = call;
  }

  get _to(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _minGasLimit(): BigInt {
    return this._call.inputValues[1].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[2].value.toBytes();
  }
}

export class DepositETHToCall__Outputs {
  _call: DepositETHToCall;

  constructor(call: DepositETHToCall) {
    this._call = call;
  }
}

export class FinalizeBridgeERC20Call extends ethereum.Call {
  get inputs(): FinalizeBridgeERC20Call__Inputs {
    return new FinalizeBridgeERC20Call__Inputs(this);
  }

  get outputs(): FinalizeBridgeERC20Call__Outputs {
    return new FinalizeBridgeERC20Call__Outputs(this);
  }
}

export class FinalizeBridgeERC20Call__Inputs {
  _call: FinalizeBridgeERC20Call;

  constructor(call: FinalizeBridgeERC20Call) {
    this._call = call;
  }

  get _localToken(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _remoteToken(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _from(): Address {
    return this._call.inputValues[2].value.toAddress();
  }

  get _to(): Address {
    return this._call.inputValues[3].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[4].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[5].value.toBytes();
  }
}

export class FinalizeBridgeERC20Call__Outputs {
  _call: FinalizeBridgeERC20Call;

  constructor(call: FinalizeBridgeERC20Call) {
    this._call = call;
  }
}

export class FinalizeBridgeETHCall extends ethereum.Call {
  get inputs(): FinalizeBridgeETHCall__Inputs {
    return new FinalizeBridgeETHCall__Inputs(this);
  }

  get outputs(): FinalizeBridgeETHCall__Outputs {
    return new FinalizeBridgeETHCall__Outputs(this);
  }
}

export class FinalizeBridgeETHCall__Inputs {
  _call: FinalizeBridgeETHCall;

  constructor(call: FinalizeBridgeETHCall) {
    this._call = call;
  }

  get _from(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _to(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[2].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[3].value.toBytes();
  }
}

export class FinalizeBridgeETHCall__Outputs {
  _call: FinalizeBridgeETHCall;

  constructor(call: FinalizeBridgeETHCall) {
    this._call = call;
  }
}

export class FinalizeERC20WithdrawalCall extends ethereum.Call {
  get inputs(): FinalizeERC20WithdrawalCall__Inputs {
    return new FinalizeERC20WithdrawalCall__Inputs(this);
  }

  get outputs(): FinalizeERC20WithdrawalCall__Outputs {
    return new FinalizeERC20WithdrawalCall__Outputs(this);
  }
}

export class FinalizeERC20WithdrawalCall__Inputs {
  _call: FinalizeERC20WithdrawalCall;

  constructor(call: FinalizeERC20WithdrawalCall) {
    this._call = call;
  }

  get _l1Token(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _l2Token(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _from(): Address {
    return this._call.inputValues[2].value.toAddress();
  }

  get _to(): Address {
    return this._call.inputValues[3].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[4].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[5].value.toBytes();
  }
}

export class FinalizeERC20WithdrawalCall__Outputs {
  _call: FinalizeERC20WithdrawalCall;

  constructor(call: FinalizeERC20WithdrawalCall) {
    this._call = call;
  }
}

export class FinalizeETHWithdrawalCall extends ethereum.Call {
  get inputs(): FinalizeETHWithdrawalCall__Inputs {
    return new FinalizeETHWithdrawalCall__Inputs(this);
  }

  get outputs(): FinalizeETHWithdrawalCall__Outputs {
    return new FinalizeETHWithdrawalCall__Outputs(this);
  }
}

export class FinalizeETHWithdrawalCall__Inputs {
  _call: FinalizeETHWithdrawalCall;

  constructor(call: FinalizeETHWithdrawalCall) {
    this._call = call;
  }

  get _from(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _to(): Address {
    return this._call.inputValues[1].value.toAddress();
  }

  get _amount(): BigInt {
    return this._call.inputValues[2].value.toBigInt();
  }

  get _extraData(): Bytes {
    return this._call.inputValues[3].value.toBytes();
  }
}

export class FinalizeETHWithdrawalCall__Outputs {
  _call: FinalizeETHWithdrawalCall;

  constructor(call: FinalizeETHWithdrawalCall) {
    this._call = call;
  }
}

export class InitializeCall extends ethereum.Call {
  get inputs(): InitializeCall__Inputs {
    return new InitializeCall__Inputs(this);
  }

  get outputs(): InitializeCall__Outputs {
    return new InitializeCall__Outputs(this);
  }
}

export class InitializeCall__Inputs {
  _call: InitializeCall;

  constructor(call: InitializeCall) {
    this._call = call;
  }

  get _messenger(): Address {
    return this._call.inputValues[0].value.toAddress();
  }

  get _superchainConfig(): Address {
    return this._call.inputValues[1].value.toAddress();
  }
}

export class InitializeCall__Outputs {
  _call: InitializeCall;

  constructor(call: InitializeCall) {
    this._call = call;
  }
}

export class SetSuperchainConfigCall extends ethereum.Call {
  get inputs(): SetSuperchainConfigCall__Inputs {
    return new SetSuperchainConfigCall__Inputs(this);
  }

  get outputs(): SetSuperchainConfigCall__Outputs {
    return new SetSuperchainConfigCall__Outputs(this);
  }
}

export class SetSuperchainConfigCall__Inputs {
  _call: SetSuperchainConfigCall;

  constructor(call: SetSuperchainConfigCall) {
    this._call = call;
  }

  get _superchainConfig(): Address {
    return this._call.inputValues[0].value.toAddress();
  }
}

export class SetSuperchainConfigCall__Outputs {
  _call: SetSuperchainConfigCall;

  constructor(call: SetSuperchainConfigCall) {
    this._call = call;
  }
}
