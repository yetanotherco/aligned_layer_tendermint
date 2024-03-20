// src/magic-string.ts
import MagicStringBase from "magic-string";
var MagicString = class extends MagicStringBase {
  removeNode(node, { offset = 0 } = {}) {
    this.remove(offset + node.start, offset + node.end);
    return this;
  }
  moveNode(node, index, { offset = 0 } = {}) {
    this.move(offset + node.start, offset + node.end, index);
    return this;
  }
  sliceNode(node, { offset = 0 } = {}) {
    return this.slice(offset + node.start, offset + node.end);
  }
  sliceNodes(nodes, { offset = 0 } = {}) {
    if (nodes.length === 0)
      return "";
    return this.slice(
      offset + nodes[0].start,
      offset + nodes.slice(-1)[0].end
    );
  }
  overwriteNode(node, content, { offset = 0, ...options } = {}) {
    const _content = typeof content === "string" ? content : this.sliceNode(content);
    this.overwrite(offset + node.start, offset + node.end, _content, options);
    return this;
  }
};

export {
  MagicString
};
