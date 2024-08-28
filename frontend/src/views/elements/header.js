import { View } from "../../common/types/index.js"

export class Header extends View {
    constructor(props) {
        super(props)
        this.element = document.createElement("header");
        this._setAttributes()
        this._appendChildren()
        this._applyStyles()
        this._setConstraints()
    }
}