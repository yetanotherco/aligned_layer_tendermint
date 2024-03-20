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
    this.channelOpenInit = this.channelOpenInit.bind(this);
    this.channelOpenTry = this.channelOpenTry.bind(this);
    this.channelOpenAck = this.channelOpenAck.bind(this);
    this.channelOpenConfirm = this.channelOpenConfirm.bind(this);
    this.channelCloseInit = this.channelCloseInit.bind(this);
    this.channelCloseConfirm = this.channelCloseConfirm.bind(this);
    this.recvPacket = this.recvPacket.bind(this);
    this.timeout = this.timeout.bind(this);
    this.timeoutOnClose = this.timeoutOnClose.bind(this);
    this.acknowledgement = this.acknowledgement.bind(this);
  }

  (0, _createClass2["default"])(MsgClientImpl, [{
    key: "channelOpenInit",
    value: function channelOpenInit(request) {
      var data = _tx.MsgChannelOpenInit.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "ChannelOpenInit", data);
      return promise.then(function (data) {
        return _tx.MsgChannelOpenInitResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "channelOpenTry",
    value: function channelOpenTry(request) {
      var data = _tx.MsgChannelOpenTry.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "ChannelOpenTry", data);
      return promise.then(function (data) {
        return _tx.MsgChannelOpenTryResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "channelOpenAck",
    value: function channelOpenAck(request) {
      var data = _tx.MsgChannelOpenAck.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "ChannelOpenAck", data);
      return promise.then(function (data) {
        return _tx.MsgChannelOpenAckResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "channelOpenConfirm",
    value: function channelOpenConfirm(request) {
      var data = _tx.MsgChannelOpenConfirm.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "ChannelOpenConfirm", data);
      return promise.then(function (data) {
        return _tx.MsgChannelOpenConfirmResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "channelCloseInit",
    value: function channelCloseInit(request) {
      var data = _tx.MsgChannelCloseInit.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "ChannelCloseInit", data);
      return promise.then(function (data) {
        return _tx.MsgChannelCloseInitResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "channelCloseConfirm",
    value: function channelCloseConfirm(request) {
      var data = _tx.MsgChannelCloseConfirm.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "ChannelCloseConfirm", data);
      return promise.then(function (data) {
        return _tx.MsgChannelCloseConfirmResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "recvPacket",
    value: function recvPacket(request) {
      var data = _tx.MsgRecvPacket.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "RecvPacket", data);
      return promise.then(function (data) {
        return _tx.MsgRecvPacketResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "timeout",
    value: function timeout(request) {
      var data = _tx.MsgTimeout.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "Timeout", data);
      return promise.then(function (data) {
        return _tx.MsgTimeoutResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "timeoutOnClose",
    value: function timeoutOnClose(request) {
      var data = _tx.MsgTimeoutOnClose.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "TimeoutOnClose", data);
      return promise.then(function (data) {
        return _tx.MsgTimeoutOnCloseResponse.decode(new _m0.Reader(data));
      });
    }
  }, {
    key: "acknowledgement",
    value: function acknowledgement(request) {
      var data = _tx.MsgAcknowledgement.encode(request).finish();

      var promise = this.rpc.request("ibc.core.channel.v1.Msg", "Acknowledgement", data);
      return promise.then(function (data) {
        return _tx.MsgAcknowledgementResponse.decode(new _m0.Reader(data));
      });
    }
  }]);
  return MsgClientImpl;
}();

exports.MsgClientImpl = MsgClientImpl;