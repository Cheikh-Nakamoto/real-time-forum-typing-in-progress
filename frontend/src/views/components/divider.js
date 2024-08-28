import { Div } from "../elements/index.js";

export class Divider{
    constructor({id,width="1px", height="1px"}){
        return new Div({
            id: id,
            className: "divider",
            style: {
                width: width,
                height:height,
                backgroundColor: "var(--bs-blue)",
            }
        })
    }
}