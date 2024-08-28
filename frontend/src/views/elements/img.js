import { View } from "../../common/types/index.js";

export class Image extends View{
    constructor(props){
        super(props);
        this.element = document.createElement("img");
        this.src = props.src
        this.alt = props.alt
        this.element.src = this.src;
        this.element.alt = this.alt;
        this._applyStyles()
        this._setAttributes()
        this._setConstraints()
    }
}