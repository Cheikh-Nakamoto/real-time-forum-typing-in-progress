import { Div } from "../../elements/index.js"

export class Dot {
    constructor(){}
    get element(){
        return new Div({
            className:'dot',
        }).element
    }
}