import { Div, Text } from "../elements/index.js";

export class Logo {
    constructor() {
        return new Div({
            className: "logo-container",
            children: [
                new Div({
                    className: "logo-text",
                    children: [
                        new Text({
                            text: "Forum",
                            className: "logo-text"
                        })
                    ]
                })
            ]
        })
    }
}