import { View } from "../../common/types/index.js"

export class Nav extends View {
    constructor(props){
        super(props)
        this.element = document.createElement("nav");
        this._setAttributes()
        this._appendChildren()
        this._applyStyles()
        this._setConstraints()
    }
}