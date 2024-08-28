import { View } from "../../common/types/index.js"
export class AuthLayout extends View {
    constructor(properties) {
        super(properties);
       this.element = document.createElement("div")
       this._appendChildren()
       this._applyStyles()
    }
}