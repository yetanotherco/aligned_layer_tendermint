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
    this.connectionOpenInit = this.connectionOpenInit.bind(this);
    this.connectionOpenTry = this.connectionOpenTry.bind(this);
    this.connectionOpenAck = this.connectionOpenAck.bind(this);
    this.connectionOpenConfirm = this.connectionOpenConfirm.bind(this);
  }

  (0, _createClass2["default"])(MsgClientImpl, [{
    key: "connectionOpenInit",
    value: function connectionOpenInit(request) {
      var data = _tx.MsgConnectionOpenInit.encode(request).finish();

      var promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenInit", data);
      return promise.then(function (data) {
        return _tx.MsgConnectionOpenInitResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "connectionOpenTry",
    value: function connectionOpenTry(request) {
      var data = _tx.MsgConnectionOpenTry.encode(request).finish();

      var promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenTry", data);
      return promise.then(function (data) {
        return _tx.MsgConnectionOpenTryResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "connectionOpenAck",
    value: function connectionOpenAck(request) {
      var data = _tx.MsgConnectionOpenAck.encode(request).finish();

      var promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenAck", data);
      return promise.then(function (data) {
        return _tx.MsgConnectionOpenAckResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "connectionOpenConfirm",
    value: function connectionOpenConfirm(request) {
      var data = _tx.MsgConnectionOpenConfirm.encode(request).finish();

      var promise = this.rpc.request("ibc.core.connection.v1.Msg", "ConnectionOpenConfirm", data);
      return promise.then(function (data) {
        return _tx.MsgConnectionOpenConfirmResponse.decode(new _m0.Reader(data));
      });
    }
  }]);
  return MsgClientImpl;
}();

exports.MsgClientImpl = MsgClientImpl;