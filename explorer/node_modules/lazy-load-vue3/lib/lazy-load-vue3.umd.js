function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); enumerableOnly && (symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; })), keys.push.apply(keys, symbols); } return keys; }

function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = null != arguments[i] ? arguments[i] : {}; i % 2 ? ownKeys(Object(source), !0).forEach(function (key) { _defineProperty(target, key, source[key]); }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)) : ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } return target; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _slicedToArray(arr, i) { return _arrayWithHoles(arr) || _iterableToArrayLimit(arr, i) || _unsupportedIterableToArray(arr, i) || _nonIterableRest(); }

function _nonIterableRest() { throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); }

function _iterableToArrayLimit(arr, i) { var _i = arr == null ? null : typeof Symbol !== "undefined" && arr[Symbol.iterator] || arr["@@iterator"]; if (_i == null) return; var _arr = []; var _n = true; var _d = false; var _s, _e; try { for (_i = _i.call(arr); !(_n = (_s = _i.next()).done); _n = true) { _arr.push(_s.value); if (i && _arr.length === i) break; } } catch (err) { _d = true; _e = err; } finally { try { if (!_n && _i["return"] != null) _i["return"](); } finally { if (_d) throw _e; } } return _arr; }

function _arrayWithHoles(arr) { if (Array.isArray(arr)) return arr; }

function _toConsumableArray(arr) { return _arrayWithoutHoles(arr) || _iterableToArray(arr) || _unsupportedIterableToArray(arr) || _nonIterableSpread(); }

