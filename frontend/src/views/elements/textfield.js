import { View } from "../../common/types/index.js"

export class TextField extends View {
    constructor(props) {
        super(props);
        this.multiLine = props.multiLine || false
        this.element = this.multiLine ? document.createElement("textarea") : document.createElement("input");
        if (!this.multiLine) {
            this.element.type = props.type || "text";
        }
        this._setAttributes()
        this._applyStyles()
        this._setConstraints()
        this._setEventListeners()
    }
}