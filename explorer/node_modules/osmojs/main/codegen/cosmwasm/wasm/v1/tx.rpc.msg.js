"use strict";

var _interopRequireDefault = require("@babel/runtime/helpers/interopRequireDefault");

var _typeof = require("@babel/runtime/helpers/typeof");

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.MsgClientImpl = void 0;

var _classCallCheck2 = _interopRequireDefault(require("@babel/runtime/helpers/classCallCheck"));

var _createClass2 = _interopRequireDefault(require("@babel/runtime/helpers/createClass"));

var _defineProperty2 = _interopRequireDefault(require("@babel/runtime/helpers/defineProperty"));

var _m0 = _interopRequireWildcard(require("protobufjs/minimal"));

var _tx = require("./tx");

function _getRequireWildcardCache(nodeInterop) { if (typeof WeakMap !== "function") return null; var cacheBabelInterop = new WeakMap(); var cacheNodeInterop = new WeakMap(); return (_getRequireWildcardCache = function _getRequireWildcardCache(nodeInterop) { return nodeInterop ? cacheNodeInterop : cacheBabelInterop; })(nodeInterop); }

function _interopRequireWildcard(obj, nodeInterop) { if (!nodeInterop && obj && obj.__esModule) { return obj; } if (obj === null || _typeof(obj) !== "object" && typeof obj !== "function") { return { "default": obj }; } var cache = _getRequireWildcardCache(nodeInterop); if (cache && cache.has(obj)) { return cache.get(obj); } var newObj = {}; var hasPropertyDescriptor = Object.defineProperty && Object.getOwnPropertyDescriptor; for (var key in obj) { if (key !== "default" && Object.prototype.hasOwnProperty.call(obj, key)) { var desc = hasPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : null; if (desc && (desc.get || desc.set)) { Object.defineProperty(newObj, key, desc); } else { newObj[key] = obj[key]; } } } newObj["default"] = obj; if (cache) { cache.set(obj, newObj); } return newObj; }

var MsgClientImpl = /*#__PURE__*/function () {
  function MsgClientImpl(rpc) {
    (0, _classCallCheck2["default"])(this, MsgClientImpl);
    (0, _defineProperty2["default"])(this, "rpc", void 0);
    this.rpc = rpc;
    this.storeCode = this.storeCode.bind(this);
    this.instantiateContract = this.instantiateContract.bind(this);
    this.executeContract = this.executeContract.bind(this);
    this.migrateContract = this.migrateContract.bind(this);
    this.updateAdmin = this.updateAdmin.bind(this);
    this.clearAdmin = this.clearAdmin.bind(this);
  }

  (0, _createClass2["default"])(MsgClientImpl, [{
    key: "storeCode",
    value: function storeCode(request) {
      var data = _tx.MsgStoreCode.encode(request).finish();

      var promise = this.rpc.request("cosmwasm.wasm.v1.Msg", "StoreCode", data);
      return promise.then(function (data) {
        return _tx.MsgStoreCodeResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "instantiateContract",
    value: function instantiateContract(request) {
      var data = _tx.MsgInstantiateContract.encode(request).finish();

      var promise = this.rpc.request("cosmwasm.wasm.v1.Msg", "InstantiateContract", data);
      return promise.then(function (data) {
        return _tx.MsgInstantiateContractResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "executeContract",
    value: function executeContract(request) {
      var data = _tx.MsgExecuteContract.encode(request).finish();

      var promise = this.rpc.request("cosmwasm.wasm.v1.Msg", "ExecuteContract", data);
      return promise.then(function (data) {
        return _tx.MsgExecuteContractResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "migrateContract",
    value: function migrateContract(request) {
      var data = _tx.MsgMigrateContract.encode(request).finish();

      var promise = this.rpc.request("cosmwasm.wasm.v1.Msg", "MigrateContract", data);
      return promise.then(function (data) {
        return _tx.MsgMigrateContractResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "updateAdmin",
    value: function updateAdmin(request) {
      var data = _tx.MsgUpdateAdmin.encode(request).finish();

      var promise = this.rpc.request("cosmwasm.wasm.v1.Msg", "UpdateAdmin", data);
      return promise.then(function (data) {
        return _tx.MsgUpdateAdminResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "clearAdmin",
    value: function clearAdmin(request) {
      var data = _tx.MsgClearAdmin.encode(request).finish();

      var promise = this.rpc.request("cosmwasm.wasm.v1.Msg", "ClearAdmin", data);
      return promise.then(function (data) {
        return _tx.MsgClearAdminResponse.decode(new _m0.Reader(data));
      });
    }
  }]);
  return MsgClientImpl;
}();

exports.MsgClientImpl = MsgClientImpl;