function _nonIterableSpread() { throw new TypeError("Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); }

function _iterableToArray(iter) { if (typeof Symbol !== "undefined" && iter[Symbol.iterator] != null || iter["@@iterator"] != null) return Array.from(iter); }

function _arrayWithoutHoles(arr) { if (Array.isArray(arr)) return _arrayLikeToArray(arr); }

function _createForOfIteratorHelper(o, allowArrayLike) { var it = typeof Symbol !== "undefined" && o[Symbol.iterator] || o["@@iterator"]; if (!it) { if (Array.isArray(o) || (it = _unsupportedIterableToArray(o)) || allowArrayLike && o && typeof o.length === "number") { if (it) o = it; var i = 0; var F = function F() {}; return { s: F, n: function n() { if (i >= o.length) return { done: true }; return { done: false, value: o[i++] }; }, e: function e(_e2) { throw _e2; }, f: F }; } throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method."); } var normalCompletion = true, didErr = false, err; return { s: function s() { it = it.call(o); }, n: function n() { var step = it.next(); normalCompletion = step.done; return step; }, e: function e(_e3) { didErr = true; err = _e3; }, f: function f() { try { if (!normalCompletion && it.return != null) it.return(); } finally { if (didErr) throw err; } } }; }

function _unsupportedIterableToArray(o, minLen) { if (!o) return; if (typeof o === "string") return _arrayLikeToArray(o, minLen); var n = Object.prototype.toString.call(o).slice(8, -1); if (n === "Object" && o.constructor) n = o.constructor.name; if (n === "Map" || n === "Set") return Array.from(o); if (n === "Arguments" || /^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)) return _arrayLikeToArray(o, minLen); }

function _arrayLikeToArray(arr, len) { if (len == null || len > arr.length) len = arr.length; for (var i = 0, arr2 = new Array(len); i < len; i++) { arr2[i] = arr[i]; } return arr2; }

function _typeof(obj) { "@babel/helpers - typeof"; return _typeof = "function" == typeof Symbol && "symbol" == typeof Symbol.iterator ? function (obj) { return typeof obj; } : function (obj) { return obj && "function" == typeof Symbol && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj; }, _typeof(obj); }

(function (global, factory) {
  (typeof exports === "undefined" ? "undefined" : _typeof(exports)) === 'object' && typeof module !== 'undefined' ? factory(exports, require('vue')) : typeof define === 'function' && define.amd ? define(['exports', 'vue'], factory) : (global = typeof globalThis !== 'undefined' ? globalThis : global || self, factory(global.LazyLoad = {}, global.Vue));
})(this, function (exports, vue) {
  'use strict';

  var defaultKey = 'default';
  var elLazyKeySetMap = new Map();
  var lazyKeyVmSetMap = new Map();
  var lazyKeyElSetMap = new Map();
  var baseConfig = {
    error: '',
    loading: '',
    errorClassList: [],
    loadingClassList: [],
    loadedClassList: [],
    onError: undefined,
    onLoad: undefined
  };
  var config = Object.assign({
    timeout: 200,
    preLoad: 0.3,
    component: false,
    sorted: true,
    debounce: false,
    afterListen: undefined
  }, baseConfig);
  var directiveConfig = Object.assign({
    src: '',
    lazyKey: defaultKey,
    watchUpdate: true
  }, baseConfig);
  var listener = throttle(function (sorted, handleType, set, top, right, bottom, left, y, x) {
    var checkFn = top !== undefined ? function (el) {
      return inParentView(el, top, right, bottom, left, y, x);
    } : inViewPort;
    handleType === 0
    /* component */
    ? handler(sorted, set, checkFn, function (e) {
      return e.$el;
    }, updateComponentVm) : handler(sorted, set, checkFn, function (e) {
      return e;
    }, updateDirectiveEl);
  });
  window.addEventListener('scroll', listener);

  function inViewPort(el) {
    var _el$getBoundingClient = el.getBoundingClientRect(),
        left = _el$getBoundingClient.left,
        right = _el$getBoundingClient.right,
        top = _el$getBoundingClient.top,
        bottom = _el$getBoundingClient.bottom;

    return top <= window.innerHeight && bottom > 0 && left <= window.innerWidth && right > 0 ? 0
    /* in */
    : left !== 0 || top !== 0 ? 1
    /* notIn */
    : 2
    /* noView */
    ;
  }

  function inParentView(el, pTop, pRight, pBottom, pLeft, y, x) {
    var _el$getBoundingClient2 = el.getBoundingClientRect(),
        left = _el$getBoundingClient2.left,
        right = _el$getBoundingClient2.right,
        top = _el$getBoundingClient2.top,
        bottom = _el$getBoundingClient2.bottom;

    return top <= pBottom + y && bottom >= pTop - y && left <= pRight + x && right >= pLeft - x ? 0
    /* in */
    : left !== 0 || top !== 0 ? 1
    /* notIn */
    : 2
    /* noView */
    ;
  }

  function handler(sorted, targetSet, checkFn, getEl, updateFn) {
    var flag = false;

    var _iterator = _createForOfIteratorHelper(targetSet),
        _step;

    try {
      for (_iterator.s(); !(_step = _iterator.n()).done;) {
        var e = _step.value;

        if (getEl(e).compareDocumentPosition(document) & Node.DOCUMENT_POSITION_DISCONNECTED) {
          // unmount
          targetSet.delete(e);
          continue;
        }

        var viewStatus = checkFn(getEl(e));

        if (viewStatus === 0
        /* in */
        ) {
          flag = true;
          updateFn(e, targetSet);
        } else if (viewStatus === 1
        /* notIn */
        && sorted && flag) break;
      }
    } catch (err) {
      _iterator.e(err);
    } finally {
      _iterator.f();
    }
  }

  function updateDirectiveEl(el, targetElSet) {
    var _el$lazy = el.lazy,
        src = _el$lazy.src,
        loadingClassList = _el$lazy.loadingClassList,
        errorClassList = _el$lazy.errorClassList,
        error = _el$lazy.error,
        loadedClassList = _el$lazy.loadedClassList,
        onError = _el$lazy.onError,
        onLoad = _el$lazy.onLoad;
    el.setAttribute('src', src);
    el.addEventListener('error', function () {
      var _el$classList, _el$classList2, _el$classList3;

      (_el$classList = el.classList).remove.apply(_el$classList, _toConsumableArray(loadingClassList));

      (_el$classList2 = el.classList).remove.apply(_el$classList2, _toConsumableArray(loadedClassList));

      (_el$classList3 = el.classList).add.apply(_el$classList3, _toConsumableArray(errorClassList));

      if (error) el.setAttribute('src', error);
      onError === null || onError === void 0 ? void 0 : onError(el, el.lazy);
    }, {
      once: true
    });
    el.addEventListener('load', function () {
      var _el$classList4, _el$classList5, _el$classList6;

      (_el$classList4 = el.classList).remove.apply(_el$classList4, _toConsumableArray(loadingClassList));

      (_el$classList5 = el.classList).remove.apply(_el$classList5, _toConsumableArray(errorClassList));

      (_el$classList6 = el.classList).add.apply(_el$classList6, _toConsumableArray(loadedClassList));

      onLoad === null || onLoad === void 0 ? void 0 : onLoad(el, el.lazy);
    }, {
      once: true
    });
    targetElSet === null || targetElSet === void 0 ? void 0 : targetElSet.delete(el);
  }

  function updateComponentVm(vm, targetVmSet) {
    vm.isLoaded = true;
    targetVmSet === null || targetVmSet === void 0 ? void 0 : targetVmSet.delete(vm);
  }

  function addComponentRecords(vm) {
    var _vm$$props$lazyKey;

    var viewStatus = inViewPort(vm.$el);
    if (viewStatus === 0
    /* in */
    ) return updateComponentVm(vm);
    if (viewStatus === 2
    /* noView */
    ) return;
    var lazyKey = (_vm$$props$lazyKey = vm.$props.lazyKey) !== null && _vm$$props$lazyKey !== void 0 ? _vm$$props$lazyKey : defaultKey;
    var vmSet = lazyKeyVmSetMap.get(lazyKey) || new Set();
    if (vmSet.has(vm)) return;
    vmSet.add(vm);
    lazyKeyVmSetMap.set(lazyKey, vmSet);
    addListener(vm.$el.parentElement, lazyKey);
  }

  function addDirectiveRecords(el, lazyKey) {
    var _lazyKey;

    var viewStatus = inViewPort(el);
    if (viewStatus === 0
    /* in */
    ) return updateDirectiveEl(el);
    if (viewStatus === 2
    /* noView */
    ) return;
    lazyKey = (_lazyKey = lazyKey) !== null && _lazyKey !== void 0 ? _lazyKey : defaultKey;
    var elSet = lazyKeyElSetMap.get(lazyKey) || new Set();
    if (elSet.has(el)) return;
    elSet.add(el);
    lazyKeyElSetMap.set(lazyKey, elSet);
    addListener(el.parentElement, lazyKey);
  }

  function addListener(parent, lazyKey) {
    while (parent) {
      var lazyKeySet = elLazyKeySetMap.get(parent) || new Set();

      if (!lazyKeySet.has(lazyKey)) {
        elLazyKeySetMap.set(parent, lazyKeySet.add(lazyKey));
        parent.addEventListener('scroll', listener);
      } else break;

      parent = parent.parentElement;
    }
  }

  function throttle(cb) {
    var flag = false,
        lastScrollLeft = 0,
        lastScrollTop = 0,
        timer;

    var handler = function handler(sorted, event) {
      if (event && ![window, document].includes(event.target)) {
        var target = event.target;
        var targetVmSets = findSet(target, lazyKeyVmSetMap);
        var targetElSets = findSet(target, lazyKeyElSetMap);

        var _target$getBoundingCl = target.getBoundingClientRect(),
            left = _target$getBoundingCl.left,
            right = _target$getBoundingCl.right,
            top = _target$getBoundingCl.top,
            bottom = _target$getBoundingCl.bottom;

        var scrollLeft = target.scrollLeft,
            scrollTop = target.scrollTop;
        var y = Math.abs(scrollTop - lastScrollTop) * config.preLoad;
        var x = Math.abs(scrollLeft - lastScrollLeft) * config.preLoad;

        var _iterator2 = _createForOfIteratorHelper(targetVmSets),
            _step2;

        try {
          for (_iterator2.s(); !(_step2 = _iterator2.n()).done;) {
            var targetVmSet = _step2.value;
            cb(sorted, 0
            /* component */
            , targetVmSet, top, right, bottom, left, y, x);
          }
        } catch (err) {
          _iterator2.e(err);
        } finally {
          _iterator2.f();
        }

        var _iterator3 = _createForOfIteratorHelper(targetElSets),
            _step3;

        try {
          for (_iterator3.s(); !(_step3 = _iterator3.n()).done;) {
            var targetElSet = _step3.value;
            cb(sorted, 1
            /* directive */
            , targetElSet, top, right, bottom, left, y, x);
          }
        } catch (err) {
          _iterator3.e(err);
        } finally {
          _iterator3.f();
        }

        lastScrollLeft = scrollLeft;
        lastScrollTop = scrollTop;
      } else {
        var _iterator4 = _createForOfIteratorHelper(lazyKeyVmSetMap),
            _step4;

        try {
          for (_iterator4.s(); !(_step4 = _iterator4.n()).done;) {
            var _step4$value = _slicedToArray(_step4.value, 2),
                vmSet = _step4$value[1];

            cb(sorted, 0
            /* component */
            , vmSet);
          }
        } catch (err) {
          _iterator4.e(err);
        } finally {
          _iterator4.f();
        }

        var _iterator5 = _createForOfIteratorHelper(lazyKeyElSetMap),
            _step5;

        try {
          for (_iterator5.s(); !(_step5 = _iterator5.n()).done;) {
            var _step5$value = _slicedToArray(_step5.value, 2),
                elSet = _step5$value[1];

            cb(sorted, 1
            /* directive */
            , elSet);
          }
        } catch (err) {
          _iterator5.e(err);
        } finally {
          _iterator5.f();
        }
      }

      flag = false;
      config.afterListen && config.afterListen(event, lazyKeyElSetMap, lazyKeyVmSetMap);
    };

    return function (event, sorted) {
      var _sorted;

      sorted = (_sorted = sorted) !== null && _sorted !== void 0 ? _sorted : config.sorted;

      if (config.debounce) {
        clearTimeout(timer);
        timer = setTimeout(function () {
          return handler(sorted, event);
        }, config.timeout + 50); // debounce
      }

      if (flag) return;
      flag = true;
      setTimeout(function () {
        return handler(sorted, event);
      }, config.timeout); // throttle
    };
  }

  function findSet(target, map) {
    var lazyKeySet = elLazyKeySetMap.get(target);
    var res = [];

    var _iterator6 = _createForOfIteratorHelper(lazyKeySet),
        _step6;

    try {
      for (_iterator6.s(); !(_step6 = _iterator6.n()).done;) {
        var lazyKey = _step6.value;
        if (map.has(lazyKey)) res.push(map.get(lazyKey));
      }
    } catch (err) {
      _iterator6.e(err);
    } finally {
      _iterator6.f();
    }

    return res;
  }

  var LazyComponent = vue.defineComponent({
    render: function render() {
      var _this = this;

      if (!this.isLoaded) {
        vue.nextTick().then(function () {
          return addComponentRecords(_this);
        });
        return this.$slots.loading ? vue.h('div', this.$slots.loading()) : vue.h('div');
      } else {
        return this.$slots.default ? vue.h(this.$slots.default) : vue.h('div');
      }
    },
    props: ['lazyKey'],
    setup: function setup() {
      return {
        isLoaded: vue.ref(false)
      };
    }
  });
  var LazyDirective = {
    beforeMount: function beforeMount(el, _ref) {
      var value = _ref.value;
      var lazyKey;

      if (typeof value === 'string') {
        el.lazy = Object.assign(_objectSpread({}, directiveConfig), {
          src: value
        });
      } else {
        var _el$classList7;

        var loading = value.loading,
            _value$loadingClassLi = value.loadingClassList,
            loadingClassList = _value$loadingClassLi === void 0 ? [] : _value$loadingClassLi;
        if (loading) el.setAttribute('src', loading);

        (_el$classList7 = el.classList).add.apply(_el$classList7, _toConsumableArray(loadingClassList));

        el.lazy = Object.assign(_objectSpread({}, directiveConfig), value);
        lazyKey = value.lazyKey;
      }

      vue.nextTick().then(function () {
        return addDirectiveRecords(el, lazyKey);
      });
    },
    updated: function updated(el, _ref2) {
      var _el$lazy2;

      var value = _ref2.value;
      if (!((_el$lazy2 = el.lazy) !== null && _el$lazy2 !== void 0 && _el$lazy2.watchUpdate)) return;
      var oldSrc = el.lazy.src;
      var lazyKey;

      if (typeof value === 'string') {
        if (value === oldSrc) return;
        el.lazy = Object.assign(_objectSpread({}, directiveConfig), {
          src: value
        });
      } else {
        var _el$classList8;

        if (value.src === oldSrc) return;
        var loading = value.loading,
            _value$loadingClassLi2 = value.loadingClassList,
            loadingClassList = _value$loadingClassLi2 === void 0 ? [] : _value$loadingClassLi2;
        if (loading) el.setAttribute('src', loading);

        (_el$classList8 = el.classList).add.apply(_el$classList8, _toConsumableArray(loadingClassList));

        el.lazy = Object.assign(_objectSpread({}, directiveConfig), value);
        lazyKey = value.lazyKey;
      }

      addDirectiveRecords(el, lazyKey);
    }
  };
  var index = {
    install: function install(app) {
      var options = arguments.length > 1 && arguments[1] !== undefined ? arguments[1] : {};
      var v;
      Object.keys(baseConfig).forEach(function (k) {
        return (v = options[k]) !== undefined && (baseConfig[k] = v);
      });
      Object.keys(config).forEach(function (k) {
        return (v = options[k]) !== undefined && (config[k] = v);
      });
      Object.assign(directiveConfig, baseConfig);
      if (config.component) app.component('lazy-component', LazyComponent);
      app.directive('lazy', LazyDirective);
    }
  };
  exports.config = config;
  exports["default"] = index;
  exports.listener = listener;
  Object.defineProperty(exports, '__esModule', {
    value: true
  });
});
