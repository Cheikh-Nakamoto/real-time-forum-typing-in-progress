import { Div } from "../../elements/index.js"
import { Dot } from "./Dot.js"

export class TypingIndicator {
    constructor() { }

    get element() {
        return new Div({
            id:"typing",
            className: "typing-indicator incoming",
            style: {
                padding: "15px",
                display:"none",
                flexDirection:"row",
                backgroundColor:"var(--bs-gray)",
                borderRadius:"15%"
            },
            children:[
                new Dot(),
                new Dot(),
                new Dot()
            ]
        }).element
    }
}
