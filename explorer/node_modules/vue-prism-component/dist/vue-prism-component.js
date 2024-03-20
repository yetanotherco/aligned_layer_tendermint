(function (global, factory) {
  typeof exports === 'object' && typeof module !== 'undefined' ? module.exports = factory(require('vue'), require('prismjs')) :
  typeof define === 'function' && define.amd ? define(['vue', 'prismjs'], factory) :
  (global = global || self, global.PrismComponent = factory(global.vue, global.Prism));
}(this, (function (vue, Prism) { 'use strict';

  Prism = Prism && Object.prototype.hasOwnProperty.call(Prism, 'default') ? Prism['default'] : Prism;

  var index = vue.defineComponent({
    props: {
      code: {
        type: String
      },
      inline: {
        type: Boolean,
        "default": false
      },
      language: {
        type: String,
        "default": 'markup'
      }
    },
    setup: function setup(props, _ref) {
      var slots = _ref.slots,
          attrs = _ref.attrs;
      var defaultSlot = slots && slots["default"] && slots["default"]() || [];
      var code = props.code || defaultSlot && defaultSlot.length ? defaultSlot[0].children : '';
      var inline = props.inline;
      var language = props.language;
      var prismLanguage = Prism.languages[language];
      var className = "language-".concat(language);

      return function () {
        if (inline) {
          return vue.h('code', {
            "class": [className],
            innerHTML: Prism.highlight(code, prismLanguage)
          });
        }

        return vue.h('pre', Object.assign({}, attrs, {
          "class": [attrs["class"], className]
        }), [vue.h('code', Object.assign({}, attrs, {
          "class": [attrs["class"], className],
          innerHTML: Prism.highlight(code, prismLanguage)
        }))]);
      };
    }
  });

  return index;

})));
