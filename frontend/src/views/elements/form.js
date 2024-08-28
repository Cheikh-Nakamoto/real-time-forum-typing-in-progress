import { View } from "../../common/types/index.js";
export class Form extends View {
  constructor(properties) {
    super(properties);
    this.element = document.createElement("form");
    this._setAttributes();
    this._appendChildren();
    this._applyStyles();
    this._setEventListeners();
    this._setConstraints();
  }
}
