import { View } from "../../common/types/index.js"

export class MaterialIcon extends View{
    constructor(props){
        props.className += ' material-symbols-outlined'
        super(props)
        this.element = document.createElement('span');
        this.element.innerText = props.iconName;
        this._setAttributes()
        this._applyStyles()
        this._setConstraints()
        this._setEventListeners()

    }
}