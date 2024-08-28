import { View } from "../../common/types/index.js"


export class RadioButton extends View { 
    constructor(props) {
        super(props);
        this.element = document.createElement("input");
        this.element.type = "radio";
        this._setAttributes();
        this._applyStyles();
        this._setConstraints();
        this._setEventListeners();
    }
}